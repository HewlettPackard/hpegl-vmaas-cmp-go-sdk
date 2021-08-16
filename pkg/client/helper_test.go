// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	consts "github.com/hpe-hcss/vmaas-cmp-go-sdk/pkg/common"
)

const mockHost = "mockhost"

func getDefaultHeaders() map[string]string {
	return map[string]string{
		"Accept": consts.ContentType,
	}
}
