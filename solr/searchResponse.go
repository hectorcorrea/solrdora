package solr

type SearchResponse struct {
	NumFound  int
	Start     int
	Documents []Document
	Facets    []FacetField
	Params    SearchParams
}

func NewSearchResponse(params SearchParams, raw responseRaw) SearchResponse {
	r := SearchResponse{
		NumFound:  raw.Data.NumFound,
		Start:     raw.Data.Start,
		Documents: raw.Data.Documents,
		Facets:    NewFacetFields(raw.FacetCounts),
		Params:    params,
	}
	return r
}
