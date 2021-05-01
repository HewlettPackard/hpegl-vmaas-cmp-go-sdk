// (C) Copyright 2021 Hewlett Packard Enterprise Development LP
// CMP API types and Client methods for Network Domains
package cmp

import (
    "fmt"
)

var (
	NetworkDomainsPath = "/api/networks/domains"
)

// NetworkDomain structures for use in request and response payloads

type NetworkDomain struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Fqdn string `json:"fqdn"`
	Description string `json:"description"`
	Active bool `json:"active"`
	Visibility string `json:"visibility"`
	PublicZone bool `json:"publicZone"`
	DomainController bool `json:"domainController"`
	DomainUsername string `json:"domainUsername"`
	DomainPassword string `json:"domainPassword"`
	DcServer string `json:"dcServer"`
	OuPath string `json:"ouPath"`
	Account *TenantAbbrev `json:"account"`
	Owner *TenantAbbrev `json:"owner"`
	RefSource string `json:"refSource"`
	RefType string `json:"refType"`
	RefId int64 `json:"refId"`
}

type ListNetworkDomainsResult struct {
    NetworkDomains *[]NetworkDomain `json:"networkDomains"`
    Meta *MetaResult `json:"meta"`
}

type GetNetworkDomainResult struct {
    NetworkDomain *NetworkDomain `json:"networkDomain"`
}

type CreateNetworkDomainResult struct {
	Success bool `json:"success"`
	Message string `json:"msg"`
	Errors map[string]string `json:"errors"`
	NetworkDomain *NetworkDomain `json:"networkDomain"`
}

type UpdateNetworkDomainResult struct {
	CreateNetworkDomainResult
}

type DeleteNetworkDomainResult struct {
	DeleteResult
}

type NetworkDomainPayload struct {
	Name string `json:"name"`
	Fqdn string `json:"fqdn"`
	Description string `json:"description"`
	Active bool `json:"active"`
	Visibility string `json:"visibility"`
	PublicZone bool `json:"publicZone"`
	DomainController bool `json:"domainController"`
	DomainUsername string `json:"domainUsername"`
	DomainPassword string `json:"domainPassword"`
	DcServer string `json:"dcServer"`
	OuPath string `json:"ouPath"`
	Account *TenantAbbrev `json:"account"`
	Owner *TenantAbbrev `json:"owner"`
}

type CreateNetworkDomainPayload struct {
	NetworkDomainPayload *NetworkDomainPayload `json:"networkDomain"`
}

type UpdateNetworkDomainPayload struct {
	CreateNetworkDomainPayload
}

// Request types

// type ListNetworkDomainsRequest struct {
// 	Request
// }

// type GetNetworkDomainRequest struct {
// 	Request
// 	ID int64
// }

// type CreateNetworkDomainRequest struct {
// 	Request
// 	Payload *CreateNetworkDomainPayload
// }

// type UpdateNetworkDomainRequest struct {
// 	Request
// 	ID int64
// 	Payload *UpdateNetworkDomainPayload
// }

// type DeleteNetworkDomainRequest struct {
// 	Request
// 	ID int64
// }

// // Response types

// type ListNetworkDomainsResponse struct {
// 	Response
// 	Result *ListNetworkDomainsResult
// }

// type GetNetworkDomainResponse struct {
// 	Response
// 	Result *GetNetworkDomainResult
// }

// type CreateNetworkDomainResponse struct {
// 	Response
// 	Result *CreateNetworkDomainResult
// }

// type UpdateNetworkDomainResponse struct {
// 	Response
// 	Result *UpdateNetworkDomainResult
// }

// type DeleteNetworkDomainResponse struct {
// 	Request
// 	Result *DeleteNetworkDomainResult
// }

// Client request methods

func (client * Client) ListNetworkDomains(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path: NetworkDomainsPath,
		QueryParams: req.QueryParams,
		Result: &ListNetworkDomainsResult{},
	})
}

func (client * Client) GetNetworkDomain(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path: fmt.Sprintf("%s/%d", NetworkDomainsPath, id),
		QueryParams: req.QueryParams,
		Result: &GetNetworkDomainResult{},
	})
}

func (client * Client) CreateNetworkDomain(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "POST",
		Path: NetworkDomainsPath,
		QueryParams: req.QueryParams,
		Body: req.Body,
		Result: &CreateNetworkDomainResult{},
	})
}

func (client * Client) UpdateNetworkDomain(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "PUT",
		Path: fmt.Sprintf("%s/%d", NetworkDomainsPath, id),
		QueryParams: req.QueryParams,
		Body: req.Body,
		Result: &UpdateNetworkDomainResult{},
	})
}


func (client * Client) DeleteNetworkDomain(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "DELETE",
		Path: fmt.Sprintf("%s/%d", NetworkDomainsPath, id),
		QueryParams: req.QueryParams,
		Body: req.Body,
		Result: &DeleteNetworkDomainResult{},
	})
}

// helper functions

func (client * Client) FindNetworkDomainByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListNetworkDomains(&Request{
		QueryParams:map[string]string{
			"name": name,
      	},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListNetworkDomainsResult)
	networkDomainsCount := len(*listResult.NetworkDomains)
	if networkDomainsCount != 1 {
		return resp, fmt.Errorf("Found %d Network Domains for %v", networkDomainsCount, name)
	}
	firstRecord := (*listResult.NetworkDomains)[0]
	networkDomainId := firstRecord.ID
	return client.GetNetworkDomain(networkDomainId, &Request{})
}
