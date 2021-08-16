// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

import (
	"encoding/json"
	"time"
)

// CreateInstanceBody
type CreateInstanceBody struct {
	// Cloud ID
	ZoneID            json.Number                           `json:"zoneId"`
	Instance          *CreateInstanceBodyInstance           `json:"instance"`
	Volumes           []CreateInstanceBodyVolumes           `json:"volumes"`
	NetworkInterfaces []CreateInstanceBodyNetworkInterfaces `json:"networkInterfaces"`
	Config            *CreateInstanceBodyConfig             `json:"config"`
	Labels            []string                              `json:"labels,omitempty"`
	Tags              []CreateInstanceBodyTag               `json:"tags,omitempty"`
	Evars             []GetInstanceResponseInstanceEvars    `json:"evars,omitempty"`
	LayoutSize        int                                   `json:"layoutSize,omitempty"`
	CloneName         string                                `json:"name,omitempty"`
	Context           string                                `json:"Context,omitempty"`
	PowerScheduleType json.Number                           `json:"powerScheduleType,omitempty"`
	ShutdownDays      json.Number                           `json:"shutdownDays,omitempty"`
	ExpireDays        json.Number                           `json:"expireDays,omitempty"`
	Ports             []CreateInstancePorts                 `json:"ports,omitempty"`
	Environment       string                                `json:"environment,omitempty"`
}

type CreateInstancePorts struct {
	Name string `json:"name,omitempty"`
	Port string `json:"port,omitempty"`
	Lb   string `json:"lb,omitempty"`
}

type CreateInstanceBodyTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CreateInstanceBodyInstance
type CreateInstanceBodyInstance struct {
	Name              string                                  `json:"name,omitempty"`
	Template          int                                     `json:"template,omitempty"`
	Type              string                                  `json:"type,omitempty"`
	Site              *CreateInstanceBodyInstanceSite         `json:"site"`
	InstanceType      *CreateInstanceBodyInstanceInstanceType `json:"instanceType"`
	Layout            *CreateInstanceBodyInstanceLayout       `json:"layout"`
	Plan              *CreateInstanceBodyInstancePlan         `json:"plan"`
	HostName          string                                  `json:"hostName,omitempty"`
	Copies            int                                     `json:"copies,omitempty"`
	EnvironmentPrefix string                                  `json:"environmentPrefix,omitempty"`
}

// CreateInstanceBodyConfig
type CreateInstanceBodyConfig struct {
	// To specify agent install (on/off)
	NoAgent              string      `json:"noAgent,omitempty"`
	Template             int         `json:"template,omitempty"`
	ResourcePoolID       json.Number `json:"resourcePoolId"`
	SmbiosAssetTag       string      `json:"smbiosAssetTag,omitempty"`
	HostID               string      `json:"hostId,omitempty"`
	VmwareDomainName     string      `json:"vmwareDomainName,omitempty"`
	VmwareCustomSpec     string      `json:"vmwareCustomSpec,omitempty"`
	NestedVirtualization string      `json:"nestedVirtualization,omitempty"`
	CreateUser           bool        `json:"createUser,omitempty"`
	VMwareFolderID       string      `json:"vmwareFolderId,omitempty"`
}

// CreateInstanceBodyInstanceInstanceType
type CreateInstanceBodyInstanceInstanceType struct {
	// Instance type code
	Code string `json:"code"`
}

// CreateInstanceBodyInstanceLayout
type CreateInstanceBodyInstanceLayout struct {
	// The layout id for the instance type that you want to provision.
	ID json.Number `json:"id"`
}

// CreateInstanceBodyInstancePlan
type CreateInstanceBodyInstancePlan struct {
	// Service Plan ID
	ID json.Number `json:"id"`
}

// CreateInstanceBodyInstanceSite
type CreateInstanceBodyInstanceSite struct {
	// Group ID
	ID   int     `json:"id"`
	Name *string `'json:"name,omitempty"`
}

// CreateInstanceBodyNetwork
type CreateInstanceBodyNetwork struct {
	ID int `json:"id"`
}

