// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

import "time"

type GetAllNetworkRouter struct {
	NetworkRouters []GetNetworkRouter `json:"networkRouters"`
}

type GetSpecificRouterResp struct {
	NetworkRouter GetNetworkRouter `json:"networkRouter"`
}

type CreateRouterResp struct {
	Success bool `json:"success"`
	ID      int  `json:"id"`
}

type GetNetworkRouter struct {
	ID              int           `json:"id"`
	Code            string        `json:"code"`
	Name            string        `json:"name"`
	Description     interface{}   `json:"description"`
	Category        string        `json:"category"`
	DateCreated     time.Time     `json:"dateCreated"`
	LastUpdated     time.Time     `json:"lastUpdated"`
	RouterType      string        `json:"routerType"`
	Status          string        `json:"status"`
	Enabled         bool          `json:"enabled"`
	ExternalIP      interface{}   `json:"externalIp"`
	ExternalID      string        `json:"externalId"`
	ProviderID      string        `json:"providerId"`
	Type            IDModel       `json:"type"`
	NetworkServer   IDModel       `json:"networkServer"`
	Zone            IDModel       `json:"zone"`
	Instance        interface{}   `json:"instance"`
	ExternalNetwork interface{}   `json:"externalNetwork"`
	Site            interface{}   `json:"site"`
	Interfaces      []interface{} `json:"interfaces"`
}

type CreateRouterRequest struct {
	NetworkRouter CreateRouterRequestRouter `json:"networkRouter"`
}

type CreateRouterRequestRouter struct {
	Name            string                    `json:"name" tf:"name"`
	Type            IDModel                   `json:"type"`
	TypeID          int                       `json:"-" tf:"type_id,computed"`
	Enabled         bool                      `json:"enabled" tf:"enable"`
	Site            IDModel                   `json:"site"`
	GroupID         int                       `json:"-" tf:"group_id"`
	NetworkServer   IDModel                   `json:"networkServer"`
	NetworkServerID int                       `json:"-" tf:"network_server_id,computed"`
	Config          CreateRouterRequestConfig `json:"config"`
	EnableBgp       string                    `json:"enableBgp" tf:"enable_bgp"`

	// for tftags parsing
	Tier0Config CreateRouterTier0Config `json:"-" tf:"tier0_config,sub"`
	Tier1Config CreateRouterTier1Config `json:"-" tf:"tier1_config,sub"`
}

type CreateRouterRequestConfig struct {
	EdgeCluster string `json:"edgeCluster,omitempty" tf:"edge_cluster"`
	HaMode      string `json:"haMode,omitempty" tf:"ha_mode"`
	FailOver    string `json:"failOver,omitempty" tf:"fail_over"`
	CreateRouterTier0Config
}

type CreateRouterTier0Config struct {
	Bgp
	RouteRedistributionTier0
	RouteRedistributionTier1
	// For tftag parsing
	BGP     Bgp                      `json:"-" tf:"bgp,sub"`
	RRTier0 RouteRedistributionTier0 `json:"-" tf:"route_redistribution_tier0,sub"`
	RRTier1 RouteRedistributionTier1 `json:"-" tf:"route_redistribution_tier1,sub"`
}

type CreateRouterTier1Config struct {
	RouteAdvertisement *RouteAdvertisement `json:"-" tf:"route_advertisement,sub"`
}

type RouteRedistributionTier0 struct {
	TIER0STATIC            bool `json:"TIER0_STATIC,omitempty" tf:"tier0_static"`
	TIER0NAT               bool `json:"TIER0_NAT,omitempty" tf:"tier0_nat"`
	TIER0IPSECLOCALIP      bool `json:"TIER0_IPSEC_LOCAL_IP,omitempty" tf:"tier0_ipsec_local_ip"`
	TIER0DNSFORWARDERIP    bool `json:"TIER0_DNS_FORWARDER_IP,omitempty" tf:"tier0_dns_forwarder_ip"`
	TIER0SERVICEINTERFACE  bool `json:"TIER0_SERVICE_INTERFACE,omitempty" tf:"tier0_service_interface"`
	TIER0EXTERNALINTERFACE bool `json:"TIER0_EXTERNAL_INTERFACE,omitempty" tf:"tier0_external_interface"`
	TIER0LOOPBACKINTERFACE bool `json:"TIER0_LOOPBACK_INTERFACE,omitempty" tf:"tier0_loopback_interface"`
	TIER0SEGMENT           bool `json:"TIER0_SEGMENT,omitempty" tf:"tier0_segment"`
}

