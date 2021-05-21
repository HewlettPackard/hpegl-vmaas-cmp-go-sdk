// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

import (
	"encoding/json"
	"time"
)

// CreateInstanceBody
type CreateInstanceBody struct {
	// Cloud ID
	ZoneId            string                                `json:"zoneId"`
	Instance          *CreateInstanceBodyInstance           `json:"instance"`
	Volumes           []CreateInstanceBodyVolumes           `json:"volumes"`
	NetworkInterfaces []CreateInstanceBodyNetworkInterfaces `json:"networkInterfaces"`
	Config            *CreateInstanceBodyConfig             `json:"config"`
	Copies            int32                                 `json:"copies,omitempty"`
	Labels            []string                              `json:"labels,omitempty"`
	Tags              []CreateInstanceBodyTag               `json:"tags,omitempty"`
}

type CreateInstanceBodyTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CreateInstanceBodyInstance
type CreateInstanceBodyInstance struct {
	Name         string                                  `json:"name"`
	Cloud        string                                  `json:"cloud"`
	Type         string                                  `json:"type"`
	Site         *CreateInstanceBodyInstanceSite         `json:"site"`
	InstanceType *CreateInstanceBodyInstanceInstanceType `json:"instanceType"`
	Layout       *CreateInstanceBodyInstanceLayout       `json:"layout"`
	Plan         *CreateInstanceBodyInstancePlan         `json:"plan"`
}

// CreateInstanceBodyConfig
type CreateInstanceBodyConfig struct {
	// Virtual Image ID(Required when VMware InstanceType is used)
	Template       int32       `json:"template,omitempty"`
	ResourcePoolId interface{} `json:"resourcePoolId"`
	// To specify agent install (on/off)
	NoAgent              string `json:"noAgent,omitempty"`
	SmbiosAssetTag       string `json:"smbiosAssetTag,omitempty"`
	HostId               string `json:"hostId,omitempty"`
	VmwareDomainName     string `json:"vmwareDomainName,omitempty"`
	VmwareCustomSpec     string `json:"vmwareCustomSpec,omitempty"`
	NestedVirtualization string `json:"nestedVirtualization,omitempty"`
	CreateUser           bool   `json:"createUser,omitempty"`
}

// CreateInstanceBodyInstanceInstanceType
type CreateInstanceBodyInstanceInstanceType struct {
	// Instance type code
	Code string `json:"code"`
}

// CreateInstanceBodyInstanceLayout
type CreateInstanceBodyInstanceLayout struct {
	// The layout id for the instance type that you want to provision.
	Id string `json:"id"`
}

// CreateInstanceBodyInstancePlan
type CreateInstanceBodyInstancePlan struct {
	// Service Plan ID
	Id string `json:"id"`
}

// CreateInstanceBodyInstanceSite
type CreateInstanceBodyInstanceSite struct {
	// Group ID
	Id int32 `json:"id"`
}

// CreateInstanceBodyNetwork
type CreateInstanceBodyNetwork struct {
	Id int32 `json:"id"`
}

// CreateInstanceBodyNetworkInterfaces
type CreateInstanceBodyNetworkInterfaces struct {
	Network                *CreateInstanceBodyNetwork `json:"network"`
	NetworkInterfaceTypeId json.Number                `json:"networkInterfaceTypeId"`
}

// CreateInstanceBodyVolumes
type CreateInstanceBodyVolumes struct {
	Id         int32 `json:"id,omitempty"`
	RootVolume bool  `json:"rootVolume,omitempty"`
	// Name/type of the LV being created
	Name        string `json:"name"`
	Size        int32  `json:"size,omitempty"`
	StorageType int32  `json:"storageType,omitempty"`
	// The ID of the specific datastore. Auto selection can be specified as auto or autoCluster (for clusters).
	DatastoreId interface{} `json:"datastoreId"`
}

type Instances struct {
	Instances []GetInstanceResponseInstance `json:"instances"`
}

// GetInstanceResponse
type GetInstanceResponse struct {
	Instance *GetInstanceResponseInstance `json:"instance"`
}

