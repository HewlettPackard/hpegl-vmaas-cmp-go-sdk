// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
)

// Linger please
var (
	_ context.Context
)

type RolesApiService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
RolesApiService
Create a Role
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param body
 * @param serviceInstanceId
@return models.GetRoleResponse
*/
func (a *RolesApiService) CreateARole(ctx context.Context, body models.CreateRoleBody) (models.GetRoleResponse,
	*http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.GetRoleResponse
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/roles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = &body

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RolesApiService
Customizing Blueprint Access
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param body
 * @param serviceInstanceId
 * @param roleId
@return models.CommonRoleUpdateResp
*/
func (a *RolesApiService) CustomizingBlueprintAccess(ctx context.Context,
	body models.UpdateBlueprintAccessBody, roleId int) (models.CommonRoleUpdateResp, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.CommonRoleUpdateResp
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/roles/{role_id}/update-blueprint"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", fmt.Sprintf("%v", roleId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = &body

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RolesApiService
Customizing Cloud Access
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param body
 * @param serviceInstanceId
 * @param roleId
@return models.CommonRoleUpdateResp
*/
func (a *RolesApiService) CustomizingCloudAccess(ctx context.Context, body models.UpdateCloudAccessBody,
	roleId int) (models.CommonRoleUpdateResp, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.CommonRoleUpdateResp
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/roles/{role_id}/update-cloud"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", fmt.Sprintf("%v", roleId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = &body

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RolesApiService
Customizing Group Access
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param body
 * @param serviceInstanceId
 * @param roleId
@return models.CommonRoleUpdateResp
*/
func (a *RolesApiService) CustomizingGroupAccess(ctx context.Context,
	body models.UpdateGroupAccessBody, roleId int) (models.CommonRoleUpdateResp, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.CommonRoleUpdateResp
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/roles/{role_id}/update-group"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", fmt.Sprintf("%v", roleId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = &body

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RolesApiService
Customizing Instance Type Access
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param body
 * @param serviceInstanceId
 * @param roleId
@return models.CommonRoleUpdateResp
*/
func (a *RolesApiService) CustomizingInstanceTypeAccess(ctx context.Context,
	body models.UpdateInstancetypeAccessBody, roleId int) (models.CommonRoleUpdateResp, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.CommonRoleUpdateResp
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/roles/{role_id}/update-instance-type"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", fmt.Sprintf("%v", roleId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = &body

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RolesApiService
Delete a Role
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param roleId

*/
func (a *RolesApiService) DeleteARole(ctx context.Context, roleId int) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/roles/{role_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", fmt.Sprintf("%v", roleId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
RolesApiService
Get a Specific Role
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param roleId
@return models.GetRoleResponse
*/
func (a *RolesApiService) GetASpecificRole(ctx context.Context, roleId int) (models.GetRoleResponse,
	*http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.GetRoleResponse
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/roles/{role_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", fmt.Sprintf("%v", roleId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, ParseError(localVarHttpResponse)
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RolesApiService
Get All Roles
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceId

*/
func (a *RolesApiService) GetAllRoles(ctx context.Context) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/roles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
RolesApiService
Updating Basic Role Settings
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param body
 * @param serviceInstanceId
 * @param roleId
@return models.GetRoleResponse
*/
func (a *RolesApiService) UpdatingBasicRoleSettings(ctx context.Context, body models.UpdateRoleBody,
	roleId int) (models.GetRoleResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.GetRoleResponse
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/roles/{role_id}"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", fmt.Sprintf("%v", roleId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = &body

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, ParseError(localVarHttpResponse)
	}
	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RolesApiService
Along with feature access global access for Group, Cloud, Instance Types and Blue Prints can be provided
using this API.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param body
 * @param serviceInstanceId
 * @param roleId
@return models.CommonRoleUpdateResp
*/
func (a *RolesApiService) UpdatingRolePermissions(ctx context.Context, body models.UpdateRoleFeaturePermission,
	roleId int) (models.CommonRoleUpdateResp, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.CommonRoleUpdateResp
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/roles/{role_id}/update-permission"
	localVarPath = strings.ReplaceAll(localVarPath, "{"+"role_id"+"}", fmt.Sprintf("%v", roleId))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = &body

	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode >= 300 {
		return localVarReturnValue, localVarHttpResponse, ParseError(localVarHttpResponse)
	}
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
