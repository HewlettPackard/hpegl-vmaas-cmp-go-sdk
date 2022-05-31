// (C) Copyright 2022 Hewlett Packard Enterprise Development LP

package models

import "encoding/json"

// Create LB Req
type CreateLoadBalancerRequest struct {
	NetworkLoadBalancer CreateNetworkLoadBalancerRequest `json:"loadBalancer"`
}

type CreateNetworkLoadBalancerRequest struct {
	Name                string                    `json:"name"`
	Type                string                    `json:"type"`
	Description         string                    `json:"description"`
	NetworkServerID     int                       `json:"networkServerId"`
	Enabled             bool                      `json:"enabled"`
	Visibility          string                    `json:"visibility"`
	Config              CreateConfig              `json:"config"`
	ResourcePermissions EnableResourcePermissions `json:"resourcePermission"`
}

type CreateConfig struct {
	AdminState bool   `json:"adminState"`
	Loglevel   string `json:"loglevel"`
	Size       string `json:"size"`
	Tier1      string `json:"tier1"`
}

type EnableResourcePermissions struct {
	All bool `json:"all"`
}

// Create LB resp
type CreateNetworkLoadBalancerResp struct {
	Success                 bool                    `json:"success"`
	NetworkLoadBalancerResp NetworkLoadBalancerResp `json:"loadBalancer"`
}

type NetworkLoadBalancerResp struct {
	ID          int          `json:"id" tf:"id,computed"`
	Name        string       `json:"name"`
	AccountId   int          `json:"accountId"`
	Cloud       CloudInfo    `json:"cloud"`
	Type        Types        `json:"type"`
	Description string       `json:"description"`
	Port        string       `json:"port"`
	SSLEnabled  bool         `json:"sslEnabled"`
	Enabled     bool         `json:"enabled"`
	Visibility  string       `json:"visibility"`
	Config      CreateConfig `json:"config"`
	DateCreated string       `json:"dateCreated"`
	LastUpdated string       `json:"lastUpdated"`
}

type CloudInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Types struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// Get LBs
type GetNetworkLoadBalancers struct {
	GetNetworkLoadBalancerResp []GetNetworkLoadBalancerResp `json:"loadBalancers"`
}
type GetNetworkLoadBalancerResp struct {
	ID          int          `json:"id" tf:"id,computed"`
	Name        string       `json:"name"`
	AccountId   int          `json:"accountId"`
	Cloud       CloudInfo    `json:"cloud"`
	Type        Types        `json:"type"`
	Description string       `json:"description"`
	Port        string       `json:"port"`
	Host        string       `json:"host"`
	IP          string       `json:"ip"`
	SSLEnabled  bool         `json:"sslEnabled"`
	Enabled     bool         `json:"enabled"`
	Visibility  string       `json:"visibility"`
	Config      CreateConfig `json:"config"`
	DateCreated string       `json:"dateCreated"`
	LastUpdated string       `json:"lastUpdated"`
	Meta        MetaInfo     `json:"meta"`
}

type MetaInfo struct {
	Max    json.Number `json:"max"`
	Offset json.Number `json:"offset"`
	Size   json.Number `json:"size"`
	Total  json.Number `json:"total"`
}

// Get Specific LB
type GetSpecificNetworkLoadBalancer struct {
	GetSpecificNetworkLoadBalancerResp GetSpecificNetworkLoadBalancerResp `json:"loadBalancer"`
}

type GetSpecificNetworkLoadBalancerResp struct {
	ID          int          `json:"id" tf:"id,computed"`
	Name        string       `json:"name"`
	AccountId   int          `json:"accountId"`
	Cloud       CloudInfo    `json:"cloud"`
	Type        Types        `json:"type"`
	Description string       `json:"description"`
	Port        string       `json:"port"`
	Host        string       `json:"host"`
	IP          string       `json:"ip"`
	SSLEnabled  bool         `json:"sslEnabled"`
	Enabled     bool         `json:"enabled"`
	Visibility  string       `json:"visibility"`
	Config      CreateConfig `json:"config"`
	DateCreated string       `json:"dateCreated"`
	LastUpdated string       `json:"lastUpdated"`
	Meta        MetaInfo     `json:"meta"`
}

