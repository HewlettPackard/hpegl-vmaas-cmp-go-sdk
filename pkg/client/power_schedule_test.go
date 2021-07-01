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

func TestPowerSchedulesAPIService_GetAllPowerSchedules(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	templateName := "test_template_get_all_power_schedules"

	tests := []struct {
		name    string
		param   map[string]string
		given   func(m *MockAPIClientHandler)
		want    models.GetAllPowerSchedules
		wantErr bool
	}{
		{
			name: "Normal Test case 1: Get all all power schedules",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/power-schedules"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"schedules": [{
							"id": 1,
							"name": "test_template_get_all_power_schedules"
						}]
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getUrlValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)

			},
			want: models.GetAllPowerSchedules{
				Schedules: []models.GetAllPowerSchedulesSchedules{
					{
						ID:   1,
						Name: templateName,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Failed Test case 2: Error in prepare request",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/power-schedules"
				method := "GET"
				headers := getDefaultHeaders()
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers,
					getUrlValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(nil, errors.New("prepare error request"))
			},
			want:    models.GetAllPowerSchedules{},
			wantErr: true,
		},
		{
			name: "Failed Test case 3: Error in callAPI",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/power-schedules"
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
					getUrlValues(map[string]string{
						"name": templateName,
					}), url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)

			},
			want:    models.GetAllPowerSchedules{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := PowerSchedulesAPIService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}
			tt.given(mockAPIClient)
			got, err := a.GetAllPowerSchedules(ctx, tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("PowerSchedulesAPIService.GetAllPowerSchedules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PowerSchedulesAPIService.GetAllPowerSchedules() = %v, want %v", got, tt.want)
			}
		})
	}
}
