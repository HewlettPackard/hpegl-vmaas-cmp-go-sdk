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
	"github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
)

type PlansAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
PlansAPIService
Get All Service Plans
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID

*/
func (a *PlansAPIService) GetAllServicePlans(ctx context.Context,
	param map[string]string) (models.ServicePlans, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFileName     string
		localVarFileBytes    []byte
		servicePlansResponse models.ServicePlans
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.ServicePlansPath)

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
		return servicePlansResponse, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return servicePlansResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return servicePlansResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return servicePlansResponse, err
	}

	if err := json.Unmarshal(localVarBody, &servicePlansResponse); err != nil {
		return servicePlansResponse, err
	}

	return servicePlansResponse, nil
}