// GetInstanceResponseInstance
type GetInstanceResponseInstance struct {
	Id                  int32                                       `json:"id,omitempty"`
	Uuid                string                                      `json:"uuid,omitempty"`
	AccountId           int32                                       `json:"accountId,omitempty"`
	Tenant              *GetInstanceResponseInstanceTenant          `json:"tenant,omitempty"`
	InstanceType        *GetInstanceResponseInstanceInstanceType    `json:"instanceType,omitempty"`
	Group               *GetInstanceResponseInstanceGroup           `json:"group,omitempty"`
	Cloud               *GetInstanceResponseInstanceCloud           `json:"cloud,omitempty"`
	Containers          []int32                                     `json:"containers,omitempty"`
	Servers             []int32                                     `json:"servers,omitempty"`
	ConnectionInfo      []GetInstanceResponseInstanceConnectionInfo `json:"connectionInfo,omitempty"`
	Layout              *GetInstanceResponseInstanceLayout          `json:"layout,omitempty"`
	Plan                *GetInstanceResponseInstancePlan            `json:"plan,omitempty"`
	Name                string                                      `json:"name,omitempty"`
	Description         string                                      `json:"description,omitempty"`
	Config              *GetInstanceResponseInstanceConfig          `json:"config,omitempty"`
	Volumes             []GetInstanceResponseInstanceVolumes        `json:"volumes,omitempty"`
	Controllers         []GetInstanceResponseInstanceController     `json:"controllers,omitempty"`
	Interfaces          []GetInstanceResponseInstanceInterfaces     `json:"interfaces,omitempty"`
	CustomOptions       *interface{}                                `json:"customOptions,omitempty"`
	InstanceVersion     string                                      `json:"instanceVersion,omitempty"`
	Labels              []string                                    `json:"labels,omitempty"`
	Tags                []GetInstanceResponseInstanceTags           `json:"tags,omitempty"`
	Evars               []GetInstanceResponseInstanceEvars          `json:"evars,omitempty"`
	MaxMemory           int64                                       `json:"maxMemory,omitempty"`
	MaxStorage          int64                                       `json:"maxStorage,omitempty"`
	MaxCores            int32                                       `json:"maxCores,omitempty"`
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
	ExpireCount         int32                                       `json:"expireCount,omitempty"`
	ExpireWarningSent   bool                                        `json:"expireWarningSent,omitempty"`
	ShutdownCount       int32                                       `json:"shutdownCount,omitempty"`
	ShutdownWarningSent bool                                        `json:"shutdownWarningSent,omitempty"`
	CreatedBy           *GetInstanceResponseInstanceCreatedBy       `json:"createdBy,omitempty"`
	Owner               *GetInstanceResponseInstanceCreatedBy       `json:"owner,omitempty"`
}

//GetInstanceResponseInstanceCloud
type GetInstanceResponseInstanceCloud struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceConfig
type GetInstanceResponseInstanceConfig struct {
	ResourcePoolID       interface{}   `json:"resourcePoolId,omitempty"`
	Template             int           `json:"template,omitempty"`
	Poolprovidertype     interface{}   `json:"poolProviderType,omitempty"`
	Isvpcselectable      bool          `json:"isVpcSelectable,omitempty"`
	Smbiosassettag       interface{}   `json:"smbiosAssetTag,omitempty"`
	Isec2                bool          `json:"isEC2,omitempty"`
	Createuser           bool          `json:"createUser,omitempty"`
	Nestedvirtualization interface{}   `json:"nestedVirtualization,omitempty"`
	Vmwarefolderid       interface{}   `json:"vmwareFolderId,omitempty"`
	Expose               []interface{} `json:"expose,omitempty"`
	Noagent              bool          `json:"noAgent,omitempty"`
	Customoptions        interface{}   `json:"customOptions,omitempty"`
	Createbackup         bool          `json:"createBackup,omitempty"`
	Memorydisplay        string        `json:"memoryDisplay,omitempty"`
	Backup               struct {
		Veeammanagedserver string `json:"veeamManagedServer,omitempty"`
		Createbackup       bool   `json:"createBackup,omitempty"`
		Jobaction          string `json:"jobAction,omitempty"`
		Jobretentioncount  int    `json:"jobRetentionCount,omitempty"`
	} `json:"backup,omitempty"`
	Layoutsize  int           `json:"layoutSize,omitempty"`
	Lbinstances []interface{} `json:"lbInstances,omitempty"`
}

