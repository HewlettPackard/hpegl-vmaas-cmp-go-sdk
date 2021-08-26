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
	models "github.com/HewlettPackard/vmaas-cmp-go-sdk/pkg/models"
)

type VirtualImagesAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
VirtualImageAPIService
Get All Virtual images
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param name/phrase optional
@return models.VirtualImages
*/
func (a *VirtualImagesAPIService) GetAllVirtualImages(ctx context.Context,
	param map[string]string) (models.VirtualImages, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFileName     string
		localVarFileBytes    []byte
		virtualImageResponse models.VirtualImages
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.VirtualImagePath)

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
		return virtualImageResponse, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return virtualImageResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return virtualImageResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return virtualImageResponse, err
	}

	if err := json.Unmarshal(localVarBody, &virtualImageResponse); err != nil {
		return virtualImageResponse, err
	}

	return virtualImageResponse, nil
}
