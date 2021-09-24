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

func (n *NetworksAPIService) GetAllNetworks(
	ctx context.Context,
	param map[string]string,
) (models.ListNetworksBody, error) {
	var networksResp models.ListNetworksBody
	networkAPI := &api{
		method: "GET",
		path:   fmt.Sprintf("%s/%s/%s", n.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath),
		client: n.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &networksResp)
		},
	}
	err := networkAPI.do(ctx, nil, param)

	return networksResp, err
}

func (n *NetworksAPIService) GetSpecificNetwork(
	ctx context.Context,
	networkID int,
) (models.GetSpecificNetworkBody, error) {
	var networksResp models.GetSpecificNetworkBody
	networkAPI := &api{
		method: "GET",
		path:   fmt.Sprintf("%s/%s/%s/%d", n.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath, networkID),
		client: n.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &networksResp)
		},
	}
	err := networkAPI.do(ctx, nil, nil)

	return networksResp, err
}

func (n *NetworksAPIService) CreateNetwork(
	ctx context.Context,
	networkReq models.CreateNetworkRequest,
) (models.CreateNetworkResponse, error) {
	var networksResp models.CreateNetworkResponse
	networkAPI := &api{
		method: "POST",
		path:   fmt.Sprintf("%s/%s/%s", n.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath),
		client: n.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &networksResp)
		},
	}
	err := networkAPI.do(ctx, networkReq, nil)

	return networksResp, err
}

func (n *NetworksAPIService) DeleteNetwork(
	ctx context.Context,
	networkID int,
) (models.SuccessOrErrorMessage, error) {
	var output models.SuccessOrErrorMessage
	networkAPI := &api{
		method: "DELETE",
		path:   fmt.Sprintf("%s/%s/%s/%d", n.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath, networkID),
		client: n.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &output)
		},
	}
	err := networkAPI.do(ctx, nil, nil)

	return output, err
}

func (n *NetworksAPIService) GetNetworkType(
	ctx context.Context,
	params map[string]string,
) (models.GetNetworkTypesResponse, error) {
	var resp models.GetNetworkTypesResponse
	networkAPI := &api{
		method: "GET",
		path:   fmt.Sprintf("%s/%s/%s", n.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworkTypePath),
		client: n.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := networkAPI.do(ctx, nil, params)

	return resp, err
}
