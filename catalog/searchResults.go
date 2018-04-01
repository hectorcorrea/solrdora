package catalog

import (
	"gosiah/solr"
)

type SearchResults struct {
	Q          string
	BibRecords []BibRecord
	Facets     solr.Facets
	NumFound   int
	Start      int
	Rows       int
	UserUrlNoQ string
}

func NewSearchResults(resp solr.SearchResponse, baseUrl string) SearchResults {
	r := SearchResults{
		NumFound:   resp.NumFound,
		Facets:     resp.Facets,
		Q:          resp.Params.Q,
		Start:      resp.Params.Start,
		Rows:       resp.Params.Rows,
		UserUrlNoQ: baseUrl + resp.ToQueryString(false),
	}
	r.Facets.SetAddRemoveUrls(baseUrl + resp.ToQueryString(true))

	for _, doc := range resp.Documents {
		record := NewBibRecord(doc)
		r.BibRecords = append(r.BibRecords, record)
	}
	return r
}
