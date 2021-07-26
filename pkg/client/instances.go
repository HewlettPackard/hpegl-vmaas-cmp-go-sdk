// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"

	consts "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/common"
	models "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
)

type InstancesAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
InstancesAPIService
Clone an instance and all VM within that instance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID
 * @param optional nil or *InstancesAPICloneAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of CloneInstanceBody) -

*/

type InstancesAPICloneAnInstanceOpts struct {
	Body optional.Interface
}

func (a *InstancesAPIService) CloneAnInstance(ctx context.Context, instanceID int,
	localVarOptionals *models.CreateInstanceBody) (models.SuccessOrErrorMessage, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	var cloneResp models.SuccessOrErrorMessage
	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/clone", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return cloneResp, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return cloneResp, err
	}
	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return cloneResp, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return cloneResp, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return cloneResp, err
	}
	if err := json.Unmarshal(localVarBody, &cloneResp); err != nil {
		return cloneResp, err
	}

	return cloneResp, nil
}

/*
InstancesAPIService
Creates an image template from an existing instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID
 * @param optional nil or *InstancesAPICloneToImageOpts - Optional Parameters:
     * @param "Body" (optional.Interface of CloneToImage) -

*/

type InstancesAPICloneToImageOpts struct {
	Body optional.Interface
}

func (a *InstancesAPIService) CloneToImage(ctx context.Context, instanceID int,
	localVarOptionals *InstancesAPICloneToImageOpts) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/clone-image", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
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
InstancesAPIService
Create an Instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param optional nil or *InstancesAPICreateAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of CreateInstanceBody) -
@return models.GetInstanceResponse
*/

func (a *InstancesAPIService) CreateAnInstance(ctx context.Context,
	localVarOptionals *models.CreateInstanceBody) (models.GetInstanceResponse, error) {
	var (
		localVarHTTPMethod     = strings.ToUpper("Post")
		localVarPostBody       interface{}
		localVarFileName       string
		localVarFileBytes      []byte
		createInstanceResponse models.GetInstanceResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return createInstanceResponse, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return createInstanceResponse, err
	}
	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return createInstanceResponse, err
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		return createInstanceResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return createInstanceResponse, err
	}

	if err = json.Unmarshal(localVarBody, &createInstanceResponse); err != nil {
		return createInstanceResponse, err
	}

	return createInstanceResponse, nil
}

/*
InstancesAPIService
Will delete an instance and all associated monitors and backups.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID
 * @param optional nil or *InstancesAPIDeleteAnIstanceOpts - Optional Parameters:
     * @param "Force" (optional.String) -

*/

type InstancesAPIDeleteAnIstanceOpts struct {
	Force optional.String
}

func (a *InstancesAPIService) DeleteAnInstance(ctx context.Context,
	instanceID int) (models.SuccessOrErrorMessage, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
		return models.SuccessOrErrorMessage{}, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return models.SuccessOrErrorMessage{}, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return models.SuccessOrErrorMessage{}, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return models.SuccessOrErrorMessage{}, err
	}

	// fmt.Println(string(localVarBody))
	var instancesResponse models.SuccessOrErrorMessage
	if err = json.Unmarshal(localVarBody, &instancesResponse); err != nil {
		return models.SuccessOrErrorMessage{}, err
	}

	return instancesResponse, nil
}

/*
InstancesAPIService
Get a Specific Instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID
@return models.GetInstanceResponse
*/
func (a *InstancesAPIService) GetASpecificInstance(ctx context.Context,
	instanceID int) (models.GetInstanceResponse, error) {
	var (
		localVarHTTPMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		getInstanceResponse models.GetInstanceResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
		return getInstanceResponse, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return getInstanceResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return getInstanceResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return getInstanceResponse, err
	}

	if err = json.Unmarshal(localVarBody, &getInstanceResponse); err != nil {
		return getInstanceResponse, err
	}

	return getInstanceResponse, nil
}

