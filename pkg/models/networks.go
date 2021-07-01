//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

// CreateNetworkIpPoolBody
type CreateNetworkIpPoolBody struct {
	Networkpool *CreateNetworkIpPoolBodyNetworkpool `json:"networkpool,omitempty"`
}

// CreateNetworkIpPoolBodyNetworkpool
type CreateNetworkIpPoolBodyNetworkpool struct {
	Name     string                                       `json:"name,omitempty"`
	Type_    *CreateNetworkIpPoolBodyNetworkpoolType      `json:"type,omitempty"`
	IpRanges []CreateNetworkIpPoolBodyNetworkpoolIpRanges `json:"ipRanges,omitempty"`
}

// CreateNetworkIpPoolBodyNetworkpoolIpRanges
type CreateNetworkIpPoolBodyNetworkpoolIpRanges struct {
	StartAddress string `json:"startAddress,omitempty"`
	EndAddress   string `json:"endAddress,omitempty"`
}

// CreateNetworkIpPoolBodyNetworkpoolType
type CreateNetworkIpPoolBodyNetworkpoolType struct {
	Code string `json:"code,omitempty"`
}

// CreateNetworkPoolResponseBody
type CreateNetworkPoolResponseBody struct {
	Success bool                       `json:"success"`
	Errors  *GetNetworkPoolResposeBody `json:"errors"`
}

// CreateNetworkProxyBody
type CreateNetworkProxyBody struct {
	NetworkProxy *CreateNetworkProxyBodyNetworkProxy `json:"networkProxy,omitempty"`
}

// CreateNetworkProxyBodyNetworkProxy
type CreateNetworkProxyBodyNetworkProxy struct {
	Name             string `json:"name,omitempty"`
	ProxyHost        string `json:"proxyHost,omitempty"`
	ProxyPort        string `json:"proxyPort,omitempty"`
	ProxyUser        string `json:"proxyUser,omitempty"`
	ProxyPassword    string `json:"proxyPassword,omitempty"`
	ProxyDomain      string `json:"proxyDomain,omitempty"`
	ProxyWorkstation string `json:"proxyWorkstation,omitempty"`
}

// CreateNetworkProxyRespose
type CreateNetworkProxyResponse struct {
}

// GetNetworkBody
type GetNetworkBody struct {
	Network          *GetNetworkBodyNetwork `json:"network"`
	Nodes            *GetNetworkBodyNodes   `json:"nodes"`
	InstanceStats    *interface{}           `json:"instanceStats"`
	InstanceClusters []string               `json:"instanceClusters"`
}

// GetNetworkBodyNetwork
type GetNetworkBodyNetwork struct {
	Id                      float64                                  `json:"id"`
	Name                    string                                   `json:"name"`
	Zone                    *ListNetworksBodyOwner                   `json:"zone"`
	Type_                   *ListNetworksBodyType                    `json:"type"`
	Owner                   *ListNetworksBodyOwner                   `json:"owner"`
	Code                    string                                   `json:"code"`
	Category                string                                   `json:"category"`
	ExternalId              string                                   `json:"externalId"`
	InternalId              string                                   `json:"internalId"`
	UniqueId                string                                   `json:"uniqueId"`
	ExternalType            string                                   `json:"externalType"`
	RefType                 string                                   `json:"refType"`
	RefId                   float64                                  `json:"refId"`
	DhcpServer              bool                                     `json:"dhcpServer"`
	Visibility              string                                   `json:"visibility"`
	EnableAdmin             bool                                     `json:"enableAdmin"`
	ScanNetwork             bool                                     `json:"scanNetwork"`
	Active                  bool                                     `json:"active"`
	DefaultNetwork          bool                                     `json:"defaultNetwork"`
	AssignPublicIp          bool                                     `json:"assignPublicIp"`
	ApplianceUrlProxyBypass bool                                     `json:"applianceUrlProxyBypass"`
	ZonePool                *ListNetworksBodyOwner                   `json:"zonePool"`
	AllowStaticOverride     bool                                     `json:"allowStaticOverride"`
	Subnets                 []string                                 `json:"subnets"`
	Tenants                 []ListNetworksBodyTenants                `json:"tenants"`
	ResourcePermission      *GetNetworkBodyNetworkResourcePermission `json:"resourcePermission"`
}

