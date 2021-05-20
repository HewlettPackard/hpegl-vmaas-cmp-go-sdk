// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

// CommonRoleUpdateResp
type CommonRoleUpdateResp struct {
	Success bool   `json:"success,omitempty"`
	Access  string `json:"access,omitempty"`
}

// CreateRoleBodyRole
type CreateRoleBodyRole struct {
	Authority   string  `json:"authority"`
	Description string  `json:"description"`
	BaseRoleId  float64 `json:"baseRoleId,omitempty"`
	RoleType    string  `json:"roleType"`
}

// CreateRoleBody
type CreateRoleBody struct {
	Role *CreateRoleBodyRole `json:"role"`
}

// GetRoleResponseAppTemplatePermissions
type GetRoleResponseAppTemplatePermissions struct {
	Id     float64 `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Access string  `json:"access,omitempty"`
}

// GetRoleResponseFeaturePermissions
type GetRoleResponseFeaturePermissions struct {
	Id     float64 `json:"id,omitempty"`
	Code   string  `json:"code,omitempty"`
	Name   string  `json:"name,omitempty"`
	Access string  `json:"access,omitempty"`
}

// GetRoleResponseInstanceTypePermissions
type GetRoleResponseInstanceTypePermissions struct {
	Id     float64 `json:"id,omitempty"`
	Code   string  `json:"code,omitempty"`
	Name   string  `json:"name,omitempty"`
	Access string  `json:"access,omitempty"`
}

// GetRoleResponseRoleOwner
type GetRoleResponseRoleOwner struct {
	Id   float64 `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
}

// GetRoleResponseRole
type GetRoleResponseRole struct {
	Id          float64 `json:"id,omitempty"`
	Authority   string  `json:"authority,omitempty"`
	Description string  `json:"description,omitempty"`
	// Account
	Scope string `json:"scope,omitempty"`
	// user
	RoleType          string                    `json:"roleType,omitempty"`
	Multitenant       bool                      `json:"multitenant,omitempty"`
	MultitenantLocked bool                      `json:"multitenantLocked,omitempty"`
	ParentRoleId      float64                   `json:"parentRoleId,omitempty"`
	Diverged          bool                      `json:"diverged,omitempty"`
	OwnerId           float64                   `json:"ownerId,omitempty"`
	Owner             *GetRoleResponseRoleOwner `json:"owner,omitempty"`
	DateCreated       string                    `json:"dateCreated,omitempty"`
	LastUpdated       string                    `json:"lastUpdated,omitempty"`
}

// GetRoleResponseSites
type GetRoleResponseSites struct {
	Id     float64 `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Access string  `json:"access,omitempty"`
}

// GetRoleResponseZones
type GetRoleResponseZones struct {
	Id     float64 `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Access string  `json:"access,omitempty"`
}

// GetRoleResponse
type GetRoleResponse struct {
	Role                        *GetRoleResponseRole                     `json:"role,omitempty"`
	FeaturePermissions          []GetRoleResponseFeaturePermissions      `json:"featurePermissions,omitempty"`
	GlobalSiteAccess            string                                   `json:"globalSiteAccess,omitempty"`
	Sites                       []GetRoleResponseSites                   `json:"sites,omitempty"`
	GlobalZoneAccess            string                                   `json:"globalZoneAccess,omitempty"`
	Zones                       []GetRoleResponseZones                   `json:"zones,omitempty"`
	GlobalInstanceTypeAccess    string                                   `json:"globalInstanceTypeAccess,omitempty"`
	InstanceTypePermissions     []GetRoleResponseInstanceTypePermissions `json:"instanceTypePermissions,omitempty"`
	GlobalAppTemplateAccess     string                                   `json:"globalAppTemplateAccess,omitempty"`
	AppTemplatePermissions      []GetRoleResponseAppTemplatePermissions  `json:"appTemplatePermissions,omitempty"`
	GlobalCatalogItemTypeAccess string                                   `json:"globalCatalogItemTypeAccess,omitempty"`
	CatalogItemTypePermissions  []GetRoleResponseAppTemplatePermissions  `json:"catalogItemTypePermissions,omitempty"`
	PersonaPermissions          []GetRoleResponseAppTemplatePermissions  `json:"personaPermissions,omitempty"`
}

// UpdateRoleBodyRole
type UpdateRoleBodyRole struct {
	Authority   string `json:"authority,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateRoleBody
type UpdateRoleBody struct {
	Role *UpdateRoleBodyRole `json:"role"`
}

// UpdateRoleFeaturePermission
type UpdateRoleFeaturePermission struct {
	PermissionCode string `json:"permissionCode"`
	Access         string `json:"access"`
}

// UpdateBlueprintAccessBody
type UpdateBlueprintAccessBody struct {
	AppTemplateId float64 `json:"appTemplateId"`
	Access        string  `json:"access"`
}

// UpdateInstancetypeAccessBody
type UpdateInstancetypeAccessBody struct {
	InstanceTypeId float64 `json:"instanceTypeId"`
	Access         string  `json:"access"`
}

// UpdateGroupAccessBody
type UpdateGroupAccessBody struct {
	GroupId float64 `json:"groupId"`
	Access  string  `json:"access"`
}

// UpdateCloudAccessBody
type UpdateCloudAccessBody struct {
	CloudId float64 `json:"cloudId"`
	Access  string  `json:"access"`
}
