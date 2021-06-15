//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func getUrlValues(query map[string]string) url.Values {
	m := make(map[string][]string)
	for k, v := range query {
		m[k] = []string{v}
	}
	return m
}

type customError struct {
	error      string
	body       map[string]interface{}
	statusCode int
}

func (c customError) Error() string {
	if c.error != "" {
		return fmt.Sprintf("status code: %d. Error: %s", c.statusCode, c.error)
	}
	details := "{"
	for k, v := range c.body {
		details = fmt.Sprintf("%s %s: %v,", details, k, v)
	}
	details = details[:len(details)-1] + "}"
	return fmt.Sprintf("status code: %d. Details: %v", c.statusCode, details)
}

func ParseError(resp *http.Response) error {
	customErr := customError{
		statusCode: resp.StatusCode,
	}
	err := json.NewDecoder(resp.Body).Decode(&customErr.body)
	if err != nil {
		customErr.error = err.Error()
	} else if len(customErr.body) == 0 {
		customErr.error = "No additional information is available"
	}
	return customErr
}
