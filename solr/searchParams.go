package solr

import (
	"fmt"
	"net/url"
)

type SearchParams struct {
	Q     string
	Fl    []string
	Rows  int
	Start int
	// fq		[]FilterQuery
	Facets  []FacetField
	Options map[string]string
}

func (params SearchParams) toSolrQueryString() string {
	qs := ""
	if params.Q != "" {
		qs += encode("q", params.Q)
	}

	qs += encodeMany("fl", params.Fl)

	if len(params.Facets) > 0 {
		qs += "facet=on&"
		for _, f := range params.Facets {
			qs += encode("facet.field", f.Name)
			qs += fmt.Sprintf("f.%s.facet.mincount=1&", url.QueryEscape(f.Name))
			// TODO account for facetLimit
		}
	}

	if params.Start > 0 {
		qs += fmt.Sprintf("start=%d&", params.Start)
	}

	if params.Rows > 0 {
		qs += fmt.Sprintf("rows=%d&", params.Rows)
	}

	for k, v := range params.Options {
		qs += encode(k, v)
	}
	return qs
}