/*
InstancesAPIService
Fetch the list of available instance types
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID

*/
func (a *InstancesAPIService) GetAllInstanceTypesForProvisioning(ctx context.Context) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstanceTypesPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
InstancesAPIService
Get All Instances
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param optional nil or *InstancesAPIGetAllInstancesOpts - Optional Parameters:
     * @param "Max" (optional.Int32) -  Max number of results to return
     * @param "Name" (optional.String) -  Filter by name
     * @param "Phrase" (optional.String) -  Filter by wildcard search of name and description
     * @param "InstanceType" (optional.Int32) -  Filter by instance type code
     * @param "CreatedBy" (optional.String) -  Filter by Created By (User) ID. Accepts multiple values
     * @param "Status" (optional.String) -  Filter by instance status
     * @param "ShowDeleted" (optional.Bool) -  If true, includes instances in pending removal status.
     * @param "Deleted" (optional.Bool) -  If true, only instances in pending removal status are returned.
     * @param "Labels" (optional.String) -  Filter by labels
     * @param "Tags" (optional.String) -  Filter by tags

*/

type InstancesAPIGetAllInstancesOpts struct {
	Max          optional.Int32
	Name         optional.String
	Phrase       optional.String
	InstanceType optional.Int32
	CreatedBy    optional.String
	Status       optional.String
	ShowDeleted  optional.Bool
	Deleted      optional.Bool
	Labels       optional.String
	Tags         optional.String
}

func (a *InstancesAPIService) GetAllInstances(ctx context.Context,
	queryParams map[string]string) (models.Instances, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		instancesResponse  models.Instances
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath)

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

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return instancesResponse, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return instancesResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return instancesResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return instancesResponse, err
	}

	if err = json.Unmarshal(localVarBody, &instancesResponse); err != nil {
		return instancesResponse, err
	}

	return instancesResponse, nil
}

/*
InstancesAPIService
This endpoint retrieves all the Service Plans available for the specified cloud and instance layout.
The response includes details about the plans and their configuration options. It may be used to get
the list of available plans when creating a new instance or resizing an existing instance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param zoneID Cloud ID
 * @param layoutID Instance Layout ID
 * @param optional nil or *InstancesAPIGetAvailableServicePlansForAnInstanceOpts - Optional Parameters:
     * @param "SiteID" (optional.Int32) -  Group ID
@return models.GetServicePlanResponse
*/

type InstancesAPIGetAvailableServicePlansForAnInstanceOpts struct {
	SiteID optional.Int32
}

func (a *InstancesAPIService) GetAvailableServicePlansForAnInstance(ctx context.Context, zoneID int, layoutID int,
	localVarOptionals *InstancesAPIGetAvailableServicePlansForAnInstanceOpts) (models.GetServicePlanResponse,
	*http.Response, error) {
	var (
		localVarHTTPMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.GetServicePlanResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/service-plans", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("zoneId", parameterToString(zoneID, ""))
	localVarQueryParams.Add("layoutId", parameterToString(layoutID, ""))
	if localVarOptionals != nil && localVarOptionals.SiteID.IsSet() {
		localVarQueryParams.Add("siteId", parameterToString(localVarOptionals.SiteID.Value(), ""))
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
InstancesAPIService
List all environment variables associated with the instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID

*/
func (a *InstancesAPIService) GetEnvVariables(ctx context.Context, instanceID int) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/envs", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
InstancesAPIService
Retrieves the process history for a specific instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID

*/
func (a *InstancesAPIService) GetInstanceHistory(ctx context.Context, instanceID int) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/history", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
InstancesAPIService
Lists VMware Snapshot of the instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID

*/
func (a *InstancesAPIService) GetListOfSnapshotsForAnInstance(ctx context.Context,
	instanceID int) (models.ListSnapshotResponse, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		snapshotResponse   models.ListSnapshotResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/snapshots", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
		return snapshotResponse, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return snapshotResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return snapshotResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return snapshotResponse, err
	}
	if err = json.Unmarshal(localVarBody, &snapshotResponse); err != nil {
		return snapshotResponse, err
	}

	return snapshotResponse, nil
}

/*
InstancesAPIService
Fetch an instance type by ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceTypeID

*/
func (a *InstancesAPIService) GetSpecificInstanceTypeForProvisioning(ctx context.Context,
	instanceTypeID int) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstanceTypesPath, instanceTypeID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
InstancesAPIService
It is possible to import a snapshot of an instance. This creates a Virtual Image of the instance as
it currently exists.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID
 * @param optional nil or *InstancesAPIImportSnapshotOfAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ImportSnapshotBody) -

*/

func (a *InstancesAPIService) ImportSnapshotOfAnInstance(ctx context.Context, instanceID int,
	localVarOptionals *models.ImportSnapshotBody) (models.SuccessOrErrorMessage, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/import-snapshot", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return models.SuccessOrErrorMessage{}, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return models.SuccessOrErrorMessage{}, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return models.SuccessOrErrorMessage{}, err
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		return models.SuccessOrErrorMessage{}, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return models.SuccessOrErrorMessage{}, err
	}

	var instanceResponse models.SuccessOrErrorMessage
	if err = json.Unmarshal(localVarBody, &instanceResponse); err != nil {
		return models.SuccessOrErrorMessage{}, err
	}

	return instanceResponse, nil
}

