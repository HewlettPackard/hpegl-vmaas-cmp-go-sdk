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

func TestInstancesApiService_CloneAnInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//templateName := "test_clone_an_instance"
	tests := []struct {
		name       string
		param      models.CreateInstanceBody
		instanceId int
		given      func(m *MockAPIClientHandler)
		want       models.SuccessOrErrorMessage
		wantErr    bool
	}{
		{
			name: "Normal Test case 1: Clone an Instance",
			param: models.CreateInstanceBody{
				ZoneId:    "1",
				CloneName: "Instance_Clone",
			},
			instanceId: 1,
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
					ZoneId:    "1",
					CloneName: "Instance_Clone",
				})
				//pBody := []byte(`{"ZoneId":"1","CloneName":"Instance_Clone"}`)
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)

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
				ZoneId:    "1",
				CloneName: "Instance_Clone",
			},
			instanceId: 1,
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
					ZoneId:    "1",
					CloneName: "Instance_Clone",
				})
				//pBody := []byte(`{"ZoneId":"1","CloneName":"Instance_Clone"}`)
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{}, url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.SuccessOrErrorMessage{},
			wantErr: true,
		},
		{
			name: "Failed test case 3: error in callAPI",
			param: models.CreateInstanceBody{
				ZoneId:    "1",
				CloneName: "Instance_Clone",
			},
			instanceId: 1,
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
					ZoneId:    "1",
					CloneName: "Instance_Clone",
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)

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
			a := InstancesApiService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.CloneAnInstance(ctx, tt.instanceId, &tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesApiService.CloneAnInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesApiService.CloneAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstancesApiService_CreateAnInstance(t *testing.T) {
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
				ZoneId:    "1",
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
					ZoneId:    "1",
					CloneName: "Instance_Create",
				})

				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)

			},
			want: models.GetInstanceResponse{
				Instance: &models.GetInstanceResponseInstance{
					Id:   1,
					Name: "test_create_an_instance",
				},
			},
			wantErr: false,
		},
		{
			name: "Failed test case 2: Error in call prepare request",
			param: models.CreateInstanceBody{
				ZoneId:    "1",
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
					ZoneId:    "1",
					CloneName: "Instance_Create",
				})
				//pBody := []byte(`{"ZoneId":"1","CloneName":"Instance_Clone"}`)
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{}, url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))

			},
			want:    models.GetInstanceResponse{},
			wantErr: true,
		},
		{
			name: "Failed test case 3: error in callAPI",
			param: models.CreateInstanceBody{
				ZoneId:    "1",
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
					ZoneId:    "1",
					CloneName: "Instance_Create",
				})
				m.EXPECT().prepareRequest(gomock.Any(), path, method, pBody, headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)

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
			a := InstancesApiService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.CreateAnInstance(ctx, &tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesApiService.CreateAnInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesApiService.CreateAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstancesApiService_DeleteAnInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		instanceId int
		given      func(m *MockAPIClientHandler)
		want       models.SuccessOrErrorMessage
		wantErr    bool
	}{
		{
			name:       "Normal Test case 1: Delete an Instance",
			instanceId: 1,
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
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)

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
			instanceId: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "DELETE"
				headers := getDefaultHeaders()
				m.EXPECT().prepareRequest(ctx, path, method, nil, headers, url.Values{}, url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.SuccessOrErrorMessage{},
			wantErr: true,
		}, {
			name:       "Failed test case 3: error in callAPI",
			instanceId: 1,
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
				m.EXPECT().prepareRequest(ctx, path, method, nil, headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)
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
			a := InstancesApiService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.DeleteAnInstance(ctx, tt.instanceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesApiService.DeleteAnInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesApiService.DeleteAnInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
