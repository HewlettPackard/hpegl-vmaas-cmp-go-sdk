//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type NetworksAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (a *NetworksAPIService) GetAllNetworks(ctx context.Context,
	param map[string]string) (models.ListNetworksBody, error) {
	var networksResp models.ListNetworksBody
	instanceClone := &api{
		method: "GET",
		path:   fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &networksResp)
		},
	}
	err := instanceClone.do(ctx, nil, param)

	return networksResp, err
}
