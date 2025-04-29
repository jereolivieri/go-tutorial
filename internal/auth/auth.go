package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key
// from the headers of a HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	// Example:
	// Authorization: ApiKey {insert apikey here}
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return vals[1], nil
}
