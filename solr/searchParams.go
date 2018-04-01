package solr

import (
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

// NewSearchParams from a query string
// 	`qs` is typically req.URL.Query()
// 	`options` to pass to Solr (e.g. defType: "edismax")
// 	`facets` to request from Solr (e.g. fieldName: "Field Name")
func NewSearchParams(qs url.Values, options map[string]string,
	facets map[string]string) SearchParams {

	params := SearchParams{
		Q:             QsGet(qs, "q", "*"),
		Rows:          QsGetInt(qs, "rows", 10),
		Start:         QsGetInt(qs, "start", 0),
		FilterQueries: NewFilterQueries(qs["fq"]),
		Options:       options,
		Facets:        NewFacets(facets),
	}
	return params
}

func (params SearchParams) toSolrQueryString() string {
	qs := ""
	qs += QsAddDefault("q", params.Q, "*")
	qs += QsAddMany("fl", params.Fl)
	qs += params.FilterQueries.toQueryString()
	qs += params.Facets.toQueryString()

	if params.Start > 0 {
		qs += QsAddInt("start", params.Start)
	}

	if params.Rows > 0 {
		qs += QsAddInt("rows", params.Rows)
	}

	for k, v := range params.Options {
		qs += QsAdd(k, v)
	}
	return qs
}
