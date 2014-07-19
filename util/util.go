package util

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func NewHTTPRequest(method string, urlString string, headers map[string]string, queries map[string]string) (*http.Request, error) {
	switch {
	case strings.ToLower(method) == "get":
		var queriesString string
		if queries != nil {
			q := url.Values{}
			for key := range queries {
				fmt.Printf("Adding query: %s, %s", key, queries[key])
				q.Add(key, queries[key])
			}
			queriesString = q.Encode()
		}
		formalUrlString := urlString
		if queriesString != "" {
			formalUrlString += "?" + queriesString
		}
		req, err := http.NewRequest("GET", formalUrlString, nil)
		if err != nil {
			return nil, err
		}
		if headers != nil {
			for key := range headers {
				req.Header.Add(key, headers[key])
			}
		}
		return req, nil
	}
	return nil, errors.New("Not supported method type")
}
