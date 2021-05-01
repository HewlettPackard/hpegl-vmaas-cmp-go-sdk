// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package cmp_test

import (
	"github.com/hpe-hcss/vmaas-cmp-go-sdk"
	"testing"
)

var (
	testNetworkName = "golangtest-network"
)

func TestListNetworks(t *testing.T) {
	client := getTestClient(t)
	req := &cmp.Request{}
	resp, err := client.ListNetworks(req)
	assertResponse(t, resp, err)
}

func TestGetNetwork(t *testing.T) {
	client := getTestClient(t)
	req := &cmp.Request{}
	resp, err := client.ListNetworks(req)
	assertResponse(t, resp, err)

	// parse JSON and fetch the first one by ID

	result := resp.Result.(*cmp.ListNetworksResult)
	recordCount := result.Meta.Total
	t.Logf("Found %d Network Domains.", recordCount)
	if recordCount != 0 {
		// Get by ID
		record := (*result.Networks)[0]
		resp, err = client.GetNetwork(record.ID, &cmp.Request{})
		assertResponse(t, resp, err)

		// List by name

	} else {
		
	}
}

func _Busted_TestNetworksCRUD(t *testing.T) {
	client := getTestClient(t)
	//create
	// this has no uniqueness check on name, it probably should..
	req := &cmp.Request{
		Body: map[string]interface{}{
			"network": map[string]interface{}{
				"name": testNetworkName,
				"description": "a test network",
				"zone": map[string]interface{}{
					"id": 1,
				},
				// what else? varies by type...heh
			},
		},
	}
	resp, err := client.CreateNetwork(req)
	result := resp.Result.(*cmp.CreateNetworkResult)
	assertResponse(t, resp, err)
	assertNotNil(t, result)
	assertEqual(t, result.Success, true)
	assertNotNil(t, result.Network)
	assertNotEqual(t, result.Network.ID, 0)
	assertEqual(t, result.Network.Name, testNetworkName)

	// update
	updateReq := &cmp.Request{
		Body: map[string]interface{}{
			"network": map[string]interface{}{
				"description": "my new description",
			},
		},
	}
	updateResp, updateErr := client.UpdateNetwork(result.Network.ID, updateReq)
	updateResult := updateResp.Result.(*cmp.UpdateNetworkResult)
	assertResponse(t, updateResp, updateErr)
	assertNotNil(t, updateResult)
	assertEqual(t, updateResult.Success, true)
	assertNotNil(t, updateResult.Network)
	assertNotEqual(t, updateResult.Network.ID, 0)
	assertEqual(t, updateResult.Network.Description, "my new description")
	
	// delete
	deleteReq := &cmp.Request{}
	deleteResp, deleteErr := client.DeleteNetwork(result.Network.ID, deleteReq)
	deleteResult := deleteResp.Result.(*cmp.DeleteNetworkResult)
	assertResponse(t, deleteResp, deleteErr)
	assertNotNil(t, deleteResult)
	assertEqual(t, deleteResult.Success, true)

}