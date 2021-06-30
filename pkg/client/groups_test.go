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
	"github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
)

func TestGroupsApiService_GetASpecificGroup(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	templateName := "test_template_get_a_group"

	tests := []struct {
		name    string
		groupId int
		given   func(m *MockAPIClientHandler)
		want    models.GroupResp
		wantErr bool
	}{
		{
			name:    "Normal Test case 1: Get a specific group",
			groupId: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/groups/1"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"group": {
							"id": 1,
							"name": "test_template_get_a_group"
						}
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)

			},
			want: models.GroupResp{
				Group: &models.Group{
					Id:   1,
					Name: templateName,
				},
			},
			wantErr: false,
		},
		{
			name:    "Failed Test case 2: Error in prepare requst",
			groupId: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/groups/1"
				method := "GET"
				headers := getDefaultHeaders()
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{}, url.Values{}, "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.GroupResp{},
			wantErr: true,
		},
		{
			name:    "Failed Test case 3: Error in callAPI",
			groupId: 1,
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/groups/1"
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
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{}, url.Values{}, "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)

			},
			want:    models.GroupResp{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := GroupsApiService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}

			tt.given(mockAPIClient)
			got, err := a.GetASpecificGroup(ctx, tt.groupId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupsApiService.GetASpecificGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupsApiService.GetASpecificGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupsApiService_GetAllGroups(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	templateName := "test_template_get_all_groups"

	tests := []struct {
		name    string
		param   map[string]string
		given   func(m *MockAPIClientHandler)
		want    models.Groups
		wantErr bool
	}{
		{
			name: "Normal Test case 1: Get all groups",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/groups"
				method := "GET"
				headers := getDefaultHeaders()
				req, _ := http.NewRequest(method, path, nil)
				respBody := ioutil.NopCloser(bytes.NewReader([]byte(`
					{
						"groups": [{
							"id": 1,
							"name": "test_template_get_all_groups"
						}]
					}
				`)))
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{}, getUrlValues(map[string]string{
					"name": templateName,
				}), "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 200,
					Body:       respBody,
				}, nil)

			},
			want: models.Groups{
				Groups: &[]models.Group{
					{
						Id:   1,
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
				path := mockHost + "/v1/groups"
				method := "GET"
				headers := getDefaultHeaders()
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{}, getUrlValues(map[string]string{
					"name": templateName,
				}), "", nil).Return(nil, errors.New("prepare request error"))
			},
			want:    models.Groups{},
			wantErr: true,
		},
		{
			name: "Failes Test case 3: Error in callAPI",
			param: map[string]string{
				"name": templateName,
			},
			given: func(m *MockAPIClientHandler) {
				path := mockHost + "/v1/groups"
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
				m.EXPECT().prepareRequest(gomock.Any(), path, method, nil, headers, url.Values{}, getUrlValues(map[string]string{
					"name": templateName,
				}), "", nil).Return(req, nil)

				m.EXPECT().callAPI(req).Return(&http.Response{
					StatusCode: 500,
					Body:       respBody,
				}, nil)

			},
			want:    models.Groups{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAPIClient := NewMockAPIClientHandler(ctrl)
			a := GroupsApiService{
				Client: mockAPIClient,
				Cfg: Configuration{
					Host: mockHost,
				},
			}

			tt.given(mockAPIClient)
			got, err := a.GetAllGroups(ctx, tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupsApiService.GetAllGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupsApiService.GetAllGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
