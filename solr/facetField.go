package solr

type FacetValue struct {
	Text  string
	Count int
	// remove_url string
	// add_url string
}

type FacetField struct {
	Name   string
	Title  string
	Values []FacetValue
}

// Converts the raw FacetCounts from Solr into
// an array of our own FacetField type
func NewFacetFields(fc facetCountsRaw) []FacetField {
	facets := []FacetField{}
	for field, tokens := range fc.Fields {
		// tokens is an array in the form [value1, count1, value2, count2]
		// here we break it into an array of FacetValue that has specific
		// value and count properties
		values := []FacetValue{}
		for i := 0; i < len(tokens); i += 2 {
			text := tokens[i].(string)
			count := int(tokens[i+1].(float64))
			value := FacetValue{Text: text, Count: count}
			values = append(values, value)
		}
		facet := FacetField{Name: field, Title: field, Values: values}
		facets = append(facets, facet)
	}
	return facets
}
