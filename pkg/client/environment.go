// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	models "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type EnvironmentAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (e *EnvironmentAPIService) GetAllEnvironment(ctx context.Context,
	param map[string]string) (models.GetAllEnvironment, error) {
	response := models.GetAllEnvironment{}

	allEnvAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s", e.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.EnvironmentPath),
		client: e.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &response)
		},
	}

	err := allEnvAPI.do(ctx, nil, param)

	return response, err
}
