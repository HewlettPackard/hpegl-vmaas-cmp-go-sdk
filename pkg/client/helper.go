//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import "net/url"

func getUrlValues(query map[string]string) url.Values {
	m := make(map[string][]string)
	for k, v := range query {
		m[k] = []string{v}
	}
	return m
}
