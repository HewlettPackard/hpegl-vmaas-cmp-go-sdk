// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	consts "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/common"
	models "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
)

var _ context.Context

type EnvironmentApiService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (e *EnvironmentApiService) GetAllEnvironment(ctx context.Context,
	param map[string]string) (models.GetAllEnvironment, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		environmentResponse models.GetAllEnvironment
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", e.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.EnvironmentPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getUrlValues(param)
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := e.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return environmentResponse, err
	}

	localVarHttpResponse, err := e.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return environmentResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return environmentResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return environmentResponse, err
	}

	if err := json.Unmarshal(localVarBody, &environmentResponse); err != nil {
		return environmentResponse, err
	}

	return environmentResponse, nil
}
