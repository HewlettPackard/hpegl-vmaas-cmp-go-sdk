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
	NetworkServerID     int                       `json:"networkServerId" tf:"network_server_id"`
	Enabled             bool                      `json:"enabled"`
	Visibility          string                    `json:"visibility"`
	Config              CreateConfig              `json:"config"`
	ResourcePermissions EnableResourcePermissions `json:"resourcePermission" tf:"resource_permission"`
}

type CreateConfig struct {
	AdminState bool   `json:"adminState" tf:"admin_state"`
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
	AccountID   int          `json:"accountId"`
	Cloud       CloudInfo    `json:"cloud"`
	Type        Types        `json:"type"`
	Description string       `json:"description"`
	Port        int          `json:"port"`
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
	AccountID   int          `json:"accountId"`
	Cloud       CloudInfo    `json:"cloud"`
	Type        Types        `json:"type"`
	Description string       `json:"description"`
	Port        int          `json:"port"`
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
	AccountID   int          `json:"accountId"`
	Cloud       CloudInfo    `json:"cloud"`
	Type        Types        `json:"type"`
	Description string       `json:"description"`
	Port        int          `json:"port"`
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

// Create LB Monitor Req
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

// Create LB Monitor Resp
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
	AliasPort          int       `json:"aliasPort"`
	InternalID         string    `json:"internalId"`
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
	ID   int           `json:"id"`
	Name string        `json:"name"`
	IP   string        `json:"ip"`
	Type LBMonitorType `json:"type"`
}

type LBMonitorType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// Create LB Profile Req
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
	HTTPIdleTimeoutName    int    `json:"httpIdleTimeout"`
	FastTCPIdleTimeout     int    `json:"fastTcpIdleTimeout"`
	ConnectionCloseTimeout int    `json:"connectionCloseTimeout"`
	HaFlowMirroring        bool   `json:"haFlowMirroring"`
	CookieMode             string `json:"cookieMode"`
	CookieName             string `json:"cookieName"`
	CookieType             string `json:"cookieType"`
	CookieFallback         bool   `json:"cookieFallback"`
	CookieGarbling         bool   `json:"cookieGarbling"`
}

// Create LB Profile Resp
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
	ExternalID          string        `json:"externalId"`
	Enabled             bool          `json:"enabled"`
	InsertXforwardedFor bool          `json:"insertXforwardedFor"`
	Editable            bool          `json:"editable"`
	LBProfileConfig     profileConfig `json:"config"`
}

type profileConfig struct {
	HTTPIdleTimeout        int    `json:"httpIdleTimeout"`
	ResponseHeaderSize     int    `json:"responseHeaderSize"`
	SharePersistence       bool   `json:"sharePersistence"`
	RequestHeaderSize      int    `json:"requestHeaderSize"`
	HaPersistenceMirroring bool   `json:"haPersistenceMirroring"`
	PreferServerCipher     bool   `json:"preferServerCipher"`
	ProfileType            string `json:"profileType"`
	CookieGarbling         bool   `json:"cookieGarbling"`
	NtlmAuthentication     bool   `json:"ntlmAuthentication"`
	HaFlowMirroring        bool   `json:"haFlowMirroring"`
	SessionCache           bool   `json:"sessionCache"`
	CookieFallback         bool   `json:"cookieFallback"`
	ResponseTimeout        int    `json:"responseTimeout"`
	PurgeEntries           bool   `json:"purgeEntries"`
	ConnectionCloseTimeout int    `json:"connectionCloseTimeout"`
	FastTCPIdleTimeout     int    `json:"fastTcpIdleTimeout"`
}

// Get LB Profile
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
	InternalID          string          `json:"internalId"`
	ExternalID          string          `json:"externalId"`
	Enabled             bool            `json:"enabled"`
	InsertXforwardedFor bool            `json:"insertXforwardedFor"`
	Editable            bool            `json:"editable"`
	DateCreated         string          `json:"dateCreated"`
	LastUpdated         string          `json:"lastUpdated"`
	LBProfileConfig     LBprofileConfig `json:"config"`
}

type LBprofileConfig struct {
	HTTPIdleTimeout    int    `json:"httpIdleTimeout"`
	NtlmAuthentication bool   `json:"ntlmAuthentication"`
	RequestHeaderSize  int    `json:"requestHeaderSize"`
	ResponseHeaderSize int    `json:"responseHeaderSize"`
	ResponseTimeout    int    `json:"responseTimeout"`
	XForwardedFor      string `json:"xForwardedFor"`
	ProfileType        string `json:"profileType"`
	ResourceType       string `json:"resource_type"`
}