// CreateInstanceBodyNetworkInterfaces
type CreateInstanceBodyNetworkInterfaces struct {
	Network                *CreateInstanceBodyNetwork `json:"network"`
	NetworkInterfaceTypeID json.Number                `json:"networkInterfaceTypeId,omitempty"`
}

// CreateInstanceBodyVolumes
type CreateInstanceBodyVolumes struct {
	ID         int  `json:"id,omitempty"`
	RootVolume bool `json:"rootVolume,omitempty"`
	// Name/type of the LV being created
	Name        string `json:"name"`
	Size        int    `json:"size,omitempty"`
	StorageType int    `json:"storageType,omitempty"`
	// The ID of the specific datastore. Auto selection can be specified as auto or autoCluster (for clusters).
	DatastoreID interface{} `json:"datastoreId,omitempty"`
}

type Instances struct {
	Instances []GetInstanceResponseInstance `json:"instances"`
	Success   bool                          `json:"success,omitempty"`
}

// GetInstanceResponse
type GetInstanceResponse struct {
	Instance *GetInstanceResponseInstance `json:"instance"`
}

// GetInstanceResponseInstance
type GetInstanceResponseInstance struct {
	ID                  int                                         `json:"id,omitempty"`
	UUID                string                                      `json:"uuid,omitempty"`
	AccountID           int                                         `json:"accountId,omitempty"`
	Tenant              *GetInstanceResponseInstanceTenant          `json:"tenant,omitempty"`
	InstanceType        *GetInstanceResponseInstanceInstanceType    `json:"instanceType,omitempty"`
	Group               *GetInstanceResponseInstanceGroup           `json:"group,omitempty"`
	Cloud               *GetInstanceResponseInstanceCloud           `json:"cloud,omitempty"`
	Containers          interface{}                                 `json:"containers,omitempty"`
	Servers             interface{}                                 `json:"servers,omitempty"`
	ConnectionInfo      []GetInstanceResponseInstanceConnectionInfo `json:"connectionInfo,omitempty"`
	Layout              *GetInstanceResponseInstanceLayout          `json:"layout,omitempty"`
	Plan                *GetInstanceResponseInstancePlan            `json:"plan,omitempty"`
	Name                string                                      `json:"name,omitempty"`
	Description         string                                      `json:"description,omitempty"`
	Config              *GetInstanceResponseInstanceConfig          `json:"config,omitempty"`
	Volumes             []GetInstanceResponseInstanceVolumes        `json:"volumes,omitempty"`
	Interfaces          []GetInstanceResponseInstanceInterfaces     `json:"interfaces,omitempty"`
	CustomOptions       *interface{}                                `json:"customOptions,omitempty"`
	InstanceVersion     string                                      `json:"instanceVersion,omitempty"`
	Labels              []string                                    `json:"labels,omitempty"`
	Tags                []CreateInstanceBodyTag                     `json:"tags,omitempty"`
	Evars               []GetInstanceResponseInstanceEvars          `json:"evars,omitempty"`
	MaxMemory           int64                                       `json:"maxMemory,omitempty"`
	MaxStorage          int64                                       `json:"maxStorage,omitempty"`
	MaxCores            int                                         `json:"maxCores,omitempty"`
	HourlyCost          float64                                     `json:"hourlyCost,omitempty"`
	HourlyPrice         float64                                     `json:"hourlyPrice,omitempty"`
	DateCreated         time.Time                                   `json:"dateCreated,omitempty"`
	LastUpdated         time.Time                                   `json:"lastUpdated,omitempty"`
	HostName            string                                      `json:"hostName,omitempty"`
	FirewallEnabled     bool                                        `json:"firewallEnabled,omitempty"`
	NetworkLevel        string                                      `json:"networkLevel,omitempty"`
	AutoScale           bool                                        `json:"autoScale,omitempty"`
	Locked              bool                                        `json:"locked,omitempty"`
	Status              string                                      `json:"status,omitempty"`
	StatusDate          string                                      `json:"statusDate,omitempty"`
	StatusMessage       string                                      `json:"statusMessage,omitempty"`
	ErrorMessage        string                                      `json:"errorMessage,omitempty"`
	ExpireCount         int                                         `json:"expireCount,omitempty"`
	ExpireWarningSent   bool                                        `json:"expireWarningSent,omitempty"`
	ShutdownCount       int                                         `json:"shutdownCount,omitempty"`
	ShutdownWarningSent bool                                        `json:"shutdownWarningSent,omitempty"`
	CreatedBy           *GetInstanceResponseInstanceCreatedBy       `json:"createdBy,omitempty"`
	Owner               *GetInstanceResponseInstanceCreatedBy       `json:"owner,omitempty"`
	EnvironmentPrefix   string                                      `json:"environmentPrefix"`
	InstanceContext     string                                      `json:"instanceContext"`
}

