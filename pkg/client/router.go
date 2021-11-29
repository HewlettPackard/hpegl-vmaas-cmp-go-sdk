// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

const routerCompatibleVersion = "5.2.10"

type RouterAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (r *RouterAPIService) GetAllRouter(
	ctx context.Context,
	queryParams map[string]string,
) (models.GetAllNetworkRouter, error) {
	routerResp := models.GetAllNetworkRouter{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "GET",
		path:              fmt.Sprintf("%s/%s", consts.NetworksPath, consts.NetworkRouterPath),
		client:            r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := routerAPI.do(ctx, nil, queryParams)

	return routerResp, err
}

func (r *RouterAPIService) GetSpecificRouter(
	ctx context.Context,
	routerID int,
) (models.GetSpecificRouterResp, error) {
	routerResp := models.GetSpecificRouterResp{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d",
			consts.NetworksPath, consts.NetworkRouterPath, routerID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := routerAPI.do(ctx, nil, nil)

	return routerResp, err
}

func (r *RouterAPIService) CreateRouter(
	ctx context.Context,
	request models.CreateRouterRequest,
) (models.CreateRouterResp, error) {
	routerResp := models.CreateRouterResp{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "POST",
		path:              fmt.Sprintf("%s/%s", consts.NetworksPath, consts.NetworkRouterPath),
		client:            r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := routerAPI.do(ctx, request, nil)

	return routerResp, err
}

func (r *RouterAPIService) UpdateRouter(
	ctx context.Context,
	routerID int,
	request models.CreateRouterRequest,
) (models.SuccessOrErrorMessage, error) {
	routerResp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "PUT",
		path: fmt.Sprintf("%s/%s/%d",
			consts.NetworksPath, consts.NetworkRouterPath, routerID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := routerAPI.do(ctx, request, nil)

	return routerResp, err
}

func (r *RouterAPIService) DeleteRouter(
	ctx context.Context,
	routerID int,
) (models.SuccessOrErrorMessage, error) {
	routerResp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d",
			consts.NetworksPath, consts.NetworkRouterPath, routerID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := routerAPI.do(ctx, nil, nil)

	return routerResp, err
}

func (r *RouterAPIService) GetRouterTypes(
	ctx context.Context,
	queryParams map[string]string,
) (models.GetNetworlRouterTypes, error) {
	routerResp := models.GetNetworlRouterTypes{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "GET",
		path:              consts.NetworkRouterTypePath,
		client:            r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := routerAPI.do(ctx, nil, queryParams)

	return routerResp, err
}

func (r *RouterAPIService) GetNetworkServices(
	ctx context.Context,
	queryParams map[string]string,
) (models.GetNetworkServicesResp, error) {
	routerResp := models.GetNetworkServicesResp{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s", consts.NetworksPath,
			consts.NetworkServicePath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &routerResp)
		},
	}
	err := routerAPI.do(ctx, nil, queryParams)

	return routerResp, err
}

func (r *RouterAPIService) CreateRouterNat(
	ctx context.Context,
	routerID int,
	request models.CreateRouterNatRequest,
) (models.SuccessOrErrorMessage, error) {
	natResp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "POST",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersNatPath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &natResp)
		},
	}
	err := routerAPI.do(ctx, request, nil)

	return natResp, err
}

func (r *RouterAPIService) GetSpecificRouterNat(
	ctx context.Context,
	routerID, natID int,
) (models.GetSpecificRouterNatResponse, error) {
	natResp := models.GetSpecificRouterNatResponse{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersNatPath, natID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &natResp)
		},
	}
	err := routerAPI.do(ctx, nil, nil)

	return natResp, err
}

func (r *RouterAPIService) UpdateRouterNat(
	ctx context.Context,
	routerID, natID int,
	req models.CreateRouterNatRequest,
) (models.SuccessOrErrorMessage, error) {
	natResp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "PUT",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersNatPath, natID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &natResp)
		},
	}
	err := routerAPI.do(ctx, req, nil)

	return natResp, err
}

func (r *RouterAPIService) DeleteRouterNat(
	ctx context.Context,
	routerID, natID int,
) (models.SuccessOrErrorMessage, error) {
	natResp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersNatPath, natID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &natResp)
		},
	}
	err := routerAPI.do(ctx, nil, nil)

	return natResp, err
}

