// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

// UpdatePolicyBodyPolicyConfig
// Supported config types: ['maxVms', 'maxMemory', 'maxCores', 'maxStorage', 'maxHosts',
// 'serverNaming', 'delayedRemoval', 'hostNaming', 'naming', 'maxNetworks', 'powerSchedule',
// 'provisionApproval', 'maxRouters', 'shutdown', 'tags', 'createUser','createUserGroup', 'workflow']
type UpdatePolicyBodyPolicyConfig struct {
	MaxVms float64 `json:"maxVms"`
}

// UpdatePolicyBodyPolicy
type UpdatePolicyBodyPolicy struct {
	Name        string                        `json:"name,omitempty"`
	Description string                        `json:"description,omitempty"`
	Config      *UpdatePolicyBodyPolicyConfig `json:"config,omitempty"`
	Enabled     bool                          `json:"enabled,omitempty"`
}

// UpdatePolicyBody
type UpdatePolicyBody struct {
	Policy *UpdatePolicyBodyPolicy `json:"policy"`
}
