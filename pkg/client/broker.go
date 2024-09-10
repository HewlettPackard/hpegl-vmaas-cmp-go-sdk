// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

// BrokerAPIService is a service that provides methods to interact with the broker API
type BrokerAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

// GetMorpheusDetails returns Morpheus details to terraform
func (a *BrokerAPIService) GetMorpheusDetails(ctx context.Context) (models.MorpheusDetails, error) {
	// Get the service instance ID and Morpheus URL
	ServiceSubscriptionDetailsResp := models.SubscriptionDetailsResponse{}
	serviceSubscriptionDetailsAPI := &api{
		method:                 "GET",
		path:                   consts.SubscriptionDetails,
		client:                 a.Client,
		removeVmaasCMPBasePath: true,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &ServiceSubscriptionDetailsResp)
		},
	}

	// Use the default query params
	if err := serviceSubscriptionDetailsAPI.do(ctx, nil, a.Cfg.DefaultQueryParams); err != nil {
		return models.MorpheusDetails{}, fmt.Errorf("error getting service subscription details: %v", err)
	}

	// Get the Morpheus token
	MorpheusTokenResp := models.MorpheusTokenResponse{}
	log.Printf(consts.MorpheusToken, ServiceSubscriptionDetailsResp.ServiceInstanceID)
	morpheusTokenAPI := &api{
		method:                 "GET",
		path:                   fmt.Sprintf(consts.MorpheusToken, ServiceSubscriptionDetailsResp.ServiceInstanceID),
		client:                 a.Client,
		removeVmaasCMPBasePath: true,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &MorpheusTokenResp)
		},
	}

	// No query params needed
	if err := morpheusTokenAPI.do(ctx, nil, nil); err != nil {
		return models.MorpheusDetails{}, fmt.Errorf("error getting Morpheus token: %v", err)
	}

	// build response
	ret := models.MorpheusDetails{
		AccessToken:  MorpheusTokenResp.AccessToken,
		RefreshToken: MorpheusTokenResp.RefreshToken,
		URL:          ServiceSubscriptionDetailsResp.URL,
	}

	return ret, nil
}
