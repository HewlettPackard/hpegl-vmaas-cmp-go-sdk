[![CircleCI](https://circleci.com/gh/hpe-hcss/vmaas-cmp-go-sdk.svg?style=svg&circle-token=18e112d5ef6c20a7dc516fc320b8cd2e329af629)](https://circleci.com/gh/hpe-hcss/vmaas-cmp-go-sdk)

# vmaas-cmp-go-sdk

This package provides the official [Go](https://golang.org/) library for the [CMP API](https://docs.greenlake.hpe.com/docs/greenlake/services/private-cloud/internal/openapi/private-cloud-cmp-latest/overview/).

This is being developed in conjunction with the [VMaaS Terraform Provider](https://github.com/hpe-hcss/vmaas-terraform-resources).

## Sample Usage
```go
package main

import (
	"fmt"
	"github.com/HewlettPackard/vmaas-cmp-go-sdk/pkg/client"
	"golang.org/x/net/context"
	"os"
)


func main(){
	ctx := context.Background()
	headers := map[string]string{
		"location":      "<location_name>",
		"space":         "<space_name>",
		"Authorization": "<iam_token>",
	}
	config := client.Configuration{
		Host: "https://vmaas-cmp.intg.hpedevops.net",
		DefaultHeader: headers,
	}
	apiClient := client.NewAPIClient(&config)
	groupsClient := client.GroupsApiService{
		 apiClient,
		 config,
	}
	resp, err := groupsClient.GetASpecificGroup(ctx, 1)
	if err != nil {
		fmt.Printf("Get Group Error %v", err)
		os.Exit(1)
	}
	fmt.Printf("Group Name: %v | Group Location: %v", resp.Group.Name, resp.Group.Location)
}
```
