// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	models "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
)

func TestInstancesAPIService_CloneAnInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//templateName := "test_clone_an_instance"
	tests := []struct {
		name       string
		param      models.CreateInstanceBody
		instanceID int
		given      func(m *MockAPIClientHandler)
		want       models.SuccessOrErrorMessage
		wantErr    bool
	}{
		{
			name: "Normal Test case 1: Clone an Instance",
			param: models.CreateInstanceBody{
				ZoneID:    "1",
				CloneName: "Instance_Clone",
			},
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/clone"
				method := "PUT"
				//headers := getDefaultHeaders()
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"zoneId": "1",
						"name": "Instance_Clone"
					}
				`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"success": true,
						"message": "Successfully cloned an instance"
					}
				`)))
				// mock the context only since it is not validated in this function
				pBody, _ := json.Marshal(models.CreateInstanceBody{
					ZoneID:    "1",
					CloneName: "Instance_Clone",
				})
				//pBody := []byte(`{"zoneId":"1","CloneName":"Instance_Clone"}`)
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)
			},
			want: models.SuccessOrErrorMessage{
				Success: true,
				Message: "Successfully cloned an instance",
			},
			wantErr: false,
		},
		{
			name: "Failed test case 2: Error in call prepare request",
			param: models.CreateInstanceBody{
				ZoneID:    "1",
				CloneName: "Instance_Clone",
			},
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/clone"
				method := "PUT"
				//headers := getDefaultHeaders()
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				// mock the context only since it is not validated in this function
				pBody, _ := json.Marshal(models.CreateInstanceBody{
					ZoneID:    "1",
					CloneName: "Instance_Clone",
				})
				//pBody := []byte(`{"ZoneID":"1","CloneName":"Instance_Clone"}`)
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.SuccessOrErrorMessage{},
			wantErr: true,
		},
		{
			name: "Failed test case 3: error in callAPI",
			param: models.CreateInstanceBody{
				ZoneID:    "1",
				CloneName: "Instance_Clone",
			},
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/clone"
				method := "PUT"
				//headers := getDefaultHeaders()
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"zoneId": "1",
						"name": "Instance_Clone"
					}
				`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				// mock the context only since it is not validated in this function
				pBody, _ := json.Marshal(models.CreateInstanceBody{
					ZoneID:    "1",
					CloneName: "Instance_Clone",
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)
			},
			want:    models.SuccessOrErrorMessage{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := InstancesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.CloneAnInstance(ctx, tt.instanceID, &tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesAPIService.CloneAnInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesAPIService.CloneAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstancesAPIService_CreateAnInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		param   models.CreateInstanceBody
		given   func(m *MockAPIClientHandler)
		want    models.GetInstanceResponse
		wantErr bool
	}{
		{
			name: "Normal Test case 1: Create an Instance",
			param: models.CreateInstanceBody{
				ZoneID:    "1",
				CloneName: "Instance_Create",
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances"
				method := "POST"
				//headers := getDefaultHeaders()
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"zoneId": "1",
						"name": "Instance_Create"
					}
				`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"instance": {
						"id": 1,
						"name": "test_create_an_instance"
					}
				}
				`)))
				pBody, _ := json.Marshal(models.CreateInstanceBody{
					ZoneID:    "1",
					CloneName: "Instance_Create",
				})

				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)

			},
			want: models.GetInstanceResponse{
				Instance: &models.GetInstanceResponseInstance{
					ID:   1,
					Name: "test_create_an_instance",
				},
			},
			wantErr: false,
		},
		{
			name: "Failed test case 2: Error in call prepare request",
			param: models.CreateInstanceBody{
				ZoneID:    "1",
				CloneName: "Instance_Create",
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances"
				method := "POST"
				//headers := getDefaultHeaders()
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				// mock the context only since it is not validated in this function
				pBody, _ := json.Marshal(models.CreateInstanceBody{
					ZoneID:    "1",
					CloneName: "Instance_Create",
				})
				//pBody := []byte(`{"zoneId":"1","CloneName":"Instance_Clone"}`)
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))

			},
			want:    models.GetInstanceResponse{},
			wantErr: true,
		},
		{
			name: "Failed test case 3: error in callAPI",
			param: models.CreateInstanceBody{
				ZoneID:    "1",
				CloneName: "Instance_Create",
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances"
				method := "POST"
				//headers := getDefaultHeaders()
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"zoneId": "1",
						"name": "Instance_Create"
					}
				`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				// mock the context only since it is not validated in this function
				pBody, _ := json.Marshal(models.CreateInstanceBody{
					ZoneID:    "1",
					CloneName: "Instance_Create",
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)

			},
			want:    models.GetInstanceResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := InstancesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.CreateAnInstance(ctx, &tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesAPIService.CreateAnInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesAPIService.CreateAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstancesAPIService_DeleteAnInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		instanceID int
		given      func(m *MockAPIClientHandler)
		want       models.SuccessOrErrorMessage
		wantErr    bool
	}{
		{
			name:       "Normal Test case 1: Delete an Instance",
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "DELETE"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"success": true,
					"message": "Successfully Deleted the instance"
				}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)
			},
			want: models.SuccessOrErrorMessage{
				Success: true,
				Message: "Successfully Deleted the instance",
			},
			wantErr: false,
		},
		{
			name:       "Failed test case 2: Error in call prepare request",
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "DELETE"
				headers := getDefaultHeaders()
				m.EXPECT().prepareRequest(ctx, path, method, nil, headers, url.Values{},
					url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.SuccessOrErrorMessage{},
			wantErr: true,
		}, {
			name:       "Failed test case 3: error in callAPI",
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "DELETE"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				m.EXPECT().prepareRequest(ctx, path, method, nil, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)
			},
			want:    models.SuccessOrErrorMessage{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := InstancesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.DeleteAnInstance(ctx, tt.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesAPIService.DeleteAnInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesAPIService.DeleteAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestInstancesAPIService_ImportSnapshotOfAnInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		instanceID int
		param      models.ImportSnapshotBody
		given      func(m *MockAPIClientHandler)
		want       models.SuccessOrErrorMessage
		wantErr    bool
	}{
		{
			name:       "Normal Test case 1: Import Snapshot Of an Instance",
			instanceID: 1,
			param: models.ImportSnapshotBody{
				StorageProviderID: 1,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/import-snapshot"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`{
					"storageProviderId" : 1
				}`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"success": true,
					"message": "Successfully imported a snapshot of an instance"
				}
				`)))
				pBody, _ := json.Marshal(models.ImportSnapshotBody{
					StorageProviderID: 1,
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)
			},
			want: models.SuccessOrErrorMessage{
				Success: true,
				Message: "Successfully imported a snapshot of an instance",
			},
			wantErr: false,
		},
		{
			name:       "Failed test case 2: Error in call prepare request",
			instanceID: 1,
			param: models.ImportSnapshotBody{
				StorageProviderID: 1,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/import-snapshot"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				pBody, _ := json.Marshal(models.ImportSnapshotBody{
					StorageProviderID: 1})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.SuccessOrErrorMessage{},
			wantErr: true,
		},
		{
			name:       "Failed test case 3: error in callAPI",
			instanceID: 1,
			param: models.ImportSnapshotBody{
				StorageProviderID: 1,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/import-snapshot"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`{
					"storageProviderId" : 1
				}`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				pBody, _ := json.Marshal(models.ImportSnapshotBody{
					StorageProviderID: 1,
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)
			},
			want:    models.SuccessOrErrorMessage{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := InstancesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.ImportSnapshotOfAnInstance(ctx, tt.instanceID, &tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesAPIService.ImportSnapshotOfAnInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesAPIService.ImportSnapshotOfAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstancesAPIService_ResizeAnInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		instanceID int
		param      models.ResizeInstanceBody
		given      func(m *MockAPIClientHandler)
		want       models.ResizeInstanceResponse
		wantErr    bool
	}{
		{
			name:       "Normal Test case 1: Resize an Instance",
			instanceID: 1,
			param: models.ResizeInstanceBody{
				Instance: &models.ResizeInstanceBodyInstance{
					ID: 1,
				},
				Volumes: []models.ResizeInstanceBodyInstanceVolumes{{
					ID:   "1",
					Name: "test_instance_volume",
				}},
				DeleteOriginalVolumes: false,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/resize"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"instance" : {"id": 1},
					"volumes" : [{"id": 1, "name": "test_instance_volume" }],
					"deleteOriginalVolumes" : false
				}`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`{
					"instance": {
						"id": 1,
						"string": "test_instance_response"
					}
				}`)))
				pBody, _ := json.Marshal(models.ResizeInstanceBody{
					Instance: &models.ResizeInstanceBodyInstance{
						ID: 1,
					},
					Volumes: []models.ResizeInstanceBodyInstanceVolumes{{
						ID:   "1",
						Name: "test_instance_volume",
					}},
					DeleteOriginalVolumes: false,
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)
			},
			want: models.ResizeInstanceResponse{
				Instance: &models.ResizeInstanceResponseInstance{
					ID:   1,
					Name: "test_instance_response",
				},
			},
			wantErr: false,
		},
		{
			name:       "Failed test case 2: error in prepare request",
			instanceID: 1,
			param: models.ResizeInstanceBody{
				Instance: &models.ResizeInstanceBodyInstance{
					ID: 1,
				},
				Volumes: []models.ResizeInstanceBodyInstanceVolumes{{
					ID:   "1",
					Name: "test_instance_volume",
				}},
				DeleteOriginalVolumes: false,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/resize"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				pBody, _ := json.Marshal(models.ResizeInstanceBody{
					Instance: &models.ResizeInstanceBodyInstance{
						ID: 1,
					},
					Volumes: []models.ResizeInstanceBodyInstanceVolumes{{
						ID:   "1",
						Name: "test_instance_volume",
					}},
					DeleteOriginalVolumes: false,
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.ResizeInstanceResponse{},
			wantErr: true,
		},
		{
			name:       "Failed test case 3: error in callAPI",
			instanceID: 1,
			param: models.ResizeInstanceBody{
				Instance: &models.ResizeInstanceBodyInstance{
					ID: 1,
				},
				Volumes: []models.ResizeInstanceBodyInstanceVolumes{{
					ID:   "1",
					Name: "test_instance_volume",
				}},
				DeleteOriginalVolumes: false,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/resize"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"instance" : {"id": 1},
					"volumes" : [{"id": 1, "name": "test_instance_volume" }],
					"deleteOriginalVolumes" : false
				}`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				pBody, _ := json.Marshal(models.ResizeInstanceBody{
					Instance: &models.ResizeInstanceBodyInstance{
						ID: 1,
					},
					Volumes: []models.ResizeInstanceBodyInstanceVolumes{{
						ID:   "1",
						Name: "test_instance_volume",
					}},
					DeleteOriginalVolumes: false,
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)
			},
			want:    models.ResizeInstanceResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := InstancesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.ResizeAnInstance(ctx, tt.instanceID, &tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesAPIService.ResizeAnInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesAPIService.ResizeAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstancesAPIService_SnapshotAnInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		instanceID int
		param      models.SnapshotBody
		given      func(m *MockAPIClientHandler)
		want       models.Instances
		wantErr    bool
	}{
		{
			name:       "Normal Test case 1: Resize an Instance",
			instanceID: 1,
			param: models.SnapshotBody{
				Snapshot: &models.SnapshotBodySnapshot{
					Name:        "test_snapshot_name",
					Description: "test_snapshot_description",
				},
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/snapshot"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"snapshot": {
						"name": "test_snapshot_name",
						"description" : "test_snapshot_description"
					}
				}`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`{
					"instances": [{
						"id": 1,
						"name": "test_snapshot_name"
					}],
					"success" : true
				}`)))
				pBody, _ := json.Marshal(models.SnapshotBody{
					Snapshot: &models.SnapshotBodySnapshot{
						Name:        "test_snapshot_name",
						Description: "test_snapshot_description",
					},
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)
			},
			want: models.Instances{
				Instances: []models.GetInstanceResponseInstance{{
					ID:   1,
					Name: "test_snapshot_name",
				}},
				Success: true,
			},
			wantErr: false,
		},
		{
			name:       "Failed test case 2: error in prepare request",
			instanceID: 1,
			param: models.SnapshotBody{
				Snapshot: &models.SnapshotBodySnapshot{
					Name:        "test_snapshot_name",
					Description: "test_snapshot_description",
				},
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/snapshot"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				pBody, _ := json.Marshal(models.SnapshotBody{
					Snapshot: &models.SnapshotBodySnapshot{
						Name:        "test_snapshot_name",
						Description: "test_snapshot_description",
					},
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.Instances{},
			wantErr: true,
		},
		{
			name:       "Failed test case 3: error in callAPII",
			instanceID: 1,
			param: models.SnapshotBody{
				Snapshot: &models.SnapshotBodySnapshot{
					Name:        "test_snapshot_name",
					Description: "test_snapshot_description",
				},
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/snapshot"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"snapshot": {
						"name": "test_snapshot_name",
						"description" : "test_snapshot_description"
					}
				}`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				pBody, _ := json.Marshal(models.SnapshotBody{
					Snapshot: &models.SnapshotBodySnapshot{
						Name:        "test_snapshot_name",
						Description: "test_snapshot_description",
					},
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)
			},
			want:    models.Instances{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := InstancesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.SnapshotAnInstance(ctx, tt.instanceID, &tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesAPIService.SnapshotAnInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesAPIService.SnapshotAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstancesAPIService_UpdatingAnInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		instanceID int
		param      models.UpdateInstanceBody
		given      func(m *MockAPIClientHandler)
		want       models.UpdateInstanceResponse
		wantErr    bool
	}{
		{
			name:       "Normal Test case 1: Update an Instance",
			instanceID: 1,
			param: models.UpdateInstanceBody{
				Instance: &models.UpdateInstanceBodyInstance{
					Name:        "test_update_instance_name",
					Description: "test_update_instance_description",
				},
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"snapshot": {
						"name": "test_update_instance_name",
						"description" : "test_update_instance_description"
					}
				}`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`{
					"instance": {
						"id": 1,
						"name": "test_update_instance_name"
					}
				}`)))
				pBody, _ := json.Marshal(models.UpdateInstanceBody{
					Instance: &models.UpdateInstanceBodyInstance{
						Name:        "test_update_instance_name",
						Description: "test_update_instance_description",
					},
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)
			},
			want: models.UpdateInstanceResponse{
				Instance: &models.UpdateInstanceResponseInstance{
					Name: "test_update_instance_name",
					ID:   1,
				},
			},
			wantErr: false,
		},
		{
			name:       "Failed test case 2: error in prepare request",
			instanceID: 1,
			param: models.UpdateInstanceBody{
				Instance: &models.UpdateInstanceBodyInstance{
					Name:        "test_update_instance_name",
					Description: "test_update_instance_description",
				},
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}

				pBody, _ := json.Marshal(models.UpdateInstanceBody{
					Instance: &models.UpdateInstanceBodyInstance{
						Name:        "test_update_instance_name",
						Description: "test_update_instance_description",
					},
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.UpdateInstanceResponse{},
			wantErr: true,
		},
		{
			name:       "Failed test case 3: error in callAPI",
			instanceID: 1,
			param: models.UpdateInstanceBody{
				Instance: &models.UpdateInstanceBodyInstance{
					Name:        "test_update_instance_name",
					Description: "test_update_instance_description",
				},
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "PUT"
				headers := map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				}
				postBody := ioutil.NopCloser(bytes.NewReader([]byte(`
				{
					"snapshot": {
						"name": "test_update_instance_name",
						"description" : "test_update_instance_description"
					}
				}`)))
				req, _ := http.NewRequest(method, path, postBody)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`{
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				}`)))
				pBody, _ := json.Marshal(models.UpdateInstanceBody{
					Instance: &models.UpdateInstanceBodyInstance{
						Name:        "test_update_instance_name",
						Description: "test_update_instance_description",
					},
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)
				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)
			},
			want:    models.UpdateInstanceResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := InstancesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.UpdatingAnInstance(ctx, tt.instanceID, &tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesAPIService.UpdatingAnInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesAPIService.UpdatingAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestInstancesAPIService_GetASpecificInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		instanceID int
		// All expectaion captures here
		given   func(m *MockAPIClientHandler)
		want    models.GetInstanceResponse
		wantErr bool
	}{
		{
			name:       "Normal Test case 1: Get a specific instance",
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"instance": {
							"id": 1,
							"name": "test_instance"
						}
					}
				`)))
				// mock the context only since it is not validated in this function
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)

			},
			want: models.GetInstanceResponse{
				Instance: &models.GetInstanceResponseInstance{
					ID:   1,
					Name: "test_instance",
				},
			},
			wantErr: false,
		},
		{
			name:       "Failed test case 2: Error in call prepare request",
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "GET"
				headers := getDefaultHeaders()
				// mock the context only since it is not validated in this function
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{},
					url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))

			},
			want:    models.GetInstanceResponse{},
			wantErr: true,
		},
		{
			name:       "Failed test case 3: error in callAPI",
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				// mock the context only since it is not validated in this function
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)

			},
			want:    models.GetInstanceResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := InstancesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}

			tt.given(mockAPIClient)
			got, err := a.GetASpecificInstance(ctx, tt.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesAPIService.GetASpecificInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesAPIService.GetASpecificInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstancesAPIService_GetAllInstances(t *testing.T) {

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	templateName := "test_all_instances"
	tests := []struct {
		name    string
		param   map[string]string
		given   func(m *MockAPIClientHandler)
		want    models.Instances
		wantErr bool
	}{
		{
			name: "Normal Test case 1: Get all instances",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"instances": [{
							"id": 1,
							"name": "test_all_instances"
						}],
						"success": true
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getURLValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)

			},
			want: models.Instances{
				Instances: []models.GetInstanceResponseInstance{
					{
						ID:   1,
						Name: "test_all_instances",
					},
				},
				Success: true,
			},
			wantErr: false,
		},
		{
			name: "Failed test case 2: Error in call prepare request",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances"
				method := "GET"
				headers := getDefaultHeaders()
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getURLValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))

			},
			want:    models.Instances{},
			wantErr: true,
		},
		{
			name: "Failed test case 3: error in callAPI",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getURLValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)

			},
			want:    models.Instances{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := InstancesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}

			tt.given(mockAPIClient)
			got, err := a.GetAllInstances(ctx, tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesAPIService.GetAllInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesAPIService.GetAllInstances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstancesAPIService_GetListOfSnapshotsForAnInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		instanceID int
		given      func(m *MockAPIClientHandler)
		want       models.ListSnapshotResponse
		wantErr    bool
	}{
		{
			name:       "Normal Test case 1: Get list of snapshots for an instance",
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/snapshots"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"snapshots": [{
							"id": 1,
							"name": "test_snapshots_for_instance"
						}]
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)

			},
			want: models.ListSnapshotResponse{
				Snapshots: []models.ListSnapshotResponseInstance{
					{
						ID:   1,
						Name: "test_snapshots_for_instance",
					},
				},
			},
			wantErr: false,
		},
		{
			name:       "Failed test case 2: Error in call prepare request",
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/snapshots"
				method := "GET"
				headers := getDefaultHeaders()
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{},
					url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))

			},
			want:    models.ListSnapshotResponse{},
			wantErr: true,
		},
		{
			name:       "Failed test case 3: error in callAPI",
			instanceID: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1/snapshots"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"message": "Internal Server Error",
						"recommendedActions": [
							"Unknown error occurred. Please contact the administrator"
						]
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{},
					url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)

			},
			want:    models.ListSnapshotResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := InstancesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.GetListOfSnapshotsForAnInstance(ctx, tt.instanceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesAPIService.GetListOfSnapshotsForAnInstance() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesAPIService.GetListOfSnapshotsForAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}