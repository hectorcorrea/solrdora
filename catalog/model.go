package catalog

import (
	"gosiah/solr"
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
	Q          string
	Params     solr.SearchParams
	BibRecords []BibRecord
	Facets     solr.Facets
	NumFound   int
	Start      int
	Rows       int
	UserUrl    string
	UserUrlNoQ string
	RawUrl     string
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
	solr := solr.New(c.coreUrl)
	resp, err := solr.Search(params)
	if err != nil {
		return SearchResults{}, err
	}

	// -- move this to its own SearchResults.method
	resp.Facets.SetAddRemoveUrls("/catalog?" + resp.UserUrl)

	results := SearchResults{
		NumFound:   resp.NumFound,
		Params:     params,
		Facets:     resp.Facets,
		Q:          params.Q,     // shortcuts
		Start:      params.Start, // shortcuts
		Rows:       params.Rows,  // shortcuts
		UserUrl:    resp.UserUrl,
		UserUrlNoQ: "/catalog?" + resp.ToUserQueryStringNoQ(),
		RawUrl:     "/catalog?" + resp.UserUrl,
	}
	// --

	for _, doc := range resp.Documents {
		record := DocToRecord(doc)
		results.BibRecords = append(results.BibRecords, record)
	}
	return results, nil
}
