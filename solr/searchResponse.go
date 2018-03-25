package solr

type SearchResponse struct {
	NumFound  int
	Start     int
	Documents []Document
	Facets    Facets
	Params    SearchParams
}

func NewSearchResponse(params SearchParams, raw responseRaw) SearchResponse {
	r := SearchResponse{
		NumFound:  raw.Data.NumFound,
		Start:     raw.Data.Start,
		Documents: raw.Data.Documents,
		Facets:    NewFacets(raw.FacetCounts, params.FilterQueries),
		Params:    params,
	}
	return r
}
