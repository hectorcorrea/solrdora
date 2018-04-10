package models

import (
	"github.com/hectorcorrea/solr"
	"net/url"
)

type Search struct {
	settings Settings
}

func NewSearch(settings Settings) Search {
	return Search{settings: settings}
}

func (search Search) Get(id string) (solr.Document, error) {
	params := solr.NewGetParams("id:"+id, search.settings.ViewOneFl, search.settings.SolrOptions)
	s := solr.New(search.settings.SolrCoreUrl, true)
	doc, err := s.Get(params)
	if err != nil {
		return solr.Document{}, err
	}
	return doc, nil
}

func (search Search) Search(qs url.Values, baseUrl string) (SearchResults, error) {
	params := solr.NewSearchParamsFromQs(qs, search.settings.SolrOptions, search.settings.SolrFacets)
	params.Fl = search.settings.SearchFl
	solr := solr.New(search.settings.SolrCoreUrl, true)
	resp, err := solr.Search(params)
	if err != nil {
		return SearchResults{}, err
	}
	results := NewSearchResults(resp, baseUrl)
	return results, nil
}
