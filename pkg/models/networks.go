//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

type IDNameModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListNetworksBody struct {
	Networks     GetSpecificNetworkBody `json:"networks"`
	NetworkCount int                    `json:"networkCount"`
}

type GetSpecificNetworkBodyType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetSpecificNetworkBody struct {
	ID                      int                        `json:"id"`
	Name                    string                     `json:"name"`
	Zone                    IDNameModel                `json:"zone"`
	Type                    GetSpecificNetworkBodyType `json:"type"`
	Owner                   IDNameModel                `json:"owner"`
	Code                    string                     `json:"code"`
	Category                string                     `json:"category"`
	ExternalID              string                     `json:"externalId"`
	InternalID              string                     `json:"internalId"`
	UniqueID                string                     `json:"uniqueId"`
	ExternalType            string                     `json:"externalType"`
	RefType                 string                     `json:"refType"`
	RefID                   int                        `json:"refId"`
	DhcpServer              bool                       `json:"dhcpServer"`
	Status                  string                     `json:"status"`
	Visibility              string                     `json:"visibility"`
	EnableAdmin             bool                       `json:"enableAdmin"`
	ScanNetwork             bool                       `json:"scanNetwork"`
	Active                  bool                       `json:"active"`
	DefaultNetwork          bool                       `json:"defaultNetwork"`
	AssignPublicIP          bool                       `json:"assignPublicIp"`
	ApplianceURLProxyBypass bool                       `json:"applianceUrlProxyBypass"`
	ZonePool                IDNameModel                `json:"zonePool"`
	AllowStaticOverride     bool                       `json:"allowStaticOverride"`
	Tenants                 []IDNameModel              `json:"tenants"`
}
