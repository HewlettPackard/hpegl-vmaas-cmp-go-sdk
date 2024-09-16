// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

const (
	testServiceInstanceID    = "18ba6409-ac59-4eac-9414-0147e72d615e"
	testAccessToken          = "2b9fba7f-7c14-4773-a970-a9ad393811ac"
	testRefreshToken         = "2b9fba7f-7c14-4773-a970-a9ad393811ac"
	testMorpheusURL          = "https://1234-mp.private.greenlake.hpe-gl-intg.com/"
	testAccessTokenExpiresIn = 3600
	testRefreshTokenExpires  = 1758034360176
)

func TestBrokerAPIService_GetMorpheusDetails(t *testing.T) {
	ctx := context.Background()
	testCtrl := gomock.NewController(t)
	defer testCtrl.Finish()

	headers := getDefaultHeaders()

	queryParams := map[string]string{
		"location":   "BLR",
		"space_name": "default",
	}

	clientCfg := Configuration{
		DefaultHeader:      headers,
		DefaultQueryParams: queryParams,
	}

	tests := []struct {
		name    string
		given   func(m *MockAPIClientHandler)
		want    models.MorpheusDetails
		wantErr bool
	}{
		{
			name: "Test GetMorpheusDetails success",
			want: models.MorpheusDetails{
				ID:                   testServiceInstanceID,
				AccessToken:          testAccessToken,
				RefreshToken:         testRefreshToken,
				AccessTokenExpiresIn: testAccessTokenExpiresIn,
				RefreshTokenExpires:  testRefreshTokenExpires,
				URL:                  testMorpheusURL,
			},
			wantErr: false,
			given: func(m *MockAPIClientHandler) {
				// Get subscription details
				m.EXPECT().getHost().Return(mockHost)
				pathSubscription := mockHost + "/" + consts.SubscriptionDetails
				method := "GET"
				reqSubscription, _ := http.NewRequest(method, pathSubscription, nil)
				respBodySubscription := io.NopCloser(bytes.NewReader([]byte(`
					{
						"ServiceInstanceID": "` + testServiceInstanceID + `",	
						"URL": "` + testMorpheusURL + `"	
					}
				`)))
				// mock the context only since it is not validated in this function
				m.EXPECT().getVersion().Return(999999)
				m.EXPECT().prepareRequest(gomock.Any(), pathSubscription, method, nil, headers,
					getURLValues(queryParams), url.Values{}, "", nil).Return(reqSubscription, nil)

				m.EXPECT().callAPI(reqSubscription).Return(&http.Response{
					StatusCode: 200,
					Body:       respBodySubscription,
				}, nil)

				// Get Morpheus token
				m.EXPECT().getHost().Return(mockHost)
				pathToken := mockHost + "/" + fmt.Sprintf(consts.MorpheusToken, testServiceInstanceID)
				reqToken, _ := http.NewRequest(method, pathToken, nil)
				tokenResp := models.MorpheusTokenResponse{
					AccessToken:          testAccessToken,
					RefreshToken:         testRefreshToken,
					AccessTokenExpiresIn: testAccessTokenExpiresIn,
					RefreshTokenExpires:  testRefreshTokenExpires,
				}
				body, err := json.Marshal(tokenResp)
				assert.NoError(t, err)
				respBodyToken := io.NopCloser(bytes.NewReader(body))
				// mock the context only since it is not validated in this function
				m.EXPECT().getVersion().Return(999999)
				m.EXPECT().prepareRequest(gomock.Any(), pathToken, method, nil, headers,
					url.Values{}, url.Values{}, "", nil).Return(reqToken, nil)

				m.EXPECT().callAPI(reqToken).Return(&http.Response{
					StatusCode: 200,
					Body:       respBodyToken,
				}, nil)
			},
		},

		{
			name:    "Test GetMorpheusDetails error in get subscription details prepare request",
			want:    models.MorpheusDetails{},
			wantErr: true,
			given: func(m *MockAPIClientHandler) {
				// Get subscription details
				m.EXPECT().getHost().Return(mockHost)
				pathSubscription := mockHost + "/" + consts.SubscriptionDetails
				method := "GET"
				// mock the context only since it is not validated in this function
				m.EXPECT().getVersion().Return(999999)
				m.EXPECT().prepareRequest(gomock.Any(), pathSubscription, method, nil, headers,
					getURLValues(queryParams), url.Values{}, "", nil).
					Return(nil, errors.New("error in prepare request"))
			},
		},

		{
			name:    "Test GetMorpheusDetails error in get subscription details call API",
			want:    models.MorpheusDetails{},
			wantErr: true,
			given: func(m *MockAPIClientHandler) {
				// Get subscription details
				m.EXPECT().getHost().Return(mockHost)
				pathSubscription := mockHost + "/" + consts.SubscriptionDetails
				method := "GET"
				reqSubscription, _ := http.NewRequest(method, pathSubscription, nil)
				respBodySubscription := io.NopCloser(bytes.NewReader([]byte(`
					{
						"ServiceInstanceID": "` + testServiceInstanceID + `",	
						"URL": "` + testMorpheusURL + `"	
					}
				`)))
				// mock the context only since it is not validated in this function
				m.EXPECT().getVersion().Return(999999)
				m.EXPECT().prepareRequest(gomock.Any(), pathSubscription, method, nil, headers,
					getURLValues(queryParams), url.Values{}, "", nil).Return(reqSubscription, nil)

				m.EXPECT().callAPI(reqSubscription).Return(&http.Response{
					StatusCode: 500,
					Body:       respBodySubscription,
				}, nil)
			},
		},

		{
			name:    "Test GetMorpheusDetails error in get Morpheus token prepare request",
			want:    models.MorpheusDetails{},
			wantErr: true,
			given: func(m *MockAPIClientHandler) {
				// Get subscription details
				m.EXPECT().getHost().Return(mockHost)
				pathSubscription := mockHost + "/" + consts.SubscriptionDetails
				method := "GET"
				reqSubscription, _ := http.NewRequest(method, pathSubscription, nil)
				respBodySubscription := io.NopCloser(bytes.NewReader([]byte(`
					{
						"ServiceInstanceID": "` + testServiceInstanceID + `",	
						"URL": "` + testMorpheusURL + `"	
					}
				`)))
				// mock the context only since it is not validated in this function
				m.EXPECT().getVersion().Return(999999)
				m.EXPECT().prepareRequest(gomock.Any(), pathSubscription, method, nil, headers,
					getURLValues(queryParams), url.Values{}, "", nil).Return(reqSubscription, nil)

				m.EXPECT().callAPI(reqSubscription).Return(&http.Response{
					StatusCode: 200,
					Body:       respBodySubscription,
				}, nil)

				// Get Morpheus token
				m.EXPECT().getHost().Return(mockHost)
				pathToken := mockHost + "/" + fmt.Sprintf(consts.MorpheusToken, testServiceInstanceID)
				// mock the context only since it is not validated in this function
				m.EXPECT().getVersion().Return(999999)
				m.EXPECT().prepareRequest(gomock.Any(), pathToken, method, nil, headers,
					url.Values{}, url.Values{}, "", nil).
					Return(nil, errors.New("error in prepare request"))
			},
		},

		{
			name:    "Test GetMorpheusDetails error in get Morpheus token call API",
			want:    models.MorpheusDetails{},
			wantErr: true,
			given: func(m *MockAPIClientHandler) {
				// Get subscription details
				m.EXPECT().getHost().Return(mockHost)
				pathSubscription := mockHost + "/" + consts.SubscriptionDetails
				method := "GET"
				reqSubscription, _ := http.NewRequest(method, pathSubscription, nil)
				respBodySubscription := io.NopCloser(bytes.NewReader([]byte(`
					{
						"ServiceInstanceID": "` + testServiceInstanceID + `",	
						"URL": "` + testMorpheusURL + `"	
					}
				`)))
				// mock the context only since it is not validated in this function
				m.EXPECT().getVersion().Return(999999)
				m.EXPECT().prepareRequest(gomock.Any(), pathSubscription, method, nil, headers,
					getURLValues(queryParams), url.Values{}, "", nil).Return(reqSubscription, nil)

				m.EXPECT().callAPI(reqSubscription).Return(&http.Response{
					StatusCode: 200,
					Body:       respBodySubscription,
				}, nil)

				// Get Morpheus token
				m.EXPECT().getHost().Return(mockHost)
				pathToken := mockHost + "/" + fmt.Sprintf(consts.MorpheusToken, testServiceInstanceID)
				reqToken, _ := http.NewRequest(method, pathToken, nil)
				tokenResp := models.MorpheusTokenResponse{
					AccessToken:          testAccessToken,
					RefreshToken:         testRefreshToken,
					AccessTokenExpiresIn: testAccessTokenExpiresIn,
					RefreshTokenExpires:  testRefreshTokenExpires,
				}
				body, err := json.Marshal(tokenResp)
				assert.NoError(t, err)
				respBodyToken := io.NopCloser(bytes.NewReader([]byte(body)))
				// mock the context only since it is not validated in this function
				m.EXPECT().getVersion().Return(999999)
				m.EXPECT().prepareRequest(gomock.Any(), pathToken, method, nil, headers,
					url.Values{}, url.Values{}, "", nil).Return(reqToken, nil)

				m.EXPECT().callAPI(reqToken).Return(&http.Response{
					StatusCode: 500,
					Body:       respBodyToken,
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := NewMockAPIClientHandler(testCtrl)
			tt.given(mockClient)
			a := &BrokerAPIService{
				Cfg:    clientCfg,
				Client: mockClient,
			}
			got, err := a.GetMorpheusDetails(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("BrokerAPIService.GetMorpheusDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BrokerAPIService.GetMorpheusDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
