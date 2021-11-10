// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type CmpStatus struct {
	Client APIClientHandler
	Cfg    Configuration
}

func (a *CmpStatus) GetSetupCheck(ctx context.Context) (models.CmpSetupCheck, error) {
	checkResp := models.CmpSetupCheck{}

	allCloudDSAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.SetupCheckPath),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &checkResp)
		},
	}
	err := allCloudDSAPI.do(ctx, nil, nil)

	return checkResp, err
}
