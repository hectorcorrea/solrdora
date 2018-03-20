package solr

import (
	"fmt"
	"net/url"
	"strings"
)

type SearchParams struct {
	Q  string
	Fl []string
	// fq		[]FilterQuery
	Facets  []FacetField
	Options map[string]string
}

func (params SearchParams) toSolrQueryString() string {
	qs := ""
	if params.Q != "" {
		qs += fmt.Sprintf("q=%s&", url.QueryEscape(params.Q))
	}

	if len(params.Fl) > 0 {
		qs += fmt.Sprintf("fl=%s&", strings.Join(params.Fl, ","))
	}

	if len(params.Facets) > 0 {
		qs += "facets=on&"
		for _, f := range params.Facets {
			qs += fmt.Sprintf("facet.field=%s&", url.QueryEscape(f.name))
			qs += fmt.Sprintf("f.%s.facet.mincount=1&", url.QueryEscape(f.name))
			// TODO account for facetLimit
		}
	}

	for k, v := range params.Options {
		qs += fmt.Sprintf("%s=%s&", k, url.QueryEscape(v))
	}
	return qs
}
