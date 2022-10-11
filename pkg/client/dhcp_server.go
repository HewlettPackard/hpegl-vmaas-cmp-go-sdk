// (C) Copyright 2022 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

const dhcpCompatibleVersion = "5.2.13"

type DhcpServerAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (lb *DhcpServerAPIService) CreateDhcpServer(
	ctx context.Context,
	serverID int,
	request models.CreateNetworkDhcpServerRequest,
) (models.CreateNetworkDhcpServerResp, error) {
	dhcpServerResp := models.CreateNetworkDhcpServerResp{}
	dhcpServerAPI := &api{
		compatibleVersion: dhcpCompatibleVersion,
		method:            "POST",
		path: fmt.Sprintf("%s/%s/%d/%s",
			consts.NetworksPath, consts.ServerPath,
			serverID, consts.DhcpServerPath),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &dhcpServerResp)
		},
	}
	err := dhcpServerAPI.do(ctx, request, nil)

	return dhcpServerResp, err
}

func (lb *DhcpServerAPIService) UpdateDhcpServer(
	ctx context.Context,
	serverID int,
	dhcpID int,
	request models.CreateNetworkDhcpServerRequest,
) (models.CreateNetworkDhcpServerResp, error) {
	dhcpServerResp := models.CreateNetworkDhcpServerResp{}
	dhcpServerAPI := &api{
		compatibleVersion: dhcpCompatibleVersion,
		method:            "PUT",
		path: fmt.Sprintf("%s/%s/%d/%s/%d",
			consts.NetworksPath, consts.ServerPath,
			serverID, consts.DhcpServerPath, dhcpID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &dhcpServerResp)
		},
	}
	err := dhcpServerAPI.do(ctx, request, nil)

	return dhcpServerResp, err
}

func (lb *DhcpServerAPIService) DeleteDhcpServer(
	ctx context.Context,
	serverID int,
	dhcpID int,
) (models.SuccessOrErrorMessage, error) {
	dhcpServerResp := models.SuccessOrErrorMessage{}
	dhcpServerAPI := &api{
		compatibleVersion: dhcpCompatibleVersion,
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d/%s/%d",
			consts.NetworksPath, consts.ServerPath,
			serverID, consts.DhcpServerPath, dhcpID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &dhcpServerResp)
		},
	}
	err := dhcpServerAPI.do(ctx, nil, nil)

	return dhcpServerResp, err
}

func (lb *DhcpServerAPIService) GetDhcpServers(
	ctx context.Context,
	serverID int,
) (models.GetNetworkDhcpServers, error) {
	dhcpServerResp := models.GetNetworkDhcpServers{}
	dhcpServerAPI := &api{
		compatibleVersion: dhcpCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s",
			consts.NetworksPath, consts.ServerPath,
			serverID, consts.DhcpServerPath),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &dhcpServerResp)
		},
	}
	err := dhcpServerAPI.do(ctx, nil, nil)

	return dhcpServerResp, err
}

func (lb *DhcpServerAPIService) GetSpecificDhcpServer(
	ctx context.Context,
	serverID int,
	dhcpID int,
) (models.GetSpecificNetworkDhcpServer, error) {
	dhcpServerResp := models.GetSpecificNetworkDhcpServer{}
	dhcpServerAPI := &api{
		compatibleVersion: dhcpCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s/%d",
			consts.NetworksPath, consts.ServerPath,
			serverID, consts.DhcpServerPath, dhcpID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &dhcpServerResp)
		},
	}
	err := dhcpServerAPI.do(ctx, nil, nil)

	return dhcpServerResp, err
}
