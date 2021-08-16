// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/common"
	models "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/models"
)

type ServersAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
ServersApiService
Get All Servers
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param name/phrase optional
@return models.Servers
*/
func (a *ServersAPIService) GetAllServers(
	ctx context.Context,
	queryParams map[string]string,
) (models.Servers, error) {
	serversResponse := models.Servers{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.ServerPath),
		client: a.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &serversResponse)
		},
	}
	err := serverAPI.do(ctx, nil, queryParams)

	return serversResponse, err
}

func (a *ServersAPIService) GetSpecificServer(
	ctx context.Context,
	serverID int) (models.GetSpecificServerResponse, error) {
	serversResponse := models.GetSpecificServerResponse{}
	serverAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.ServerPath, serverID),
		client: a.Client,
		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &serversResponse)
		},
		validations: []validationFunc{
			func() error {
				if serverID < 1 {
					return fmt.Errorf("%s", "server id should be greater than or equal to 1")
				}

				return nil
			},
		},
	}
	err := serverAPI.do(ctx, nil, nil)

	return serversResponse, err
}
