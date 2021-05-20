# vmaas-cmp-go-sdk

This package provides the official [Go](https://golang.org/) library for the [CMP API](https://docs.greenlake.hpe.com/docs/greenlake/services/private-cloud/internal/openapi/private-cloud-cmp-latest/overview/).

This is being developed in conjunction with the [VMaaS Terraform Provider](https://github.com/hpe-hcss/vmaas-terraform-resources).

## Sample Usage
```
package main

import (
	"fmt"
	"github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/client"
	"golang.org/x/net/context"
	"os"
)


func main(){
	ctx := context.Background()
	config := client.Configuration{
		Host: "https://vmaas-cmp.intg.hpedevops.net",
	}
	apiClient := client.NewAPIClient(&config)
	groupsClient := client.GroupsApiService{
		 apiClient,
		 config,
	}
	resp, err := groupsClient.GetASpecificGroup(ctx, "0123456789abcdef", 1)
	if err != nil {
		fmt.Printf("List Group Error %v", err)
		os.Exit(1)
	}
	fmt.Printf("Group Name: %v | Group Location: %v", resp.Group.Name, resp.Group.Location)
}
```