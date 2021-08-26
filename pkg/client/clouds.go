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

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type CloudsAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
CloudsAPIService
Get a Specific Cloud
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param cloudID The cloud ID

*/
func (a *CloudsAPIService) GetASpecificCloud(ctx context.Context, cloudID int) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/zones/{cloud_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"cloud_id"+"}", fmt.Sprintf("%v", cloudID))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if cloudID < 1 {
		return nil, reportError("cloudID must be greater than 1")
	}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarHTTPResponse, ParseError(localVarHTTPResponse)
	}

	return localVarHTTPResponse, nil
}

/*
CloudsAPIService
Get a Specific Cloud Data Store
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param cloudID The cloud ID
 * @param datastoreID The cloud datastore ID

*/
func (a *CloudsAPIService) GetASpecificCloudDataStore(ctx context.Context, cloudID int,
	datastoreID int) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/zones/{cloud_id}/data-stores/{datastore_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"cloud_id"+"}", fmt.Sprintf("%v", cloudID))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"datastore_id"+"}", fmt.Sprintf("%v", datastoreID))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if cloudID < 1 {
		return nil, reportError("cloudID must be greater than 1")
	}
	if datastoreID < 1 {
		return nil, reportError("datastoreID must be greater than 1")
	}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarHTTPResponse, ParseError(localVarHTTPResponse)
	}

	return localVarHTTPResponse, nil
}

/*
CloudsAPIService
Get a Specific Cloud Resource Pool
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param cloudID The cloud ID
 * @param funID The Cloud Resourcepool ID

*/
func (a *CloudsAPIService) GetASpecificCloudResourcePool(ctx context.Context, cloudID int,
	resourcepoolID int) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/zones/{cloud_id}/resource-pools/{resourcepool_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"cloud_id"+"}", fmt.Sprintf("%v", cloudID))
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"resourcepool_id"+"}", fmt.Sprintf("%v", resourcepoolID))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if cloudID < 1 {
		return nil, reportError("cloudID must be greater than 1")
	}
	if resourcepoolID < 1 {
		return nil, reportError("resourcepoolID must be greater than 1")
	}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarHTTPResponse, ParseError(localVarHTTPResponse)
	}

	return localVarHTTPResponse, nil
}

/*
CloudsAPIService
Get All Cloud Data Stores
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param cloudID The cloud ID

*/
func (a *CloudsAPIService) GetAllCloudDataStores(ctx context.Context, cloudID int,
	queryParams map[string]string) (models.DataStoresResp, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/data-stores", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.ZonePath, cloudID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getURLValues(queryParams)
	localVarFormParams := url.Values{}
	var datastoresResp models.DataStoresResp
	if cloudID < 1 {
		return datastoresResp, reportError("cloudID must be greater than 1")
	}

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
		return datastoresResp, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return datastoresResp, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return datastoresResp, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return datastoresResp, err
	}

	if err = json.Unmarshal(localVarBody, &datastoresResp); err != nil {
		return datastoresResp, err
	}

	return datastoresResp, nil
}

/*
CloudsAPIService
Get All Cloud Resource Pools
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param cloudID The cloud ID

*/
func (a *CloudsAPIService) GetAllCloudResourcePools(ctx context.Context, cloudID int,
	queryParams map[string]string) (models.ResourcePoolsResp, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/resource-pools", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.ZonePath, cloudID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getURLValues(queryParams)
	localVarFormParams := url.Values{}
	var resourcePoolsResp models.ResourcePoolsResp
	if cloudID < 1 {
		return resourcePoolsResp, reportError("cloudID must be greater than 1")
	}

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
		return resourcePoolsResp, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return resourcePoolsResp, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return resourcePoolsResp, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return resourcePoolsResp, err
	}

	if err = json.Unmarshal(localVarBody, &resourcePoolsResp); err != nil {
		return resourcePoolsResp, err
	}

	return resourcePoolsResp, nil
}

/*
CloudsAPIService
Get All Clouds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID

*/
func (a *CloudsAPIService) GetAllClouds(ctx context.Context,
	queryParams map[string]string) (models.CloudsResp, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.ZonePath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getURLValues(queryParams)
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

	var cloudsResp models.CloudsResp
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return cloudsResp, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return cloudsResp, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return cloudsResp, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return cloudsResp, err
	}

	if err = json.Unmarshal(localVarBody, &cloudsResp); err != nil {
		return cloudsResp, err
	}

	return cloudsResp, nil
}

func (a *CloudsAPIService) GetAllCloudNetworks(ctx context.Context, cloudID,
	provisionTypeID int) (models.GetAllCloudNetworks, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.OptionsPath, consts.ZoneNetworkOptionsPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := getURLValues(map[string]string{
		"zoneId":          strconv.Itoa(cloudID),
		"provisionTypeId": strconv.Itoa(provisionTypeID),
	})
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

	var networkResp models.GetAllCloudNetworks
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return networkResp, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return networkResp, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return networkResp, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return networkResp, err
	}

	if err = json.Unmarshal(localVarBody, &networkResp); err != nil {
		return networkResp, err
	}

	return networkResp, nil
}

func (a *CloudsAPIService) GetAllCloudFolders(
	ctx context.Context,
	cloudID int,
	queryParams map[string]string,
) (models.GetAllCloudFolders, error) {
	folders := models.GetAllCloudFolders{}

	folderAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%d/%s",
			a.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.ZonePath, cloudID, consts.FolderPath),
		client: a.Client,
		// jsonParser also stores folder response, since folders is not a local variable
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &folders)
		},
		// validate cloud id > 1
		validations: []validationFunc{
			func() error {
				if cloudID < 1 {
					return fmt.Errorf("%s", "cloud id should be greater than or equal to 1")
				}

				return nil
			},
		},
	}

	err := folderAPI.do(ctx, nil, queryParams)

	return folders, err
}

func (a *CloudsAPIService) GetSpecificCloudFolder(
	ctx context.Context,
	cloudID int,
	folderID int,
) (models.GetSpecificCloudFolder, error) {
	folder := models.GetSpecificCloudFolder{}

	folderAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%d/%s/%d",
			a.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.ZonePath, cloudID, consts.FolderPath, folderID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &folder)
		},
		// validate cloud id > 1
		validations: []validationFunc{
			func() error {
				if cloudID < 1 {
					return fmt.Errorf("%s", "cloud id should be greater than or equal to 1")
				}

				return nil
			},
		},
	}

	err := folderAPI.do(ctx, nil, nil)

	return folder, err
}