// GetNetworkBodyNetworkResourcePermission
type GetNetworkBodyNetworkResourcePermission struct {
	AllPlans bool     `json:"allPlans"`
	All      bool     `json:"all"`
	Sites    []string `json:"sites"`
	Plans    []string `json:"plans"`
}

// GetNetworkBodyNodes
type GetNetworkBodyNodes struct {
	Containers []string `json:"containers"`
	Instances  []string `json:"instances"`
	Apps       []string `json:"apps"`
}

// GetNetworkPoolResposeBody
type GetNetworkPoolResposeBody struct {
	NetworkPool *GetNetworkPoolResposeBodyNetworkPool `json:"networkPool"`
}

// GetNetworkPoolResposeBodyNetworkPool
type GetNetworkPoolResposeBodyNetworkPool struct {
	Id            float64                           `json:"id"`
	Type_         *ListNetworksBodyType             `json:"type"`
	Account       *ListNetworksBodyOwner            `json:"account"`
	Name          string                            `json:"name"`
	DnsServers    []string                          `json:"dnsServers"`
	DnsSuffixList []string                          `json:"dnsSuffixList"`
	DhcpServer    bool                              `json:"dhcpServer"`
	IpCount       float64                           `json:"ipCount"`
	FreeCount     float64                           `json:"freeCount"`
	PoolEnabled   bool                              `json:"poolEnabled"`
	IpRanges      []GetNetworkPoolsResponseIpRanges `json:"ipRanges"`
}

// GetNetworkPoolsResponse
type GetNetworkPoolsResponse struct {
	NetworkPools     []GetNetworkPoolsResponseNetworkPools `json:"networkPools"`
	NetworkPoolCount float64                               `json:"networkPoolCount"`
	Meta             *ListNetworksBodyMeta                 `json:"meta"`
}

// GetNetworkPoolsResponseIpRanges
type GetNetworkPoolsResponseIpRanges struct {
	Id           float64 `json:"id,omitempty"`
	StartAddress string  `json:"startAddress,omitempty"`
	EndAddress   string  `json:"endAddress,omitempty"`
	AddressCount float64 `json:"addressCount,omitempty"`
	Active       bool    `json:"active,omitempty"`
	DateCreated  string  `json:"dateCreated,omitempty"`
	LastUpdated  string  `json:"lastUpdated,omitempty"`
}

// GetNetworkPoolsResponseNetworkPools
type GetNetworkPoolsResponseNetworkPools struct {
	Id            float64                           `json:"id,omitempty"`
	Type_         *ListNetworksBodyType             `json:"type,omitempty"`
	Account       *ListNetworksBodyOwner            `json:"account,omitempty"`
	Name          string                            `json:"name,omitempty"`
	DnsServers    []string                          `json:"dnsServers,omitempty"`
	DnsSuffixList []string                          `json:"dnsSuffixList,omitempty"`
	DhcpServer    bool                              `json:"dhcpServer,omitempty"`
	IpCount       float64                           `json:"ipCount,omitempty"`
	FreeCount     float64                           `json:"freeCount,omitempty"`
	PoolEnabled   bool                              `json:"poolEnabled,omitempty"`
	IpRanges      []GetNetworkPoolsResponseIpRanges `json:"ipRanges,omitempty"`
}

// GetNetworkProxy
type GetNetworkProxy struct {
	NetworkProxy *GetNetworkProxyNetworkProxy `json:"networkProxy"`
}

// GetNetworkProxyNetworkProxy
type GetNetworkProxyNetworkProxy struct {
	Id         float64                `json:"id"`
	Name       string                 `json:"name"`
	ProxyHost  string                 `json:"proxyHost"`
	ProxyPort  float64                `json:"proxyPort"`
	Visibility string                 `json:"visibility"`
	Account    *ListNetworksBodyOwner `json:"account"`
	Owner      *ListNetworksBodyOwner `json:"owner"`
}

// GetNetworkRouter
type GetNetworkRouter struct {
	NetworkRouter *ListNetworkRoutersProperties `json:"networkRouter"`
}

// ListNetworkProxies
type ListNetworkProxies struct {
	NetworkProxies    []ListNetworkProxiesNetworkProxies `json:"networkProxies"`
	NetworkProxyCount float64                            `json:"networkProxyCount"`
	Meta              *ListNetworksBodyMeta              `json:"meta"`
}

