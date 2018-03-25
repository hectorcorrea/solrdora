package solr

import (
	"fmt"
	"net/url"
)

type FacetValue struct {
	Text   string
	Count  int
	Active bool
}

type FacetField struct {
	Field  string
	Title  string
	Values []FacetValue
}

type Facets []FacetField

// Converts the raw FacetCounts from Solr into an array of our own
// with a few extra data.
//
// `fc` contains the facet data as reported by Solr.
// `fq` contains the `fq` values (field/value) passed to Solr during the search.
func NewFacets(fc facetCountsRaw, fq FilterQueries) Facets {
	facets := Facets{}
	for field, tokens := range fc.Fields {
		// Tokens is an array in the form [value1, count1, value2, count2]
		// here we break it into an array of FacetValue that has specific
		// value and count properties. We consider the FacetValue "active"
		// if it was used in the "fq" parameters.
		facet := FacetField{Field: field, Title: field}
		for i := 0; i < len(tokens); i += 2 {
			text := tokens[i].(string)
			count := int(tokens[i+1].(float64))
			active := fq.HasFieldValue(field, text)
			facet.AddValue(text, count, active)
		}
		facets = append(facets, facet)
	}
	return facets
}

func (ff *FacetField) AddValue(text string, count int, active bool) {
	value := FacetValue{
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
