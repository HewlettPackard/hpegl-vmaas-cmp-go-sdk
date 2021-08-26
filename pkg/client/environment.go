// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	models "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type EnvironmentAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (e *EnvironmentAPIService) GetAllEnvironment(ctx context.Context,
	param map[string]string) (models.GetAllEnvironment, error) {
	var (
		localVarHTTPMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		environmentResponse models.GetAllEnvironment
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", e.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.EnvironmentPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getURLValues(param)
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := e.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return environmentResponse, err
	}

	localVarHTTPResponse, err := e.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return environmentResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return environmentResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return environmentResponse, err
	}

	if err := json.Unmarshal(localVarBody, &environmentResponse); err != nil {
		return environmentResponse, err
	}

	return environmentResponse, nil
}
