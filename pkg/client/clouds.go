// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	consts "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/common"
	"github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
)

// Linger please
var (
	_ context.Context
)

type CloudsApiService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
CloudsApiService
Get a Specific Cloud
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param cloudId The cloud ID

*/
func (a *CloudsApiService) GetASpecificCloud(ctx context.Context, cloudId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/zones/{cloud_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"cloud_id"+"}", fmt.Sprintf("%v", cloudId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if cloudId < 1 {
		return nil, reportError("cloudId must be greater than 1")
	}

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

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
CloudsApiService
Get a Specific Cloud Data Store
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param cloudId The cloud ID
 * @param datastoreId The cloud datastore ID

*/
func (a *CloudsApiService) GetASpecificCloudDataStore(ctx context.Context, cloudId int, datastoreId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/zones/{cloud_id}/data-stores/{datastore_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"cloud_id"+"}", fmt.Sprintf("%v", cloudId))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"datastore_id"+"}", fmt.Sprintf("%v", datastoreId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if cloudId < 1 {
		return nil, reportError("cloudId must be greater than 1")
	}
	if datastoreId < 1 {
		return nil, reportError("datastoreId must be greater than 1")
	}

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

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
CloudsApiService
Get a Specific Cloud Resource Pool
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
	Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param cloudId The cloud ID
 * @param funId The Cloud Resourcepool ID

*/
func (a *CloudsApiService) GetASpecificCloudResourcePool(ctx context.Context, cloudId int, resourcepoolId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/zones/{cloud_id}/resource-pools/{resourcepool_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"cloud_id"+"}", fmt.Sprintf("%v", cloudId))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"resourcepool_id"+"}", fmt.Sprintf("%v", resourcepoolId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if cloudId < 1 {
		return nil, reportError("cloudId must be greater than 1")
	}
	if resourcepoolId < 1 {
		return nil, reportError("resourcepoolId must be greater than 1")
	}

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

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	return localVarHttpResponse, nil
}

/*
CloudsApiService
Get All Cloud Data Stores
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param cloudId The cloud ID

*/
func (a *CloudsApiService) GetAllCloudDataStores(ctx context.Context, cloudId int, queryParams map[string]string) (models.DataStoresResp, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/data-stores", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.ZonePath, cloudId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getUrlValues(queryParams)
	localVarFormParams := url.Values{}
	var datastoresResp models.DataStoresResp
	if cloudId < 1 {
		return datastoresResp, reportError("cloudId must be greater than 1")
	}

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

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return datastoresResp, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return datastoresResp, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return datastoresResp, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return datastoresResp, err
	}

	if err = json.Unmarshal(localVarBody, &datastoresResp); err != nil {
		return datastoresResp, err
	}

	return datastoresResp, nil
}

/*
CloudsApiService
Get All Cloud Resource Pools
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param cloudId The cloud ID

*/
func (a *CloudsApiService) GetAllCloudResourcePools(ctx context.Context, cloudId int, queryParams map[string]string) (models.ResourcePoolsResp, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/resource-pools", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.ZonePath, cloudId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getUrlValues(queryParams)
	localVarFormParams := url.Values{}
	var resourcePoolsResp models.ResourcePoolsResp
	if cloudId < 1 {
		return resourcePoolsResp, reportError("cloudId must be greater than 1")
	}

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

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return resourcePoolsResp, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return resourcePoolsResp, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return resourcePoolsResp, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return resourcePoolsResp, err
	}

	if err = json.Unmarshal(localVarBody, &resourcePoolsResp); err != nil {
		return resourcePoolsResp, err
	}

	return resourcePoolsResp, nil
}

/*
CloudsApiService
Get All Clouds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceId

*/
func (a *CloudsApiService) GetAllClouds(ctx context.Context, queryParams map[string]string) (models.CloudsResp, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.ZonePath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getUrlValues(queryParams)
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

	var cloudsResp models.CloudsResp
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return cloudsResp, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return cloudsResp, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return cloudsResp, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return cloudsResp, err
	}

	if err = json.Unmarshal(localVarBody, &cloudsResp); err != nil {
		return cloudsResp, err
	}

	return cloudsResp, nil
}

func (a *CloudsApiService) GetAllFolders(ctx context.Context, cloudId int, queryParams map[string]string) (models.GetFolders, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.ZonePath, cloudId, consts.FolderPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getUrlValues(queryParams)
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

	var folderResp models.GetFolders
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return folderResp, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return folderResp, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return folderResp, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return folderResp, err
	}

	if err = json.Unmarshal(localVarBody, &folderResp); err != nil {
		return folderResp, err
	}

	return folderResp, nil
}

func (a *CloudsApiService) GetAllCloudNetworks(ctx context.Context, cloudId, provisionTypeId int) (models.GetAllCloudNetworks, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.OptionsPath, consts.ZoneNetworkOptionsPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getUrlValues(map[string]string{
		"zoneId":          strconv.Itoa(cloudId),
		"provisionTypeId": strconv.Itoa(provisionTypeId),
	})
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

	var networkResp models.GetAllCloudNetworks
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams,
		localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return networkResp, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return networkResp, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return networkResp, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return networkResp, err
	}

	if err = json.Unmarshal(localVarBody, &networkResp); err != nil {
		return networkResp, err
	}

	return networkResp, nil
}
