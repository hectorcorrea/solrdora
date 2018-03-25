// The *Raw structs are used to unmarshall the JSON from Solr
// via Go's built-in functions. They are not exposed outside
// the solr package.
package solr

type headerRaw struct {
	Status int `json:"status"`
	QTime  int `json:"QTime"`
	// Use interface{} because some params are strings and
	// others (e.g. fq) are arrays of strings.
	Params map[string]interface{} `json:"params"`
}

type dataRaw struct {
	NumFound  int        `json:"numFound"`
	Start     int        `json:"start"`
	Documents []Document `json:"docs"`
}

type errorRaw struct {
	Trace string `json:"trace"`
	Code  int    `json:"code"`
}

type responseRaw struct {
	Header      headerRaw      `json:"responseHeader"`
	Data        dataRaw        `json:"response"`
	Error       errorRaw       `json:"error"`
	FacetCounts facetCountsRaw `json:"facet_counts"`
}

type facetCountsRaw struct {
	Queries interface{}              `json:"facet_queries"`
	Fields  map[string][]interface{} `json:"facet_fields"`
}
