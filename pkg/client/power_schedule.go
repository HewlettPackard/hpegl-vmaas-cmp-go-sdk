// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	models "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type PowerSchedulesAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
VirtualImageApiService
Get All Virtual images
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param name/phrase optional
@return models.VirtualImages
*/
func (a *PowerSchedulesAPIService) GetAllPowerSchedules(ctx context.Context,
	param map[string]string) (models.GetAllPowerSchedules, error) {
	response := models.GetAllPowerSchedules{}

	allPowerScheduleAPI := &api{
		method: "GET",
		path:   consts.PowerSchedulPath,
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &response)
		},
	}
	err := allPowerScheduleAPI.do(ctx, nil, param)

	return response, err
}
