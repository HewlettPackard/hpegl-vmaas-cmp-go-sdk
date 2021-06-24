// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

const mockHost = "mockhost"

func getDefaultHeaders() map[string]string {
	return map[string]string{
		"Accept": "application/json",
	}
}
