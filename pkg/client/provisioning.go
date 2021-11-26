// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	models "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type ProvisioningAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (a *ProvisioningAPIService) GetAllProvisioningTypes(ctx context.Context,
	param map[string]string) (models.GetAllProvisioningTypes, error) {
	response := models.GetAllProvisioningTypes{}

	allProvisionAPI := &api{
		method: "GET",
		path:   consts.ProvisionTypesPath,
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &response)
		},
	}

	err := allProvisionAPI.do(ctx, nil, param)

	return response, err
}