// Get LB Specific Profile
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
	InternalID          string          `json:"internalId"`
	ExternalID          string          `json:"externalId"`
	Enabled             bool            `json:"enabled"`
	InsertXforwardedFor bool            `json:"insertXforwardedFor"`
	Editable            bool            `json:"editable"`
	DateCreated         string          `json:"dateCreated"`
	LastUpdated         string          `json:"lastUpdated"`
	LBProfileConfig     LBprofileConfig `json:"config"`
}

// Get LB Monitors
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
	InternalID         string    `json:"internalId"`
	ExternalID         string    `json:"externalId"`
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

// Get Specific LB Monitor
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
	InternalID         string    `json:"internalId"`
	ExternalID         string    `json:"externalId"`
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
	TCPMultiplexing       bool        `json:"tcpMultiplexing"`
	TCPMultiplexingNumber int         `json:"tcpMultiplexingNumber"`
	SnatIPAddress         string      `json:"snatIpAddress"`
	MemberGroup           MemberGroup `json:"memberGroup"`
}

type MemberGroup struct {
	Name             string `json:"name"`
	Path             string `json:"path"`
	IPRevisionFilter string `json:"ipRevisionFilter"`
	Port             int    `json:"port"`
}

// Create LB Pool Resp
type CreateLBPoolResp struct {
	Success    bool       `json:"success"`
	LBPoolResp LBPoolResp `json:"loadBalancerPool"`
}

type LBPoolResp struct {
	ID               int          `json:"id"`
	Name             string       `json:"name"`
	Category         string       `json:"category"`
	Visibility       string       `json:"visibility"`
	Description      string       `json:"description"`
	InternalID       string       `json:"internalId"`
	ExternalID       string       `json:"externalId"`
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
	SnatIPAddresses       []string    `json:"snatIpAddresses"`
	TCPMultiplexingNumber int         `json:"tcpMultiplexingNumber"`
	PassiveMonitorPath    int         `json:"passiveMonitorPath"`
	TCPMultiplexing       bool        `json:"tcpMultiplexing"`
	SnatIPAddress         string      `json:"snatIpAddress"`
	SnatTranslationType   string      `json:"snatTranslationType"`
	MemberGroup           MemberGroup `json:"memberGroup"`
}

// Get LB Pools
type GetLBPools struct {
	GetLBPoolsResp []GetLBPoolsResp `json:"loadBalancerPools"`
}