// ListNetworkProxiesNetworkProxies
type ListNetworkProxiesNetworkProxies struct {
	Id         float64                `json:"id,omitempty"`
	Name       string                 `json:"name,omitempty"`
	ProxyHost  string                 `json:"proxyHost,omitempty"`
	ProxyPort  float64                `json:"proxyPort,omitempty"`
	Visibility string                 `json:"visibility,omitempty"`
	Account    *ListNetworksBodyOwner `json:"account,omitempty"`
	Owner      *ListNetworksBodyOwner `json:"owner,omitempty"`
}

// ListNetworkRouters
type ListNetworkRouters struct {
	NetworkRouters *[]interface{}                `json:"networkRouters"`
	Items          *interface{}                  `json:"items,omitempty"`
	Properties     *ListNetworkRoutersProperties `json:"properties,omitempty"`
}

// ListNetworkRoutersPropertiesConfig
type ListNetworkRoutersPropertiesConfig struct {
	DhcpRelay *interface{} `json:"dhcpRelay"`
}

// ListNetworkRoutersPropertiesFirewallDefaultPolicy
type ListNetworkRoutersPropertiesFirewallDefaultPolicy struct {
	Action         string `json:"action"`
	LoggingEnabled bool   `json:"loggingEnabled"`
}

// ListNetworkRoutersPropertiesFirewallGlobal
type ListNetworkRoutersPropertiesFirewallGlobal struct {
	TcpPickOngoingConnections     bool    `json:"tcpPickOngoingConnections"`
	EnableFtpLooseMode            bool    `json:"enableFtpLooseMode"`
	TcpAllowOutOfWindowPackets    bool    `json:"tcpAllowOutOfWindowPackets"`
	TcpSendResetForClosedVsePorts bool    `json:"tcpSendResetForClosedVsePorts"`
	DropInvalidTraffic            bool    `json:"dropInvalidTraffic"`
	LogInvalidTraffic             bool    `json:"logInvalidTraffic"`
	TcpTimeoutOpen                float64 `json:"tcpTimeoutOpen"`
	TcpTimeoutEstablished         float64 `json:"tcpTimeoutEstablished"`
	TcpTimeoutClose               float64 `json:"tcpTimeoutClose"`
	UdpTimeout                    float64 `json:"udpTimeout"`
	IcmpTimeout                   float64 `json:"icmpTimeout"`
	Icmp6Timeout                  float64 `json:"icmp6Timeout"`
	IpGenericTimeout              float64 `json:"ipGenericTimeout"`
	EnableSynFloodProtection      bool    `json:"enableSynFloodProtection"`
	LogIcmpErrors                 bool    `json:"logIcmpErrors"`
	DropIcmpReplays               bool    `json:"dropIcmpReplays"`
	EnableSnmpAlg                 bool    `json:"enableSnmpAlg"`
	EnableFtpAlg                  bool    `json:"enableFtpAlg"`
	EnableTftpAlg                 bool    `json:"enableTftpAlg"`
}

// ListNetworkRoutersPropertiesFirewall
type ListNetworkRoutersPropertiesFirewall struct {
	Enabled       bool                                               `json:"enabled"`
	Version       float64                                            `json:"version"`
	DefaultPolicy *ListNetworkRoutersPropertiesFirewallDefaultPolicy `json:"defaultPolicy"`
	Global        *ListNetworkRoutersPropertiesFirewallGlobal        `json:"global"`
	Rules         []ListNetworkRoutersPropertiesFirewallRules        `json:"rules"`
}

