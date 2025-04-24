// (C) Copyright 2021-2025 Hewlett Packard Enterprise Development LP
package models

type GetAllProvisioningTypes struct {
	ProvisionTypes []ProvisionTypes `json:"provisionTypes"`
}

type ProvisionTypes struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	Code             string         `json:"code"`
	StorageTypes     []StorageTypes `json:"storageTypes"`
	RootStorageTypes []StorageTypes `json:"rootStorageTypes"`
}

type StorageTypes struct {
	ID   int    `json:"id"`   // 133
	Code string `json:"code"` // vmware-thin
	Name string `json:"name"` // Thin
}
