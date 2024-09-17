// (C) Copyright 2024 Hewlett Packard Enterprise Development LP

package models

// Broker structs go here

// SubscriptionDetailsResponse is the response for Subscription Details from the broker
type SubscriptionDetailsResponse struct {
	ServiceInstanceID string `json:"ServiceInstanceID"`
	URL               string `json:"URL"`
}

// MorpheusTokenResponse is the response for Morpheus Token from the broker
type MorpheusTokenResponse struct {
	AccessToken        string `json:"access_token"`
	AccessTokenExpires int    `json:"expires"`
}

// MorpheusDetails is what we return to terraform
type MorpheusDetails struct {
	ID                 string `json:"id"` // This is the ServiceInstanceID, added here for use by the provider
	AccessToken        string `json:"access_token"`
	AccessTokenExpires int    `json:"access_token_expires"` // This is the Unix timestamp of when the token expires
	URL                string `json:"URL"`
}