// ListNetworkRoutersPropertiesFirewallRules
type ListNetworkRoutersPropertiesFirewallRules struct {
	Id              float64  `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Direction       string   `json:"direction,omitempty"`
	RuleType        string   `json:"ruleType,omitempty"`
	Policy          string   `json:"policy,omitempty"`
	SourceType      string   `json:"sourceType,omitempty"`
	DestinationType string   `json:"destinationType,omitempty"`
	ApplicationType string   `json:"applicationType,omitempty"`
	Applications    []string `json:"applications,omitempty"`
}

// ListNetworkRoutersProperties
type ListNetworkRoutersProperties struct {
	Id              float64                                                              `json:"id"`
	Code            string                                                               `json:"code"`
	Name            string                                                               `json:"name"`
	Category        string                                                               `json:"category"`
	DateCreated     string                                                               `json:"dateCreated"`
	LastUpdated     string                                                               `json:"lastUpdated"`
	RouterType      string                                                               `json:"routerType"`
	Status          string                                                               `json:"status"`
	Enabled         bool                                                                 `json:"enabled"`
	ExternalIp      string                                                               `json:"externalIp"`
	ProviderId      string                                                               `json:"providerId"`
	Type_           *ListNetworkRoutersPropertiesType                                    `json:"type"`
	NetworkServer   *ListNetworkRoutersPropertiesNetworkServer                           `json:"networkServer"`
	Zone            *ListNetworkRoutersPropertiesNetworkServerIntegrationIntegrationType `json:"zone"`
	ExternalNetwork *ListNetworkRoutersPropertiesNetworkServerIntegrationIntegrationType `json:"externalNetwork"`
	Interfaces      []ListNetworkRoutersPropertiesInterfaces                             `json:"interfaces"`
	Firewall        *ListNetworkRoutersPropertiesFirewall                                `json:"firewall"`
	Routes          []ListNetworkRoutersPropertiesRoutes                                 `json:"routes"`
	Config          *ListNetworkRoutersPropertiesConfig                                  `json:"config"`
	Permissions     *ListNetworkRoutersPropertiesPermissions                             `json:"permissions"`
}

// ListNetworkRoutersPropertiesInterfaces
type ListNetworkRoutersPropertiesInterfaces struct {
	Id              float64               `json:"id,omitempty"`
	Name            string                `json:"name,omitempty"`
	Code            string                `json:"code,omitempty"`
	InterfaceType   string                `json:"interfaceType,omitempty"`
	NetworkPosition string                `json:"networkPosition,omitempty"`
	IpAddress       string                `json:"ipAddress,omitempty"`
	Cidr            string                `json:"cidr,omitempty"`
	ExternalLink    string                `json:"externalLink,omitempty"`
	Enabled         bool                  `json:"enabled,omitempty"`
	Network         *ListNetworksBodyType `json:"network,omitempty"`
}

// ListNetworkRoutersPropertiesNetworkServer
type ListNetworkRoutersPropertiesNetworkServer struct {
	Id          float64                                               `json:"id"`
	Name        string                                                `json:"name"`
	Integration *ListNetworkRoutersPropertiesNetworkServerIntegration `json:"integration"`
}

// ListNetworkRoutersPropertiesNetworkServerIntegration
type ListNetworkRoutersPropertiesNetworkServerIntegration struct {
	Id              float64                                                              `json:"id"`
	Name            string                                                               `json:"name"`
	Enabled         bool                                                                 `json:"enabled"`
	Status          string                                                               `json:"status"`
	StatusDate      string                                                               `json:"statusDate"`
	IntegrationType *ListNetworkRoutersPropertiesNetworkServerIntegrationIntegrationType `json:"integrationType"`
}

// ListNetworkRoutersPropertiesNetworkServerIntegrationIntegrationType
type ListNetworkRoutersPropertiesNetworkServerIntegrationIntegrationType struct {
	Id   float64 `json:"id"`
	Code string  `json:"code"`
	Name string  `json:"name"`
}

// ListNetworkRoutersPropertiesPermissions
type ListNetworkRoutersPropertiesPermissions struct {
	Visibility string `json:"visibility"`
}

// ListNetworkRoutersPropertiesRoutes
type ListNetworkRoutersPropertiesRoutes struct {
	Id                float64 `json:"id,omitempty"`
	Code              string  `json:"code,omitempty"`
	RouteType         string  `json:"routeType,omitempty"`
	Source            string  `json:"source,omitempty"`
	SourceType        string  `json:"sourceType,omitempty"`
	Destination       string  `json:"destination,omitempty"`
	DestinationType   string  `json:"destinationType,omitempty"`
	DefaultRoute      bool    `json:"defaultRoute,omitempty"`
	ExternalInterface string  `json:"externalInterface,omitempty"`
	ExternalId        string  `json:"externalId,omitempty"`
	Enabled           bool    `json:"enabled,omitempty"`
	Visible           bool    `json:"visible,omitempty"`
}

// ListNetworkRoutersPropertiesType
type ListNetworkRoutersPropertiesType struct {
	Id               float64  `json:"id"`
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Enabled          bool     `json:"enabled"`
	Creatable        bool     `json:"creatable"`
	Selectable       bool     `json:"selectable"`
	HasFirewall      bool     `json:"hasFirewall"`
	HasDhcp          bool     `json:"hasDhcp"`
	HasRouting       bool     `json:"hasRouting"`
	HasNetworkServer bool     `json:"hasNetworkServer"`
	OptionTypes      []string `json:"optionTypes"`
	RuleOptionTypes  []string `json:"ruleOptionTypes"`
}

// ListNetworksBody
type ListNetworksBody struct {
	Networks     []ListNetworksBodyNetworks `json:"networks"`
	NetworkCount int                        `json:"networkCount"`
	Meta         *ListNetworksBodyMeta      `json:"meta"`
}

// ListNetworksBodyMeta
type ListNetworksBodyMeta struct {
	Size   float64 `json:"size"`
	Total  float64 `json:"total"`
	Offset float64 `json:"offset"`
	Max    float64 `json:"max"`
}

// ListNetworksBodyNetworks
type ListNetworksBodyNetworks struct {
	Id                      int                       `json:"id,omitempty"`
	Name                    string                    `json:"name,omitempty"`
	DisplayName             string                    `json:"displayName,omitempty"`
	Zone                    *ListNetworksBodyZone     `json:"zone,omitempty"`
	Type_                   *ListNetworksBodyType     `json:"type,omitempty"`
	Owner                   *ListNetworksBodyOwner    `json:"owner,omitempty"`
	Code                    string                    `json:"code,omitempty"`
	Category                string                    `json:"category,omitempty"`
	ExternalId              string                    `json:"externalId,omitempty"`
	InternalId              string                    `json:"internalId,omitempty"`
	UniqueId                string                    `json:"uniqueId,omitempty"`
	ExternalType            string                    `json:"externalType,omitempty"`
	RefType                 string                    `json:"refType,omitempty"`
	RefId                   float64                   `json:"refId,omitempty"`
	DhcpServer              bool                      `json:"dhcpServer,omitempty"`
	Visibility              string                    `json:"visibility,omitempty"`
	EnableAdmin             bool                      `json:"enableAdmin,omitempty"`
	ScanNetwork             bool                      `json:"scanNetwork,omitempty"`
	Active                  bool                      `json:"active,omitempty"`
	DefaultNetwork          bool                      `json:"defaultNetwork,omitempty"`
	AssignPublicIp          bool                      `json:"assignPublicIp,omitempty"`
	ApplianceUrlProxyBypass bool                      `json:"applianceUrlProxyBypass,omitempty"`
	ZonePool                *ListNetworksBodyOwner    `json:"zonePool,omitempty"`
	AllowStaticOverride     bool                      `json:"allowStaticOverride,omitempty"`
	Subnets                 []interface{}             `json:"subnets,omitempty"`
	Tenants                 []ListNetworksBodyTenants `json:"tenants,omitempty"`
}

// ListNetworksBodyOwner
type ListNetworksBodyOwner struct {
	Id   float64 `json:"id"`
	Name string  `json:"name"`
}

// ListNetworksBodyTenants
type ListNetworksBodyTenants struct {
	Id   float64 `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
}

