// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package cmp_test

import (
	"github.com/hpe-hcss/vmaas-cmp-go-sdk"
	"testing"
)

func TestWhoami(t *testing.T) {
	client := getTestClient(t)
	resp, err := client.Whoami()
	assertResponse(t, resp, err)
	result := resp.Result.(*cmp.WhoamiResult)
	assertNotNil(t, result.User)
	assertNotNil(t, result.User.ID)
	assertEqual(t, result.User.Username, testUsername)
}
