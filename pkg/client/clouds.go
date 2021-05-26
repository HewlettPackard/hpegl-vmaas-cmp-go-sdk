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
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param cloudId The cloud ID

*/
func (a *CloudsApiService) GetASpecificCloud(ctx context.Context, serviceInstanceId string, cloudId int32) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/{service_instance_id}/zones/{cloud_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_instance_id"+"}", fmt.Sprintf("%v", serviceInstanceId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cloud_id"+"}", fmt.Sprintf("%v", cloudId), -1)

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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v models.ErrUnauthorized
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v models.ErrNotFound
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v models.ErrInternalError
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CloudsApiService
Get a Specific Cloud Data Store
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param cloudId The cloud ID
 * @param datastoreId The cloud datastore ID

*/
func (a *CloudsApiService) GetASpecificCloudDataStore(ctx context.Context, serviceInstanceId string, cloudId int32, datastoreId int32) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/{service_instance_id}/zones/{cloud_id}/data-stores/{datastore_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_instance_id"+"}", fmt.Sprintf("%v", serviceInstanceId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cloud_id"+"}", fmt.Sprintf("%v", cloudId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"datastore_id"+"}", fmt.Sprintf("%v", datastoreId), -1)

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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v models.ErrUnauthorized
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v models.ErrNotFound
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v models.ErrInternalError
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CloudsApiService
Get a Specific Cloud Resource Pool
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param cloudId The cloud ID
 * @param funId The Cloud Resourcepool ID

*/
func (a *CloudsApiService) GetASpecificCloudResourcePool(ctx context.Context, serviceInstanceId string, cloudId int32, resourcepoolId int32) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.Cfg.BasePath + "/v1/{service_instance_id}/zones/{cloud_id}/resource-pools/{resourcepool_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_instance_id"+"}", fmt.Sprintf("%v", serviceInstanceId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cloud_id"+"}", fmt.Sprintf("%v", cloudId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourcepool_id"+"}", fmt.Sprintf("%v", resourcepoolId), -1)

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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v models.ErrUnauthorized
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v models.ErrNotFound
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v models.ErrInternalError
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
CloudsApiService
Get All Cloud Data Stores
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param cloudId The cloud ID

*/
func (a *CloudsApiService) GetAllCloudDataStores(ctx context.Context, serviceInstanceId string, cloudId int32) (models.DataStoresResp, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%s/%d/data-stores", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		serviceInstanceId, consts.CloudsPath, cloudId)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return datastoresResp, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return datastoresResp, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return datastoresResp, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v models.ErrUnauthorized
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return datastoresResp, newErr
			}
			newErr.model = v
			return datastoresResp, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v models.ErrNotFound
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return datastoresResp, newErr
			}
			newErr.model = v
			return datastoresResp, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v models.ErrInternalError
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return datastoresResp, newErr
			}
			newErr.model = v
			return datastoresResp, newErr
		}
		return datastoresResp, newErr
	}
	if err = json.Unmarshal(localVarBody, &datastoresResp); err != nil {
		return datastoresResp, err
	}
	return datastoresResp, nil
}

/*
CloudsApiService
Get All Cloud Resource Pools
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId
 * @param cloudId The cloud ID

*/
func (a *CloudsApiService) GetAllCloudResourcePools(ctx context.Context, serviceInstanceId string, cloudId int, queryParams map[string]string) (models.ResourcePoolsResp, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%s/%d/resource-pools", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		serviceInstanceId, consts.CloudsPath, cloudId)

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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return resourcePoolsResp, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return resourcePoolsResp, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return resourcePoolsResp, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v models.ErrUnauthorized
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return resourcePoolsResp, newErr
			}
			newErr.model = v
			return resourcePoolsResp, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v models.ErrNotFound
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return resourcePoolsResp, newErr
			}
			newErr.model = v
			return resourcePoolsResp, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v models.ErrInternalError
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return resourcePoolsResp, newErr
			}
			newErr.model = v
			return resourcePoolsResp, newErr
		}
		return resourcePoolsResp, newErr
	}
	if err = json.Unmarshal(localVarBody, &resourcePoolsResp); err != nil {
		return resourcePoolsResp, err
	}
	return resourcePoolsResp, nil
}

/*
CloudsApiService
Get All Clouds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceInstanceId

*/
func (a *CloudsApiService) GetAllClouds(ctx context.Context, serviceInstanceId string) (models.CloudsResp, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := fmt.Sprintf("%s/%s/%s/%s", a.Cfg.Host, consts.VmaasCmpApiBasePath,
		serviceInstanceId, consts.CloudsPath)

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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key

		}
	}
	var cloudsResp models.CloudsResp
	r, err := a.Client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return cloudsResp, err
	}

	localVarHttpResponse, err := a.Client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return cloudsResp, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return cloudsResp, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v models.ErrUnauthorized
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return cloudsResp, newErr
			}
			newErr.model = v
			return cloudsResp, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v models.ErrInternalError
			err = a.Client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return cloudsResp, newErr
			}
			newErr.model = v
			return cloudsResp, newErr
		}
		return cloudsResp, newErr
	}
	if err = json.Unmarshal(localVarBody, &cloudsResp); err != nil {
		return cloudsResp, err
	}
	return cloudsResp, nil
}
