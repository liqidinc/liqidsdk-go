// File: enums_gen.go
//
// Copyright (c) 2022 Liqid, Inc. All rights reserved.
// 
// Redistribution and use in source and binary forms, with or without
// modification, are not permitted without prior consent.
//
// Liqid SDK - Version 3.2.0
// This file was automatically generated - do not modify it directly.
//

package liqidsdk

// BackupDestination
// Indicates the backup destination
type BackupDestination int32
const (
	// BackupDestinationLocal
	// Local backup
	BackupDestinationLocal BackupDestination = 0

	// BackupDestinationRemote
	// Remote backup
	BackupDestinationRemote BackupDestination = 1
)

// ConfigurationFileSection
// Describes a type of configuration file section
type ConfigurationFileSection string
const (
	// ConfigurationFileSectionEnclosure
	// Describes enclosure configuration
	ConfigurationFileSectionEnclosure ConfigurationFileSection = "enclosure"

	// ConfigurationFileSectionCPU
	// Describes CPU configuration
	ConfigurationFileSectionCPU ConfigurationFileSection = "cpu"
)

// ConfigurationFileType
// Describes a type of configuration file
type ConfigurationFileType string
const (
	// ConfigurationFileTypeSlurm
	// Slurm configuration file
	ConfigurationFileTypeSlurm ConfigurationFileType = "slurm"

	// ConfigurationFileTypeKubernetes
	// Kubernetes configuration file
	ConfigurationFileTypeKubernetes ConfigurationFileType = "kubernetes"

	// ConfigurationFileTypeRedfish
	// Redfish configuration file
	ConfigurationFileTypeRedfish ConfigurationFileType = "redfish"

	// ConfigurationFileTypeVAPI
	// VAPI configuration file
	ConfigurationFileTypeVAPI ConfigurationFileType = "vapi"

	// ConfigurationFileTypeIPMI
	// IPMI configuration file
	ConfigurationFileTypeIPMI ConfigurationFileType = "ipmi"
)

// DeviceType
// Represents a particular device type as a short string.
// Note that redundant but differing values are not represented here.
// This includes iba, fch, ramd, scm, vic, enc, nof, gof, pcpu, ncpu, gcpu
type DeviceType string
const (
	// DeviceTypeEnclosure
	// Represents an enclosure
	DeviceTypeEnclosure DeviceType = "encl"

	// DeviceTypeEthernet
	// Represents an ethernet NIC
	DeviceTypeEthernet DeviceType = "eth"

	// DeviceTypeLink
	// Represents a general network link
	DeviceTypeLink DeviceType = "link"

	// DeviceTypeMemory
	// Represents configurable/switchable memory
	DeviceTypeMemory DeviceType = "mem"

	// DeviceTypeNetwork
	// Represents a general network link
	DeviceTypeNetwork DeviceType = "nic"

	// DeviceTypeInfiniband1
	// Represents an infiniband link
	DeviceTypeInfiniband1 DeviceType = "infi"

	// DeviceTypeInfiniband2
	// Represents an infiniband link
	DeviceTypeInfiniband2 DeviceType = "iba"

	// DeviceTypeSSD
	// Represents a switchable SSD resource
	DeviceTypeSSD DeviceType = "ssd"

	// DeviceTypeTarget
	// Represents a target resource
	DeviceTypeTarget DeviceType = "targ"

	// DeviceTypePLX
	// Represents a PLX chip resource
	DeviceTypePLX DeviceType = "plx"

	// DeviceTypeCompute
	// Represents a CPU resource.
	// Returned from CPUDeviceInfo
	DeviceTypeCompute DeviceType = "compute"

	// DeviceTypeCPU
	// Returned from CPUDeviceInfo
	DeviceTypeCPU DeviceType = "cpu"

	// DeviceTypeGPU
	// Represents a graphics processing unit resource
	DeviceTypeGPU DeviceType = "gpu"

	// DeviceTypeGCPU
	// Represents a graphics processing unit resource
	DeviceTypeGCPU DeviceType = "gcpu"

	// DeviceTypeFiberChannel1
	// Represents a fiber channel link
	DeviceTypeFiberChannel1 DeviceType = "fibr"

	// DeviceTypeFiberChannel2
	// Represents a fiber channel link
	DeviceTypeFiberChannel2 DeviceType = "fch"

	// DeviceTypeFPGA
	// Represents an FPGA resource
	DeviceTypeFPGA DeviceType = "fpga"
)

