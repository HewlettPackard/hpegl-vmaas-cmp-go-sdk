// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type PlansAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
PlansAPIService
Get All Service Plans
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
*/
func (a *PlansAPIService) GetAllServicePlans(ctx context.Context,
	param map[string]string) (models.ServicePlans, error) {
	response := models.ServicePlans{}

	allServicePlansAPI := &api{
		method: "GET",
		path:   consts.ServicePlansPath,
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &response)
		},
	}
	err := allServicePlansAPI.do(ctx, nil, param)

	return response, err
}

func (a *PlansAPIService) GetSpecificServicePlan(
	ctx context.Context,
	planID int,
) (models.GetSpecificServicePlan, error) {
	response := models.GetSpecificServicePlan{}

	allServicePlansAPI := &api{
		method: "GET",
		path:   fmt.Sprintf("%s/%d", consts.ServicePlansPath, planID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &response)
		},
	}
	err := allServicePlansAPI.do(ctx, nil, nil)

	return response, err
}
