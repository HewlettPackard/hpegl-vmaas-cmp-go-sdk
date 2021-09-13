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

func (a *NetworksAPIService) GetAllNetworks(
	ctx context.Context,
	param map[string]string,
) (models.ListNetworksBody, error) {
	var networksResp models.ListNetworksBody
	networkAPI := &api{
		method: "GET",
		path:   fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &networksResp)
		},
	}
	err := networkAPI.do(ctx, nil, param)

	return networksResp, err
}

func (a *NetworksAPIService) GetSpecificNetwork(
	ctx context.Context,
	networkID int,
) (models.GetSpecificNetworkBody, error) {
	var networksResp models.GetSpecificNetworkBody
	networkAPI := &api{
		method: "GET",
		path:   fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath, networkID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &networksResp)
		},
	}
	err := networkAPI.do(ctx, nil, nil)

	return networksResp, err
}

func (a *NetworksAPIService) CreateNetwork(
	ctx context.Context,
	networkReq models.CreateNetworkRequest,
) (models.GetSpecificNetworkBody, error) {
	var networksResp models.GetSpecificNetworkBody
	networkAPI := &api{
		method: "POST",
		path:   fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &networksResp)
		},
	}
	err := networkAPI.do(ctx, networkReq, nil)

	return networksResp, err
}

func (a *NetworksAPIService) DeleteNetwork(
	ctx context.Context,
) (models.SuccessOrErrorMessage, error) {
	var output models.SuccessOrErrorMessage
	networkAPI := &api{
		method: "DELETE",
		path:   fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &output)
		},
	}
	err := networkAPI.do(ctx, nil, nil)

	return output, err
}