func (r *RouterAPIService) CreateRouterFirewallRuleGroup(
	ctx context.Context,
	routerID int,
	request models.CreateRouterFirewallRuleGroupRequest,
) (models.SuccessOrErrorMessage, error) {
	firewallGroupResp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "POST",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersFirewallRuleGroupPath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &firewallGroupResp)
		},
	}
	err := routerAPI.do(ctx, request, nil)

	return firewallGroupResp, err
}

func (r *RouterAPIService) GetSpecificRouterFirewallRuleGroup(
	ctx context.Context,
	routerID, firewallGroupID int,
) (models.GetSpecificRouterFirewallRuleGroupResponse, error) {
	firewallGroupResp := models.GetSpecificRouterFirewallRuleGroupResponse{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersFirewallRuleGroupPath, firewallGroupID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &firewallGroupResp)
		},
	}
	err := routerAPI.do(ctx, nil, nil)

	return firewallGroupResp, err
}

func (r *RouterAPIService) DeleteRouterFirewallRuleGroup(
	ctx context.Context,
	routerID, firewallGroupID int,
) (models.SuccessOrErrorMessage, error) {
	firewallGroupResp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: routerCompatibleVersion,
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RoutersFirewallRuleGroupPath, firewallGroupID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &firewallGroupResp)
		},
	}
	err := routerAPI.do(ctx, nil, nil)

	return firewallGroupResp, err
}

func (r *RouterAPIService) CreateRouterRoute(
	ctx context.Context,
	routerID int,
	req models.CreateRouterRoute,
) (models.SuccessOrErrorMessage, error) {
	resp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: "5.2.12",
		method:            "POST",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RouterRoutePath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := routerAPI.do(ctx, req, nil)

	return resp, err
}

func (r *RouterAPIService) GetSpecificRouterRoute(
	ctx context.Context,
	routerID, routeID int,
) (models.GetSpecificRouterRoute, error) {
	resp := models.GetSpecificRouterRoute{}
	routerAPI := &api{
		compatibleVersion: "5.2.12",
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RouterRoutePath, routeID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := routerAPI.do(ctx, nil, nil)

	return resp, err
}

func (r *RouterAPIService) DeleteRouterRoute(
	ctx context.Context,
	routerID, routeID int,
) (models.SuccessOrErrorMessage, error) {
	resp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: "5.2.12",
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RouterRoutePath, routeID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := routerAPI.do(ctx, nil, nil)

	return resp, err
}

func (r *RouterAPIService) CreateRouterBgpNeighbor(
	ctx context.Context,
	routerID int,
	req models.CreateNetworkRouterBgpNeighborRequest,
) (models.SuccessOrErrorMessage, error) {
	resp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: "5.2.12",
		method:            "POST",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RouterBgpNeighborPath),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := routerAPI.do(ctx, req, nil)

	return resp, err
}

func (r *RouterAPIService) GetSpecificRouterBgpNeighbor(
	ctx context.Context,
	routerID, bgpNeighborID int,
) (models.GetSpecificNetworkRouterBgpNeighbor, error) {
	resp := models.GetSpecificNetworkRouterBgpNeighbor{}
	routerAPI := &api{
		compatibleVersion: "5.2.12",
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RouterBgpNeighborPath, bgpNeighborID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := routerAPI.do(ctx, nil, nil)

	return resp, err
}

func (r *RouterAPIService) UpdateRouterBgpNeighbor(
	ctx context.Context,
	routerID, bgpNeighborID int,
	req models.CreateNetworkRouterBgpNeighborRequest,
) (models.SuccessOrErrorMessage, error) {
	resp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: "5.2.12",
		method:            "PUT",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RouterBgpNeighborPath, bgpNeighborID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := routerAPI.do(ctx, req, nil)

	return resp, err
}

func (r *RouterAPIService) DeleteRouterBgpNeighbor(
	ctx context.Context,
	routerID, bgpNeighborID int,
) (models.SuccessOrErrorMessage, error) {
	resp := models.SuccessOrErrorMessage{}
	routerAPI := &api{
		compatibleVersion: "5.2.12",
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworksPath,
			consts.NetworkRouterPath, routerID, consts.RouterBgpNeighborPath, bgpNeighborID),
		client: r.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &resp)
		},
	}
	err := routerAPI.do(ctx, nil, nil)

	return resp, err
}
