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

type ProvisioningAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (a *ProvisioningAPIService) GetAllProvisioningTypes(ctx context.Context,
	param map[string]string) (models.GetAllProvisioningTypes, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		provisions         models.GetAllProvisioningTypes
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.ProvisionTypesPath)

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

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return provisions, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return provisions, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return provisions, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return provisions, err
	}

	if err := json.Unmarshal(localVarBody, &provisions); err != nil {
		return provisions, err
	}

	return provisions, nil
}
