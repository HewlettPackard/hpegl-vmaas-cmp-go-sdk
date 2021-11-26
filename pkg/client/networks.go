//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

const networkCompatibleVersion = "5.2.10"

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
		path:   consts.NetworksPath,
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
		path:   fmt.Sprintf("%s/%d", consts.NetworksPath, networkID),
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
		compatibleVersion: "5.2.13",
		method:            "POST",
		path:              consts.NetworksPath,
		client:            n.Client,

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
		compatibleVersion: networkCompatibleVersion,
		method:            "DELETE",
		path:              fmt.Sprintf("%s/%d", consts.NetworksPath, networkID),
		client:            n.Client,

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
		compatibleVersion: networkCompatibleVersion,
		method:            "GET",
		path:              consts.NetworkTypePath,
		client:            n.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := networkAPI.do(ctx, nil, params)

	return resp, err
}

func (n *NetworksAPIService) GetNetworkPool(
	ctx context.Context,
	params map[string]string,
) (models.GetNetworkPoolsResp, error) {
	var resp models.GetNetworkPoolsResp
	networkAPI := &api{
		method: "GET",
		path:   fmt.Sprintf("%s/%s", consts.NetworksPath, consts.NetworkPoolPath),
		client: n.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := networkAPI.do(ctx, nil, params)

	return resp, err
}

func (n *NetworksAPIService) GetSpecificNetworkPool(
	ctx context.Context,
	networkPoolID int,
) (models.GetSpecificNetworkPool, error) {
	var resp models.GetSpecificNetworkPool
	networkAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%d",
			consts.NetworksPath, consts.NetworkPoolPath, networkPoolID),
		client: n.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := networkAPI.do(ctx, nil, nil)

	return resp, err
}

func (n *NetworksAPIService) UpdateNetwork(
	ctx context.Context,
	networkID int,
	request models.CreateNetworkRequest,
) (models.SuccessOrErrorMessage, error) {
	var output models.SuccessOrErrorMessage
	networkAPI := &api{
		compatibleVersion: "5.2.13",
		method:            "PUT",
		path:              fmt.Sprintf("%s/%d", consts.NetworksPath, networkID),
		client:            n.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &output)
		},
	}
	err := networkAPI.do(ctx, request, nil)

	return output, err
}

func (n *NetworksAPIService) GetNetworkProxy(
	ctx context.Context,
	params map[string]string,
) (models.GetAllNetworkProxies, error) {
	var proxyResp models.GetAllNetworkProxies
	networkAPI := &api{
		method: "GET",
		path:   fmt.Sprintf("%s/%s", consts.NetworksPath, consts.NetworkProxyPath),
		client: n.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &proxyResp)
		},
	}
	err := networkAPI.do(ctx, nil, params)

	return proxyResp, err
}
