package solr

import (
	"fmt"
	"net/url"
	"strings"
)

// Encodes a single value as a query string parameter.
func encode(param, value string) string {
	if value == "" {
		return ""
	}
	return fmt.Sprintf("%s=%s&", param, url.QueryEscape(value))
}

func encodeDefault(param, value, defaultValue string) string {
	if value == "" {
		return encode(param, defaultValue)
	}
	return encode(param, value)
}

// Encodes an array of values as a query string parameter (the values
// are separated by commas).
func encodeMany(param string, values []string) string {
	if len(values) == 0 {
		return ""
	}
	valuesEnc := []string{}
	for _, value := range values {
		valuesEnc = append(valuesEnc, url.QueryEscape(value))
	}
	return fmt.Sprintf("%s=%s&", param, strings.Join(valuesEnc, ","))
}
