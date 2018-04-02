package catalog

import (
	"gosiah/solr"
)

type SearchResults struct {
	Q           string
	BibRecords  []BibRecord
	Facets      solr.Facets
	NumFound    int
	Start       int
	Rows        int
	Url         string
	UrlNoQ      string
	NextPageUrl string
	PrevPageUrl string
}

func NewSearchResults(resp solr.SearchResponse, baseUrl string) SearchResults {
	results := SearchResults{
		NumFound:    resp.NumFound,
		Facets:      resp.Facets,
		Start:       resp.Start,
		Rows:        resp.Rows,
		Url:         baseUrl + resp.Url,
		PrevPageUrl: baseUrl + resp.PrevPageUrl,
		NextPageUrl: baseUrl + resp.NextPageUrl,
	}

	results.Facets.SetAddRemoveUrls(results.Url)

	if resp.Q != "*" {
		results.Q = resp.Q
		results.UrlNoQ = baseUrl + resp.UrlNoQ
	}

	for _, doc := range resp.Documents {
		record := NewBibRecord(doc)
		results.BibRecords = append(results.BibRecords, record)
	}

	return results
}
