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
	Type          IDModel                   `json:"type"`
	Enabled       bool                      `json:"enabled"`
	Site          IDModel                   `json:"site"`
	NetworkServer IDModel                   `json:"networkServer"`
	Name          string                    `json:"name"`
	Config        CreateRouterRequestConfig `json:"config"`
	EnableBgp     string                    `json:"enableBgp"`
}

type CreateRouterRequestConfig struct {
	EdgeCluster             string `json:"edgeCluster"`
	HaMode                  string `json:"haMode"`
	FailOver                string `json:"failOver"`
	TIER0STATIC             string `json:"TIER0_STATIC"`
	TIER0NAT                string `json:"TIER0_NAT"`
	TIER0IPSECLOCALIP       string `json:"TIER0_IPSEC_LOCAL_IP"`
	TIER0DNSFORWARDERIP     string `json:"TIER0_DNS_FORWARDER_IP"`
	TIER0SERVICEINTERFACE   string `json:"TIER0_SERVICE_INTERFACE"`
	TIER0EXTERNALINTERFACE  string `json:"TIER0_EXTERNAL_INTERFACE"`
	TIER0LOOPBACKINTERFACE  string `json:"TIER0_LOOPBACK_INTERFACE"`
	TIER0SEGMENT            string `json:"TIER0_SEGMENT"`
	TIER1DNSFORWARDERIP     string `json:"TIER1_DNS_FORWARDER_IP"`
	TIER1STATIC             string `json:"TIER1_STATIC"`
	TIER1LBVIP              string `json:"TIER1_LB_VIP"`
	TIER1NAT                string `json:"TIER1_NAT"`
	TIER1LBSNAT             string `json:"TIER1_LB_SNAT"`
	TIER1IPSECLOCALENDPOINT string `json:"TIER1_IPSEC_LOCAL_ENDPOINT"`
	TIER1SERVICEINTERFACE   string `json:"TIER1_SERVICE_INTERFACE"`
	TIER1SEGMENT            string `json:"TIER1_SEGMENT"`
	LOCALASNUM              string `json:"LOCAL_AS_NUM"`
	ECMP                    string `json:"ECMP"`
	MULTIPATHRELAX          string `json:"MULTIPATH_RELAX"`
	INTERSRIBGP             string `json:"INTER_SR_IBGP"`
	RESTARTMODE             string `json:"RESTART_MODE"`
	RESTARTTIME             int    `json:"RESTART_TIME"`
	STALEROUTETIME          int    `json:"STALE_ROUTE_TIME"`
}
