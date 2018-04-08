package models

import (
	"fmt"
	"github.com/hectorcorrea/solr"
	"reflect"
)

type Result struct {
	Document solr.Document
}

type SearchResults struct {
	Q           string
	Results     []Result
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

	for _, doc := range resp.Documents {
		record := NewResult(doc)
		results.Results = append(results.Results, record)
	}

	return results
}

func NewResult(doc solr.Document) Result {
	return Result{Document: doc}
}

func (r Result) IsMultiValue(field string) bool {
	value := reflect.ValueOf(r.Document[field])
	return value.Kind() == reflect.Slice
}

func (r Result) Id() string {
	return fmt.Sprintf("%s", r.Document["id"])
}