//Create LB Monitor Req
type CreateLBMonitor struct {
	CreateLBMonitorReq CreateLBMonitorReq `json:"loadBalancerMonitor"`
}

type CreateLBMonitorReq struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	MonitorType        string `json:"monitorType"`
	MonitorTimeout     int    `json:"monitorTimeout"`
	MonitorInterval    int    `json:"monitorInterval"`
	SendVersion        string `json:"sendVersion"`
	SendType           string `json:"sendType"`
	ReceiveCode        string `json:"receiveCode"`
	MonitorDestination string `json:"monitorDestination"`
	MonitorReverse     bool   `json:"monitorReverse"`
	MonitorTransparent bool   `json:"monitorTransparent"`
	MonitorAdaptive    bool   `json:"monitorAdaptive"`
	FallCount          int    `json:"fallCount"`
	RiseCount          int    `json:"riseCount"`
	AliasPort          int    `json:"aliasPort"`
}

//Create LB Monitor Resp
type CreateLBMonitorResp struct {
	Success       bool          `json:"success"`
	LBMonitorResp LBMonitorResp `json:"loadBalancerMonitor"`
}

type LBMonitorResp struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Visibility         string    `json:"visibility"`
	Description        string    `json:"description"`
	MonitorType        string    `json:"monitorType"`
	MonitorInterval    int       `json:"monitorInterval"`
	MonitorTimeout     int       `json:"monitorTimeout"`
	SendVersion        string    `json:"sendVersion"`
	SendType           string    `json:"sendType"`
	ReceiveCode        string    `json:"receiveCode"`
	MonitorDestination string    `json:"monitorDestination"`
	MonitorReverse     bool      `json:"monitorReverse"`
	MonitorTransparent bool      `json:"monitorTransparent"`
	MonitorAdaptive    bool      `json:"monitorAdaptive"`
	AliasPort          string    `json:"aliasPort"`
	InternalId         string    `json:"internalId"`
	MonitorSource      string    `json:"monitorSource"`
	Status             string    `json:"status"`
	Enabled            bool      `json:"enabled"`
	FallCount          int       `json:"fallCount"`
	RiseCount          int       `json:"riseCount"`
	DateCreated        string    `json:"dateCreated"`
	LastUpdated        string    `json:"lastUpdated"`
	LoadBalancer       LBMonitor `json:"loadBalancer"`
}

type LBMonitor struct {
	ID   string        `json:"id"`
	Name string        `json:"name"`
	IP   string        `json:"ip"`
	Type LBMonitorType `json:"type"`
}

type LBMonitorType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

//Create LB Profile Req
type CreateLBProfile struct {
	CreateLBProfileReq CreateLBProfileReq `json:"loadBalancerProfile"`
}

type CreateLBProfileReq struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	ServiceType   string    `json:"serviceType"`
	ProfileConfig LBProfile `json:"config"`
}

type LBProfile struct {
	ProfileType            string `json:"profileType"`
	RequestHeaderSize      int    `json:"requestHeaderSize"`
	ResponseHeaderSize     int    `json:"responseHeaderSize"`
	ResponseTimeout        int    `json:"responseTimeout"`
	HttpIdleTimeoutName    int    `json:"httpIdleTimeout"`
	FastTcpIdleTimeout     int    `json:"fastTcpIdleTimeout"`
	ConnectionCloseTimeout int    `json:"connectionCloseTimeout"`
	HaFlowMirroring        bool   `json:"haFlowMirroring"`
}

//Create LB Profile Resp
type CreateLBProfileResp struct {
	Success       bool          `json:"success"`
	LBProfileResp LBProfileResp `json:"loadBalancerProfile"`
}

type LBProfileResp struct {
	ID                  int           `json:"id"`
	Name                string        `json:"name"`
	Category            string        `json:"category"`
	ServiceType         string        `json:"serviceType"`
	ServiceTypeDisplay  string        `json:"serviceTypeDisplay"`
	Visibility          string        `json:"visibility"`
	Description         string        `json:"description"`
	InternalID          string        `json:"internalId"`
	ExternalID          string        `json:"ExternalID"`
	Enabled             string        `json:"enabled"`
	InsertXforwardedFor string        `json:"insertXforwardedFor"`
	Editable            string        `json:"editable"`
	LBProfileConfig     profileConfig `json:"config"`
}

