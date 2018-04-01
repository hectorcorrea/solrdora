package solr

import "log"

type SearchResponse struct {
	Params    SearchParams
	Q         string
	NumFound  int
	Start     int
	Documents []Document
	Facets    Facets
}

func NewSearchResponse(params SearchParams, raw responseRaw) SearchResponse {
	r := SearchResponse{
		Params:    params,
		NumFound:  raw.Data.NumFound,
		Start:     raw.Data.Start,
		Documents: raw.Data.Documents,
	}
	r.Facets = r.facetsFromResponse(raw.FacetCounts)
	return r
}

// Creates a new Facets object from the raw FacetCounts from Solr.
// 	`counts` contains the facet data as reported by Solr.
func (r SearchResponse) facetsFromResponse(counts facetCountsRaw) Facets {
	facets := Facets{}
	for field, tokens := range counts.Fields {

		// Create the facetField object for the field
		facet := facetField{Field: field, Title: field}
		for _, def := range r.Params.Facets {
			if def.Field == facet.Field {
				// if the field is on the facets indicated on the params
				// (it should always be) grab the Title from it
				facet.Title = def.Title
				break
			}
		}

		// Tokens is an array in the form [value1, count1, value2, count2].
		// Here we break it into an array of FacetValues that has specific
		// value and count properties.
		for i := 0; i < len(tokens); i += 2 {
			text := tokens[i].(string)
			count := int(tokens[i+1].(float64))
			// Mark the facet for this value as active if it is
			// present on the FilterQueries
			active := r.Params.FilterQueries.HasFieldValue(field, text)
			facet.addValue(text, count, active)
		}

		facets = append(facets, facet)
	}
	return facets
}

func (r SearchResponse) ToQueryString(q bool) string {
	qs := ""
	if q {
		qs += qsAddRaw("q", r.Params.Q)
	}
	for _, facet := range r.Facets {
		for _, value := range facet.Values {
			if value.Active {
				qs += qsAddRaw("fq", facet.Field+"|"+value.Text)
			}
		}
	}
	log.Printf("qs= %s", qs)
	return qs
}
