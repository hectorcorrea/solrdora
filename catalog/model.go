package catalog

import (
	"github.com/hectorcorrea/solr"
)

type Catalog struct {
	coreUrl string
}

func New(coreUrl string) Catalog {
	return Catalog{coreUrl: coreUrl}
}

func (c Catalog) Get(params solr.GetParams) (Result, error) {
	s := solr.New(c.coreUrl, true)
	doc, err := s.Get(params)
	if err != nil {
		return Result{}, err
	}
	return NewResult(doc), nil
}

func (c Catalog) Search(params solr.SearchParams, baseUrl string) (SearchResults, error) {
	solr := solr.New(c.coreUrl, true)
	resp, err := solr.Search(params)
	if err != nil {
		return SearchResults{}, err
	}

	results := NewSearchResults(resp, baseUrl)
	return results, nil
}