type profileConfig struct {
	HttpIdleTimeout        string `json:"httpIdleTimeout"`
	ResponseHeaderSize     string `json:"responseHeaderSize"`
	SharePersistence       string `json:"sharePersistence"`
	RequestHeaderSize      string `json:"requestHeaderSize"`
	HaPersistenceMirroring string `json:"haPersistenceMirroring"`
	PreferServerCipher     string `json:"preferServerCipher"`
	ProfileType            string `json:"profileType"`
	CookieGarbling         string `json:"cookieGarbling"`
	NtlmAuthentication     string `json:"ntlmAuthentication"`
	HaFlowMirroring        string `json:"haFlowMirroring"`
	SessionCache           string `json:"sessionCache"`
	CookieFallback         string `json:"cookieFallback"`
	ResponseTimeout        string `json:"responseTimeout"`
	PurgeEntries           string `json:"purgeEntries"`
	ConnectionCloseTimeout string `json:"connectionCloseTimeout"`
	FastTcpIdleTimeout     string `json:"fastTcpIdleTimeout"`
}

//Get LB Profile
type GetLBProfile struct {
	GetLBProfilesResp []GetLBProfilesResp `json:"loadBalancerProfiles"`
}

type GetLBProfilesResp struct {
	ID                  int             `json:"id"`
	Name                string          `json:"name"`
	Category            string          `json:"category"`
	ServiceType         string          `json:"serviceType"`
	ServiceTypeDisplay  string          `json:"serviceTypeDisplay"`
	Visibility          string          `json:"visibility"`
	Description         string          `json:"description"`
	InternalId          string          `json:"internalId"`
	ExternalId          string          `json:"externalId"`
	Enabled             string          `json:"enabled"`
	InsertXforwardedFor string          `json:"insertXforwardedFor"`
	Editable            string          `json:"editable"`
	DateCreated         string          `json:"dateCreated"`
	LastUpdated         string          `json:"lastUpdated"`
	LBProfileConfig     LBprofileConfig `json:"config"`
}

type LBprofileConfig struct {
	HttpIdleTimeout    string `json:"httpIdleTimeout"`
	NtlmAuthentication string `json:"ntlmAuthentication"`
	RequestHeaderSize  string `json:"requestHeaderSize"`
	ResponseHeaderSize string `json:"responseHeaderSize"`
	ResponseTimeout    string `json:"responseTimeout"`
	XForwardedFor      string `json:"xForwardedFor"`
	ProfileType        string `json:"profileType"`
	Resource_type      string `json:"resource_type"`
}

//Get LB Specific Profile
type GetLBSpecificProfile struct {
	GetLBSpecificProfilesResp GetLBSpecificProfilesResp `json:"loadBalancerProfile"`
}

type GetLBSpecificProfilesResp struct {
	ID                  int             `json:"id"`
	Name                string          `json:"name"`
	Category            string          `json:"category"`
	ServiceType         string          `json:"serviceType"`
	ServiceTypeDisplay  string          `json:"serviceTypeDisplay"`
	Visibility          string          `json:"visibility"`
	Description         string          `json:"description"`
	InternalId          string          `json:"internalId"`
	ExternalId          string          `json:"externalId"`
	Enabled             string          `json:"enabled"`
	InsertXforwardedFor string          `json:"insertXforwardedFor"`
	Editable            string          `json:"editable"`
	DateCreated         string          `json:"dateCreated"`
	LastUpdated         string          `json:"lastUpdated"`
	LBProfileConfig     LBprofileConfig `json:"config"`
}

//Get LB Monitors
type GetLBMonitors struct {
	GetLBMonitorsResp []GetLBMonitorsResp `json:"loadBalancerMonitors"`
}

