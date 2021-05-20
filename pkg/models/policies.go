// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

// CreatePolicyBodyPolicyConfig
type CreatePolicyBodyPolicyConfig struct {
	MaxVms float64 `json:"maxVms"`
}

// CreatePolicyBodyPolicyPolicyType
type CreatePolicyBodyPolicyPolicyType struct {
	Code string `json:"code"`
}

// CreatePolicyBodyPolicy
type CreatePolicyBodyPolicy struct {
	Name        string                            `json:"name,omitempty"`
	Description string                            `json:"description,omitempty"`
	PolicyType  *CreatePolicyBodyPolicyPolicyType `json:"policyType,omitempty"`
	Config      *CreatePolicyBodyPolicyConfig     `json:"config,omitempty"`
	Enabled     bool                              `json:"enabled,omitempty"`
	RefId       float64                           `json:"refId,omitempty"`
	RefType     string                            `json:"refType,omitempty"`
}

// CreatePolicyBody
type CreatePolicyBody struct {
	Policy *CreatePolicyBodyPolicy `json:"policy"`
}

// UpdatePolicyBodyPolicyConfig
// Supported config types: ['maxVms', 'maxMemory', 'maxCores', 'maxStorage', 'maxHosts', 'serverNaming', 'delayedRemoval', 'hostNaming', 'naming', 'maxNetworks', 'powerSchedule', 'provisionApproval', 'maxRouters', 'shutdown', 'tags', 'createUser', 'createUserGroup', 'workflow']
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
