// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

import "time"

type GetAllNetworkRouter struct {
	NetworkRouters []GetNetworkRouter `json:"networkRouters"`
}

type GetSpecificRouterResp struct {
	NetworkRouter GetNetworkRouter `json:"networkRouter"`
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
	Tier0Config *CreateRouterTier0Config `json:"-" tf:"tier0_config,sub"`
	Tier1Config *CreateRouterTier1Config `json:"-" tf:"tier1_config,sub"`
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
	TIER0STATIC            bool `json:"TIER0_STATIC,omitempty" tf:"TIER0_STATIC"`
	TIER0NAT               bool `json:"TIER0_NAT,omitempty" tf:"TIER0_NAT"`
	TIER0IPSECLOCALIP      bool `json:"TIER0_IPSEC_LOCAL_IP,omitempty" tf:"TIER0_IPSEC_LOCAL_IP"`
	TIER0DNSFORWARDERIP    bool `json:"TIER0_DNS_FORWARDER_IP,omitempty" tf:"TIER0_DNS_FORWARDER_IP"`
	TIER0SERVICEINTERFACE  bool `json:"TIER0_SERVICE_INTERFACE,omitempty" tf:"TIER0_SERVICE_INTERFACE"`
	TIER0EXTERNALINTERFACE bool `json:"TIER0_EXTERNAL_INTERFACE,omitempty" tf:"TIER0_EXTERNAL_INTERFACE"`
	TIER0LOOPBACKINTERFACE bool `json:"TIER0_LOOPBACK_INTERFACE,omitempty" tf:"TIER0_LOOPBACK_INTERFACE"`
	TIER0SEGMENT           bool `json:"TIER0_SEGMENT,omitempty" tf:"TIER0_SEGMENT"`
}

type RouteAdvertisement struct {
	Tier1Connected          bool   `json:"TIER1_CONNECTED,omitempty" tf:"TIER1_CONNECTED"`
	Tier1StaticRoutes       bool   `json:"TIER1_STATIC_ROUTES,omitempty" tf:"TIER1_STATIC_ROUTES"`
	Tier0Gateway            string `json:"tier0Gateway,omitempty" tf:"tier0Gateway"`
	TIER1DNSFORWARDERIP     bool   `json:"TIER1_DNS_FORWARDER_IP,omitempty" tf:"TIER1_DNS_FORWARDER_IP"`
	TIER1STATIC             bool   `json:"TIER1_STATIC,omitempty" tf:"TIER1_STATIC"`
	TIER1LBVIP              bool   `json:"TIER1_LB_VIP,omitempty" tf:"TIER1_LB_VIP"`
	TIER1NAT                bool   `json:"TIER1_NAT,omitempty" tf:"TIER1_NAT"`
	TIER1LBSNAT             bool   `json:"TIER1_LB_SNAT,omitempty" tf:"TIER1_LB_SNAT"`
	TIER1IPSECLOCALENDPOINT bool   `json:"TIER1_IPSEC_LOCAL_ENDPOINT,omitempty" tf:"TIER1_IPSEC_LOCAL_ENDPOINT"`
}

type RouteRedistributionTier1 struct {
	TIER1SERVICEINTERFACE bool `json:"TIER1_SERVICE_INTERFACE,omitempty" tf:"TIER1_SERVICE_INTERFACE"`
	TIER1SEGMENT          bool `json:"TIER1_SEGMENT,omitempty" tf:"TIER1_SEGMENT"`
	RouteAdvertisement
}

type Bgp struct {
	LOCALASNUM     bool   `json:"LOCAL_AS_NUM,omitempty" tf:"LOCAL_AS_NUM"`
	ECMP           bool   `json:"ECMP,omitempty" tf:"ECMP"`
	MULTIPATHRELAX bool   `json:"MULTIPATH_RELAX,omitempty" tf:"MULTIPATH_RELAX"`
	INTERSRIBGP    bool   `json:"INTER_SR_IBGP,omitempty" tf:"INTER_SR_IBGP"`
	RESTARTMODE    string `json:"RESTART_MODE,omitempty" tf:"RESTART_MODE"`
	RESTARTTIME    int    `json:"RESTART_TIME,omitempty" tf:"RESTART_TIME"`
	STALEROUTETIME int    `json:"STALE_ROUTE_TIME,omitempty" tf:"STALE_ROUTE_TIME"`
}