// GetInstanceResponseInstanceConfigBackup
type GetInstanceResponseInstanceConfigBackup struct {
	ProviderBackupType int32  `json:"providerBackupType,omitempty"`
	JobAction          string `json:"jobAction,omitempty"`
	JobName            string `json:"jobName,omitempty"`
	Name               string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceConfigRemovalOptions
type GetInstanceResponseInstanceConfigRemovalOptions struct {
	Force           bool  `json:"force,omitempty"`
	KeepBackups     bool  `json:"keepBackups,omitempty"`
	ReleaseEIPs     bool  `json:"releaseEIPs,omitempty"`
	RemoveVolumes   bool  `json:"removeVolumes,omitempty"`
	RemoveResources bool  `json:"removeResources,omitempty"`
	UserId          int32 `json:"userId,omitempty"`
}

// GetInstanceResponseInstanceConnectionInfo
type GetInstanceResponseInstanceConnectionInfo struct {
	Ip string `json:"ip,omitempty"`
}

// GetInstanceResponseInstanceCreatedBy
type GetInstanceResponseInstanceCreatedBy struct {
	Id       int32  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

// GetInstanceResponseInstanceEvars
type GetInstanceResponseInstanceEvars struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Export bool   `json:"export,omitempty"`
	Masked bool   `json:"masked,omitempty"`
}

// GetInstanceResponseInstanceGroup
type GetInstanceResponseInstanceGroup struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceInstanceType
type GetInstanceResponseInstanceInstanceType struct {
	Id       int32  `json:"id,omitempty"`
	Code     string `json:"code,omitempty"`
	Category string `json:"category,omitempty"`
	Name     string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceInterfaces
type GetInstanceResponseInstanceInterfaces struct {
	Id      string                              `json:"id,omitempty"`
	Network *GetInstanceResponseInstanceNetwork `json:"network,omitempty"`
}

// GetInstanceResponseInstanceLayout
type GetInstanceResponseInstanceLayout struct {
	Id                int32  `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	ProvisionTypeCode string `json:"provisionTypeCode,omitempty"`
}

// GetInstanceResponseInstanceNetwork
type GetInstanceResponseInstanceNetwork struct {
	Id                     int32                                   `json:"id,omitempty"`
	Subnet                 string                                  `json:"subnet,omitempty"`
	Group                  string                                  `json:"group,omitempty"`
	DhcpServer             bool                                    `json:"dhcpServer,omitempty"`
	Name                   string                                  `json:"name,omitempty"`
	Pool                   *GetInstanceResponseInstanceNetworkPool `json:"pool,omitempty"`
	IpAddress              string                                  `json:"ipAddress,omitempty"`
	IpMode                 string                                  `json:"ipMode,omitempty"`
	NetworkInterfaceTypeId int32                                   `json:"networkInterfaceTypeId,omitempty"`
}

// GetInstanceResponseInstanceNetworkPool
type GetInstanceResponseInstanceNetworkPool struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetInstanceResponseInstancePlan
type GetInstanceResponseInstancePlan struct {
	Id   int32  `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceTags
type GetInstanceResponseInstanceTags struct {
	Id    int32  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// GetInstanceResponseInstanceTenant
type GetInstanceResponseInstanceTenant struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetInstanceResponseInstanceVolumes
type GetInstanceResponseInstanceVolumes struct {
	Size              int32   `json:"size,omitempty"`
	Name              string  `json:"name,omitempty"`
	RootVolume        bool    `json:"rootVolume,omitempty"`
	Id                int     `json:"id,omitempty"`
	DatastoreId       string  `json:"datastoreId,omitempty"`
	MaxStorage        float64 `json:"maxStorage,omitempty"`
	DeviceDisplayName string  `json:"deviceDisplayName,omitempty"`
}

type GetInstanceResponseInstanceController struct {
	Editable           bool   `json:"editable,omitempty"`
	Typename           string `json:"typeName,omitempty"`
	Maxdevices         int    `json:"maxDevices,omitempty"`
	Displayorder       int    `json:"displayOrder,omitempty"`
	Active             bool   `json:"active,omitempty"`
	Unitnumber         string `json:"unitNumber,omitempty"`
	Reservedunitnumber int    `json:"reservedUnitNumber,omitempty"`
	Busnumber          string `json:"busNumber,omitempty"`
	Removable          bool   `json:"removable,omitempty"`
	Name               string `json:"name,omitempty"`
	Typeid             int    `json:"typeId,omitempty"`
	ID                 int    `json:"id,omitempty"`
	Category           string `json:"category,omitempty"`
}

// ResizeInstanceBody
type ResizeInstanceBody struct {
	Instance *ResizeInstanceBodyInstance `json:"instance,omitempty"`
}

// ResizeInstanceBodyInstance
type ResizeInstanceBodyInstance struct {
	// Instance ID
	Id   int32                           `json:"id,omitempty"`
	Plan *ResizeInstanceBodyInstancePlan `json:"plan,omitempty"`
	// Can be used to grow just the logical volume of the instance instead of choosing a plan
	Volumes               []ResizeInstanceBodyInstanceVolumes `json:"volumes,omitempty"`
	DeleteOriginalVolumes bool                                `json:"deleteOriginalVolumes,omitempty"`
}

// ResizeInstanceBodyInstancePlan
type ResizeInstanceBodyInstancePlan struct {
	// Service Plan ID
	Id int32 `json:"id,omitempty"`
}

// ResizeInstanceBodyInstanceVolumes
type ResizeInstanceBodyInstanceVolumes struct {
	Id          int32  `json:"id,omitempty"`
	RootVolume  bool   `json:"rootVolume,omitempty"`
	Name        string `json:"name,omitempty"`
	Size        int32  `json:"size,omitempty"`
	StorageType int32  `json:"storageType,omitempty"`
	DatastoreId int32  `json:"datastoreId,omitempty"`
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
	AddTags *interface{} `json:"addTags,omitempty"`
	// Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed
	RemoveTags *interface{} `json:"removeTags,omitempty"`
}

// GetServicePlanResponse
type GetServicePlanResponse struct {
	Plans []GetServicePlanResponsePlans `json:"plans,omitempty"`
}

// GetServicePlanResponsePlans
type GetServicePlanResponsePlans struct {
	Id                    int32                                `json:"id,omitempty"`
	Name                  string                               `json:"name,omitempty"`
	Value                 int32                                `json:"value,omitempty"`
	Code                  string                               `json:"code,omitempty"`
	MaxStorage            int32                                `json:"maxStorage,omitempty"`
	MaxMemory             int32                                `json:"maxMemory,omitempty"`
	MaxCores              int32                                `json:"maxCores,omitempty"`
	CustomCpu             bool                                 `json:"customCpu,omitempty"`
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
	MinDisk               int32                                `json:"minDisk,omitempty"`
	LvmSupported          bool                                 `json:"lvmSupported,omitempty"`
	Datastores            *GetServicePlanResponseDatastores    `json:"datastores,omitempty"`
	SupportsAutoDatastore bool                                 `json:"supportsAutoDatastore,omitempty"`
	AutoOptions           string                               `json:"autoOptions,omitempty"`
	CpuOptions            string                               `json:"cpuOptions,omitempty"`
	CoreOptions           string                               `json:"coreOptions,omitempty"`
	MemoryOptions         string                               `json:"memoryOptions,omitempty"`
	RootCustomSizeOptions *interface{}                         `json:"rootCustomSizeOptions,omitempty"`
	CustomSizeOptions     *interface{}                         `json:"customSizeOptions,omitempty"`
	CustomCores           bool                                 `json:"customCores,omitempty"`
	MaxDisks              int32                                `json:"maxDisks,omitempty"`
	MemorySizeType        string                               `json:"memorySizeType,omitempty"`
}

// GetServicePlanResponseStorageTypes
type GetServicePlanResponseStorageTypes struct {
	Id               int32    `json:"id,omitempty"`
	Editable         bool     `json:"editable,omitempty"`
	OptionTypes      []string `json:"optionTypes,omitempty"`
	DisplayOrder     int32    `json:"displayOrder,omitempty"`
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
