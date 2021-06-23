// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	models "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
)

func TestInstancesApiService_GetASpecificInstance(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		instanceId int
		// All expectaion captures here
		given   func(m *MockAPIClientHandler)
		want    models.GetInstanceResponse
		wantErr bool
	}{
		{
			name:       "Normal Test case 1: Get a specific instance",
			instanceId: 1,
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
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)

			},
			want: models.GetInstanceResponse{
				Instance: &models.GetInstanceResponseInstance{
					Id:   1,
					Name: "test_instance",
				},
			},
			wantErr: false,
		},
		{
			name:       "Failed test case 2: Error in call prepare request",
			instanceId: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/instances/1"
				method := "GET"
				headers := getDefaultHeaders()
				// mock the context only since it is not validated in this function
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{}, url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))

			},
			want:    models.GetInstanceResponse{},
			wantErr: true,
		},
		{
			name:       "Failed test case 3: error in callAPI",
			instanceId: 1,
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
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)

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
			got, err := a.GetASpecificInstance(ctx, tt.instanceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstancesApiService.GetASpecificInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstancesApiService.GetASpecificInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