type RouteAdvertisement struct {
	Tier1Connected          bool   `json:"TIER1_CONNECTED,omitempty" tf:"tier1_connected"`
	Tier1StaticRoutes       bool   `json:"TIER1_STATIC_ROUTES,omitempty" tf:"tier1_static_routes"`
	Tier0Gateway            string `json:"tier0Gateway,omitempty" tf:"tier0gateway"`
	TIER1DNSFORWARDERIP     bool   `json:"TIER1_DNS_FORWARDER_IP,omitempty" tf:"tier1_dns_forwarder_ip"`
	TIER1STATIC             bool   `json:"TIER1_STATIC,omitempty" tf:"tier1_static"`
	TIER1LBVIP              bool   `json:"TIER1_LB_VIP,omitempty" tf:"tier1_lb_vip"`
	TIER1NAT                bool   `json:"TIER1_NAT,omitempty" tf:"tier1_nat"`
	TIER1LBSNAT             bool   `json:"TIER1_LB_SNAT,omitempty" tf:"tier1_lb_snat"`
	TIER1IPSECLOCALENDPOINT bool   `json:"TIER1_IPSEC_LOCAL_ENDPOINT,omitempty" tf:"tier1_ipsec_local_endpoint"`
}

type RouteRedistributionTier1 struct {
	TIER1SERVICEINTERFACE bool `json:"TIER1_SERVICE_INTERFACE,omitempty" tf:"tier1_service_interface"`
	TIER1SEGMENT          bool `json:"TIER1_SEGMENT,omitempty" tf:"tier1_segment"`
	RouteAdvertisement
}

type Bgp struct {
	LOCALASNUM     bool   `json:"LOCAL_AS_NUM,omitempty" tf:"local_as_num"`
	ECMP           bool   `json:"ECMP,omitempty" tf:"ecmp"`
	MULTIPATHRELAX bool   `json:"MULTIPATH_RELAX,omitempty" tf:"multipath_relax"`
	INTERSRIBGP    bool   `json:"INTER_SR_IBGP,omitempty" tf:"inter_sr_ibgp"`
	RESTARTMODE    string `json:"RESTART_MODE,omitempty" tf:"restart_mode"`
	RESTARTTIME    int    `json:"RESTART_TIME,omitempty" tf:"restart_time"`
	STALEROUTETIME int    `json:"STALE_ROUTE_TIME,omitempty" tf:"stale_route_time"`
}

type GetNetworlRouterTypes struct {
	NetworkRouterTypes []NetworkRouterTypes `json:"networkRouterTypes"`
}

type NetworkRouterTypes struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type GetNetworkServicesResp struct {
	NetworkServices []GetNetworkServices `json:"networkServices"`
}

type GetNetworkServices struct {
	ServiceType               string      `json:"serviceType"`
	ServiceTypeName           string      `json:"serviceTypeName"`
	Type                      string      `json:"type"`
	TypeName                  string      `json:"typeName"`
	Name                      string      `json:"name"`
	ID                        int         `json:"id"`
	IntegrationID             int         `json:"integrationId"`
	CanEdit                   bool        `json:"canEdit"`
	CanDelete                 bool        `json:"canDelete"`
	Status                    string      `json:"status"`
	LastUpdated               time.Time   `json:"lastUpdated"`
	BrandingImageName         interface{} `json:"brandingImageName"`
	SupportsTenantPermissions bool        `json:"supportsTenantPermissions"`
}
