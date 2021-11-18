// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type LibraryAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
LibrariesAPIService
Get All layouts
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID

*/
func (a *LibraryAPIService) GetAllLayouts(ctx context.Context,
	param map[string]string) (models.LayoutsResp, error) {
	response := models.LayoutsResp{}

	allLayoutsAPI := &api{
		method: "GET",
		path:   consts.LibraryLayoutPath,
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &response)
		},
	}
	err := allLayoutsAPI.do(ctx, nil, param)

	return response, err
}

func (a *LibraryAPIService) GetAllInstanceTypes(ctx context.Context,
	param map[string]string) (models.InstanceTypesResp, error) {
	response := models.InstanceTypesResp{}

	allInstTypeAPI := &api{
		method: "GET",
		path:   consts.LibraryInstanceTypesPath,
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &response)
		},
	}
	err := allInstTypeAPI.do(ctx, nil, param)

	return response, err
}
