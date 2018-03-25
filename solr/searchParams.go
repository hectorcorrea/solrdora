package solr

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
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
// `qs` is typically req.URL.Query()
func NewSearchParams(qs url.Values) SearchParams {
	params := SearchParams{
		Q:             qsToString("q", qs, "*"),
		Rows:          qsToInt("rows", qs, 10),
		Start:         qsToInt("start", qs, 0),
		FilterQueries: NewFilterQueries(qs["fq"]),
	}
	return params
}

func qsToInt(key string, qs url.Values, defValue int) int {
	if len(qs[key]) == 0 {
		return defValue
	}

	i, err := strconv.Atoi(qs[key][0])
	if err != nil {
		return defValue
	}
	return i
}

func qsToString(key string, qs url.Values, defValue string) string {
	if len(qs[key]) == 0 {
		return defValue
	}

	value := strings.TrimSpace(qs[key][0])
	if value == "" {
		return defValue
	}
	return value
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
