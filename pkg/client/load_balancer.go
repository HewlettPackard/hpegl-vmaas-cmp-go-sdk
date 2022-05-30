// (C) Copyright 2022 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

const lbCompatibleVersion = "5.4.6"

type loadBalancerAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (lb *loadBalancerAPIService) CreateLoadBalancer(
	ctx context.Context,
	request models.CreateLoadBalancerRequest,
) (models.CreateNetworkLoadBalancerResp, error) {
	loadBalancerResp := models.CreateNetworkLoadBalancerResp{}
	loadBalancerAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "POST",
		path:              fmt.Sprintf("%s/%s", consts.NetworkLoadBalancerPath, consts.LoadBalancerPath),
		client:            lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &loadBalancerResp)
		},
	}
	err := loadBalancerAPI.do(ctx, request, nil)

	return loadBalancerResp, err
}

func (lb *loadBalancerAPIService) DeleteLoadBalancer(
	ctx context.Context,
	lbID int,
) (models.SuccessOrErrorMessage, error) {
	loadBalancerResp := models.SuccessOrErrorMessage{}
	loadBalancerAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d",
			consts.NetworkLoadBalancerPath, consts.LoadBalancerPath, lbID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &loadBalancerResp)
		},
	}
	err := loadBalancerAPI.do(ctx, nil, nil)

	return loadBalancerResp, err
}

func (lb *loadBalancerAPIService) GetLoadBalancers(
	ctx context.Context,
) (models.GetNetworkLoadBalancers, error) {
	loadBalancerResp := models.GetNetworkLoadBalancers{}
	loadBalancerAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "GET",
		path:              fmt.Sprintf("%s/%s", consts.NetworkLoadBalancerPath, consts.LoadBalancerPath),
		client:            lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &loadBalancerResp)
		},
	}
	err := loadBalancerAPI.do(ctx, nil, nil)

	return loadBalancerResp, err
}

func (lb *loadBalancerAPIService) GetSpecificLoadBalancers(
	ctx context.Context,
	lbID int,
) (models.GetSpecificNetworkLoadBalancer, error) {
	loadBalancerResp := models.GetSpecificNetworkLoadBalancer{}
	loadBalancerAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &loadBalancerResp)
		},
	}
	err := loadBalancerAPI.do(ctx, nil, nil)

	return loadBalancerResp, err
}

func (lb *loadBalancerAPIService) CreateLBMonitor(
	ctx context.Context,
	request models.CreateLBMonitorReq,
	lbID int,
) (models.CreateLBMonitorResp, error) {
	LBMonitorResp := models.CreateLBMonitorResp{}
	LBMonitorAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "POST",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerMonitorPath),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBMonitorResp)
		},
	}
	err := LBMonitorAPI.do(ctx, request, nil)

	return LBMonitorResp, err
}

func (lb *loadBalancerAPIService) DeleteLBMonitor(
	ctx context.Context,
	lbID int,
	lbMonitorID int,
) (models.SuccessOrErrorMessage, error) {
	LBMonitorResp := models.SuccessOrErrorMessage{}
	LBMonitorAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d/%s/%d",
			consts.NetworkLoadBalancerPath, consts.LoadBalancerPath, lbID,
			consts.LoadBalancerMonitorPath, lbMonitorID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBMonitorResp)
		},
	}
	err := LBMonitorAPI.do(ctx, nil, nil)

	return LBMonitorResp, err
}

func (lb *loadBalancerAPIService) GetLBMonitors(
	ctx context.Context,
	lbID int,
) (models.GetLBMonitorsResp, error) {
	LBMonitorResp := models.GetLBMonitorsResp{}
	LBMonitorAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerMonitorPath),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBMonitorResp)
		},
	}
	err := LBMonitorAPI.do(ctx, nil, nil)

	return LBMonitorResp, err
}