/*
InstancesAPIService
&#x27;This will lock the instance. While locked, instances may not be removed.&#x27;
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID

*/
func (a *InstancesAPIService) LockAnInstance(ctx context.Context, instanceID int) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/lock", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
InstancesAPIService
It is possible to resize VMs within an instance by increasing their memory plan or storage limit.
This is done by assigning a new service plan to the VM.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID
 * @param optional nil or *InstancesAPIResizeAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ResizeInstanceBody) -
@return models.GetInstanceResponse
*/

func (a *InstancesAPIService) ResizeAnInstance(ctx context.Context, instanceID int,
	localVarOptionals *models.ResizeInstanceBody) (models.ResizeInstanceResponse, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		resizeResponse     models.ResizeInstanceResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/resize", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return resizeResponse, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return resizeResponse, err
	}
	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return resizeResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return resizeResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return resizeResponse, err
	}

	if err = json.Unmarshal(localVarBody, &resizeResponse); err != nil {
		return resizeResponse, err
	}

	return resizeResponse, nil
}

/*
InstancesAPIService
Restarts all VM running within an instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID

*/
func (a *InstancesAPIService) RestartAnInstance(ctx context.Context, instanceID int) (models.InstancePowerResponse, error) {
	var (
		localVarHTTPMethod    = strings.ToUpper("Put")
		localVarPostBody      interface{}
		localVarFileName      string
		localVarFileBytes     []byte
		instanceStateResponse models.InstancePowerResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/restart", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
		return instanceStateResponse, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return instanceStateResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return instanceStateResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return instanceStateResponse, err
	}

	if err = json.Unmarshal(localVarBody, &instanceStateResponse); err != nil {
		return instanceStateResponse, err
	}

	return instanceStateResponse, nil
}

/*
InstancesAPIService
Creates VMware Snapshot of the instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID
 * @param optional nil or *InstancesAPISnapshotAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of SnapshotBody) -

*/

func (a *InstancesAPIService) SnapshotAnInstance(ctx context.Context, instanceID int,
	localVarOptionals *models.SnapshotBody) (models.Instances, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFileName     string
		localVarFileBytes    []byte
		snapshotInstanceresp models.Instances
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/snapshot", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return snapshotInstanceresp, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return snapshotInstanceresp, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return snapshotInstanceresp, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return snapshotInstanceresp, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return snapshotInstanceresp, err
	}

	if err := json.Unmarshal(localVarBody, &snapshotInstanceresp); err != nil {
		return snapshotInstanceresp, err
	}

	return snapshotInstanceresp, nil
}