// ListNetworksBodyType
type ListNetworksBodyType struct {
	Id   float64 `json:"id"`
	Name string  `json:"name"`
	Code string  `json:"code"`
}

// ListNetworksBodyZone
type ListNetworksBodyZone struct {
	Id   float64 `json:"id"`
	Name string  `json:"name"`
}

// UpdateNetworkBody
type UpdateNetworkBody struct {
	Network *UpdateNetworkBodyNetwork `json:"network"`
}

// UpdateNetworkBodyNetwork
type UpdateNetworkBodyNetwork struct {
	Description string `json:"description,omitempty"`
	Active      bool   `json:"active,omitempty"`
	DhcpServer  bool   `json:"dhcpServer,omitempty"`
	// Supported values \"on\" or \"off\"
	ApplianceUrlProxyBypass string                                       `json:"applianceUrlProxyBypass,omitempty"`
	DnsPrimary              string                                       `json:"dnsPrimary,omitempty"`
	DnsSecondary            string                                       `json:"dnsSecondary,omitempty"`
	NetworkProxy            *UpdateNetworkBodyNetworkNetworkProxy        `json:"networkProxy,omitempty"`
	Pool                    *UpdateNetworkBodyNetworkPool                `json:"pool,omitempty"`
	ResourcePermissions     *UpdateNetworkBodyNetworkResourcePermissions `json:"resourcePermissions,omitempty"`
}

