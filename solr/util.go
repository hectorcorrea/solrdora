package solr

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Adds a parameter and its value to a query string
func qsAdd(param, value string) string {
	if value == "" {
		return ""
	}
	return fmt.Sprintf("%s=%s&", param, value)
}

// Extracts a value from a query string (uses defValue if not found)
func qsGet(qs url.Values, key string, defValue string) string {
	if len(qs[key]) == 0 {
		return defValue
	}

	value := strings.TrimSpace(qs[key][0])
	if value == "" {
		return defValue
	}
	return value
}

// Extracts an integer value from a query string (uses defValue if not found)
func qsGetInt(qs url.Values, key string, defValue int) int {
	if len(qs[key]) == 0 {
		return defValue
	}

	i, err := strconv.Atoi(qs[key][0])
	if err != nil {
		return defValue
	}
	return i
}

// Encodes a single value as a query string parameter.
func encode(param, value string) string {
	return qsAdd(param, url.QueryEscape(value))
}

func encodeInt(param string, value int) string {
	return qsAdd(param, fmt.Sprintf("%d", value))
}

func encodeDefault(param, value, defaultValue string) string {
	if value == "" {
		return encode(param, defaultValue)
	}
	return encode(param, value)
}

// Encodes an array of values as a query string parameter (the values
// in the return value are separated by commas).
func encodeMany(param string, values []string) string {
	if len(values) == 0 {
		return ""
	}
	valuesEnc := []string{}
	for _, value := range values {
		valuesEnc = append(valuesEnc, url.QueryEscape(value))
	}
	return qsAdd(param, strings.Join(valuesEnc, ","))
}
