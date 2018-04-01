package catalog

import (
	"gosiah/solr"
)

type Catalog struct {
	coreUrl string
}

func New(coreUrl string) Catalog {
	return Catalog{coreUrl: coreUrl}
}

func (c Catalog) Get(id string) (BibRecord, error) {
	s := solr.New(c.coreUrl)
	doc, err := s.Get(id, []string{})
	if err != nil {
		return BibRecord{}, err
	}
	return NewBibRecord(doc), nil
}

func (c Catalog) Search(params solr.SearchParams) (SearchResults, error) {
	solr := solr.New(c.coreUrl)
	resp, err := solr.Search(params)
	if err != nil {
		return SearchResults{}, err
	}

	results := NewSearchResults(resp, "/catalog?")
	return results, nil
}
