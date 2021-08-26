// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	consts "github.com/HewlettPackard/vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/vmaas-cmp-go-sdk/pkg/models"
)

type LibraryAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
LibrariesAPIService
Get All layouts
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID

*/
func (a *LibraryAPIService) GetAllLayouts(ctx context.Context,
	param map[string]string) (models.LayoutsResp, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.LibraryLayoutPath)

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
		return models.LayoutsResp{}, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return models.LayoutsResp{}, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return models.LayoutsResp{}, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return models.LayoutsResp{}, err
	}

	var layoutsResp models.LayoutsResp
	if err = json.Unmarshal(localVarBody, &layoutsResp); err != nil {
		return models.LayoutsResp{}, err
	}

	return layoutsResp, nil
}

func (a *LibraryAPIService) GetAllInstanceTypes(ctx context.Context,
	param map[string]string) (models.InstanceTypesResp, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.LibraryInstanceTypesPath)

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
		return models.InstanceTypesResp{}, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return models.InstanceTypesResp{}, err
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		return models.InstanceTypesResp{}, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return models.InstanceTypesResp{}, err
	}
	var instanceTypesResp models.InstanceTypesResp
	if err = json.Unmarshal(localVarBody, &instanceTypesResp); err != nil {
		return models.InstanceTypesResp{}, err
	}

	return instanceTypesResp, nil
}
