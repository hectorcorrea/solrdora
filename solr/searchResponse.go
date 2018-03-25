package solr

import "log"

type SearchResponse struct {
	Q         string
	NumFound  int
	Start     int
	Documents []Document
	Facets    Facets
	Params    SearchParams
	UserUrl   string
}

func NewSearchResponse(params SearchParams, raw responseRaw) SearchResponse {
	r := SearchResponse{
		NumFound:  raw.Data.NumFound,
		Start:     raw.Data.Start,
		Documents: raw.Data.Documents,
		Facets:    NewFacetsFromResponse(raw.FacetCounts, params.FilterQueries),
		Params:    params,
	}
	r.UserUrl = r.toUserQueryString()
	return r
}

func (r SearchResponse) toUserQueryString() string {
	qs := ""
	qs += qsAdd("q", r.Params.Q)
	for _, facet := range r.Facets {
		for _, value := range facet.Values {
			if value.Active {
				qs += qsAdd("fq", facet.Field+"|"+value.Text)
			}
		}
	}
	log.Printf("qs= %s", qs)
	return qs
}

func (r SearchResponse) ToUserQueryStringNoQ() string {
	qs := ""
	for _, facet := range r.Facets {
		for _, value := range facet.Values {
			if value.Active {
				qs += qsAdd("fq", facet.Field+"|"+value.Text)
			}
		}
	}
	return qs
}