type GetLBMonitorsResp struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Visibility         string    `json:"visibility"`
	Description        string    `json:"description"`
	MonitorType        string    `json:"monitorType"`
	MonitorInterval    int       `json:"monitorInterval"`
	MonitorTimeout     int       `json:"monitorTimeout"`
	MonitorReverse     bool      `json:"monitorReverse"`
	MonitorTransparent bool      `json:"monitorTransparent"`
	MonitorAdaptive    bool      `json:"monitorAdaptive"`
	InternalId         string    `json:"internalId"`
	ExternalId         string    `json:"externalId"`
	MonitorSource      string    `json:"monitorSource"`
	Status             string    `json:"status"`
	Enabled            bool      `json:"enabled"`
	MaxRetry           int       `json:"maxRetry"`
	FallCount          int       `json:"fallCount"`
	RiseCount          int       `json:"riseCount"`
	DataLength         int       `json:"dataLength"`
	DateCreated        string    `json:"dateCreated"`
	LastUpdated        string    `json:"lastUpdated"`
	LoadBalancer       LBMonitor `json:"loadBalancer"`
}

//Get Specific LB Monitor
type GetSpecificLBMonitor struct {
	GetSpecificLBMonitorResp GetSpecificLBMonitorResp `json:"loadBalancerMonitor"`
}

type GetSpecificLBMonitorResp struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Visibility         string    `json:"visibility"`
	Description        string    `json:"description"`
	MonitorType        string    `json:"monitorType"`
	MonitorInterval    int       `json:"monitorInterval"`
	MonitorTimeout     int       `json:"monitorTimeout"`
	SendVersion        string    `json:"sendVersion"`
	SendType           string    `json:"sendType"`
	ReceiveCode        string    `json:"receiveCode"`
	MonitorReverse     bool      `json:"monitorReverse"`
	MonitorTransparent bool      `json:"monitorTransparent"`
	MonitorAdaptive    bool      `json:"monitorAdaptive"`
	InternalId         string    `json:"internalId"`
	ExternalId         string    `json:"externalId"`
	MonitorSource      string    `json:"monitorSource"`
	Status             string    `json:"status"`
	Enabled            bool      `json:"enabled"`
	MaxRetry           int       `json:"maxRetry"`
	FallCount          int       `json:"fallCount"`
	RiseCount          int       `json:"riseCount"`
	DateCreated        string    `json:"dateCreated"`
	LastUpdated        string    `json:"lastUpdated"`
	LoadBalancer       LBMonitor `json:"loadBalancer"`
}

// Create LB Pool
type CreateLBPool struct {
	CreateLBPoolReq CreateLBPoolReq `json:"loadBalancerPool"`
}

type CreateLBPoolReq struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	VipBalance  string     `json:"vipBalance"`
	MinActive   int        `json:"minActive"`
	PoolConfig  PoolConfig `json:"config"`
}

type PoolConfig struct {
	SnatTranslationType   string      `json:"snatTranslationType"`
	PassiveMonitorPath    int         `json:"passiveMonitorPath"`
	ActiveMonitorPaths    int         `json:"activeMonitorPaths"`
	TcpMultiplexing       bool        `json:"tcpMultiplexing"`
	TcpMultiplexingNumber int         `json:"tcpMultiplexingNumber"`
	SnatIpAddress         string      `json:"snatIpAddress"`
	MemberGroup           MemberGroup `json:"memberGroup"`
}

type MemberGroup struct {
	Name             string `json:"name"`
	Path             string `json:"path"`
	IpRevisionFilter string `json:"ipRevisionFilter"`
	Port             int    `json:"port"`
}

//Create LB Pool Resp
type CreateLBPoolResp struct {
	Success    bool       `json:"success"`
	LBPoolResp LBPoolResp `json:"loadBalancerPool"`
}

