package solr

import (
	"fmt"
)

type Document map[string]interface{}

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

func (d Document) FirstValue(fieldName string) string {
	values := d.Values(fieldName)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}
