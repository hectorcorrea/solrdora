package solr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Solr struct {
	CoreUrl string
}

func New(coreUrl string) Solr {
	return Solr{CoreUrl: coreUrl}
}

func (s Solr) Get(id string, fl []string) (Document, error) {
	url := s.CoreUrl + "/select?"
	url += encode("q", "id:"+id)
	url += encodeMany("fl", fl)

	raw, err := s.httpGet(url)
	if err != nil {
		return Document{}, err
	}

	count := len(raw.Data.Documents)
	if count == 0 {
		return Document{}, nil
	} else if count > 1 {
		msg := fmt.Sprintf("More than one document with ID %s was found", id)
		return Document{}, errors.New(msg)
	}
	return raw.Data.Documents[0], err
}

func (s Solr) Search(params SearchParams) (SearchResponse, error) {
	url := s.CoreUrl + "/select?" + params.toSolrQueryString()
	raw, err := s.httpGet(url)
	if err != nil {
		return SearchResponse{}, err
	}
	return NewSearchResponse(params, raw), err
}

func (s Solr) httpGet(url string) (responseRaw, error) {
	log.Printf("Solr URL: %s", url)
	r, err := http.Get(url)
	if err != nil {
		return responseRaw{}, err
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return responseRaw{}, err
	}

	if r.StatusCode < 200 || r.StatusCode > 299 {
		msg := fmt.Sprintf("HTTP Status: %s. ", r.Status)
		if len(body) > 0 {
			msg += fmt.Sprintf("Body: %s", body)
		}
		return responseRaw{}, errors.New(msg)
	}

	// log.Printf("Body: %s", body)

	var response responseRaw
	err = json.Unmarshal([]byte(body), &response)
	if err == nil {
		// HTTP request was successful but Solr reported an error.
		if response.Error.Trace != "" {
			msg := fmt.Sprintf("Solr Error. %#v", response.Error)
			err = errors.New(msg)
		}
	}
	return response, err
}