type GetLBPoolsResp struct {
	ID               int          `json:"id"`
	Name             string       `json:"name"`
	Visibility       string       `json:"visibility"`
	Description      string       `json:"description"`
	InternalID       string       `json:"internalId"`
	ExternalID       string       `json:"externalId"`
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

// Get Specific LB Pools
type GetSpecificLBPool struct {
	GetSpecificLBPoolResp GetSpecificLBPoolResp `json:"loadBalancerPool"`
}

type GetSpecificLBPoolResp struct {
	ID               int          `json:"id"`
	Name             string       `json:"name"`
	Visibility       string       `json:"visibility"`
	Description      string       `json:"description"`
	InternalID       string       `json:"internalId"`
	ExternalID       string       `json:"externalId"`
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

// CREATE LB Virtual servers

type CreateLBVirtualServers struct {
	CreateLBVirtualServersReq CreateLBVirtualServersReq `json:"loadBalancerInstance"`
}

type CreateLBVirtualServersReq struct {
	Description         string              `json:"description"`
	VipName             string              `json:"vipName"`
	VipAddress          string              `json:"vipAddress"`
	VipProtocol         string              `json:"vipProtocol"`
	VipPort             string              `json:"vipPort"`
	Pool                int                 `json:"pool"`
	SSLServerCert       int                 `json:"sslServerCert"`
	SSLCert             int                 `json:"sslCert"`
	VirtualServerConfig VirtualServerConfig `json:"config"`
}

type VirtualServerConfig struct {
	Persistence        string `json:"persistence"`
	PersistenceProfile int    `json:"persistenceProfile"`
	ApplicationProfile string `json:"applicationProfile"`
	SSLClientProfile   string `json:"sslClientProfile"`
	SSLServerProfile   string `json:"sslServerProfile"`
}

// CREATE LB Virtual Server Resp
type LBVirtualServersResp struct {
	CreateLBVirtualServersResp CreateLBVirtualServersResp `json:"loadBalancerInstance"`
	Success                    bool                       `json:"success"`
}

type CreateLBVirtualServersResp struct {
	ID                 int           `json:"id"`
	Name               string        `json:"name"`
	Description        string        `json:"description"`
	Active             bool          `json:"active"`
	Sticky             bool          `json:"sticky"`
	SslEnabled         bool          `json:"sslEnabled"`
	ExternalAddress    bool          `json:"externalAddress"`
	VipName            string        `json:"vipName"`
	VipAddress         string        `json:"vipAddress"`
	VipProtocol        string        `json:"vipProtocol"`
	VipMode            string        `json:"vipMode"`
	VipPort            int           `json:"vipPort"`
	VipShared          bool          `json:"vipShared"`
	VipDirectAddress   string        `json:"vipDirectAddress"`
	Removing           bool          `json:"removing"`
	VirtualServiceName string        `json:"virtualServiceName"`
	VipSource          string        `json:"vipSource"`
	DateCreated        string        `json:"dateCreated"`
	LastUpdated        string        `json:"lastUpdated"`
	LoadBalancer       LBMonitor     `json:"loadBalancer"`
	VSConfig           VSConfig      `json:"config"`
	VSPool             VSPool        `json:"pool"`
	SSLCert            SSLCert       `json:"sslCert"`
	SSLServerCert      SSLServerCert `json:"sslServerCert"`
}

// GET LB Virtual servers

type GetLBVirtualServers struct {
	GetLBVirtualServersResp []GetLBVirtualServersResp `json:"loadBalancerInstances"`
}

type GetLBVirtualServersResp struct {
	ID                 int           `json:"id"`
	Name               string        `json:"name"`
	Description        string        `json:"description"`
	Active             bool          `json:"active"`
	Sticky             bool          `json:"sticky"`
	SslEnabled         bool          `json:"sslEnabled"`
	ExternalAddress    bool          `json:"externalAddress"`
	VipName            string        `json:"vipName"`
	VipAddress         string        `json:"vipAddress"`
	VipProtocol        string        `json:"vipProtocol"`
	VipMode            string        `json:"vipMode"`
	VipPort            int           `json:"vipPort"`
	VipShared          bool          `json:"vipShared"`
	VipDirectAddress   string        `json:"vipDirectAddress"`
	Removing           bool          `json:"removing"`
	VirtualServiceName string        `json:"virtualServiceName"`
	VipSource          string        `json:"vipSource"`
	DateCreated        string        `json:"dateCreated"`
	LastUpdated        string        `json:"lastUpdated"`
	LoadBalancer       LBMonitor     `json:"loadBalancer"`
	VSConfig           VSConfig      `json:"config"`
	VSPool             VSPool        `json:"pool"`
	SSLCert            SSLCert       `json:"sslCert"`
	SSLServerCert      SSLServerCert `json:"sslServerCert"`
	Meta               MetaInfo      `json:"meta"`
}

type VSConfig struct {
	Persistence        string   `json:"persistence"`
	PersistenceProfile int      `json:"persistenceProfile"`
	SslServerProfile   string   `json:"sslServerProfile"`
	SslClientProfile   string   `json:"sslClientProfile"`
	ApplicationProfile string   `json:"applicationProfile"`
	Monitors           []string `json:"monitors"`
}

type VSPool struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SSLCert struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SSLServerCert struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GET LB Specific Virtual servers

type GetSpecificLBVirtualServers struct {
	GetSpecificLBVirtualServersResp GetSpecificLBVirtualServersResp `json:"loadBalancerInstance"`
}

type GetSpecificLBVirtualServersResp struct {
	ID                 int           `json:"id"`
	Name               string        `json:"name"`
	Description        string        `json:"description"`
	Active             bool          `json:"active"`
	Sticky             bool          `json:"sticky"`
	SslEnabled         bool          `json:"sslEnabled"`
	ExternalAddress    bool          `json:"externalAddress"`
	VipName            string        `json:"vipName"`
	VipAddress         string        `json:"vipAddress"`
	VipProtocol        string        `json:"vipProtocol"`
	VipMode            string        `json:"vipMode"`
	VipPort            int           `json:"vipPort"`
	VipShared          bool          `json:"vipShared"`
	VipDirectAddress   string        `json:"vipDirectAddress"`
	Removing           bool          `json:"removing"`
	VirtualServiceName string        `json:"virtualServiceName"`
	VipSource          string        `json:"vipSource"`
	DateCreated        string        `json:"dateCreated"`
	LastUpdated        string        `json:"lastUpdated"`
	LoadBalancer       LBMonitor     `json:"loadBalancer"`
	VSConfig           VSConfig      `json:"config"`
	VSPool             VSPool        `json:"pool"`
	SSLCert            SSLCert       `json:"sslCert"`
	SSLServerCert      SSLServerCert `json:"sslServerCert"`
}
