//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

type IDNameModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListNetworksBody struct {
	Networks     []GetSpecificNetwork `json:"networks"`
	NetworkCount int                  `json:"networkCount"`
}

type CreateNetworkResponse struct {
	Success bool               `json:"success"`
	Errors  interface{}        `json:"errors"`
	Network GetSpecificNetwork `json:"network"`
}

type GetSpecificNetworkBody struct {
	Network GetSpecificNetwork `json:"network"`
}

type GetSpecificNetwork struct {
	ID                      int                  `json:"id" tf:"id,computed"`
	Name                    string               `json:"name" tf:"name"`
	Zone                    IDNameModel          `json:"zone"`
	DisplayName             string               `json:"displayName" tf:"display_name"`
	Type                    IDModel              `json:"type"`
	TypeID                  int                  `json:"-" tf:"type_id,computed"`
	Owner                   IDNameModel          `json:"owner"`
	Code                    string               `json:"code" tf:"code,computed"`
	Category                string               `json:"category"`
	ExternalID              string               `json:"externalId" tf:"external_id,computed"`
	InternalID              string               `json:"internalId" tf:"internal_id,computed"`
	UniqueID                string               `json:"uniqueId" tf:"unique_id,computed"`
	ExternalType            string               `json:"externalType"`
	RefType                 string               `json:"refType"`
	RefID                   int                  `json:"refId"`
	DhcpServer              bool                 `json:"dhcpServer"`
	Status                  string               `json:"status" tf:"status,computed"`
	Visibility              string               `json:"visibility"`
	EnableAdmin             bool                 `json:"enableAdmin"`
	ScanNetwork             bool                 `json:"scanNetwork" tf:"scan_network"`
	Active                  bool                 `json:"active" tf:"active"`
	DefaultNetwork          bool                 `json:"defaultNetwork"`
	AssignPublicIP          bool                 `json:"assignPublicIp"`
	ApplianceURLProxyBypass bool                 `json:"applianceUrlProxyBypass" tf:"appliance_url_proxy_bypass"`
	ZonePool                IDNameModel          `json:"zonePool"`
	AllowStaticOverride     bool                 `json:"allowStaticOverride"`
	Tenants                 []IDNameModel        `json:"tenants"`
	ResourcePermissions     NetworkResPermission `json:"resourcePermissions" tf:"resource_permissions,sub"`
}

type NetworkResPermission struct {
	All   bool                        `json:"all,omitempty" tf:"all"`
	Sites []NetworkResPermissionSites `json:"sites,omitempty" tf:"sites"`
}

type NetworkResPermissionSites struct {
	ID      int  `json:"id,omitempty" tf:"id"`
	Default bool `json:"default,omitempty" tf:"default"`
}

type CreateNetworkRequest struct {
	Network CreateNetwork `json:"network"`
}

type CreateNetwork struct {
	ID                  int                          `json:"-" tf:"id,computed"`
	Type                IDModel                      `json:"type,omitempty"`
	NetworkServer       IDModel                      `json:"networkServer,omitempty"`
	ResourcePermissions NetworkResPermission         `json:"resourcePermissions"`
	Name                string                       `json:"name" tf:"name"`
	Description         string                       `json:"description,omitempty"`
	DisplayName         string                       `json:"displayName,omitempty"`
	CloudID             int                          `json:"-" tf:"cloud_id"`
	GroupID             string                       `json:"-" tf:"group_id"`
	TypeID              int                          `json:"-" tf:"type_id"`
	PoolID              int                          `json:"pool,omitempty"`
	NetworkDomainID     int                          `json:"-" tf:"domain_id"`
	Site                *IDStringModel               `json:"site,omitempty"`
	NetworkDomain       *IDModel                     `json:"networkDomain,omitempty"`
	NetworkProxy        *IDModel                     `json:"networkProxy,omitempty"`
	NetworkPool         *PoolModel                   `json:"networkPool,omitempty"`
	NetworkProxyID      int                          `json:"-" tf:"proxy_id"`
	ProxyID             int                          `json:"-" tf:"proxy_id"`
	SearchDomains       string                       `json:"searchDomains,omitempty"`
	Cidr                string                       `json:"cidr,omitempty"`
	Gateway             string                       `json:"gateway,omitempty"`
	DNSPrimary          string                       `json:"dnsPrimary,omitempty"`
	DNSSecondary        string                       `json:"dnsSecondary,omitempty"`
	Active              bool                         `json:"active"`
	DhcpServer          bool                         `json:"dhcpServer"`
	ScanNetwork         bool                         `json:"scanNetwork"`
	AllowStaticOverride bool                         `json:"allowStaticOverride"`
	AppURLProxyBypass   bool                         `json:"applianceUrlProxyBypass,omitempty"`
	NoProxy             string                       `json:"noProxy,omitempty"`
	ScopeID             string                       `json:"scopeId,omitempty"`
	TfStaticNetwork     *CreateStaticNetwork         `json:"-" tf:"static_network,sub"`
	TfDhcpNetwork       *CreateDhcpNetwork           `json:"-" tf:"dhcp_network,sub"`
	TfDhcpConfig        *TfCreateDhcpNetworkConfig   `json:"-" tf:"config,sub"`
	TfStaticConfig      *TfCreateStaticNetworkConfig `json:"-" tf:"config,sub"`
	Config              *CreateNetworkConfig         `json:"config"`
}

type TfCreateStaticNetworkConfig struct {
	ConnectedGateway string `json:"" tf:"connected_gateway"`
	VlanIDs          string `json:"" tf:"vlan_ids"`
}

