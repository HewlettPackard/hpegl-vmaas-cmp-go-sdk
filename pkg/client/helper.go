//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func getURLValues(query map[string]string) url.Values {
	m := make(map[string][]string)
	for k, v := range query {
		m[k] = []string{v}
	}

	return m
}

type CustomError struct {
	Errors             string                 `json:"error,omitempty"`
	Body               map[string]interface{} `json:"body,omitempty"`
	StatusCode         int                    `json:"statuscode,omitempty"`
	RecommendedActions []string               `json:"recommendedActions,omitempty"`
}

func (c CustomError) Error() string {
	jsonObj, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(jsonObj)
}

func ParseError(resp *http.Response) error {
	customErr := CustomError{
		StatusCode: resp.StatusCode,
	}
	err := json.NewDecoder(resp.Body).Decode(&customErr.Body)
	if err != nil {
		customErr.Errors = err.Error()
	} else if len(customErr.Body) == 0 {
		customErr.Errors = "No additional information is available"
	}

	if rAction, ok := customErr.Body["recommendedActions"]; ok {
		delete(customErr.Body, "recommendedActions")
		rActions := rAction.([]interface{})
		customErr.RecommendedActions = make([]string, 0, len(rActions))
		for _, a := range rActions {
			customErr.RecommendedActions = append(customErr.RecommendedActions, a.(string))
		}
	}

	return customErr
}
