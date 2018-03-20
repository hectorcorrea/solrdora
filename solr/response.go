package solr

import (
	"fmt"
)

type ResponseHeader struct {
	Status int               `json:"status"`
	QTime  int               `json:"QTime"`
	Params map[string]string `json:"params"`
}

type Document map[string]interface{}

type Response struct {
	NumFound  int        `json:"numFound"`
	Start     int        `json:"start"`
	Documents []Document `json:"docs"`
}

type SolrResponse struct {
	ResponseHeader ResponseHeader `json:"responseHeader"`
	Response       Response       `json:"response"`
}

func (d Document) Value(fieldName string) string {
	value, ok := d[fieldName].(string)
	if ok {
		return value
	}
	return ""
}

func (d Document) Values(fieldName string) []string {
	var values []string
	valuesRaw, ok := d[fieldName].([]interface{})
	if ok {
		for _, v := range valuesRaw {
			values = append(values, fmt.Sprintf("%s", v))
		}
	}
	return values
}
