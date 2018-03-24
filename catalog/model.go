package catalog

import (
	"gosiah/solr"
	// "log"
)

type Catalog struct {
	coreUrl string
}

type BibRecord struct {
	Bib     string
	Title   string
	Version float64
	Authors []string
}

type SearchResults struct {
	Params     solr.SearchParams
	BibRecords []BibRecord
	Facets     []solr.FacetField
	NumFound   int
	Start      int
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
	return DocToRecord(doc), nil
}

func DocToRecord(doc solr.Document) BibRecord {
	id := doc.Value("id")
	title := doc.Value("title")
	version := doc.ValueFloat("_version_")
	authors := doc.Values("authorsAll")
	return BibRecord{Bib: id, Title: title, Version: version, Authors: authors}
}

func (c Catalog) Search(params solr.SearchParams) (SearchResults, error) {
	s := solr.New(c.coreUrl)
	r, err := s.Search(params)
	if err != nil {
		return SearchResults{}, err
	}

	results := SearchResults{
		NumFound: r.NumFound,
		Params:   params,
		Facets:   r.Facets,
	}
	for _, doc := range r.Documents {
		record := DocToRecord(doc)
		results.BibRecords = append(results.BibRecords, record)
	}
	return results, nil
}
