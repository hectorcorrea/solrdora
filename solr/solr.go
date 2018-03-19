package solr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Solr struct {
	CoreUrl string
}

func New(coreUrl string) Solr {
	return Solr{CoreUrl: coreUrl}
}

func (s Solr) Get(id, fl string) (Document, error) {
	url := s.CoreUrl + "/select?"
	url += "q=id:" + id + "&"

	if fl != "" {
		url += "fl=" + fl + "&"
	}

	r, err := s.httpGet(url)
	if err != nil {
		return Document{}, err
	}

	if len(r.Response.Documents) == 0 {
		return Document{}, errors.New("no doc found")
	}
	return r.Response.Documents[0], err
}

func (s Solr) Search(q, fl string) (SolrResponse, error) {
	url := s.CoreUrl + "/select?" // q=*&fl=id,author"

	if q != "" {
		url += "q=" + q + "&"
	}

	if fl != "" {
		url += "fl=" + fl + "&"
	}

	r, err := s.httpGet(url)
	return r, err
}

func (s Solr) httpGet(url string) (SolrResponse, error) {
	r, err := http.Get(url)
	if err != nil {
		return SolrResponse{}, err
	}

	if r.StatusCode < 200 || r.StatusCode > 299 {
		return SolrResponse{}, errors.New(fmt.Sprintf("HTTP Status %s", r.Status))
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return SolrResponse{}, err
	}

	var solrResponse SolrResponse
	err = json.Unmarshal([]byte(body), &solrResponse)
	return solrResponse, err
}
