// (C) Copyright 2021 Hewlett Packard Enterprise Development LP
package cmp_test

import (
	"testing"
	"github.com/hpe-hcss/vmaas-cmp-go-sdk"
)

func TestListInstances(t *testing.T) {
	client := getTestClient(t)
	resp, err := client.ListInstances(&cmp.Request{})
	assertResponse(t, resp, err)
}

func TestGetInstance(t *testing.T) {
	client := getTestClient(t)
	resp, err := client.ListInstances(&cmp.Request{})
	assertResponse(t, resp, err)
	// parse JSON and fetch the first one by ID
	listInstancesResult := resp.Result.(*cmp.ListInstancesResult)
	instancesCount := listInstancesResult.Meta.Total
	t.Logf("Found %d Instances.", instancesCount)
	// if instancesCount != 0 {
		firstInstance := (*listInstancesResult.Instances)[0]	
		// log.Printf(fmt.Sprintf("First Instance: [%d] %v: ", firstInstance.ID, firstInstance.Name))
		resp, err = client.GetInstance(firstInstance.ID, &cmp.Request{})
		assertResponse(t, resp, err)
	// }
	
}


// this requires params zoneId&layoutId&siteId, heh
// func TestListInstancePlans(t *testing.T) {
// 	client := getTestClient(t)
// 	resp, err := client.ListInstancePlans(&cmp.Request{})
// 	assertResponse(t, resp, err)
// }