/*
InstancesAPIService
Starts all VM running within an instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID

*/
func (a *InstancesAPIService) StartAnInstance(ctx context.Context, instanceID int) (models.InstancePowerResponse, error) {
	var (
		localVarHTTPMethod    = strings.ToUpper("Put")
		localVarPostBody      interface{}
		localVarFileName      string
		localVarFileBytes     []byte
		instanceStateResponse models.InstancePowerResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/start", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
		return instanceStateResponse, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return instanceStateResponse, err
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		return instanceStateResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return instanceStateResponse, err
	}

	if err = json.Unmarshal(localVarBody, &instanceStateResponse); err != nil {
		return instanceStateResponse, err
	}

	return instanceStateResponse, nil
}

/*
InstancesAPIService
Stops all VM running within an instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID

*/
func (a *InstancesAPIService) StopAnInstance(ctx context.Context, instanceID int) (models.InstancePowerResponse, error) {
	var (
		localVarHTTPMethod    = strings.ToUpper("Put")
		localVarPostBody      interface{}
		localVarFileName      string
		localVarFileBytes     []byte
		instanceStateResponse models.InstancePowerResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/stop", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
		return instanceStateResponse, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return instanceStateResponse, err
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		return instanceStateResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return instanceStateResponse, err
	}

	if err = json.Unmarshal(localVarBody, &instanceStateResponse); err != nil {
		return instanceStateResponse, err
	}

	return instanceStateResponse, nil
}

/*
InstancesAPIService
Suspends all VM running within an instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID

*/
func (a *InstancesAPIService) SuspendAnInstance(ctx context.Context, instanceID int) (models.InstancePowerResponse, error) {
	var (
		localVarHTTPMethod    = strings.ToUpper("Put")
		localVarPostBody      interface{}
		localVarFileName      string
		localVarFileBytes     []byte
		instanceStateResponse models.InstancePowerResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/suspend", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
		return instanceStateResponse, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return instanceStateResponse, err
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		return instanceStateResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return instanceStateResponse, err
	}

	if err = json.Unmarshal(localVarBody, &instanceStateResponse); err != nil {
		return instanceStateResponse, err
	}

	return instanceStateResponse, nil
}

/*
InstancesAPIService
Undo the delete of an instance that is in pending removal state
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID
@return models.GetInstanceResponse
*/
func (a *InstancesAPIService) UndoDeleteOfAnInstance(ctx context.Context,
	instanceID int) (models.GetInstanceResponse,
	*http.Response, error) {
	var (
		localVarHTTPMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.GetInstanceResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/cancel-removal", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHTTPResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
InstancesAPIService
Unlocks the instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID

*/
func (a *InstancesAPIService) UnlockAnInstance(ctx context.Context, instanceID int) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d/unlock", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
InstancesAPIService
Updating an Instance
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param instanceID
 * @param optional nil or *InstancesAPIUpdatingAnInstanceOpts - Optional Parameters:
     * @param "Body" (optional.Interface of UpdateInstanceBody) -
@return models.GetInstanceResponse
*/

type InstancesAPIUpdatingAnInstanceOpts struct {
	Body optional.Interface
}

func (a *InstancesAPIService) UpdatingAnInstance(ctx context.Context, instanceID int,
	localVarOptionals *models.UpdateInstanceBody) (models.UpdateInstanceResponse, error) {
	var (
		localVarHTTPMethod     = strings.ToUpper("Put")
		localVarPostBody       interface{}
		localVarFileName       string
		localVarFileBytes      []byte
		updateInstanceResponse models.UpdateInstanceResponse
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.InstancesPath, instanceID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	if localVarOptionals != nil {
		var err error
		localVarPostBody, err = json.Marshal(localVarOptionals)
		if err != nil {
			return updateInstanceResponse, err
		}
	}

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return updateInstanceResponse, err
	}
	localVarHTTPResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return updateInstanceResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return updateInstanceResponse, ParseError(localVarHTTPResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return updateInstanceResponse, err
	}
	if err = json.Unmarshal(localVarBody, &updateInstanceResponse); err != nil {
		return updateInstanceResponse, err
	}

	return updateInstanceResponse, nil
}
