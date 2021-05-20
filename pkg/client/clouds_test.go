//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCloudsApiService_GetASpecificCloud(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatusOk := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString("{}")),
	}

	mockAPIClient := NewMockAPIClientHandler(ctrl)
	mockAPIClient.EXPECT().prepareRequest(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
	mockAPIClient.EXPECT().callAPI(gomock.Any()).Return(mockStatusOk, nil)
	a := &CloudsApiService{
		Client: mockAPIClient,
		Cfg:    Configuration{},
	}
	_, err := a.GetASpecificCloud(context.Background(), "", 1)
	assert.Nil(t, err)
}