// FabricToggleCompositeOption
// TODO
type FabricToggleCompositeOption string
const (
	// FabricToggleCompositeOptionAdd
	// The accompanying information is to be added
	FabricToggleCompositeOptionAdd FabricToggleCompositeOption = "add"
)

// FabricType
// Describes a particular fabric topology type
type FabricType string
const (
	// FabricTypeGen3
	// Fabric is based on gen3 pci protocol and hardware
	FabricTypeGen3 FabricType = "gen3"

	// FabricTypeGen4
	// Fabric is based on gen4 pci protocol and hardware
	FabricTypeGen4 FabricType = "gen4"

	// FabricTypeNVOverFabric
	// NVMe over fabric
	FabricTypeNVOverFabric FabricType = "nvof"

	// FabricTypeGPUOverFabric
	// GPU over fabric
	FabricTypeGPUOverFabric FabricType = "gpuof"

	// FabricTypeGenZ
	// Gen Z protocol and hardware
	FabricTypeGenZ FabricType = "genz"

	// FabricTypeNVMEOverFabric
	// NVMe over fabric
	FabricTypeNVMEOverFabric FabricType = "nvmeof"

	// FabricTypeIPBased
	// IP based fabric
	FabricTypeIPBased FabricType = "ip_based"
)

// ImageType
// Describes a type of software image file
type ImageType string
const (
	// ImageTypeBoot
	// Boot image
	ImageTypeBoot ImageType = "boot"

	// ImageTypeFPGA
	// FPGA software image
	ImageTypeFPGA ImageType = "fpga"

	// ImageTypeVNIC
	// VNIC image
	ImageTypeVNIC ImageType = "vnic"

	// ImageTypeUpgrade
	// Liqid software upgrade image
	ImageTypeUpgrade ImageType = "upgrade"

	// ImageTypeApplication
	// Application software image
	ImageTypeApplication ImageType = "application"
)

// LogType
// Describes a log file type
type LogType string
const (
	// LogTypeSystem
	// General system logs
	LogTypeSystem LogType = "system"

	// LogTypeLiqid
	// Liqid software component logs
	LogTypeLiqid LogType = "liqid"

	// LogTypeCombined
	// All available logs
	LogTypeCombined LogType = "combined"
)

// ManageableType
// Describes a type of manageable entity
type ManageableType string
const (
	// ManageableTypeDevice
	// Describes a manageable device
	ManageableTypeDevice ManageableType = "ManageableDevice"

	// ManageableTypeCpuIpmiNetwork
	// Describes an IPMI-manageable CPU
	ManageableTypeCpuIpmiNetwork ManageableType = "ManageableCpuIpmiNetworkConfig"

	// ManageableTypeEnclosureIpmiNetwork
	// Describes an IPMI-manageable enclosure
	ManageableTypeEnclosureIpmiNetwork ManageableType = "ManageableDeviceIpmiNetworkConfig"
)

