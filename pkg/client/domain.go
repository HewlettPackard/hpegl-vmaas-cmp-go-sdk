//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type DomainAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (d *DomainAPIService) GetAllDomains(
	ctx context.Context,
	param map[string]string,
) (models.GetAllDomains, error) {
	var domainResp models.GetAllDomains
	networkAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%s",
			d.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath, consts.DomainPath),
		client: d.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &domainResp)
		},
	}
	err := networkAPI.do(ctx, nil, param)

	return domainResp, err
}

func (d *DomainAPIService) GetSpecificDomain(
	ctx context.Context,
	domainID int,
) (models.GetSpecificDomain, error) {
	var domainResp models.GetSpecificDomain
	networkAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%s/%d",
			d.Cfg.Host, consts.VmaasCmpAPIBasePath, consts.NetworksPath, consts.DomainPath, domainID),
		client: d.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &domainResp)
		},
	}
	err := networkAPI.do(ctx, nil, nil)

	return domainResp, err
}
