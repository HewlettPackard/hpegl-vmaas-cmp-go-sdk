// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"io/ioutil"
	"net/url"
	"strings"

	consts "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/common"
)

// all the required validation should be provided as validationFunc
type validationFunc func() error

// parse response as proper model
type jsonPareserFunc func(body []byte) error

type api struct {
	method      string
	path        string
	client      APIClientHandler
	jsonParser  jsonPareserFunc
	validations []validationFunc
}

// do will call the API provided. this function will not return any response, but
// response should be catched from jsonParser function itself
func (a *api) do(ctx context.Context, request interface{}, queryParams map[string]string) error {
	var (
		localVarHTTPMethod = strings.ToUpper(a.method)
		localVarFileName   string
		localVarFileBytes  []byte
	)
	if a.path == "" || a.method == "" || a.client == nil || a.jsonParser == nil {
		panic("api not properly configured")
	}

	for _, validations := range a.validations {
		err := validations()
		if err != nil {
			return err
		}
	}
	// create path and map variables
	localVarPath := a.path

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getURLValues(queryParams)
	localVarFormParams := url.Values{}

	// set Content-Type header
	localVarHeaderParams["Content-Type"] = consts.ContentType
	// set Accept header
	localVarHeaderParams["Accept"] = consts.ContentType

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, request,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return err
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		return ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return err
	}

	return a.jsonParser(localVarBody)
}
