// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/antihax/optional"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	models "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
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
	createInstanceResp := models.GetInstanceResponse{}

	// Pre-pending 'pool-' to ResourcePoolId in 6.0.3 and above
	if v, _ := parseVersion("6.0.3"); v <= a.Client.getVersion() {
		localVarOptionalsValue := reflect.ValueOf(localVarOptionals)
		if field := localVarOptionalsValue.FieldByName("Config"); field.IsValid() {
			configValue := reflect.ValueOf(localVarOptionals.Config)
			if configField := configValue.FieldByName("ResourcePoolID"); configField.IsValid() {
				localVarOptionals.Config.ResourcePoolID = fmt.Sprintf("pool-%v",
					localVarOptionals.Config.ResourcePoolID)
			}
		}
	}
	createInstanceAPI := &api{
		method: "POST",
		path:   consts.InstancesPath,
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &createInstanceResp)
		},
	}

	err := createInstanceAPI.do(ctx, localVarOptionals, nil)

	return createInstanceResp, err
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
	delInstanceResp := models.SuccessOrErrorMessage{}

	delInstanceAPI := &api{
		method: "DELETE",
		path: fmt.Sprintf("%s/%d",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &delInstanceResp)
		},
	}
	err := delInstanceAPI.do(ctx, nil, nil)

	return delInstanceResp, err
}

/*
InstancesAPIService
Get a Specific Instance
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
    Passed from http.Request or context.Background().
  - @param serviceInstanceID
  - @param instanceID

@return models.GetInstanceResponse
*/
func (a *InstancesAPIService) GetASpecificInstance(ctx context.Context,
	instanceID int) (models.GetInstanceResponse, error) {
	specificInstResp := models.GetInstanceResponse{}

	specificInstanceAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%d",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &specificInstResp)
		},
	}
	err := specificInstanceAPI.do(ctx, nil, nil)

	return specificInstResp, err
}

/*
InstancesAPIService
Get All Instances
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
*/

func (a *InstancesAPIService) GetAllInstances(ctx context.Context,
	queryParams map[string]string) (models.Instances, error) {
	getAllInstance := models.Instances{}

	instanceAPI := &api{
		method: "GET",
		path:   consts.InstancesPath,
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &getAllInstance)
		},
	}
	err := instanceAPI.do(ctx, nil, queryParams)

	return getAllInstance, err
}

/*
InstancesAPIService
Lists VMware Snapshot of the instance
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
    Passed from http.Request or context.Background().
  - @param serviceInstanceID
  - @param instanceID
*/
func (a *InstancesAPIService) GetListOfSnapshotsForAnInstance(ctx context.Context,
	instanceID int) (models.ListSnapshotResponse, error) {
	listSnapshotResp := models.ListSnapshotResponse{}

	listSnapshotAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%d/snapshots",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &listSnapshotResp)
		},
	}
	err := listSnapshotAPI.do(ctx, nil, nil)

	return listSnapshotResp, err
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
	importSnapshotResp := models.SuccessOrErrorMessage{}

	importSnapshotAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%d/import-snapshot",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &importSnapshotResp)
		},
	}
	err := importSnapshotAPI.do(ctx, localVarOptionals, nil)

	return importSnapshotResp, err
}

/*
InstancesAPIService
Restarts all VM running within an instance
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
    Passed from http.Request or context.Background().
  - @param serviceInstanceID
  - @param instanceID
*/
func (a *InstancesAPIService) RestartAnInstance(ctx context.Context,
	instanceID int) (models.InstancePowerResponse, error) {
	restartInstResp := models.InstancePowerResponse{}

	restartInstAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%d/restart",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &restartInstResp)
		},
	}
	err := restartInstAPI.do(ctx, nil, nil)

	return restartInstResp, err
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
	snapshotInstResp := models.Instances{}

	instanceAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%d/snapshot",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &snapshotInstResp)
		},
	}
	err := instanceAPI.do(ctx, localVarOptionals, nil)

	return snapshotInstResp, err
}

/*
InstancesAPIService
Starts all VM running within an instance
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
    Passed from http.Request or context.Background().
  - @param serviceInstanceID
  - @param instanceID
*/
func (a *InstancesAPIService) StartAnInstance(ctx context.Context,
	instanceID int) (models.InstancePowerResponse, error) {
	startInstanceResp := models.InstancePowerResponse{}

	startInstanceAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%d/start",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &startInstanceResp)
		},
	}
	err := startInstanceAPI.do(ctx, nil, nil)

	return startInstanceResp, err
}

/*
InstancesAPIService
Stops all VM running within an instance
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
    Passed from http.Request or context.Background().
  - @param serviceInstanceID
  - @param instanceID
*/
func (a *InstancesAPIService) StopAnInstance(ctx context.Context,
	instanceID int) (models.InstancePowerResponse, error) {
	stopInstanceResp := models.InstancePowerResponse{}

	stopInstanceAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%d/stop",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &stopInstanceResp)
		},
	}
	err := stopInstanceAPI.do(ctx, nil, nil)

	return stopInstanceResp, err
}

/*
InstancesAPIService
Suspends all VM running within an instance
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
    Passed from http.Request or context.Background().
  - @param serviceInstanceID
  - @param instanceID
*/
func (a *InstancesAPIService) SuspendAnInstance(ctx context.Context,
	instanceID int) (models.InstancePowerResponse, error) {
	suspendResp := models.InstancePowerResponse{}
	suspendInstanceAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%d/suspend",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &suspendResp)
		},
	}
	err := suspendInstanceAPI.do(ctx, nil, nil)

	return suspendResp, err
}

func (a *InstancesAPIService) ResizeAnInstance(ctx context.Context, instanceID int,
	request *models.ResizeInstanceBody) (models.ResizeInstanceResponse, error) {
	resizeResp := models.ResizeInstanceResponse{}

	instanceAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%d/resize",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resizeResp)
		},
	}
	err := instanceAPI.do(ctx, request, nil)

	return resizeResp, err
}

func (a *InstancesAPIService) UpdatingAnInstance(
	ctx context.Context,
	instanceID int,
	request *models.UpdateInstanceBody,
) (models.UpdateInstanceResponse, error) {
	instance := models.UpdateInstanceResponse{}

	instanceAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%d",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &instance)
		},
	}
	err := instanceAPI.do(ctx, request, nil)

	return instance, err
}

func (a *InstancesAPIService) GetInstanceHistory(
	ctx context.Context, instanceID int) (models.GetInstanceHistory, error) {
	history := models.GetInstanceHistory{}

	historyAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%d/history",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &history)
		},
	}
	err := historyAPI.do(ctx, nil, nil)

	return history, err
}

func (a *InstancesAPIService) CloneAnInstance(ctx context.Context, instanceID int,
	cloneRequest models.CreateInstanceCloneBody) (models.SuccessOrErrorMessage, error) {
	var cloneResp models.SuccessOrErrorMessage
	if v, _ := parseVersion("5.2.12"); v <= a.Client.getVersion() {
		cloneRequest.Tags = cloneRequest.Metadata
		cloneRequest.Metadata = nil
		cloneRequest.Instance.Labels = cloneRequest.Instance.Tags
		cloneRequest.Instance.Tags = nil
	}
	instanceClone := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%d/clone",
			consts.InstancesPath, instanceID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &cloneResp)
		},
	}
	err := instanceClone.do(ctx, cloneRequest, nil)

	return cloneResp, err
}
