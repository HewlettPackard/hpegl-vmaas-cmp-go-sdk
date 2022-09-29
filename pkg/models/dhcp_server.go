package models

import (
	"encoding/json"
	"time"
)

type CreateNetworkDhcpServerRequest struct {
	NetworkDhcpServer CreateNetworkDhcpServer `json:"networkDhcpServer"`
}

type CreateNetworkDhcpServer struct {
	ID              int         `json:"-" tf:"id,computed"`
	ServerID        int         `json:"-" tf:"server_id"`
	Name            string      `json:"name" tf:"name"`
	LeaseTime       int         `json:"leaseTime" tf:"lease_time"`
	ServerIPAddress string      `json:"serverIpAddress" tf:"server_address"`
	Config          *DhcpConfig `json:"config" tf:"config,sub"`
}

type DhcpConfig struct {
	EdgeCluster        string `json:"edgeCluster"  tf:"edge_cluster"`
	PreferredEdgeNode1 string `json:"preferredEdgeNode1"  tf:"active_edge_node"`
	PreferredEdgeNode2 string `json:"preferredEdgeNode2"  tf:"standby_edge_node"`
}

type CreateNetworkDhcpServerResp struct {
	Success bool `json:"success"`
	ID      int  `json:"id" tf:"id,computed"`
}

type GetNetworkDhcpServers struct {
	GetNetworkDhcpServerResp []GetNetworkDhcpServerResp `json:"networkDhcpServers"`
}

type GetNetworkDhcpServerResp struct {
	ID              int           `json:"id" tf:"id,computed"`
	DateCreated     time.Time     `json:"dateCreated"`
	ProviderID      string        `json:"providerId"`
	LastUpdated     time.Time     `json:"lastUpdated"`
	LeaseTime       int           `json:"leaseTime"`
	Name            string        `json:"name"`
	ExternalID      string        `json:"externalId"`
	ServerIPAddress string        `json:"serverIpAddress"`
	NetworkServerID NetworkServer `json:"networkServer"`
	Meta            MetaData      `json:"meta"`
}

type NetworkServer struct {
	ID int `json:"id"`
}

type MetaData struct {
	Max    json.Number `json:"max"`
	Offset json.Number `json:"offset"`
	Size   json.Number `json:"size"`
	Total  json.Number `json:"total"`
}

type GetSpecificNetworkDhcpServer struct {
	GetSpecificNetworkDhcpServerResp GetSpecificNetworkDhcpServerResp `json:"networkDhcpServer"`
}

type GetSpecificNetworkDhcpServerResp struct {
	ID              int           `json:"id" tf:"id,computed"`
	DateCreated     time.Time     `json:"dateCreated"`
	ProviderID      string        `json:"providerId"`
	LastUpdated     time.Time     `json:"lastUpdated"`
	LeaseTime       int           `json:"leaseTime"`
	Name            string        `json:"name"`
	ExternalID      string        `json:"externalId"`
	ServerIPAddress string        `json:"serverIpAddress"`
	NetworkServerID NetworkServer `json:"networkServer"`
}
