package solr

import (
	"fmt"
	"log"
	"net/url"
)

type SearchParams struct {
	Q             string
	Fl            []string
	Rows          int
	Start         int
	FilterQueries FilterQueries
	Facets        Facets
	Options       map[string]string
}

func (params SearchParams) toSolrQueryString() string {
	qs := ""
	qs += encodeDefault("q", params.Q, "*")
	qs += encodeMany("fl", params.Fl)
	qs += params.FilterQueries.toQueryString()

	if len(params.Facets) > 0 {
		qs += "facet=on&"
		for _, f := range params.Facets {
			qs += encode("facet.field", f.Field)
			qs += fmt.Sprintf("f.%s.facet.mincount=1&", url.QueryEscape(f.Field))
			// TODO account for facetLimit
		}
	} else {
		log.Printf("no facets")
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
