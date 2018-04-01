package solr

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Adds a parameter and its value to a query string
func QsAddRaw(param, value string) string {
	if value == "" {
		return ""
	}
	return fmt.Sprintf("%s=%s&", param, value)
}

// Extracts a value from a query string (uses defValue if not found)
func QsGet(qs url.Values, key string, defValue string) string {
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
func QsGetInt(qs url.Values, key string, defValue int) int {
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
func QsAdd(param, value string) string {
	return QsAddRaw(param, url.QueryEscape(value))
}

func QsAddInt(param string, value int) string {
	return QsAddRaw(param, fmt.Sprintf("%d", value))
}

func QsAddDefault(param, value, defaultValue string) string {
	if value == "" {
		return QsAdd(param, defaultValue)
	}
	return QsAdd(param, value)
}

// Encodes an array of values as a query string parameter (the values
// in the return value are separated by commas).
func QsAddMany(param string, values []string) string {
	if len(values) == 0 {
		return ""
	}
	encodedValues := []string{}
	for _, value := range values {
		encodedValues = append(encodedValues, url.QueryEscape(value))
	}
	return QsAddRaw(param, strings.Join(encodedValues, ","))
}
