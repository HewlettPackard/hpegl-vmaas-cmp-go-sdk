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

type LoadBalancerAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (lb *LoadBalancerAPIService) CreateLoadBalancer(
	ctx context.Context,
	request models.CreateLoadBalancerRequest,
) (models.CreateNetworkLoadBalancerResp, error) {
	loadBalancerResp := models.CreateNetworkLoadBalancerResp{}
	loadBalancerAPI := &api{
		//compatibleVersion: lbCompatibleVersion,
		method: "POST",
		path:   fmt.Sprintf("%s/%s", consts.NetworkLoadBalancerPath, consts.LoadBalancerPath),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &loadBalancerResp)
		},
	}
	err := loadBalancerAPI.do(ctx, request, nil)

	return loadBalancerResp, err
}

func (lb *LoadBalancerAPIService) DeleteLoadBalancer(
	ctx context.Context,
	lbID int,
) (models.SuccessOrErrorMessage, error) {
	loadBalancerResp := models.SuccessOrErrorMessage{}
	loadBalancerAPI := &api{
		//	compatibleVersion: lbCompatibleVersion,
		method: "DELETE",
		path: fmt.Sprintf("%s/%s/%d",
			consts.NetworkLoadBalancerPath, consts.LoadBalancerPath, lbID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &loadBalancerResp)
		},
	}
	fmt.Println("1111111", loadBalancerAPI)
	err := loadBalancerAPI.do(ctx, nil, nil)

	return loadBalancerResp, err
}

func (lb *LoadBalancerAPIService) GetLoadBalancers(
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

func (lb *LoadBalancerAPIService) GetSpecificLoadBalancers(
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

func (lb *LoadBalancerAPIService) CreateLBMonitor(
	ctx context.Context,
	request models.CreateLBMonitor,
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

func (lb *LoadBalancerAPIService) DeleteLBMonitor(
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

func (lb *LoadBalancerAPIService) GetLBMonitors(
	ctx context.Context,
	lbID int,
) (models.GetLBMonitors, error) {
	LBMonitorResp := models.GetLBMonitors{}
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

func (lb *LoadBalancerAPIService) GetSpecificLBMonitor(
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

func (lb *LoadBalancerAPIService) CreateLBProfile(
	ctx context.Context,
	request models.CreateLBProfile,
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

func (lb *LoadBalancerAPIService) DeleteLBProfile(
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

func (lb *LoadBalancerAPIService) GetLBProfiles(
	ctx context.Context,
	lbID int,
) (models.GetLBProfile, error) {
	LBProfileResp := models.GetLBProfile{}
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

func (lb *LoadBalancerAPIService) GetSpecificLBProfile(
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

func (lb *LoadBalancerAPIService) CreateLBPool(
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

func (lb *LoadBalancerAPIService) DeleteLBPool(
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

func (lb *LoadBalancerAPIService) GetLBPools(
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

func (lb *LoadBalancerAPIService) GetSpecificLBPool(
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

func (lb *LoadBalancerAPIService) DeleteLBVirtualServers(
	ctx context.Context,
	lbID int,
	lbVirtualServerID int,
) (models.SuccessOrErrorMessage, error) {
	LBVSResp := models.SuccessOrErrorMessage{}
	LBVSAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "DELETE",
		path: fmt.Sprintf("%s/%s/%d/%s/%d",
			consts.NetworkLoadBalancerPath, consts.LoadBalancerPath, lbID,
			consts.LoadBalancerVirtualServersPath, lbVirtualServerID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBVSResp)
		},
	}
	err := LBVSAPI.do(ctx, nil, nil)

	return LBVSResp, err
}

func (lb *LoadBalancerAPIService) GetLBVirtualServers(
	ctx context.Context,
	lbID int,
) (models.GetLBVirtualServers, error) {
	LBVSResp := models.GetLBVirtualServers{}
	LBVSAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerVirtualServersPath),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBVSResp)
		},
	}
	err := LBVSAPI.do(ctx, nil, nil)

	return LBVSResp, err
}

func (lb *LoadBalancerAPIService) GetSpecificLBVirtualServer(
	ctx context.Context,
	lbID int,
	lbVSID int,
) (models.GetSpecificLBVirtualServers, error) {
	LBVSResp := models.GetSpecificLBVirtualServers{}
	LBVSAPI := &api{
		compatibleVersion: lbCompatibleVersion,
		method:            "GET",
		path: fmt.Sprintf("%s/%s/%d/%s/%d", consts.NetworkLoadBalancerPath,
			consts.LoadBalancerPath, lbID, consts.LoadBalancerVirtualServersPath, lbVSID),
		client: lb.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &LBVSResp)
		},
	}
	err := LBVSAPI.do(ctx, nil, nil)

	return LBVSResp, err
}