type LBPoolResp struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Category         string       `json:"category"`
	Visibility       string       `json:"visibility"`
	Description      string       `json:"description"`
	InternalId       string       `json:"internalId"`
	ExternalId       string       `json:"externalId"`
	Enabled          bool         `json:"enabled"`
	VipBalance       string       `json:"vipBalance"`
	MinActive        int          `json:"minActive"`
	NumberActive     int          `json:"numberActive"`
	NumberInService  int          `json:"numberInService"`
	HealthScore      float32      `json:"healthScore"`
	PerformanceScore float32      `json:"performanceScore"`
	HealthPenalty    float32      `json:"healthPenalty"`
	SecurityPenalty  float32      `json:"securityPenalty"`
	ErrorPenalty     float32      `json:"errorPenalty"`
	Status           string       `json:"status"`
	DateCreated      string       `json:"dateCreated"`
	LastUpdated      string       `json:"lastUpdated"`
	Nodes            []string     `json:"nodes"`
	Monitors         []string     `json:"monitors"`
	Members          []string     `json:"members"`
	LoadBalancer     LBMonitor    `json:"loadBalancer"`
	LBPoolConfig     LBPoolConfig `json:"config"`
}

type LBPoolConfig struct {
	SnatIpAddresses       []string    `json:"snatIpAddresses"`
	TcpMultiplexingNumber int         `json:"tcpMultiplexingNumber"`
	PassiveMonitorPath    int         `json:"passiveMonitorPath"`
	TcpMultiplexing       bool        `json:"tcpMultiplexing"`
	SnatIpAddress         string      `json:"snatIpAddress"`
	SnatTranslationType   string      `json:"snatTranslationType"`
	MemberGroup           MemberGroup `json:"memberGroup"`
}

//Get LB Pools
type GetLBPools struct {
	GetLBPoolsResp []GetLBPoolsResp `json:"loadBalancerPools"`
}

type GetLBPoolsResp struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Visibility       string       `json:"visibility"`
	Description      string       `json:"description"`
	InternalId       string       `json:"internalId"`
	ExternalId       string       `json:"externalId"`
	Enabled          bool         `json:"enabled"`
	VipBalance       string       `json:"vipBalance"`
	MinActive        int          `json:"minActive"`
	NumberActive     int          `json:"numberActive"`
	NumberInService  int          `json:"numberInService"`
	HealthScore      float32      `json:"healthScore"`
	PerformanceScore float32      `json:"performanceScore"`
	HealthPenalty    float32      `json:"healthPenalty"`
	SecurityPenalty  float32      `json:"securityPenalty"`
	ErrorPenalty     float32      `json:"errorPenalty"`
	Status           string       `json:"status"`
	DateCreated      string       `json:"dateCreated"`
	LastUpdated      string       `json:"lastUpdated"`
	Nodes            []string     `json:"nodes"`
	Monitors         []string     `json:"monitors"`
	Members          []string     `json:"members"`
	LoadBalancer     LBMonitor    `json:"loadBalancer"`
	LBPoolConfig     LBPoolConfig `json:"config"`
	Meta             MetaInfo     `json:"meta"`
}

//Get Specific LB Pools
type GetSpecificLBPool struct {
	GetSpecificLBPoolResp GetSpecificLBPoolResp `json:"loadBalancerPool"`
}

type GetSpecificLBPoolResp struct {
	ID               int          `json:"id"`
	Name             string       `json:"name"`
	Visibility       string       `json:"visibility"`
	Description      string       `json:"description"`
	InternalId       string       `json:"internalId"`
	ExternalId       string       `json:"externalId"`
	Enabled          bool         `json:"enabled"`
	VipBalance       string       `json:"vipBalance"`
	MinActive        int          `json:"minActive"`
	NumberActive     int          `json:"numberActive"`
	NumberInService  int          `json:"numberInService"`
	HealthScore      float32      `json:"healthScore"`
	PerformanceScore float32      `json:"performanceScore"`
	HealthPenalty    float32      `json:"healthPenalty"`
	SecurityPenalty  float32      `json:"securityPenalty"`
	ErrorPenalty     float32      `json:"errorPenalty"`
	Status           string       `json:"status"`
	DateCreated      string       `json:"dateCreated"`
	LastUpdated      string       `json:"lastUpdated"`
	Nodes            []string     `json:"nodes"`
	Monitors         []string     `json:"monitors"`
	Members          []string     `json:"members"`
	LoadBalancer     LBMonitor    `json:"loadBalancer"`
	LBPoolConfig     LBPoolConfig `json:"config"`
}
