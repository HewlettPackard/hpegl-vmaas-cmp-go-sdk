// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type GroupsAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
GroupsAPIService
Creates a group
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param optional nil or *GroupsAPICreateGroupOpts - Optional Parameters:
     * @param "Body" (optional.Interface of CreateGroupBody) -

*/

type GroupsAPICreateGroupOpts struct {
	Body optional.Interface
}

func (a *GroupsAPIService) CreateGroup(ctx context.Context,
	localVarOptionals *GroupsAPICreateGroupOpts) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	// Client.greenlake.hpe.com/api/vmaas/
	localVarPath := fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.GroupsPath)

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
GroupsAPIService
If a group has zones or servers still tied to it, a delete action will fail
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param groupID The Group ID

*/
func (a *GroupsAPIService) DeleteAGroup(ctx context.Context, groupID int) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.GroupsPath, groupID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if groupID < 1 {
		return nil, reportError("groupID must be greater than 1")
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
GroupsAPIService
Get a Specific Group
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param groupID The Group ID

*/
func (a *GroupsAPIService) GetASpecificGroup(ctx context.Context, groupID int) (models.GroupResp, error) {
	specificGrpResp := models.GroupResp{}

	specificGrpRespAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.GroupsPath, groupID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &specificGrpResp)
		},
		validations: []validationFunc{
			func() error {
				if groupID < 1 {
					return fmt.Errorf("%s", "group id should be greater than or equal to 1")
				}

				return nil
			},
		},
	}
	err := specificGrpRespAPI.do(ctx, nil, nil)

	return specificGrpResp, err
}

/*
GroupsAPIService
This endpoint retrieves all groups and a list of zones associated with the group by id.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID

*/
func (a *GroupsAPIService) GetAllGroups(ctx context.Context,
	queryParams map[string]string) (models.Groups, error) {
	allGrpResp := models.Groups{}

	allGrpRespAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.GroupsPath),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &allGrpResp)
		},
	}
	err := allGrpRespAPI.do(ctx, nil, queryParams)

	return allGrpResp, err
}

/*
GroupsAPIService
This will update the zones(clouds) that are assigned to the group. Any zones that are not passed in
the zones parameter will be removed from the group.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param groupID The Group ID
 * @param optional nil or *GroupsAPIGroupZoneUpdateOpts - Optional Parameters:
     * @param "Body" (optional.Interface of UpdateGroupZone) -

*/

type GroupsAPIGroupZoneUpdateOpts struct {
	Body optional.Interface
}

func (a *GroupsAPIService) GroupZoneUpdate(ctx context.Context, groupID int,
	localVarOptionals *GroupsAPIGroupZoneUpdateOpts) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.GroupsPath, groupID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if groupID < 1 {
		return nil, reportError("groupID must be greater than 1")
	}

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
GroupsAPIService
Updating a group name
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param groupID The Group ID
 * @param optional nil or *GroupsAPIUpdatingAGroupNameOpts - Optional Parameters:
     * @param "Body" (optional.Interface of UpdateGroupName) -

*/

type GroupsAPIUpdatingAGroupNameOpts struct {
	Body optional.Interface
}

func (a *GroupsAPIService) UpdatingAGroupName(ctx context.Context, groupID int,
	localVarOptionals *GroupsAPIUpdatingAGroupNameOpts) (*http.Response, error) {
	var (
		localVarHTTPMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
		consts.GroupsPath, groupID)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if groupID < 1 {
		return nil, reportError("groupID must be greater than 1")
	}

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