func (lb *loadBalancerAPIService) GetSpecificLBMonitor(
	ctx context.Context,
	lbID int,
	lbmonitorID int,
) (models.GetSpecificNetworkLoadBalancer, error) {
	LBMonitorResp := models.GetSpecificNetworkLoadBalancer{}
	LBMonitorAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerMonitorPath, lbmonitorID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBMonitorResp)
		},
	}
	err := LBMonitorAPI.do(ctx, nil, nil)

	return LBMonitorResp, err
}

func (lb *loadBalancerAPIService) CreateLBProfile(
	ctx context.Context,
	request models.CreateLBProfileReq,
	lbID int,
) (models.CreateLBProfileResp, error) {
	LBProfileResp := models.CreateLBProfileResp{}
	LBProfileAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "POST",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerProfilePath),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBProfileResp)
		},
	}
	err := LBProfileAPI.do(ctx, request, nil)

	return LBProfileResp, err
}

func (lb *loadBalancerAPIService) DeleteLBProfile(
	ctx context.Context,
	lbID int,
	lbProfileID int,
) (models.SuccessOrErrorMessage, error) {
	LBProfileResp := models.SuccessOrErrorMessage{}
	LBProfileAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d/%s/%d",
			consts.NetworkLoadBalancerPath, consts.LoadBalancerPath, lbID,
			consts.LoadBalancerProfilePath, lbProfileID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBProfileResp)
		},
	}
	err := LBProfileAPI.do(ctx, nil, nil)

	return LBProfileResp, err
}

func (lb *loadBalancerAPIService) GetLBProfiles(
	ctx context.Context,
	lbID int,
) (models.GetLBProfilesResp, error) {
	LBProfileResp := models.GetLBProfilesResp{}
	LBProfileAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerProfilePath),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBProfileResp)
		},
	}
	err := LBProfileAPI.do(ctx, nil, nil)

	return LBProfileResp, err
}

func (lb *loadBalancerAPIService) GetSpecificLBProfile(
	ctx context.Context,
	lbID int,
	lbProfileID int,
) (models.GetLBSpecificProfile, error) {
	LBProfileResp := models.GetLBSpecificProfile{}
	LBProfileAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerProfilePath, lbProfileID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBProfileResp)
		},
	}
	err := LBProfileAPI.do(ctx, nil, nil)

	return LBProfileResp, err
}

func (lb *loadBalancerAPIService) CreateLBPool(
	ctx context.Context,
	request models.CreateLBPool,
	lbID int,
) (models.CreateLBPoolResp, error) {
	LBPoolResp := models.CreateLBPoolResp{}
	LBPoolAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "POST",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerPoolPath),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBPoolResp)
		},
	}
	err := LBPoolAPI.do(ctx, request, nil)

	return LBPoolResp, err
}

func (lb *loadBalancerAPIService) DeleteLBPool(
	ctx context.Context,
	lbID int,
	lbPoolID int,
) (models.SuccessOrErrorMessage, error) {
	LBPoolResp := models.SuccessOrErrorMessage{}
	LBPoolAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d/%s/%d",
			consts.NetworkLoadBalancerPath, consts.LoadBalancerPath, lbID,
			consts.LoadBalancerPoolPath, lbPoolID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBPoolResp)
		},
	}
	err := LBPoolAPI.do(ctx, nil, nil)

	return LBPoolResp, err
}

func (lb *loadBalancerAPIService) GetLBPools(
	ctx context.Context,
	lbID int,
) (models.GetLBPools, error) {
	LBPoolResp := models.GetLBPools{}
	LBPoolAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerPoolPath),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBPoolResp)
		},
	}
	err := LBPoolAPI.do(ctx, nil, nil)

	return LBPoolResp, err
}

func (lb *loadBalancerAPIService) GetSpecificLBPool(
	ctx context.Context,
	lbID int,
	lbPoolID int,
) (models.GetSpecificLBPool, error) {
	LBPoolResp := models.GetSpecificLBPool{}
	LBPoolAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerPoolPath, lbPoolID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBPoolResp)
		},
	}
	err := LBPoolAPI.do(ctx, nil, nil)

	return LBPoolResp, err
}