// UpdateNetworkBodyNetworkNetworkProxy
type UpdateNetworkBodyNetworkNetworkProxy struct {
	Id int `json:"id,omitempty"`
}

// UpdateNetworkBodyNetworkPool
type UpdateNetworkBodyNetworkPool struct {
	Id int `json:"id,omitempty"`
}

// UpdateNetworkBodyNetworkResourcePermissions
type UpdateNetworkBodyNetworkResourcePermissions struct {
	All   bool                                               `json:"all,omitempty"`
	Sites []UpdateNetworkBodyNetworkResourcePermissionsSites `json:"sites,omitempty"`
}

// UpdateNetworkBodyNetworkResourcePermissionsSites
type UpdateNetworkBodyNetworkResourcePermissionsSites struct {
	Id int `json:"id,omitempty"`
}

// UpdateNetworkIpPoolBody
type UpdateNetworkIpPoolBody struct {
	Networkpool *UpdateNetworkIpPoolBodyNetworkpool `json:"networkpool,omitempty"`
}

// UpdateNetworkIpPoolBodyNetworkpool
type UpdateNetworkIpPoolBodyNetworkpool struct {
	Name string `json:"name,omitempty"`
}

// UpdateNetworkProxyBody
type UpdateNetworkProxyBody struct {
	NetworkProxy *CreateNetworkProxyBodyNetworkProxy `json:"networkProxy,omitempty"`
}

// UpdateNetworkProxyRespose
type UpdateNetworkProxyResponse struct {
}

// UpdateNetworkRespose
type UpdateNetworkRespose struct {
	Success bool                         `json:"success"`
	Errors  *interface{}                 `json:"errors"`
	Network *UpdateNetworkResposeNetwork `json:"network"`
}

// UpdateNetworkResposeNetwork
type UpdateNetworkResposeNetwork struct {
	Id                      float64                                        `json:"id"`
	Name                    string                                         `json:"name"`
	Zone                    *ListNetworksBodyOwner                         `json:"zone"`
	Type_                   *ListNetworksBodyType                          `json:"type"`
	Owner                   *ListNetworksBodyOwner                         `json:"owner"`
	Code                    string                                         `json:"code"`
	Category                string                                         `json:"category"`
	ExternalId              string                                         `json:"externalId"`
	InternalId              string                                         `json:"internalId"`
	UniqueId                string                                         `json:"uniqueId"`
	ExternalType            string                                         `json:"externalType"`
	RefType                 string                                         `json:"refType"`
	RefId                   float64                                        `json:"refId"`
	DhcpServer              bool                                           `json:"dhcpServer"`
	Visibility              string                                         `json:"visibility"`
	EnableAdmin             bool                                           `json:"enableAdmin"`
	ScanNetwork             bool                                           `json:"scanNetwork"`
	Active                  bool                                           `json:"active"`
	DefaultNetwork          bool                                           `json:"defaultNetwork"`
	AssignPublicIp          bool                                           `json:"assignPublicIp"`
	ApplianceUrlProxyBypass bool                                           `json:"applianceUrlProxyBypass"`
	ZonePool                *ListNetworksBodyOwner                         `json:"zonePool"`
	AllowStaticOverride     bool                                           `json:"allowStaticOverride"`
	Subnets                 []string                                       `json:"subnets"`
	Tenants                 []ListNetworksBodyTenants                      `json:"tenants"`
	ResourcePermission      *UpdateNetworkResposeNetworkResourcePermission `json:"resourcePermission"`
}

// UpdateNetworkResposeNetworkResourcePermission
type UpdateNetworkResposeNetworkResourcePermission struct {
	AllPlans bool                                                 `json:"allPlans"`
	All      bool                                                 `json:"all"`
	Sites    []UpdateNetworkResposeNetworkResourcePermissionSites `json:"sites"`
	Plans    []interface{}                                        `json:"plans"`
}

// UpdateNetworkResposeNetworkResourcePermissionSites
type UpdateNetworkResposeNetworkResourcePermissionSites struct {
	Id       float64 `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Default_ bool    `json:"default,omitempty"`
}