type TfCreateDhcpNetworkConfig struct {
	ConnectedGateway        string `json:"-" tf:"connected_gateway"`
	VlanIDs                 string `json:"-" tf:"vlan"`
	SubnetIPManagementType  string `json:"-" tf:"dhcp_type"`
	SubnetIPServerID        string `json:"-" tf:"dhcp_server"`
	SubnetDhcpServerAddress string `json:"-" tf:"dhcp_server_address"`
	DhcpRange               string `json:"-" tf:"dhcp_range"`
	SubnetDhcpLeaseTime     string `json:"-" tf:"dhcp_lease_time"`
}

type CreateNetworkConfig struct {
	ConnectedGateway        string `json:"connectedGateway,omitempty"`
	VlanIDs                 string `json:"vlanIDs,omitempty"`
	SubnetIPManagementType  string `json:"subnetIpManagementType"`
	SubnetIPServerID        string `json:"subnetIpServerId"`
	SubnetDhcpServerAddress string `json:"subnetDhcpServerAddress"`
	DhcpRange               string `json:"dhcpRange"`
	SubnetDhcpLeaseTime     string `json:"subnetDhcpLeaseTime"`
}

type CreateStaticNetwork struct {
	Name                string                `json:"-" tf:"name"`
	Description         string                `json:"-" tf:"description"`
	DisplayName         string                `json:"-" tf:"display_name"`
	CloudID             int                   `json:"-" tf:"cloud_id"`
	GroupID             string                `json:"-" tf:"group_id"`
	TypeID              int                   `json:"-" tf:"type_id"`
	PoolID              int                   `json:"-" tf:"pool_id"`
	NetworkDomainID     int                   `json:"-" tf:"domain_id"`
	NetworkProxyID      int                   `json:"-" tf:"proxy_id"`
	ProxyID             int                   `json:"-" tf:"proxy_id"`
	SearchDomains       string                `json:"-" tf:"search_domains"`
	Cidr                string                `json:"-" tf:"cidr"`
	Gateway             string                `json:"-" tf:"gateway"`
	DNSPrimary          string                `json:"-" tf:"primary_dns"`
	DNSSecondary        string                `json:"-" tf:"secondary_dns"`
	Active              bool                  `json:"-" tf:"active"`
	ScanNetwork         bool                  `json:"-" tf:"scan_network"`
	AllowStaticOverride bool                  `json:"-" tf:"allow_static_override"`
	AppURLProxyBypass   bool                  `json:"-" tf:"appliance_url_proxy_bypass"`
	NoProxy             string                `json:"-" tf:"no_proxy"`
	ScopeID             string                `json:"-" tf:"scode_id"`
	ResourcePermissions *NetworkResPermission `json:"-" tf:"resource_permissions,sub"`
}

type CreateDhcpNetwork struct {
	Name                string                `json:"-" tf:"name"`
	Description         string                `json:"-" tf:"description"`
	DisplayName         string                `json:"-" tf:"display_name"`
	Site                *IDStringModel        `json:"-" tf:"group"`
	NetworkDomain       *IDModel              `json:"-" tf:"domain"`
	NetworkProxy        *IDModel              `json:"-" tf:"network_proxy"`
	SearchDomains       string                `json:"-" tf:"search_domains"`
	Cidr                string                `json:"-" tf:"gateway_cidr"`
	Gateway             string                `json:"-" tf:"gateway"`
	DNSPrimary          string                `json:"-" tf:"primary_dns"`
	DNSSecondary        string                `json:"-" tf:"secondary_dns"`
	Active              bool                  `json:"-" tf:"active"`
	DhcpServer          bool                  `json:"-" tf:"dhcp_enabled"`
	ScanNetwork         bool                  `json:"-" tf:"scan_network"`
	AllowStaticOverride bool                  `json:"-" tf:"allow_ip_override"`
	AppURLProxyBypass   bool                  `json:"-" tf:"bypass_proxy_for_appliance_url"`
	NoProxy             string                `json:"-" tf:"no_proxy"`
	ScopeID             string                `json:"-" tf:"transport_zone"`
	ResourcePermissions *NetworkResPermission `json:"-" tf:"resource_permissions,sub"`
}

type PoolModel struct {
	Pool int `json:"pool,omitempty"`
}

type GetNetworkTypesResponse struct {
	NetworkTypes []GetSpecificNetworkType `json:"networkTypes"`
}

type GetaNetworkType struct {
	NetworkTypes GetSpecificNetworkType `json:"networkType"`
}

type GetSpecificNetworkType struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type GetNetworkPoolsResp struct {
	NetworkPools     []GetNetworkPool `json:"networkPools"`
	NetworkPoolCount int              `json:"networkPoolCount"`
	Meta             Meta             `json:"meta"`
}

type GetSpecificNetworkPool struct {
	NetworkPool GetNetworkPool `json:"networkPool"`
}

type GetNetworkPool struct {
	ID          int                      `json:"id"`
	Type        IDModel                  `json:"type"`
	Account     IDModel                  `json:"account"`
	Category    string                   `json:"category"`
	Code        string                   `json:"code"`
	Name        string                   `json:"name" tf:"name"`
	DisplayName string                   `json:"displayName" tf:"display_name,computed"`
	InternalID  interface{}              `json:"internalId"`
	ExternalID  string                   `json:"externalId"`
	PoolGroup   interface{}              `json:"poolGroup"`
	IPRanges    []GetNetworkPoolIPRanges `json:"ipRanges"`
}

type GetNetworkPoolIPRanges struct {
	ID           int    `json:"id"`
	StartAddress string `json:"startAddress"`
	EndAddress   string `json:"endAddress"`
}
