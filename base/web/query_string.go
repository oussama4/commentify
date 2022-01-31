package web

import (
	"net/url"
	"strconv"
)

// ReadInt reads a string value from a query string and converts it to int
func ReadInt(qs url.Values, key string, defaultValue int) (int, error) {
	v := qs.Get(key)

	if v == "" {
		return defaultValue, nil
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue, err
	}

	return i, nil
}

// ReadString reads value from a query string
func ReadString(qs url.Values, key string, defaultValue string) string {
	v := qs.Get(key)

	if v == "" {
		return defaultValue
	}

	return v
}
