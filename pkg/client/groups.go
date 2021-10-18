// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	"context"
	"encoding/json"
	"fmt"

	consts "github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/common"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
)

type GroupsAPIService struct {
	Client APIClientHandler
	Cfg    Configuration
}

/*
GroupsAPIService
Get a Specific Group
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID
 * @param groupID The Group ID

*/
func (a *GroupsAPIService) GetASpecificGroup(ctx context.Context,
	groupID int) (models.GroupResp, error) {
	specificGrpResp := models.GroupResp{}

	specificGrpRespAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s/%d", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.GroupsPath, groupID),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &specificGrpResp)
		},
		validations: []validationFunc{
			func() error {
				if groupID < 1 {
					return fmt.Errorf("%s", "group id should be greater than or equal to 1")
				}

				return nil
			},
		},
	}
	err := specificGrpRespAPI.do(ctx, nil, nil)

	return specificGrpResp, err
}

/*
GroupsAPIService
This endpoint retrieves all groups and a list of zones associated with the group by id.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc.
 	Passed from http.Request or context.Background().
 * @param serviceInstanceID

*/
func (a *GroupsAPIService) GetAllGroups(ctx context.Context,
	queryParams map[string]string) (models.Groups, error) {
	allGrpResp := models.Groups{}

	allGrpRespAPI := &api{
		method: "GET",
		path: fmt.Sprintf("%s/%s/%s", a.Cfg.Host, consts.VmaasCmpAPIBasePath,
			consts.GroupsPath),
		client: a.Client,

		jsonParser: func(body []byte) error {
			return json.Unmarshal(body, &allGrpResp)
		},
	}
	err := allGrpRespAPI.do(ctx, nil, queryParams)

	return allGrpResp, err
}
