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
) (models.GetSpecificRouterResp, error) {
	routerResp := models.GetSpecificRouterResp{}
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
) (models.GetSpecificRouterResp, error) {
	routerResp := models.GetSpecificRouterResp{}
	serverAPI := &api{
		method: "GET",
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
