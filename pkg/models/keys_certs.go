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

// ListAllKeyPairMeta
type ListAllKeyPairMeta struct {
	Offset float64 `json:"offset"`
	Max    float64 `json:"max"`
	Size   float64 `json:"size"`
	Total  float64 `json:"total"`
}
