// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

type Environments []struct {
	ID          int         `json:"id"`
	Visibility  string      `json:"visibility"`
	Sortorder   int         `json:"sortOrder"`
	Code        string      `json:"code"`
	Name        string      `json:"name"`
	Active      bool        `json:"active"`
	Account     interface{} `json:"account"`
	Version     interface{} `json:"version"`
	Description string      `json:"description"`
}

type EnvironmentsZones struct {
	ID             int         `json:"id"`
	Accountid      int         `json:"accountId"`
	Groups         []int       `json:"groups"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	Location       string      `json:"location"`
	Visibility     string      `json:"visibility"`
	Zonetypeid     int         `json:"zoneTypeId"`
	Networkserver  interface{} `json:"networkServer"`
	Securityserver interface{} `json:"securityServer"`
}
type EnvironmentsGroups struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	EnvironmentsZones []EnvironmentsZones `json:"zones"`
	Integrations      []interface{}       `json:"integrations"`
	Networkserver     interface{}         `json:"networkServer"`
	Securityserver    interface{}         `json:"securityServer"`
}
type GetAllEnvironment struct {
	Environments Environments         `json:"environments"`
	Groups       []EnvironmentsGroups `json:"groups"`
}
