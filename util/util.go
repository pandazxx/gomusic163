package util

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func NewHTTPRequest(
	method string,
	urlString string,
	headers map[string]string,
	queries map[string]string) (req *http.Request, err error) {
	var queriesString string
	if queries != nil {
		q := url.Values{}
		for key := range queries {
			fmt.Printf("Adding query: %s, %s\n", key, queries[key])
			q.Add(key, queries[key])
		}
		queriesString = q.Encode()
	}

	switch {
	case strings.ToLower(method) == "get":
		formalUrlString := urlString
		if queriesString != "" {
			formalUrlString += "?" + queriesString
		}
		req, err = http.NewRequest("GET", formalUrlString, nil)
		if err != nil {
			return nil, err
		}
	case strings.ToLower(method) == "post":
		req, err = http.NewRequest("POST", urlString, strings.NewReader(queriesString))
		if err != nil {
			return nil, err
		}
	}
	if req != nil {
		if headers != nil {
			for key := range headers {
				req.Header.Add(key, headers[key])
			}
		}
		return req, nil
	}
	return nil, errors.New("Not supported method type")
}
