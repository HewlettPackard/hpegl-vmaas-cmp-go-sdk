// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type RouterAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (r *RouterAPIService) GetAllRouter(
	ctx context.Context,
	queryParams map[string]string,
) (models.GetAllNetworkRouter, error) {
	routerResp := models.GetAllNetworkRouter{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworksPath, consts.NetworkRouterPath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, nil, queryParams)

	return routerResp, err
}

func (r *RouterAPIService) GetSpecificRouter(
	ctx context.Context,
	routerID int,
) (models.GetNetworkRouter, error) {
	routerResp := models.GetNetworkRouter{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%s/%d", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworksPath, consts.NetworkRouterPath, routerID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, nil, nil)

	return routerResp, err
}

func (r *RouterAPIService) CreateRouter(
	ctx context.Context,
	request models.CreateRouterRequest,
) (models.CreateRouterResp, error) {
	routerResp := models.CreateRouterResp{}
	serverAPI := &api{
		method: "POST",
		path: fmt.Sprintf("%s/%s/%s/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworksPath, consts.NetworkRouterPath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, request, nil)

	return routerResp, err
}

func (r *RouterAPIService) UpdateRouter(
	ctx context.Context,
	routerID int,
	request models.CreateRouterRequest,
) (models.SuccessOrErrorMessage, error) {
	routerResp := models.SuccessOrErrorMessage{}
	serverAPI := &api{
		method: "PUT",
		path: fmt.Sprintf("%s/%s/%s/%s/%d", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworksPath, consts.NetworkRouterPath, routerID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, request, nil)

	return routerResp, err
}

func (r *RouterAPIService) DeleteRouter(
	ctx context.Context,
	routerID int,
) (models.SuccessOrErrorMessage, error) {
	routerResp := models.SuccessOrErrorMessage{}
	serverAPI := &api{
		method: "DELETE",
		path: fmt.Sprintf("%s/%s/%s/%s/%d", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworksPath, consts.NetworkRouterPath, routerID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, nil, nil)

	return routerResp, err
}

func (r *RouterAPIService) GetRouterTypes(
	ctx context.Context,
	queryParams map[string]string,
) (models.GetNetworlRouterTypes, error) {
	routerResp := models.GetNetworlRouterTypes{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.NetworkRouterTypePath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, nil, queryParams)

	return routerResp, err
}

func (r *RouterAPIService) GetNetworkServices(
	ctx context.Context,
	queryParams map[string]string,
) (models.GetNetworkServicesResp, error) {
	routerResp := models.GetNetworkServicesResp{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath,
			consts.NetworkServicePath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := serverAPI.do(ctx, nil, queryParams)

	return routerResp, err
}

func (r *RouterAPIService) CreateRouterNat(
	ctx context.Context,
	routerID int,
	request models.CreateRouterNatRequest,
) (models.CreateRouterNatResponse, error) {
	natResp := models.CreateRouterNatResponse{}
	serverAPI := &api{
		method: "POST",
		path: fmt.Sprintf("%s/%s/%s/%s/%d/%s", r.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersNatPath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &natResp)
		},
	}
	err := serverAPI.do(ctx, request, nil)

	return natResp, err
}