// NorthboundApplicationType
// Describes a northbound application type
type NorthboundApplicationType string
const (
	// NorthboundApplicationTypeSlurm
	// Slurm application
	NorthboundApplicationTypeSlurm NorthboundApplicationType = "slurm"

	// NorthboundApplicationTypeKubernetes
	// Kubernetes application
	NorthboundApplicationTypeKubernetes NorthboundApplicationType = "kubernetes"

	// NorthboundApplicationTypeRedfish
	// Redfish application
	NorthboundApplicationTypeRedfish NorthboundApplicationType = "redfish"

	// NorthboundApplicationTypeVAPI
	// VAPI application
	NorthboundApplicationTypeVAPI NorthboundApplicationType = "vapi"

	// NorthboundApplicationTypeIPMI
	// IPMI application
	NorthboundApplicationTypeIPMI NorthboundApplicationType = "ipmi"
)

// OperatingSystemType
// Describes an operating system
type OperatingSystemType string
const (
	// OperatingSystemTypeLinux
	// Any Linux-derived operating system
	OperatingSystemTypeLinux OperatingSystemType = "linux"

	// OperatingSystemTypeFreeBSD
	// Any FreeBSD distribution
	OperatingSystemTypeFreeBSD OperatingSystemType = "FreeBSD"

	// OperatingSystemTypeWindows
	// Any Microsoft Windows(TM) operating system
	OperatingSystemTypeWindows OperatingSystemType = "windows"

	// OperatingSystemTypeLiqidOS
	// Liqid Inc. operating system
	OperatingSystemTypeLiqidOS OperatingSystemType = "liqidos"

	// OperatingSystemTypeMaxOSX
	// Any darwin-based operating system
	OperatingSystemTypeMaxOSX OperatingSystemType = "Apple Inc. Mac(TM) OSX operating system"

	// OperatingSystemTypeUnknown
	// Operating system is unrecognized or cannot be determined
	OperatingSystemTypeUnknown OperatingSystemType = "unknown"
)

// P2PType
// Indicates whether P2P is enabled
type P2PType string
const (
	// P2PTypeOff
	// Indicates that P2P is not enabled
	P2PTypeOff P2PType = "off"

	// P2PTypeOn
	// Indicates that P2P is enabled
	P2PTypeOn P2PType = "on"
)

// RunType
// Describes the various run states for liqid
type RunType string
const (
	// RunTypeUp
	// Entity is up and running
	RunTypeUp RunType = "up"

	// RunTypeOn
	// Entity is powered-on but not necessarily running
	RunTypeOn RunType = "on"

	// RunTypeOff
	// Entity is not powered-on
	RunTypeOff RunType = "off"

	// RunTypeUnknown
	// Entity state cannot be determined
	RunTypeUnknown RunType = "unknown"
)

// ToggleFlag
// TODO
type ToggleFlag string
const (
	// ToggleFlagPermanent
	// Associated entity is permanent
	ToggleFlagPermanent ToggleFlag = "perm"

	// ToggleFlagDisappear
	// Associated entity should disappear
	ToggleFlagDisappear ToggleFlag = "disappear"

	// ToggleFlagActive
	// Associated entity is active
	ToggleFlagActive ToggleFlag = "active"

	// ToggleFlagHidden
	// Associated entity is hidden
	ToggleFlagHidden ToggleFlag = "hidden"
)

// ToggleState
// Describes a binary state of enablement/disablement
type ToggleState string
const (
	// ToggleStateEnabled
	// Indicates that the associated entity is enabled.
	ToggleStateEnabled ToggleState = "enabled"

	// ToggleStateDisabled
	// Indicates that the associated entity is disabled.
	ToggleStateDisabled ToggleState = "disabled"
)

// TokenType
// A subset of DeviceType values used exclusively with tokens
type TokenType string
const (
	// TokenTypeTarget
	// Indicates a target type
	TokenTypeTarget TokenType = "targ"

	// TokenTypeLink
	// Indicates a network link type
	TokenTypeLink TokenType = "link"

	// TokenTypeGPU
	// Indicates a GPU
	TokenTypeGPU TokenType = "gpu"

	// TokenTypeFPGA
	// Indicates an FPGA device
	TokenTypeFPGA TokenType = "fpga"
)
