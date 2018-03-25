package solr

import (
	"fmt"
	"net/url"
)

type facetValue struct {
	Text   string
	Count  int
	Active bool
}

type facetField struct {
	Field  string
	Title  string
	Values []facetValue
}

type Facets []facetField

// Creates a new Facets object from the raw FacetCounts from Solr.
//
// `fc` contains the facet data as reported by Solr.
// `fq` contains the `fq` values (field/value) passed to Solr during the search.
func NewFacets(counts facetCountsRaw, fq FilterQueries) Facets {
	facets := Facets{}
	for field, tokens := range counts.Fields {
		// Tokens is an array in the form [value1, count1, value2, count2]
		// here we break it into an array of FacetValue that has specific
		// value and count properties. We consider the FacetValue "active"
		// if it was used in the "fq" parameters.
		facet := facetField{Field: field, Title: field}
		for i := 0; i < len(tokens); i += 2 {
			text := tokens[i].(string)
			count := int(tokens[i+1].(float64))
			// Mark the facet for this value as active if it is also present
			// on the FilterQueries
			active := fq.HasFieldValue(field, text)
			facet.addValue(text, count, active)
		}
		facets = append(facets, facet)
	}
	return facets
}

func (facets *Facets) Add(field, title string) {
	facet := facetField{Field: field, Title: field}
	*facets = append(*facets, facet)
}

func (ff *facetField) addValue(text string, count int, active bool) {
	value := facetValue{
		Text:   text,
		Count:  count,
		Active: active,
	}
	ff.Values = append(ff.Values, value)
}

func (facets Facets) toQueryString() string {
	qs := ""
	if len(facets) > 0 {
		qs += encode("facet", "on")
		for _, facet := range facets {
			qs += encode("facet.field", facet.Field)
			min_count := fmt.Sprintf("f.%s.facet.mincount", url.QueryEscape(facet.Field))
			qs += encode(min_count, "1")
			// TODO account for facetLimit
		}
	}
	return qs
}