// GetInstanceResponseInstanceCloud
type GetInstanceResponseInstanceCloud struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceConfig
type GetInstanceResponseInstanceConfig struct {
	ResourcePoolID       json.Number   `json:"resourcePoolId,omitempty"`
	Template             int           `json:"template,omitempty"`
	Poolprovidertype     interface{}   `json:"poolProviderType,omitempty"`
	Isvpcselectable      bool          `json:"isVpcSelectable,omitempty"`
	Smbiosassettag       string        `json:"smbiosAssetTag,omitempty"`
	Isec2                bool          `json:"isEC2,omitempty"`
	Createuser           bool          `json:"createUser,omitempty"`
	Nestedvirtualization interface{}   `json:"nestedVirtualization,omitempty"`
	Vmwarefolderid       string        `json:"vmwareFolderId,omitempty"`
	Expose               []interface{} `json:"expose,omitempty"`
	Noagent              interface{}   `json:"noAgent,omitempty"`
	Customoptions        interface{}   `json:"customOptions,omitempty"`
	Createbackup         bool          `json:"createBackup,omitempty"`
	Memorydisplay        string        `json:"memoryDisplay,omitempty"`
	Backup               struct {
		Veeammanagedserver string `json:"veeamManagedServer,omitempty"`
		Createbackup       bool   `json:"createBackup,omitempty"`
		Jobaction          string `json:"jobAction,omitempty"`
		Jobretentioncount  int    `json:"jobRetentionCount,omitempty"`
	} `json:"backup,omitempty"`
	Layoutsize        int           `json:"layoutSize,omitempty"`
	Lbinstances       []interface{} `json:"lbInstances,omitempty"`
	PowerScheduleType int           `json:"powerScheduleType"`
}

