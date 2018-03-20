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

func (s Solr) Search(params SearchParams) (SolrResponse, error) {
	url := s.CoreUrl + "/select?" + params.toSolrQueryString()
	r, err := s.httpGet(url)
	return r, err
}

func (s Solr) httpGet(url string) (SolrResponse, error) {
	r, err := http.Get(url)
	if err != nil {
		return SolrResponse{}, err
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return SolrResponse{}, err
	}

	if r.StatusCode < 200 || r.StatusCode > 299 {
		msg := fmt.Sprintf("HTTP Status: %s. ", r.Status)
		if len(body) > 0 {
			msg += fmt.Sprintf("Body: %s", body)
		}
		return SolrResponse{}, errors.New(msg)
	}

	var solrResponse SolrResponse
	err = json.Unmarshal([]byte(body), &solrResponse)
	if err == nil {
		// HTTP request was successful but Solr reported an error.
		if solrResponse.Error.Trace != "" {
			msg := fmt.Sprintf("Solr Error. %#v", solrResponse.Error)
			err = errors.New(msg)
		}
	}
	return solrResponse, err
}
