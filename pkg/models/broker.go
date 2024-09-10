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
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// MorpheusDetails is what we return to terraform
type MorpheusDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	URL          string `json:"URL"`
}
