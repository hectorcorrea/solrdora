package solr

import (
	"fmt"
	// "log"
)

type Document map[string]interface{}

// Returns the value in a single-value field
func (d Document) Value(fieldName string) string {
	// Casting to string would have been cleaner but it _only_ works for strings.
	// Casting to interface{} allows us to fetch the value even if it is not
	// a string (e.g a float). The downsie is that fmt.Sprintf() returns a
	// funny value for non-strings, but at least we fetch the value.
	value, ok := d[fieldName].(interface{})
	if ok {
		return fmt.Sprintf("%s", value)
	}
	return ""
}

// Returns all the values in multi-value field
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

// Returns the first value in a multi-value field
func (d Document) FirstValue(fieldName string) string {
	values := d.Values(fieldName)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func (d Document) ValueFloat(fieldName string) float64 {
	value, ok := d[fieldName].(float64)
	if ok {
		return value
	}
	return 0.0
}
