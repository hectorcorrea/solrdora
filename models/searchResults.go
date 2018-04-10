package models

import (
	"github.com/hectorcorrea/solr"
)

type SearchResults struct {
	Q           string
	Documents   []solr.Document
	Facets      solr.Facets
	NumFound    int
	Start       int
	Rows        int
	First       int
	Last        int
	Url         string
	UrlNoQ      string
	NextPageUrl string
	PrevPageUrl string
	Response    solr.SearchResponse
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
		Response:    resp,
		Documents:   resp.Documents,
	}

	if results.NumFound > 0 {
		results.First = results.Start + 1
		results.Last = results.First + results.Rows
		if results.Last > results.NumFound {
			results.Last = results.NumFound
		}
	}
	results.Facets.SetAddRemoveUrls(results.Url)

	if resp.Q != "*" {
		results.Q = resp.Q
		results.UrlNoQ = baseUrl + resp.UrlNoQ
	}

	return results
}
