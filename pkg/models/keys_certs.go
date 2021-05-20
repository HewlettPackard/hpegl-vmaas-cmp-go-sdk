// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

// CreateKeyPair
type CreateKeyPair struct {
	KeyPair *CreateKeyPairKeyPair `json:"keyPair,omitempty"`
}

// CreateKeyPairKeyPair
type CreateKeyPairKeyPair struct {
	Name      string `json:"name,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	// privateKey Optional for most cases
	PrivateKey string `json:"privateKey,omitempty"`
}

// CreateKeyPairResponseBody
type CreateKeyPairResponseBody struct {
	KeyPair *CreateKeyPairResponseBodyKeyPair `json:"keyPair"`
	Success bool                              `json:"success"`
}

// CreateKeyPairResponseBodyKeyPair
type CreateKeyPairResponseBodyKeyPair struct {
	Id            float64 `json:"id"`
	AccountId     float64 `json:"accountId"`
	HasPrivateKey bool    `json:"hasPrivateKey"`
	DateCreated   string  `json:"dateCreated"`
	LastUpdated   string  `json:"lastUpdated"`
}

// ListAllKeyPair
type ListAllKeyPair struct {
	KeyPairs []ListAllKeyPairKeyPairs `json:"keyPairs"`
	Meta     *ListAllKeyPairMeta      `json:"meta"`
}

// ListAllKeyPairKeyPairs
type ListAllKeyPairKeyPairs struct {
	Id            float64 `json:"id,omitempty"`
	AccountId     float64 `json:"accountId,omitempty"`
	HasPrivateKey bool    `json:"hasPrivateKey,omitempty"`
	DateCreated   string  `json:"dateCreated,omitempty"`
	LastUpdated   string  `json:"lastUpdated,omitempty"`
}

// ListAllKeyPairMeta
type ListAllKeyPairMeta struct {
	Offset float64 `json:"offset"`
	Max    float64 `json:"max"`
	Size   float64 `json:"size"`
	Total  float64 `json:"total"`
}
