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
	baseUrl    string
}

func NewSearchResults(resp solr.SearchResponse, baseUrl string) SearchResults {
	r := SearchResults{
		NumFound: resp.NumFound,
		Facets:   resp.Facets,
		Q:        resp.Params.Q,
		Start:    resp.Params.Start,
		Rows:     resp.Params.Rows,
		baseUrl:  baseUrl,
	}

	r.Facets.SetAddRemoveUrls(r.ToUrl())

	for _, doc := range resp.Documents {
		record := NewBibRecord(doc)
		r.BibRecords = append(r.BibRecords, record)
	}
	return r
}

func (r SearchResults) ToUrl() string {
	return r.queryString(r.Q, r.Start)
}

func (r SearchResults) ToUrlNoQ() string {
	return r.queryString("", r.Start)
}

func (r SearchResults) NextPageUrl() string {
	return r.queryString(r.Q, r.Start+r.Rows)
}

func (r SearchResults) PrevPageUrl() string {
	return r.queryString(r.Q, r.Start-r.Rows)
}

func (r SearchResults) queryString(q string, start int) string {
	qs := r.baseUrl

	if q != "" {
		qs += solr.QsAddRaw("q", q)
	}

	for _, facet := range r.Facets {
		for _, value := range facet.Values {
			if value.Active {
				qs += solr.QsAddRaw("fq", facet.Field+"|"+value.Text)
			}
		}
	}

	if start > 0 {
		qs += solr.QsAddInt("start", start)
	}

	if r.Rows != 10 {
		qs += solr.QsAddInt("rows", r.Rows)
	}
	return qs
}