// GetInstanceResponseInstanceConfigBackup
type GetInstanceResponseInstanceConfigBackup struct {
	ProviderBackupType int    `json:"providerBackupType,omitempty"`
	JobAction          string `json:"jobAction,omitempty"`
	JobName            string `json:"jobName,omitempty"`
	Name               string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceConfigRemovalOptions
type GetInstanceResponseInstanceConfigRemovalOptions struct {
	Force           bool `json:"force,omitempty"`
	KeepBackups     bool `json:"keepBackups,omitempty"`
	ReleaseEIPs     bool `json:"releaseEIPs,omitempty"`
	RemoveVolumes   bool `json:"removeVolumes,omitempty"`
	RemoveResources bool `json:"removeResources,omitempty"`
	UserID          int  `json:"userId,omitempty"`
}

// GetInstanceResponseInstanceConnectionInfo
type GetInstanceResponseInstanceConnectionInfo struct {
	IP string `json:"ip,omitempty"`
}

// GetInstanceResponseInstanceCreatedBy
type GetInstanceResponseInstanceCreatedBy struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

// GetInstanceResponseInstanceEvars
type GetInstanceResponseInstanceEvars struct {
	Name   string      `json:"name,omitempty"`
	Value  interface{} `json:"value,omitempty"`
	Export bool        `json:"export,omitempty"`
	Masked bool        `json:"masked,omitempty"`
}

// GetInstanceResponseInstanceGroup
type GetInstanceResponseInstanceGroup struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceInstanceType
type GetInstanceResponseInstanceInstanceType struct {
	ID       int    `json:"id,omitempty"`
	Code     string `json:"code,omitempty"`
	Category string `json:"category,omitempty"`
	Name     string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceInterfaces
type GetInstanceResponseInstanceInterfaces struct {
	ID                     interface{}                         `json:"id,omitempty"`
	Network                *GetInstanceResponseInstanceNetwork `json:"network,omitempty"`
	NetworkInterfaceTypeID int                                 `json:"networkInterfaceTypeId,omitempty"`
}

// GetInstanceResponseInstanceLayout
type GetInstanceResponseInstanceLayout struct {
	ID                int    `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	ProvisionTypeCode string `json:"provisionTypeCode,omitempty"`
}

// GetInstanceResponseInstanceNetwork
type GetInstanceResponseInstanceNetwork struct {
	ID         json.Number                             `json:"id,omitempty"`
	Subnet     string                                  `json:"subnet,omitempty"`
	Group      string                                  `json:"group,omitempty"`
	DhcpServer bool                                    `json:"dhcpServer,omitempty"`
	Name       string                                  `json:"name,omitempty"`
	Pool       *GetInstanceResponseInstanceNetworkPool `json:"pool,omitempty"`
	IPAddress  string                                  `json:"ipAddress,omitempty"`
	IPMode     string                                  `json:"ipMode,omitempty"`
}

// GetInstanceResponseInstanceNetworkPool
type GetInstanceResponseInstanceNetworkPool struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetInstanceResponseInstancePlan
type GetInstanceResponseInstancePlan struct {
	ID   int    `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceTags
type GetInstanceResponseInstanceTags struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// GetInstanceResponseInstanceTenant
type GetInstanceResponseInstanceTenant struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceVolumes
type GetInstanceResponseInstanceVolumes struct {
	Size              int         `json:"size,omitempty"`
	Name              string      `json:"name,omitempty"`
	RootVolume        bool        `json:"rootVolume,omitempty"`
	ID                int         `json:"id,omitempty"`
	DatastoreID       interface{} `json:"datastoreId,omitempty"`
	MaxStorage        float64     `json:"maxStorage,omitempty"`
	DeviceDisplayName string      `json:"deviceDisplayName,omitempty"`
}

// ResizeInstanceBody
type ResizeInstanceBody struct {
	Instance              *ResizeInstanceBodyInstance         `json:"instance"`
	Volumes               []ResizeInstanceBodyInstanceVolumes `json:"volumes"`
	DeleteOriginalVolumes bool                                `json:"deleteOriginalVolumes,omitempty"`
}

type ResizeInstanceBodyInstance struct {
	ID   int                             `json:"id,omitempty"`
	Plan *ResizeInstanceBodyInstancePlan `json:"plan"`
}

// ResizeInstanceBodyInstancePlan
type ResizeInstanceBodyInstancePlan struct {
	// Service Plan ID
	ID int `json:"id"`
}

// ResizeInstanceBodyInstanceVolumes
type ResizeInstanceBodyInstanceVolumes struct {
	ID          json.Number `json:"id"`
	RootVolume  bool        `json:"rootVolume"`
	Name        string      `json:"name"`
	Size        int         `json:"size"`
	SizeID      interface{} `json:"sizeId,omitempty"`
	StorageType interface{} `json:"storageType,omitempty"`
	DatastoreID interface{} `json:"datastoreId,omitempty"`
}

type ResizeInstanceResponse struct {
	Instance *ResizeInstanceResponseInstance `json:"instance"`
	Success  bool                            `json:"success"`
}

type ResizeInstanceResponseInstance struct {
	ID        int                                `json:"id"`
	Name      string                             `json:"string"`
	Cloud     *GetInstanceResponseInstanceCloud  `json:"cloud,omitempty"`
	Plan      *ResizeInstanceBodyInstancePlan    `json:"plan"`
	Volumes   []GetInstanceResposeResizeVolumes  `json:"volumes"`
	AccountID int                                `json:"accountId,omitempty"`
	Tenant    *GetInstanceResponseInstanceTenant `json:"tenant,omitempty"`
}

type GetInstanceResposeResizeVolumes struct {
	ID          json.Number `json:"id,omitempty"`
	RootVolume  interface{} `json:"rootVolume,omitempty"`
	Name        string      `json:"name,omitempty"`
	Size        json.Number `json:"size,omitempty"`
	StorageType json.Number `json:"storageType,omitempty"`
	DatastoreID interface{} `json:"datastoreId,omitempty"`
}

// SnapshotBody
type SnapshotBody struct {
	Snapshot *SnapshotBodySnapshot `json:"snapshot,omitempty"`
}

// SnapshotBodySnapshot
type SnapshotBodySnapshot struct {
	// Optional name for the snapshot being created
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ListSnapshotResponse struct {
	Snapshots []ListSnapshotResponseInstance `json:"snapshots"`
}

type ListSnapshotResponseInstance struct {
	ID              int         `json:"id,omitempty"`
	Name            string      `json:"name,omitempty"`
	Description     interface{} `json:"description,omitempty"`
	ExternalID      string      `json:"externalId,omitempty"`
	Status          string      `json:"status,omitempty"`
	State           interface{} `json:"state,omitempty"`
	SnapshotType    string      `json:"snapshotType,omitempty"`
	SnapshotCreated time.Time   `json:"snapshotCreated,omitempty"`
	Zone            struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"zone,omitempty"`
	Datastore       interface{} `json:"datastore,omitempty"`
	ParentSnapshot  interface{} `json:"parentSnapshot,omitempty"`
	CurrentlyActive bool        `json:"currentlyActive,omitempty"`
	DateCreated     time.Time   `json:"dateCreated,omitempty"`
}

type ImportSnapshotBody struct {
	StorageProviderID int `json:"storageProviderId,omitempty"`
}

// UpdateInstanceBody
type UpdateInstanceBody struct {
	Instance *UpdateInstanceBodyInstance `json:"instance,omitempty"`
}

// UpdateInstanceBodyInstance
type UpdateInstanceBodyInstance struct {
	// Unique name scoped to your account for the instance
	Name string `json:"name,omitempty"`
	// Optional description field
	Description string `json:"description,omitempty"`
	// Add or update value of Metadata tags, Array of objects having a name and value
	AddTags []CreateInstanceBodyTag `json:"addTags,omitempty"`
	// Remove Metadata tags, Array of objects having a name and an optional value. If value is passed,
	// it must match to be removed
	RemoveTags        []CreateInstanceBodyTag               `json:"removeTags,omitempty"`
	Site              *CreateInstanceBodyInstanceSite       `json:"site"`
	Owner             *GetInstanceResponseInstanceCreatedBy `json:"owner,omitempty"`
	PowerScheduleType json.Number                           `json:"powerScheduleType,omitempty"`
	Labels            []string                              `json:"labels,omitempty"`
	Tags              []CreateInstanceBodyTag               `json:"tags,omitempty"`
	InstanceContext   string                                `json:"instanceContext,omitempty"`
}

type UpdateInstanceResponse struct {
	Instance *UpdateInstanceResponseInstance `json:"instance"`
}

type UpdateInstanceResponseInstance struct {
	Name    string                                `json:"name,omitempty"`
	ID      int                                   `json:"id,omitempty"`
	Group   *CreateInstanceBodyInstanceSite       `json:"group"`
	Owner   *GetInstanceResponseInstanceCreatedBy `json:"owner,omitempty"`
	Labels  []string                              `json:"labels,omitempty"`
	Tags    []CreateInstanceBodyTag               `json:"tags,omitempty"`
	Cloud   *GetInstanceResponseInstanceCloud     `json:"cloud,omitempty"`
	Success bool                                  `json:"success"`
}

// GetServicePlanResponse
type GetServicePlanResponse struct {
	Plans []GetServicePlanResponsePlans `json:"plans,omitempty"`
}

// GetServicePlanResponsePlans
type GetServicePlanResponsePlans struct {
	ID                    int                                  `json:"id,omitempty"`
	Name                  string                               `json:"name,omitempty"`
	Value                 int                                  `json:"value,omitempty"`
	Code                  string                               `json:"code,omitempty"`
	MaxStorage            int                                  `json:"maxStorage,omitempty"`
	MaxMemory             int                                  `json:"maxMemory,omitempty"`
	MaxCores              int                                  `json:"maxCores,omitempty"`
	CustomCPU             bool                                 `json:"customCpu,omitempty"`
	CustomMaxMemory       bool                                 `json:"customMaxMemory,omitempty"`
	CustomMaxStorage      bool                                 `json:"customMaxStorage,omitempty"`
	CustomMaxDataStorage  bool                                 `json:"customMaxDataStorage,omitempty"`
	CustomCoresPerSocket  bool                                 `json:"customCoresPerSocket,omitempty"`
	StorageTypes          []GetServicePlanResponseStorageTypes `json:"storageTypes,omitempty"`
	RootStorageTypes      []GetServicePlanResponseStorageTypes `json:"rootStorageTypes,omitempty"`
	AddVolumes            bool                                 `json:"addVolumes,omitempty"`
	CustomizeVolume       bool                                 `json:"customizeVolume,omitempty"`
	RootDiskCustomizable  bool                                 `json:"rootDiskCustomizable,omitempty"`
	NoDisks               bool                                 `json:"noDisks,omitempty"`
	HasDatastore          bool                                 `json:"hasDatastore,omitempty"`
	MinDisk               int                                  `json:"minDisk,omitempty"`
	LvmSupported          bool                                 `json:"lvmSupported,omitempty"`
	Datastores            *GetServicePlanResponseDatastores    `json:"datastores,omitempty"`
	SupportsAutoDatastore bool                                 `json:"supportsAutoDatastore,omitempty"`
	AutoOptions           string                               `json:"autoOptions,omitempty"`
	CPUOptions            string                               `json:"cpuOptions,omitempty"`
	CoreOptions           string                               `json:"coreOptions,omitempty"`
	MemoryOptions         string                               `json:"memoryOptions,omitempty"`
	RootCustomSizeOptions *interface{}                         `json:"rootCustomSizeOptions,omitempty"`
	CustomSizeOptions     *interface{}                         `json:"customSizeOptions,omitempty"`
	CustomCores           bool                                 `json:"customCores,omitempty"`
	MaxDisks              int                                  `json:"maxDisks,omitempty"`
	MemorySizeType        string                               `json:"memorySizeType,omitempty"`
}

// GetServicePlanResponseStorageTypes
type GetServicePlanResponseStorageTypes struct {
	ID               int      `json:"id,omitempty"`
	Editable         bool     `json:"editable,omitempty"`
	OptionTypes      []string `json:"optionTypes,omitempty"`
	DisplayOrder     int      `json:"displayOrder,omitempty"`
	Code             string   `json:"code,omitempty"`
	VolumeType       string   `json:"volumeType,omitempty"`
	Deletable        bool     `json:"deletable,omitempty"`
	DefaultType      bool     `json:"defaultType,omitempty"`
	CreateDatastore  bool     `json:"createDatastore,omitempty"`
	Resizable        bool     `json:"resizable,omitempty"`
	StorageType      string   `json:"storageType,omitempty"`
	AllowSearch      bool     `json:"allowSearch,omitempty"`
	DisplayName      string   `json:"displayName,omitempty"`
	HasDatastore     bool     `json:"hasDatastore,omitempty"`
	CustomSize       bool     `json:"customSize,omitempty"`
	AutoDelete       bool     `json:"autoDelete,omitempty"`
	Name             string   `json:"name,omitempty"`
	ConfigurableIOPS bool     `json:"configurableIOPS,omitempty"`
	CustomLabel      bool     `json:"customLabel,omitempty"`
	Enabled          bool     `json:"enabled,omitempty"`
	Description      string   `json:"description,omitempty"`
	VolumeCategory   string   `json:"volumeCategory,omitempty"`
	NameEditable     bool     `json:"nameEditable,omitempty"`
}

// GetServicePlanResponseDatastores
type GetServicePlanResponseDatastores struct {
	Cluster string `json:"cluster,omitempty"`
	Store   string `json:"store,omitempty"`
}

type InstancePowerResponse struct {
	Success bool `json:"success"`
}

type GetInstanceHistoryProcessType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type GetInstanceHistoryProcesses struct {
	ID            int                           `json:"id"`
	AccountID     int                           `json:"accountId"`
	UniqueID      string                        `json:"uniqueId"`
	ProcessType   GetInstanceHistoryProcessType `json:"processType"`
	DisplayName   string                        `json:"displayName"`
	Description   interface{}                   `json:"description"`
	SubType       interface{}                   `json:"subType"`
	SubID         interface{}                   `json:"subId"`
	ZoneID        interface{}                   `json:"zoneId"`
	IntegrationID interface{}                   `json:"integrationId"`
	AppID         interface{}                   `json:"appId"`
	InstanceID    int                           `json:"instanceId"`
	ContainerID   interface{}                   `json:"containerId"`
	ServerID      interface{}                   `json:"serverId"`
	ContainerName interface{}                   `json:"containerName"`
	Status        string                        `json:"status"`
	Reason        interface{}                   `json:"reason"`
	Percent       float64                       `json:"percent"`
	StatusEta     int                           `json:"statusEta"`
	Message       interface{}                   `json:"message"`
	Output        interface{}                   `json:"output"`
	Error         interface{}                   `json:"error"`
	StartDate     time.Time                     `json:"startDate"`
	EndDate       time.Time                     `json:"endDate"`
	Duration      int                           `json:"duration"`
	DateCreated   time.Time                     `json:"dateCreated"`
	LastUpdated   time.Time                     `json:"lastUpdated"`
	CreatedBy     struct {
		Username    string `json:"username"`
		DisplayName string `json:"displayName"`
	} `json:"createdBy"`
	UpdatedBy struct {
		Username    string `json:"username"`
		DisplayName string `json:"displayName"`
	} `json:"updatedBy"`
	Events []interface{} `json:"events"`
}

type GetInstanceHistory struct {
	Processes []GetInstanceHistoryProcesses `json:"processes"`
}

type IDModel struct {
	ID int `json:"id,omitempty"`
}

type CreateInstanceCloneInstanceTypeBody struct {
	Code string `json:"code"`
}

type CreateInstanceCloneInstanceBody struct {
	Tags              []string `json:"tags,omitempty"`
	InstanceContext   string   `json:"instanceContext,omitempty"`
	EnvironmentPrefix string   `json:"environmentPrefix,omitempty"`
	PowerScheduleType int      `json:"powerScheduleType,omitempty"`
}

type CreateInstanceCloneBody struct {
	Name              string                                `json:"name,omitempty"`
	Cloud             IDModel                               `json:"cloud,omitempty"`
	Group             IDModel                               `json:"group,omitempty"`
	Type              string                                `json:"type,omitempty"`
	HostName          string                                `json:"hostname,omitempty"`
	InstanceType      CreateInstanceCloneInstanceTypeBody   `json:"instanceType,omitempty"`
	Description       string                                `json:"description,omitempty"`
	Instance          CreateInstanceCloneInstanceBody       `json:"instance,omitempty"`
	Layout            IDModel                               `json:"layout,omitempty"`
	Plan              IDModel                               `json:"plan,omitempty"`
	LayoutSize        int                                   `json:"layoutSize,omitempty"`
	Config            CreateInstanceBodyConfig              `json:"config,omitempty"`
	Volumes           []CreateInstanceBodyVolumes           `json:"volumes,omitempty"`
	NetworkInterfaces []CreateInstanceBodyNetworkInterfaces `json:"networkInterfaces,omitempty"`
	Evars             []GetInstanceResponseInstanceEvars    `json:"evars,omitempty"`
	Metadata          []CreateInstanceBodyTag               `json:"metadata,omitempty"`
}
