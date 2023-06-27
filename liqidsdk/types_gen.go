// File: types_gen.go
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

// AsynchronousInfo
// Reports an identifier of an asynchronous operation
type AsynchronousInfo struct {
	// AsynchronousId
	// Identifier to be used for polling the state of an asynchronous task
	AsynchronousId string `json:"async_id"`

	// DeviceId
	// Device identifier which is associated with the asynchronous task
	DeviceId string `json:"device_id"`
}

// NewAsynchronousInfo initializer for AsynchronousInfo struct
func NewAsynchronousInfo() AsynchronousInfo {
	obj := AsynchronousInfo{}
	return obj
}

// AsynchronousInfoWrapper JSON body wrapper for AsynchronousInfo
type AsynchronousInfoWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []AsynchronousInfo `json:"data"`
	} `json:"response"`
}

// AsynchronousStatus
// Reports the status of an asynchronous operation
type AsynchronousStatus struct {
	// ExecutionState
	// execution state of the asynchronous operation
	ExecutionState string `json:"command_execution_state"`

	// DeviceId
	// Identifier of the device to which this operation applies
	DeviceId string `json:"device_id"`

	// ExecutionDateTime
	// Timestamp of when the process was initiated
	ExecutionDateTime string `json:"execution_date"`
}

// NewAsynchronousStatus initializer for AsynchronousStatus struct
func NewAsynchronousStatus() AsynchronousStatus {
	obj := AsynchronousStatus{}
	return obj
}

// AsynchronousStatusWrapper JSON body wrapper for AsynchronousStatus
type AsynchronousStatusWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []AsynchronousStatus `json:"data"`
	} `json:"response"`
}

// AvailableCoordinates
// A description of an available REST target including IP addressing information and Liqid Coordinates
type AvailableCoordinates struct {
	// IPAddress
	// DNS name or dotted-decimal IP address of the REST target
	IPAddress string `json:"address"`

	// PortNumber
	// UDP port number of the REST target
	PortNumber int32 `json:"port"`

	// Coordinates
	// Liqid coordinates of the REST target
	Coordinates Coordinates `json:"coordinates"`
}

// NewAvailableCoordinates initializer for AvailableCoordinates struct
func NewAvailableCoordinates() AvailableCoordinates {
	obj := AvailableCoordinates{}
	obj.Coordinates = Coordinates{}
	return obj
}

// AvailableCoordinatesWrapper JSON body wrapper for AvailableCoordinates
type AvailableCoordinatesWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []AvailableCoordinates `json:"data"`
	} `json:"response"`
}

// BackupResult
// Backups up the Director configuration
type BackupResult struct {
	// CommandLine
	// The command line used to perform the backup operation
	CommandLine string `json:"commandLine"`

	// StandardOutput
	// Content written to stdout during the backup operation
	StandardOutput string `json:"standardOutput"`

	// StandardError
	// Content written to stderr during the backup operation
	StandardError string `json:"standardError"`

	// ExitValue
	// Resulting value of the backup operation; zero indicates success
	ExitValue int32 `json:"exitValue"`
}

// NewBackupResult initializer for BackupResult struct
func NewBackupResult() BackupResult {
	obj := BackupResult{}
	return obj
}

// BackupResultWrapper JSON body wrapper for BackupResult
type BackupResultWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []BackupResult `json:"data"`
	} `json:"response"`
}

// Coordinates
// Describes a unique Liqid coordinate.
// Most of the members of this entity are obsolete.
type Coordinates struct {
	// Rack
	// Obsolete - should always be zero
	Rack int32 `json:"rack"`

	// Shelf
	// Obsolete - should always be zero
	Shelf int32 `json:"shelf"`

	// Node
	// Describes the relative position of a particular coordinate-addressable Liqid entity.
	Node int32 `json:"node"`
}

// NewCoordinates initializer for Coordinates struct
func NewCoordinates() Coordinates {
	obj := Coordinates{}
	obj.Rack = 0
	obj.Shelf = 0
	return obj
}

// CoordinatesWrapper JSON body wrapper for Coordinates
type CoordinatesWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []Coordinates `json:"data"`
	} `json:"response"`
}

// ComputeDeviceInfo
// Contains non-status information regarding a compute device
type ComputeDeviceInfo struct {
	// BusGeneration
	// Bus Generation
	BusGeneration string `json:"busgen"`

	// BusWidth
	// Bus Width
	BusWidth string `json:"buswidth"`

	// DeviceClass
	// Device class
	DeviceClass string `json:"class"`

	// ConnectionType
	// Connection Type
	ConnectionType string `json:"conn_type"`

	// DeviceIdentifier
	// Device Identifier
	DeviceIdentifier string `json:"dev_id"`

	// DeviceState
	// Device State
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device Type
	DeviceType DeviceType `json:"device_type"`

	// FabricGlobalId
	// Fabric global identifier
	FabricGlobalId string `json:"fabr_gid"`

	// FabricType
	// Fabric Type
	FabricType FabricType `json:"fabric_type"`

	// Family
	// Family
	Family string `json:"family"`

	// Flags
	// Flags
	Flags string `json:"flags"`

	// FirmwareRevision
	// Device firmware revision
	FirmwareRevision string `json:"fw_rev"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// Location
	// Liqid Coordinates for this device
	Location Coordinates `json:"location"`

	// Model
	// Model of this device
	Model string `json:"model"`

	// Name
	// Device name
	Name string `json:"name"`

	// Owner
	// Liqid Coordinates of the device above this one in the hierarchy
	Owner Coordinates `json:"owner"`

	// PartNumber
	// Device name
	PartNumber string `json:"part_num"`

	// PCIDeviceId
	// Vendor-unique device identifier expressed in hex as a '0x'-prefixed string
	PCIDeviceId string `json:"pci_device"`

	// PCIVendorId
	// PCI Vendor identifier expressed in hex as a '0x'-prefixed string
	PCIVendorId string `json:"pci_vendor"`

	// PodId
	// Obsolete value - should always be -1
	PodId int32 `json:"pod_id"`

	// SerialNumber
	// Device serial number
	SerialNumber string `json:"serial_num"`

	// SledId
	// Obsolete
	SledId string `json:"sled_id"`

	// UserDescription
	// User-supplied description
	UserDescription string `json:"udesc"`

	// Unique
	// Device-specific information
	Unique string `json:"unique"`

	// Vendor
	// Vendor name
	Vendor string `json:"vendor"`

	// CoreCount
	// Number of cores for this device
	CoreCount string `json:"core_cnt"`

	// CoreMHz
	// Cycle rate for this device
	CoreMHz float64 `json:"core_mhz"`

	// DRAMSize
	// Dynamic RAM size
	DRAMSize string `json:"dram_size"`

	// DRAMType
	// Dynamic RAM type
	DRAMType string `json:"dram_type"`

	// InstructionSet
	// Instruction set of this device
	InstructionSet string `json:"inst_set"`

	// Socket
	// Socket description of this device
	Socket string `json:"socket"`

	// NumberOfThreads
	// Number of threads supported by this device
	NumberOfThreads int32 `json:"thrd_cnt"`
}

// NewComputeDeviceInfo initializer for ComputeDeviceInfo struct
func NewComputeDeviceInfo() ComputeDeviceInfo {
	obj := ComputeDeviceInfo{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// ComputeDeviceInfoWrapper JSON body wrapper for ComputeDeviceInfo
type ComputeDeviceInfoWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []ComputeDeviceInfo `json:"data"`
	} `json:"response"`
}

// ComputeDeviceStatus
// Compute Device Status Information
type ComputeDeviceStatus struct {
	// ConnectionType
	// Connection type for the device
	ConnectionType string `json:"conn_type"`

	// DeviceId
	// Unique identifier for the device
	DeviceId string `json:"dev_id"`

	// DeviceState
	// State of this device
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device type for this device
	DeviceType DeviceType `json:"device_type"`

	// PCIDeviceId
	// PCI Vendor-Unique Device identifier
	PCIDeviceId string `json:"did"`

	// GlobalId
	// Fabric global identifier
	GlobalId string `json:"fabr_gid"`

	// FabricId
	// Fabric identifier
	FabricId int32 `json:"fabr_id"`

	// FabricType
	// Fabric type
	FabricType string `json:"fabric_type"`

	// Flags
	// Hardware flags displayed as a hex string prefixed by 'ox'
	Flags string `json:"flags"`

	// Flags2
	// Additional hardware flags displayed as a hex string prefixed by 'ox'
	Flags2 string `json:"flags2"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// PCILaneCount
	// Number of PCI lanes for this device
	PCILaneCount int32 `json:"lanes"`

	// Location
	// Liqid coordinates for this component
	Location Coordinates `json:"location"`

	// Name
	// Name of this component
	Name string `json:"name"`

	// Owner
	// Liqid coordinates for the component directly above this in the containing fabric topology
	Owner Coordinates `json:"owner"`

	// PodId
	// Pod identifier - obsolete value which is generally always -1
	PodId int32 `json:"pod_id"`

	// PortGlobalId
	// Port Global Identifier
	PortGlobalId string `json:"port_gid"`

	// SledId
	// Obsolete value
	SledId string `json:"sled_id"`

	// SwitchGlobalId
	// Switch Global Identifier
	SwitchGlobalId string `json:"swit_gid"`

	// DeviceStatusType
	// Type of DeviceStatus entity
	DeviceStatusType string `json:"type"`

	// PCIVendorId
	// Unique PCI Vendor Identifier
	PCIVendorId string `json:"vid"`

	// HConn
	// Internal value
	HConn string `json:"hconn"`

	// Unique
	// Internal value
	Unique string `json:"unique"`
}

// NewComputeDeviceStatus initializer for ComputeDeviceStatus struct
func NewComputeDeviceStatus() ComputeDeviceStatus {
	obj := ComputeDeviceStatus{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// ComputeDeviceStatusWrapper JSON body wrapper for ComputeDeviceStatus
type ComputeDeviceStatusWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []ComputeDeviceStatus `json:"data"`
	} `json:"response"`
}

// ConnectionHistory
// Describes one connection to a machine
type ConnectionHistory struct {
	// AttachTime
	// Time which device was attached
	AttachTime int64 `json:"atime"`

	// DeviceType
	// Type of the connecting device
	DeviceType DeviceType `json:"dev_type"`

	// DetachTime
	// Time which device was detached
	DetachTime int64 `json:"dtime"`

	// FabricGlobalId
	// Fabric global identifier
	FabricGlobalId string `json:"fabr_gid"`

	// Free
	// 
	Free bool `json:"free"`

	// Name
	// 
	Name string `json:"name"`

	// OwnerGlobalId
	// Owner global identifier
	OwnerGlobalId string `json:"owner_gid"`

	// UserDescription
	// User-specified description
	UserDescription string `json:"udesc"`
}

// NewConnectionHistory initializer for ConnectionHistory struct
func NewConnectionHistory() ConnectionHistory {
	obj := ConnectionHistory{}
	return obj
}

// ConnectionHistoryWrapper JSON body wrapper for ConnectionHistory
type ConnectionHistoryWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []ConnectionHistory `json:"data"`
	} `json:"response"`
}

// Credentials
// Contains credentials necessary for managing some managed entity within the configuration (such as via IPMI)
type Credentials struct {
	// Principal
	// Also known as username, user-id, etc
	Principal string `json:"principal"`

	// Password
	// The password associated with the given principal
	Password string `json:"password"`
}

// NewCredentials initializer for Credentials struct
func NewCredentials() Credentials {
	obj := Credentials{}
	return obj
}

// CredentialsWrapper JSON body wrapper for Credentials
type CredentialsWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []Credentials `json:"data"`
	} `json:"response"`
}

// DeviceCounters
// Counts of discovered devices by device type
type DeviceCounters struct {
	// CPUCount
	// Number of discovered compute devices
	CPUCount int32 `json:"comp_cnt"`

	// FPGACount
	// Number of discovered FPGA devices
	FPGACount int32 `json:"fpga_cnt"`

	// GPUCount
	// Number of discovered GPU devices
	GPUCount int32 `json:"gpu_cnt"`

	// LinkCount
	// Number of discovered link (e.g., network) devices
	LinkCount int32 `json:"link_cnt"`

	// PLXCount
	// Number of discovered PLX devices
	PLXCount int32 `json:"plx_cnt"`

	// TargetCount
	// Number of discovered target devices
	TargetCount int32 `json:"targ_cnt"`
}

// NewDeviceCounters initializer for DeviceCounters struct
func NewDeviceCounters() DeviceCounters {
	obj := DeviceCounters{}
	return obj
}

// DeviceCountersWrapper JSON body wrapper for DeviceCounters
type DeviceCountersWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []DeviceCounters `json:"data"`
	} `json:"response"`
}

// DeviceInfo
// All information other than status, for a given device
type DeviceInfo struct {
	// BusGeneration
	// Bus Generation
	BusGeneration string `json:"busgen"`

	// BusWidth
	// Bus Width
	BusWidth string `json:"buswidth"`

	// DeviceClass
	// Device class
	DeviceClass string `json:"class"`

	// ConnectionType
	// Connection Type
	ConnectionType string `json:"conn_type"`

	// DeviceIdentifier
	// Device Identifier
	DeviceIdentifier string `json:"dev_id"`

	// DeviceState
	// Device State
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device Type
	DeviceType DeviceType `json:"device_type"`

	// FabricGlobalId
	// Fabric global identifier
	FabricGlobalId string `json:"fabr_gid"`

	// FabricType
	// Fabric Type
	FabricType FabricType `json:"fabric_type"`

	// Family
	// Family
	Family string `json:"family"`

	// Flags
	// Flags
	Flags string `json:"flags"`

	// FirmwareRevision
	// Device firmware revision
	FirmwareRevision string `json:"fw_rev"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// Location
	// Liqid Coordinates for this device
	Location Coordinates `json:"location"`

	// Model
	// Model of this device
	Model string `json:"model"`

	// Name
	// Device name
	Name string `json:"name"`

	// Owner
	// Liqid Coordinates of the device above this one in the hierarchy
	Owner Coordinates `json:"owner"`

	// PartNumber
	// Device name
	PartNumber string `json:"part_num"`

	// PCIDeviceId
	// Vendor-unique device identifier expressed in hex as a '0x'-prefixed string
	PCIDeviceId string `json:"pci_device"`

	// PCIVendorId
	// PCI Vendor identifier expressed in hex as a '0x'-prefixed string
	PCIVendorId string `json:"pci_vendor"`

	// PodId
	// Obsolete value - should always be -1
	PodId int32 `json:"pod_id"`

	// SerialNumber
	// Device serial number
	SerialNumber string `json:"serial_num"`

	// SledId
	// Obsolete
	SledId string `json:"sled_id"`

	// UserDescription
	// User-supplied description
	UserDescription string `json:"udesc"`

	// Unique
	// Device-specific information
	Unique string `json:"unique"`

	// Vendor
	// Vendor name
	Vendor string `json:"vendor"`
}

// NewDeviceInfo initializer for DeviceInfo struct
func NewDeviceInfo() DeviceInfo {
	obj := DeviceInfo{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// DeviceInfoWrapper JSON body wrapper for DeviceInfo
type DeviceInfoWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []DeviceInfo `json:"data"`
	} `json:"response"`
}

// DeviceStatus
// Status of a manageable device
type DeviceStatus struct {
	// ConnectionType
	// Connection type for the device
	ConnectionType string `json:"conn_type"`

	// DeviceId
	// Unique identifier for the device
	DeviceId string `json:"dev_id"`

	// DeviceState
	// State of this device
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device type for this device
	DeviceType DeviceType `json:"device_type"`

	// PCIDeviceId
	// PCI Vendor-Unique Device identifier
	PCIDeviceId string `json:"did"`

	// GlobalId
	// Fabric global identifier
	GlobalId string `json:"fabr_gid"`

	// FabricId
	// Fabric identifier
	FabricId int32 `json:"fabr_id"`

	// FabricType
	// Fabric type
	FabricType string `json:"fabric_type"`

	// Flags
	// Hardware flags displayed as a hex string prefixed by 'ox'
	Flags string `json:"flags"`

	// Flags2
	// Additional hardware flags displayed as a hex string prefixed by 'ox'
	Flags2 string `json:"flags2"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// PCILaneCount
	// Number of PCI lanes for this device
	PCILaneCount int32 `json:"lanes"`

	// Location
	// Liqid coordinates for this component
	Location Coordinates `json:"location"`

	// Name
	// Name of this component
	Name string `json:"name"`

	// Owner
	// Liqid coordinates for the component directly above this in the containing fabric topology
	Owner Coordinates `json:"owner"`

	// PodId
	// Pod identifier - obsolete value which is generally always -1
	PodId int32 `json:"pod_id"`

	// PortGlobalId
	// Port Global Identifier
	PortGlobalId string `json:"port_gid"`

	// SledId
	// Obsolete value
	SledId string `json:"sled_id"`

	// SwitchGlobalId
	// Switch Global Identifier
	SwitchGlobalId string `json:"swit_gid"`

	// DeviceStatusType
	// Type of DeviceStatus entity
	DeviceStatusType string `json:"type"`

	// PCIVendorId
	// Unique PCI Vendor Identifier
	PCIVendorId string `json:"vid"`
}

// NewDeviceStatus initializer for DeviceStatus struct
func NewDeviceStatus() DeviceStatus {
	obj := DeviceStatus{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// DeviceStatusWrapper JSON body wrapper for DeviceStatus
type DeviceStatusWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []DeviceStatus `json:"data"`
	} `json:"response"`
}

// DeviceTypeAndAttributes
// Describes the various attributes for a particular device type
type DeviceTypeAndAttributes struct {
	// DeviceType
	// A particular data type
	DeviceType DeviceType `json:"deviceType"`

	// Attributes
	// Array of string descriptors
	Attributes map[string][]string `json:"attributes"`
}

// NewDeviceTypeAndAttributes initializer for DeviceTypeAndAttributes struct
func NewDeviceTypeAndAttributes() DeviceTypeAndAttributes {
	obj := DeviceTypeAndAttributes{}
	return obj
}

// DeviceTypeAndAttributesWrapper JSON body wrapper for DeviceTypeAndAttributes
type DeviceTypeAndAttributesWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []DeviceTypeAndAttributes `json:"data"`
	} `json:"response"`
}

// DeviceUserDescription
// Wraps a user-provided description for a particular device
type DeviceUserDescription struct {
	// UserDescription
	// The actual description
	UserDescription string `json:"udesc"`
}

// NewDeviceUserDescription initializer for DeviceUserDescription struct
func NewDeviceUserDescription() DeviceUserDescription {
	obj := DeviceUserDescription{}
	return obj
}

// DeviceUserDescriptionWrapper JSON body wrapper for DeviceUserDescription
type DeviceUserDescriptionWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []DeviceUserDescription `json:"data"`
	} `json:"response"`
}

// DigestInfo
// Describes information related to a newly-created digest message
type DigestInfo struct {
	// DigestId
	// A unique (internal) identifier for this digest message
	DigestId int32 `json:"digest_id"`

	// DigestKey
	// The digest key generated by the director to be associated with the given label
	// This is the value which should be used for authenticating subsequent REST API invocations.
	DigestKey string `json:"digest_key"`

	// DigestLabel
	// The digest label which is associated with the generated digest key.
	// This value should be used for deleting the digest message at the end of a client REST session.
	DigestLabel string `json:"digest_label"`
}

// NewDigestInfo initializer for DigestInfo struct
func NewDigestInfo() DigestInfo {
	obj := DigestInfo{}
	return obj
}

// DigestInfoWrapper JSON body wrapper for DigestInfo
type DigestInfoWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []DigestInfo `json:"data"`
	} `json:"response"`
}

// DiscoveryToken
// Describes a single discovery token
type DiscoveryToken struct {
	// Token
	// The actual token
	Token string `json:"discovery_token"`

	// Index
	// Internal index value for the token
	Index int32 `json:"index"`
}

// NewDiscoveryToken initializer for DiscoveryToken struct
func NewDiscoveryToken() DiscoveryToken {
	obj := DiscoveryToken{}
	return obj
}

// DiscoveryTokenWrapper JSON body wrapper for DiscoveryToken
type DiscoveryTokenWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []DiscoveryToken `json:"data"`
	} `json:"response"`
}

// FabricItem
// Describes a Liqid entity, the aggregate of which comprises the fabric.
type FabricItem struct {
	// Name
	// Name associated with this entity.
	Name string `json:"name"`

	// Id
	// Node identifier of this entity.
	Id int32 `json:"id"`

	// ParentId
	// Node identifier of the entity directly above this node in the fabric hierarchy.
	ParentId int32 `json:"parentId"`

	// DeviceType
	// Describes the particular type of this device.
	DeviceType DeviceType `json:"deviceType"`

	// HardwareComponent
	// Describes the hardware characteristics of this device.
	HardwareComponent HardwareComponent `json:"hardwareComponent"`
}

// NewFabricItem initializer for FabricItem struct
func NewFabricItem() FabricItem {
	obj := FabricItem{}
	obj.HardwareComponent = HardwareComponent{}
	return obj
}

// FabricItemWrapper JSON body wrapper for FabricItem
type FabricItemWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []FabricItem `json:"data"`
	} `json:"response"`
}

// FabricToggleComposite
// Describes the result of a fabric change operation
type FabricToggleComposite struct {
	// Coordinates
	// Describes the LIQID coordinates of the director for the updated fabric
	Coordinates Coordinates `json:"coordinates"`

	// ControlFlag
	// Describes the value added to the fabric
	ControlFlag NameValuePair `json:"flag"`

	// Options
	// Describes the operation which was requested
	Options FabricToggleCompositeOption `json:"options"`
}

// NewFabricToggleComposite initializer for FabricToggleComposite struct
func NewFabricToggleComposite() FabricToggleComposite {
	obj := FabricToggleComposite{}
	obj.Coordinates = Coordinates{}
	obj.ControlFlag = NameValuePair{}
	return obj
}

// FabricToggleCompositeWrapper JSON body wrapper for FabricToggleComposite
type FabricToggleCompositeWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []FabricToggleComposite `json:"data"`
	} `json:"response"`
}

// FPGADeviceInfo
// Contains non-status information regarding an FPGA device
type FPGADeviceInfo struct {
	// BusGeneration
	// Bus Generation
	BusGeneration string `json:"busgen"`

	// BusWidth
	// Bus Width
	BusWidth string `json:"buswidth"`

	// DeviceClass
	// Device class
	DeviceClass string `json:"class"`

	// ConnectionType
	// Connection Type
	ConnectionType string `json:"conn_type"`

	// DeviceIdentifier
	// Device Identifier
	DeviceIdentifier string `json:"dev_id"`

	// DeviceState
	// Device State
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device Type
	DeviceType DeviceType `json:"device_type"`

	// FabricGlobalId
	// Fabric global identifier
	FabricGlobalId string `json:"fabr_gid"`

	// FabricType
	// Fabric Type
	FabricType FabricType `json:"fabric_type"`

	// Family
	// Family
	Family string `json:"family"`

	// Flags
	// Flags
	Flags string `json:"flags"`

	// FirmwareRevision
	// Device firmware revision
	FirmwareRevision string `json:"fw_rev"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// Location
	// Liqid Coordinates for this device
	Location Coordinates `json:"location"`

	// Model
	// Model of this device
	Model string `json:"model"`

	// Name
	// Device name
	Name string `json:"name"`

	// Owner
	// Liqid Coordinates of the device above this one in the hierarchy
	Owner Coordinates `json:"owner"`

	// PartNumber
	// Device name
	PartNumber string `json:"part_num"`

	// PCIDeviceId
	// Vendor-unique device identifier expressed in hex as a '0x'-prefixed string
	PCIDeviceId string `json:"pci_device"`

	// PCIVendorId
	// PCI Vendor identifier expressed in hex as a '0x'-prefixed string
	PCIVendorId string `json:"pci_vendor"`

	// PodId
	// Obsolete value - should always be -1
	PodId int32 `json:"pod_id"`

	// SerialNumber
	// Device serial number
	SerialNumber string `json:"serial_num"`

	// SledId
	// Obsolete
	SledId string `json:"sled_id"`

	// UserDescription
	// User-supplied description
	UserDescription string `json:"udesc"`

	// Unique
	// Device-specific information
	Unique string `json:"unique"`

	// Vendor
	// Vendor name
	Vendor string `json:"vendor"`

	// CoreCount
	// Number of cores for this device
	CoreCount string `json:"core_cnt"`

	// CoreMHz
	// Cycle rate for this device
	CoreMHz float64 `json:"core_mhz"`

	// DRAMSize
	// Dynamic RAM size
	DRAMSize string `json:"dram_size"`

	// DRAMType
	// Dynamic RAM type
	DRAMType string `json:"dram_type"`

	// FPGASpeed
	// Speed for this device
	FPGASpeed string `json:"fpga_speed"`

	// NumberOfThreads
	// Number of threads supported by this device
	NumberOfThreads int32 `json:"thrd_cnt"`
}

// NewFPGADeviceInfo initializer for FPGADeviceInfo struct
func NewFPGADeviceInfo() FPGADeviceInfo {
	obj := FPGADeviceInfo{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// FPGADeviceInfoWrapper JSON body wrapper for FPGADeviceInfo
type FPGADeviceInfoWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []FPGADeviceInfo `json:"data"`
	} `json:"response"`
}

// FPGADeviceStatus
// FPGA Device Status Information
type FPGADeviceStatus struct {
	// ConnectionType
	// Connection type for the device
	ConnectionType string `json:"conn_type"`

	// DeviceId
	// Unique identifier for the device
	DeviceId string `json:"dev_id"`

	// DeviceState
	// State of this device
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device type for this device
	DeviceType DeviceType `json:"device_type"`

	// PCIDeviceId
	// PCI Vendor-Unique Device identifier
	PCIDeviceId string `json:"did"`

	// GlobalId
	// Fabric global identifier
	GlobalId string `json:"fabr_gid"`

	// FabricId
	// Fabric identifier
	FabricId int32 `json:"fabr_id"`

	// FabricType
	// Fabric type
	FabricType string `json:"fabric_type"`

	// Flags
	// Hardware flags displayed as a hex string prefixed by 'ox'
	Flags string `json:"flags"`

	// Flags2
	// Additional hardware flags displayed as a hex string prefixed by 'ox'
	Flags2 string `json:"flags2"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// PCILaneCount
	// Number of PCI lanes for this device
	PCILaneCount int32 `json:"lanes"`

	// Location
	// Liqid coordinates for this component
	Location Coordinates `json:"location"`

	// Name
	// Name of this component
	Name string `json:"name"`

	// Owner
	// Liqid coordinates for the component directly above this in the containing fabric topology
	Owner Coordinates `json:"owner"`

	// PodId
	// Pod identifier - obsolete value which is generally always -1
	PodId int32 `json:"pod_id"`

	// PortGlobalId
	// Port Global Identifier
	PortGlobalId string `json:"port_gid"`

	// SledId
	// Obsolete value
	SledId string `json:"sled_id"`

	// SwitchGlobalId
	// Switch Global Identifier
	SwitchGlobalId string `json:"swit_gid"`

	// DeviceStatusType
	// Type of DeviceStatus entity
	DeviceStatusType string `json:"type"`

	// PCIVendorId
	// Unique PCI Vendor Identifier
	PCIVendorId string `json:"vid"`

	// Unique
	// Internal value
	Unique string `json:"unique"`
}

// NewFPGADeviceStatus initializer for FPGADeviceStatus struct
func NewFPGADeviceStatus() FPGADeviceStatus {
	obj := FPGADeviceStatus{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// FPGADeviceStatusWrapper JSON body wrapper for FPGADeviceStatus
type FPGADeviceStatusWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []FPGADeviceStatus `json:"data"`
	} `json:"response"`
}

// GPUDeviceInfo
// Contains non-status information regarding a GPU device
type GPUDeviceInfo struct {
	// BusGeneration
	// Bus Generation
	BusGeneration string `json:"busgen"`

	// BusWidth
	// Bus Width
	BusWidth string `json:"buswidth"`

	// DeviceClass
	// Device class
	DeviceClass string `json:"class"`

	// ConnectionType
	// Connection Type
	ConnectionType string `json:"conn_type"`

	// DeviceIdentifier
	// Device Identifier
	DeviceIdentifier string `json:"dev_id"`

	// DeviceState
	// Device State
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device Type
	DeviceType DeviceType `json:"device_type"`

	// FabricGlobalId
	// Fabric global identifier
	FabricGlobalId string `json:"fabr_gid"`

	// FabricType
	// Fabric Type
	FabricType FabricType `json:"fabric_type"`

	// Family
	// Family
	Family string `json:"family"`

	// Flags
	// Flags
	Flags string `json:"flags"`

	// FirmwareRevision
	// Device firmware revision
	FirmwareRevision string `json:"fw_rev"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// Location
	// Liqid Coordinates for this device
	Location Coordinates `json:"location"`

	// Model
	// Model of this device
	Model string `json:"model"`

	// Name
	// Device name
	Name string `json:"name"`

	// Owner
	// Liqid Coordinates of the device above this one in the hierarchy
	Owner Coordinates `json:"owner"`

	// PartNumber
	// Device name
	PartNumber string `json:"part_num"`

	// PCIDeviceId
	// Vendor-unique device identifier expressed in hex as a '0x'-prefixed string
	PCIDeviceId string `json:"pci_device"`

	// PCIVendorId
	// PCI Vendor identifier expressed in hex as a '0x'-prefixed string
	PCIVendorId string `json:"pci_vendor"`

	// PodId
	// Obsolete value - should always be -1
	PodId int32 `json:"pod_id"`

	// SerialNumber
	// Device serial number
	SerialNumber string `json:"serial_num"`

	// SledId
	// Obsolete
	SledId string `json:"sled_id"`

	// UserDescription
	// User-supplied description
	UserDescription string `json:"udesc"`

	// Unique
	// Device-specific information
	Unique string `json:"unique"`

	// Vendor
	// Vendor name
	Vendor string `json:"vendor"`

	// CacheSize
	// Cache size for this device
	CacheSize string `json:"cache_size"`

	// CoreCount
	// Number of cores for this device
	CoreCount string `json:"core_cnt"`

	// CoreSpeed
	// Core speed for this device
	CoreSpeed string `json:"core_speed"`

	// DRAMSize
	// Dynamic RAM size
	DRAMSize string `json:"dram_size"`

	// DRAMType
	// Dynamic RAM type
	DRAMType string `json:"dram_type"`
}

// NewGPUDeviceInfo initializer for GPUDeviceInfo struct
func NewGPUDeviceInfo() GPUDeviceInfo {
	obj := GPUDeviceInfo{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// GPUDeviceInfoWrapper JSON body wrapper for GPUDeviceInfo
type GPUDeviceInfoWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []GPUDeviceInfo `json:"data"`
	} `json:"response"`
}

// GPUDeviceStatus
// GPU Device Status Information
type GPUDeviceStatus struct {
	// ConnectionType
	// Connection type for the device
	ConnectionType string `json:"conn_type"`

	// DeviceId
	// Unique identifier for the device
	DeviceId string `json:"dev_id"`

	// DeviceState
	// State of this device
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device type for this device
	DeviceType DeviceType `json:"device_type"`

	// PCIDeviceId
	// PCI Vendor-Unique Device identifier
	PCIDeviceId string `json:"did"`

	// GlobalId
	// Fabric global identifier
	GlobalId string `json:"fabr_gid"`

	// FabricId
	// Fabric identifier
	FabricId int32 `json:"fabr_id"`

	// FabricType
	// Fabric type
	FabricType string `json:"fabric_type"`

	// Flags
	// Hardware flags displayed as a hex string prefixed by 'ox'
	Flags string `json:"flags"`

	// Flags2
	// Additional hardware flags displayed as a hex string prefixed by 'ox'
	Flags2 string `json:"flags2"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// PCILaneCount
	// Number of PCI lanes for this device
	PCILaneCount int32 `json:"lanes"`

	// Location
	// Liqid coordinates for this component
	Location Coordinates `json:"location"`

	// Name
	// Name of this component
	Name string `json:"name"`

	// Owner
	// Liqid coordinates for the component directly above this in the containing fabric topology
	Owner Coordinates `json:"owner"`

	// PodId
	// Pod identifier - obsolete value which is generally always -1
	PodId int32 `json:"pod_id"`

	// PortGlobalId
	// Port Global Identifier
	PortGlobalId string `json:"port_gid"`

	// SledId
	// Obsolete value
	SledId string `json:"sled_id"`

	// SwitchGlobalId
	// Switch Global Identifier
	SwitchGlobalId string `json:"swit_gid"`

	// DeviceStatusType
	// Type of DeviceStatus entity
	DeviceStatusType string `json:"type"`

	// PCIVendorId
	// Unique PCI Vendor Identifier
	PCIVendorId string `json:"vid"`

	// Unique
	// Internal value
	Unique string `json:"unique"`
}

// NewGPUDeviceStatus initializer for GPUDeviceStatus struct
func NewGPUDeviceStatus() GPUDeviceStatus {
	obj := GPUDeviceStatus{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// GPUDeviceStatusWrapper JSON body wrapper for GPUDeviceStatus
type GPUDeviceStatusWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []GPUDeviceStatus `json:"data"`
	} `json:"response"`
}

// Group
// Describes a configured group which contains a free device pool and some number of configured machines.
// This struct does not contain information regarding the related entities; that information must be obtained via other functions/methods.
type Group struct {
	// FabricId
	// The identifier of the fabric to which this group belongs
	FabricId int32 `json:"fabr_id"`

	// GroupId
	// The unique (among the contained fabric) identifier of this group
	GroupId int32 `json:"grp_id"`

	// GroupName
	// The unique (among the contained fabric) name of this group
	GroupName string `json:"group_name"`

	// PodId
	// Obsolete - should always be -1
	PodId int32 `json:"pod_id"`
}

// NewGroup initializer for Group struct
func NewGroup() Group {
	obj := Group{}
	obj.PodId = -1
	return obj
}

// GroupWrapper JSON body wrapper for Group
type GroupWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []Group `json:"data"`
	} `json:"response"`
}

// GroupComputeDeviceRelator
// Describes a relation between a group and a compute device
type GroupComputeDeviceRelator struct {
	// DeviceStatus
	// A ComputeDeviceStatus entity for the device in the relation
	DeviceStatus ComputeDeviceStatus `json:"deviceStatus"`

	// Group
	// Group entity for the group in the relation
	Group Group `json:"group"`
}

// NewGroupComputeDeviceRelator initializer for GroupComputeDeviceRelator struct
func NewGroupComputeDeviceRelator() GroupComputeDeviceRelator {
	obj := GroupComputeDeviceRelator{}
	obj.DeviceStatus = ComputeDeviceStatus{}
	obj.Group = Group{}
	return obj
}

// GroupComputeDeviceRelatorWrapper JSON body wrapper for GroupComputeDeviceRelator
type GroupComputeDeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []GroupComputeDeviceRelator `json:"data"`
	} `json:"response"`
}

// GroupDetails
// Contains statistical information which is accumulated for the group.
// Does not contain relations with devices; that information resides in the various device relators.
type GroupDetails struct {
	// GroupId
	// System-unique identifier for the group
	GroupId int32 `json:"grp_id"`

	// GroupName
	// System-unique human-readable name of the group
	GroupName string `json:"group_name"`

	// CpuFrequency
	// Frequency of the CPUs in the group
	CpuFrequency float64 `json:"cpu-frequency"`

	// CpuCount
	// Number of CPUs in the group
	CpuCount int32 `json:"cpu-count"`

	// CpuLanes
	// Number of PCI lanes dedicated to CPUs for this group
	CpuLanes int32 `json:"cpu-lanes"`

	// CpuCoreCount
	// Number of CPU cores in the group
	CpuCoreCount int32 `json:"cpu-core-count"`

	// TotalDRAM
	// Total amount of dynamic RAM in the group
	TotalDRAM int32 `json:"total-dram"`

	// NetworkAdapterCount
	// Number of network adapters in the group
	NetworkAdapterCount int32 `json:"network-adapter-count"`

	// TotalThroughput
	// Total throughput for this group
	TotalThroughput string `json:"total-throughput"`

	// StorageDriveCount
	// Number of SSDs in the group
	StorageDriveCount int32 `json:"storage-drive-count"`

	// TotalCapacity
	// Total capacity of SSD storage in the group
	TotalCapacity int64 `json:"total-capacity"`

	// GPUCores
	// Number of GPU cores in the group
	GPUCores int32 `json:"gpu-cores"`

	// TotalMachines
	// Number of configured machines in the group
	TotalMachines int32 `json:"total-machines"`
}

// NewGroupDetails initializer for GroupDetails struct
func NewGroupDetails() GroupDetails {
	obj := GroupDetails{}
	return obj
}

// GroupDetailsWrapper JSON body wrapper for GroupDetails
type GroupDetailsWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []GroupDetails `json:"data"`
	} `json:"response"`
}

// GroupFPGADeviceRelator
// Describes a relation between a group and an FPGA device
type GroupFPGADeviceRelator struct {
	// DeviceStatus
	// FPGADeviceStatus entity for the device in the relation
	DeviceStatus FPGADeviceStatus `json:"deviceStatus"`

	// Group
	// Group entity for the group in the relation
	Group Group `json:"group"`
}

// NewGroupFPGADeviceRelator initializer for GroupFPGADeviceRelator struct
func NewGroupFPGADeviceRelator() GroupFPGADeviceRelator {
	obj := GroupFPGADeviceRelator{}
	obj.DeviceStatus = FPGADeviceStatus{}
	obj.Group = Group{}
	return obj
}

// GroupFPGADeviceRelatorWrapper JSON body wrapper for GroupFPGADeviceRelator
type GroupFPGADeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []GroupFPGADeviceRelator `json:"data"`
	} `json:"response"`
}

// GroupGPUDeviceRelator
// Describes a relation between a group and an GPU device
type GroupGPUDeviceRelator struct {
	// DeviceStatus
	// GPUDeviceStatus entity for the device in the relation
	DeviceStatus GPUDeviceStatus `json:"deviceStatus"`

	// Group
	// Group entity for the group in the relation
	Group Group `json:"group"`
}

// NewGroupGPUDeviceRelator initializer for GroupGPUDeviceRelator struct
func NewGroupGPUDeviceRelator() GroupGPUDeviceRelator {
	obj := GroupGPUDeviceRelator{}
	obj.DeviceStatus = GPUDeviceStatus{}
	obj.Group = Group{}
	return obj
}

// GroupGPUDeviceRelatorWrapper JSON body wrapper for GroupGPUDeviceRelator
type GroupGPUDeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []GroupGPUDeviceRelator `json:"data"`
	} `json:"response"`
}

// GroupPool
// Describes a group pool
type GroupPool struct {
	// Coordinates
	// Liqid coordinates for the switch which controls this group pool
	Coordinates Coordinates `json:"coordinates"`

	// FabricId
	// Fabric identifier for the fabric which contains the group
	FabricId int32 `json:"fabr_id"`

	// GroupId
	// Unique identifier of the group
	GroupId int32 `json:"grp_id"`
}

// NewGroupPool initializer for GroupPool struct
func NewGroupPool() GroupPool {
	obj := GroupPool{}
	obj.Coordinates = Coordinates{}
	return obj
}

// GroupPoolWrapper JSON body wrapper for GroupPool
type GroupPoolWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []GroupPool `json:"data"`
	} `json:"response"`
}

// GroupMemoryDeviceRelator
// Describes a relation between a group and a memory device
type GroupMemoryDeviceRelator struct {
	// DeviceStatus
	// MemoryDeviceStatus entity for the device in the relation
	DeviceStatus MemoryDeviceStatus `json:"deviceStatus"`

	// Group
	// Group entity for the group in the relation
	Group Group `json:"group"`
}

// NewGroupMemoryDeviceRelator initializer for GroupMemoryDeviceRelator struct
func NewGroupMemoryDeviceRelator() GroupMemoryDeviceRelator {
	obj := GroupMemoryDeviceRelator{}
	obj.DeviceStatus = MemoryDeviceStatus{}
	obj.Group = Group{}
	return obj
}

// GroupMemoryDeviceRelatorWrapper JSON body wrapper for GroupMemoryDeviceRelator
type GroupMemoryDeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []GroupMemoryDeviceRelator `json:"data"`
	} `json:"response"`
}

// GroupNetworkDeviceRelator
// Describes a relation between a group and a network device
type GroupNetworkDeviceRelator struct {
	// DeviceStatus
	// NetworkDeviceStatus entity for the device in the relation
	DeviceStatus NetworkDeviceStatus `json:"deviceStatus"`

	// Group
	// Group entity for the group in the relation
	Group Group `json:"group"`
}

// NewGroupNetworkDeviceRelator initializer for GroupNetworkDeviceRelator struct
func NewGroupNetworkDeviceRelator() GroupNetworkDeviceRelator {
	obj := GroupNetworkDeviceRelator{}
	obj.DeviceStatus = NetworkDeviceStatus{}
	obj.Group = Group{}
	return obj
}

// GroupNetworkDeviceRelatorWrapper JSON body wrapper for GroupNetworkDeviceRelator
type GroupNetworkDeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []GroupNetworkDeviceRelator `json:"data"`
	} `json:"response"`
}

// GroupStorageDeviceRelator
// Describes a relation between a group and a storage device
type GroupStorageDeviceRelator struct {
	// DeviceStatus
	// StorageDeviceStatus entity for the device in the relation
	DeviceStatus StorageDeviceStatus `json:"deviceStatus"`

	// Group
	// Group entity for the group in the relation
	Group Group `json:"group"`
}

// NewGroupStorageDeviceRelator initializer for GroupStorageDeviceRelator struct
func NewGroupStorageDeviceRelator() GroupStorageDeviceRelator {
	obj := GroupStorageDeviceRelator{}
	obj.DeviceStatus = StorageDeviceStatus{}
	obj.Group = Group{}
	return obj
}

// GroupStorageDeviceRelatorWrapper JSON body wrapper for GroupStorageDeviceRelator
type GroupStorageDeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []GroupStorageDeviceRelator `json:"data"`
	} `json:"response"`
}

// HardwareComponent
// Describes the hardware details of a component
type HardwareComponent struct {
	// Type
	// Hardware type
	Type string `json:"type"`

	// Name
	// Hardware name
	Name string `json:"name"`

	// Flags
	// Hardware flags displayed as a hex string prefixed by 'ox'
	Flags string `json:"flags"`

	// State
	// Hardware state
	State string `json:"state"`

	// FabricId
	// Fabric identifier of the fabric to which this entity is connected.
	FabricId int32 `json:"fabr_id"`

	// FabricGlobalId
	// Fabric-unique global identifier to which this entity is connected.
	FabricGlobalId string `json:"fabr_gid"`

	// SwitchGlobalId
	// TODO
	SwitchGlobalId string `json:"swit_gid"`

	// PCIVendorId
	// Unique Vendor identifier
	PCIVendorId string `json:"vid"`

	// PCIDeviceId
	// Vendor-unique device identifier
	PCIDeviceId string `json:"did"`

	// Revision
	// Hardware revision string
	Revision string `json:"rev"`

	// PortCount
	// Number of ports for this component
	PortCount int32 `json:"port_cnt"`

	// Ports
	// Descriptions of the component ports
	Ports []Port `json:"ports"`

	// PCILaneCount
	// Number of PCI lanes for this component
	PCILaneCount int32 `json:"lanes"`

	// ReceiverErrorCount
	// Running count of errors received for this component
	ReceiverErrorCount int32 `json:"rcv_errs"`

	// BadTLPCount
	// Running count of bad TLPs for this component
	BadTLPCount int32 `json:"bad_tlp"`

	// BadDLLPCount
	// Running count of bad DLLPs for this component
	BadDLLPCount int32 `json:"bad_dllp"`

	// ErrorCount
	// Running count of all errors for this component
	ErrorCount int32 `json:"errs"`

	// IngressBytes
	// Running count of incoming bytes
	IngressBytes int32 `json:"ibytes"`

	// EgressBytes
	// Running count of outgoing bytes
	EgressBytes int32 `json:"ebytes"`

	// Location
	// Liqid coordinates for this component
	Location Coordinates `json:"location"`

	// Owner
	// Liqid coordinates for the component directly above this in the containing fabric topology
	Owner Coordinates `json:"owner"`
}

// NewHardwareComponent initializer for HardwareComponent struct
func NewHardwareComponent() HardwareComponent {
	obj := HardwareComponent{}
	obj.Ports = []Port{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	return obj
}

// HardwareComponentWrapper JSON body wrapper for HardwareComponent
type HardwareComponentWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []HardwareComponent `json:"data"`
	} `json:"response"`
}

// Machine
// Describes a configured machine
type Machine struct {
	// Index
	// Internal value
	Index int32 `json:"index"`

	// MachineId
	// Unique identifier for this particular machine
	MachineId int32 `json:"mach_id"`

	// GroupId
	// Unique identifier of the group to which this machine belongs
	GroupId int32 `json:"grp_id"`

	// FabricId
	// Unique identifier of the fabric to which this machine belongs
	FabricId int32 `json:"fabr_id"`

	// FabricGlobalId
	// Fabric global identifier expressed in hexadecimal
	FabricGlobalId string `json:"fabr_gid"`

	// PortGlobalId
	// Port global identifier expressed in hexadecimal
	PortGlobalId string `json:"port_gid"`

	// SwitchGlobalId
	// Switch global identifier expressed in hexadecimal
	SwitchGlobalId string `json:"swit_gid"`

	// ComputeName
	// Name of the compute device associated with this machine
	ComputeName string `json:"comp_name"`

	// MachineName
	// Name of this machine
	MachineName string `json:"mach_name"`

	// P2PEnabled
	// Name of this machine
	P2PEnabled P2PType `json:"p2p"`

	// CreatedTimestamp
	// Date and time that this machine was created
	CreatedTimestamp int64 `json:"mtime"`

	// ConnectionHistory
	// Connection history for this machine
	// Expressed as an array of ConnectionHistory entities
	ConnectionHistory []ConnectionHistory `json:"connection_history"`
}

// NewMachine initializer for Machine struct
func NewMachine() Machine {
	obj := Machine{}
	obj.P2PEnabled = P2PTypeOff
	obj.ConnectionHistory = []ConnectionHistory{}
	return obj
}

// MachineWrapper JSON body wrapper for Machine
type MachineWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []Machine `json:"data"`
	} `json:"response"`
}

// MachineComputeDeviceRelator
// Describes a relation between a machine and a compute device
type MachineComputeDeviceRelator struct {
	// GroupDeviceRelator
	// A GroupComputeDeviceRelator entity for the device in the relation
	GroupDeviceRelator GroupComputeDeviceRelator `json:"groupDeviceRelator"`

	// Machine
	// Machine entity for the machine in the relation
	Machine Machine `json:"machine"`
}

// NewMachineComputeDeviceRelator initializer for MachineComputeDeviceRelator struct
func NewMachineComputeDeviceRelator() MachineComputeDeviceRelator {
	obj := MachineComputeDeviceRelator{}
	obj.GroupDeviceRelator = GroupComputeDeviceRelator{}
	obj.Machine = Machine{}
	return obj
}

// MachineComputeDeviceRelatorWrapper JSON body wrapper for MachineComputeDeviceRelator
type MachineComputeDeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []MachineComputeDeviceRelator `json:"data"`
	} `json:"response"`
}

// MachineDetails
// Additional details for a particular machine
type MachineDetails struct {
	// CPUSocketsField
	// TODO
	CPUSocketsField string `json:"cpuSocketsField"`

	// MachineId
	// Unique machine identifier
	MachineId int32 `json:"mach_id"`

	// MachineName
	// Machine name
	MachineName string `json:"mach_name"`

	// CPUThreadCount
	// Number of CPU threads for the machine
	CPUThreadCount int32 `json:"cpu-threads"`

	// CPUFrequency
	// Cycle frequency for the CPU in this machine
	CPUFrequency string `json:"cpu-frequency"`

	// CPUCoreCount
	// Number of CPU cores for the machine
	CPUCoreCount int32 `json:"cpu-cores"`

	// CPUSockets
	// Describes the CPU sockets for this machine
	CPUSockets string `json:"cpu-sockets"`

	// DynamicRAM
	// Describes the dynamic RAM for this machine
	DynamicRAM string `json:"dram-memory"`

	// FabricConnect
	// TODO
	FabricConnect string `json:"fabric-connect"`

	// NetworkAdapterCount
	// Number of network adapters in the machine
	NetworkAdapterCount int32 `json:"network-adapter-count"`

	// TotalThroughput
	// TODO
	TotalThroughput string `json:"total-throughput"`

	// StorageDriveCount
	// Number of storage drives in the machine
	StorageDriveCount int32 `json:"storage-drive-count"`

	// TotalCapacity
	// Total capacity in this machine
	TotalCapacity int64 `json:"total-capacity"`

	// GPUCount
	// Number of GPUs connected to the machine
	GPUCount int32 `json:"gpu-count"`

	// GPUCoreCount
	// Number of GPU cores for the machine
	GPUCoreCount int32 `json:"gpu-cores"`

	// OperatingSystem
	// Operating system name
	OperatingSystem string `json:"os_name"`

	// BootImage
	// Description of the boot image
	BootImage string `json:"boot_image"`

	// BootDevice
	// Boot device
	BootDevice string `json:"boot_device"`

	// IPAddress
	// IP Address (or DNS name) of the machine
	IPAddress string `json:"ip_address"`

	// IPMIAddress
	// IP Address (or DNS name) of the IPMI management port for the machine
	IPMIAddress string `json:"ipmi_address"`

	// FPGACount
	// Number of FPGAs connected to the machine
	FPGACount int32 `json:"fpga-count"`

	// FPGASpeed
	// FPGA speed
	FPGASpeed string `json:"fpga-speed"`

	// FPGADRAMSize
	// FPGA dynamic RAM size
	FPGADRAMSize int32 `json:"fpga-dram-size"`

	// CreatedAt
	// Timestamp for when the machine was created
	CreatedAt int64 `json:"created"`

	// LastModifiedAt
	// Timestamp for when the machine was last modified
	LastModifiedAt int64 `json:"modified"`
}

// NewMachineDetails initializer for MachineDetails struct
func NewMachineDetails() MachineDetails {
	obj := MachineDetails{}
	return obj
}

// MachineDetailsWrapper JSON body wrapper for MachineDetails
type MachineDetailsWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []MachineDetails `json:"data"`
	} `json:"response"`
}

// MachineFPGADeviceRelator
// Describes a relation between a machine and a FPGA device
type MachineFPGADeviceRelator struct {
	// GroupDeviceRelator
	// A GroupFPGADeviceRelator entity for the device in the relation
	GroupDeviceRelator GroupFPGADeviceRelator `json:"groupDeviceRelator"`

	// Machine
	// Machine entity for the machine in the relation
	Machine Machine `json:"machine"`
}

// NewMachineFPGADeviceRelator initializer for MachineFPGADeviceRelator struct
func NewMachineFPGADeviceRelator() MachineFPGADeviceRelator {
	obj := MachineFPGADeviceRelator{}
	obj.GroupDeviceRelator = GroupFPGADeviceRelator{}
	obj.Machine = Machine{}
	return obj
}

// MachineFPGADeviceRelatorWrapper JSON body wrapper for MachineFPGADeviceRelator
type MachineFPGADeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []MachineFPGADeviceRelator `json:"data"`
	} `json:"response"`
}

// MachineGPUDeviceRelator
// Describes a relation between a machine and a GPU device
type MachineGPUDeviceRelator struct {
	// GroupDeviceRelator
	// A GroupGPUDeviceRelator entity for the device in the relation
	GroupDeviceRelator GroupGPUDeviceRelator `json:"groupDeviceRelator"`

	// Machine
	// Machine entity for the machine in the relation
	Machine Machine `json:"machine"`
}

// NewMachineGPUDeviceRelator initializer for MachineGPUDeviceRelator struct
func NewMachineGPUDeviceRelator() MachineGPUDeviceRelator {
	obj := MachineGPUDeviceRelator{}
	obj.GroupDeviceRelator = GroupGPUDeviceRelator{}
	obj.Machine = Machine{}
	return obj
}

// MachineGPUDeviceRelatorWrapper JSON body wrapper for MachineGPUDeviceRelator
type MachineGPUDeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []MachineGPUDeviceRelator `json:"data"`
	} `json:"response"`
}

// MachineMemoryDeviceRelator
// Describes a relation between a machine and a memory device
type MachineMemoryDeviceRelator struct {
	// GroupDeviceRelator
	// A GroupMemoryDeviceRelator entity for the device in the relation
	GroupDeviceRelator GroupMemoryDeviceRelator `json:"groupDeviceRelator"`

	// Machine
	// Machine entity for the machine in the relation
	Machine Machine `json:"machine"`
}

// NewMachineMemoryDeviceRelator initializer for MachineMemoryDeviceRelator struct
func NewMachineMemoryDeviceRelator() MachineMemoryDeviceRelator {
	obj := MachineMemoryDeviceRelator{}
	obj.GroupDeviceRelator = GroupMemoryDeviceRelator{}
	obj.Machine = Machine{}
	return obj
}

// MachineMemoryDeviceRelatorWrapper JSON body wrapper for MachineMemoryDeviceRelator
type MachineMemoryDeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []MachineMemoryDeviceRelator `json:"data"`
	} `json:"response"`
}

// MachineNetworkDeviceRelator
// Describes a relation between a machine and a network device
type MachineNetworkDeviceRelator struct {
	// GroupDeviceRelator
	// A GroupNetworkDeviceRelator entity for the device in the relation
	GroupDeviceRelator GroupNetworkDeviceRelator `json:"groupDeviceRelator"`

	// Machine
	// Machine entity for the machine in the relation
	Machine Machine `json:"machine"`
}

// NewMachineNetworkDeviceRelator initializer for MachineNetworkDeviceRelator struct
func NewMachineNetworkDeviceRelator() MachineNetworkDeviceRelator {
	obj := MachineNetworkDeviceRelator{}
	obj.GroupDeviceRelator = GroupNetworkDeviceRelator{}
	obj.Machine = Machine{}
	return obj
}

// MachineNetworkDeviceRelatorWrapper JSON body wrapper for MachineNetworkDeviceRelator
type MachineNetworkDeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []MachineNetworkDeviceRelator `json:"data"`
	} `json:"response"`
}

// MachineStorageDeviceRelator
// Describes a relation between a machine and a storage device
type MachineStorageDeviceRelator struct {
	// GroupDeviceRelator
	// A GroupStorageDeviceRelator entity for the device in the relation
	GroupDeviceRelator GroupStorageDeviceRelator `json:"groupDeviceRelator"`

	// Machine
	// Machine entity for the machine in the relation
	Machine Machine `json:"machine"`
}

// NewMachineStorageDeviceRelator initializer for MachineStorageDeviceRelator struct
func NewMachineStorageDeviceRelator() MachineStorageDeviceRelator {
	obj := MachineStorageDeviceRelator{}
	obj.GroupDeviceRelator = GroupStorageDeviceRelator{}
	obj.Machine = Machine{}
	return obj
}

// MachineStorageDeviceRelatorWrapper JSON body wrapper for MachineStorageDeviceRelator
type MachineStorageDeviceRelatorWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []MachineStorageDeviceRelator `json:"data"`
	} `json:"response"`
}

// ManagedEntity
// Describes information regarding a particular vendor and model of a manageable device
type ManagedEntity struct {
	// Description
	// Name of the managed entity
	Description string `json:"description"`

	// PCIVendorId
	// Unique identity of the PCI vendor
	PCIVendorId string `json:"vid"`

	// PCIDeviceId
	// Vendor-unique identity of the PCI device
	PCIDeviceId string `json:"did"`

	// Model
	// Describes the model of the device
	Model string `json:"model"`

	// NumberOfCores
	// Number of cores for CPU entities
	NumberOfCores int32 `json:"cores"`

	// NumberOfThreads
	// Number of threads for CPU entities
	NumberOfThreads int32 `json:"threads"`

	// Speed
	// Cycle speed of the entity
	Speed int32 `json:"speed"`

	// Capacity
	// Capacity of the entity
	Capacity int32 `json:"capacity"`

	// SRIOVEnabled
	// Indicates whether Single-Root IO Virtualization (SRIOV) is enabled for this entity
	SRIOVEnabled bool `json:"sriov"`

	// ManagedEntityType
	// Indicates the type of this managed entity
	ManagedEntityType string `json:"type"`

	// Unique
	// TODO
	Unique string `json:"unique"`

	// CoreMHZ
	// Cycle speed of cpu entity
	CoreMHZ int32 `json:"core_mhz"`

	// DRAMSize
	// Size of Dynamic RAM
	DRAMSize int32 `json:"dram_size"`

	// DRAMType
	// Type of Dynamic RAM
	DRAMType string `json:"dram_type"`

	// Manufacturer
	// Manufacturer/vendor name
	Manufacturer string `json:"device_manufacturer"`

	// DeviceType
	// Type of the device
	DeviceType DeviceType `json:"device_type"`

	// DiscoveryToken
	// A portion of the PCI device identification string which can be used to determine device type
	DiscoveryToken string `json:"discovery_token"`

	// CompanionDevice
	// Known values are 'no' and 'multi'
	CompanionDevice string `json:"companion_device"`

	// EntryDescription
	// Describes the multi-variate state of this entry
	EntryDescription ManagedEntityState `json:"entry_description"`
}

// NewManagedEntity initializer for ManagedEntity struct
func NewManagedEntity() ManagedEntity {
	obj := ManagedEntity{}
	obj.EntryDescription = ManagedEntityState{}
	return obj
}

// ManagedEntityWrapper JSON body wrapper for ManagedEntity
type ManagedEntityWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []ManagedEntity `json:"data"`
	} `json:"response"`
}

// ManagedEntityState
// Describes the state of a managed entity entry
type ManagedEntityState struct {
	// Active
	// Indicates whether the entity entry is active
	Active bool `json:"active"`

	// Required
	// Indicates whether the entity entry is required
	Required bool `json:"required"`
}

// NewManagedEntityState initializer for ManagedEntityState struct
func NewManagedEntityState() ManagedEntityState {
	obj := ManagedEntityState{}
	return obj
}

// ManagedEntityStateWrapper JSON body wrapper for ManagedEntityState
type ManagedEntityStateWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []ManagedEntityState `json:"data"`
	} `json:"response"`
}

// MemoryDeviceInfo
// Contains non-status information regarding a memory device
type MemoryDeviceInfo struct {
	// BusGeneration
	// Bus Generation
	BusGeneration string `json:"busgen"`

	// BusWidth
	// Bus Width
	BusWidth string `json:"buswidth"`

	// DeviceClass
	// Device class
	DeviceClass string `json:"class"`

	// ConnectionType
	// Connection Type
	ConnectionType string `json:"conn_type"`

	// DeviceIdentifier
	// Device Identifier
	DeviceIdentifier string `json:"dev_id"`

	// DeviceState
	// Device State
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device Type
	DeviceType DeviceType `json:"device_type"`

	// FabricGlobalId
	// Fabric global identifier
	FabricGlobalId string `json:"fabr_gid"`

	// FabricType
	// Fabric Type
	FabricType FabricType `json:"fabric_type"`

	// Family
	// Family
	Family string `json:"family"`

	// Flags
	// Flags
	Flags string `json:"flags"`

	// FirmwareRevision
	// Device firmware revision
	FirmwareRevision string `json:"fw_rev"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// Location
	// Liqid Coordinates for this device
	Location Coordinates `json:"location"`

	// Model
	// Model of this device
	Model string `json:"model"`

	// Name
	// Device name
	Name string `json:"name"`

	// Owner
	// Liqid Coordinates of the device above this one in the hierarchy
	Owner Coordinates `json:"owner"`

	// PartNumber
	// Device name
	PartNumber string `json:"part_num"`

	// PCIDeviceId
	// Vendor-unique device identifier expressed in hex as a '0x'-prefixed string
	PCIDeviceId string `json:"pci_device"`

	// PCIVendorId
	// PCI Vendor identifier expressed in hex as a '0x'-prefixed string
	PCIVendorId string `json:"pci_vendor"`

	// PodId
	// Obsolete value - should always be -1
	PodId int32 `json:"pod_id"`

	// SerialNumber
	// Device serial number
	SerialNumber string `json:"serial_num"`

	// SledId
	// Obsolete
	SledId string `json:"sled_id"`

	// UserDescription
	// User-supplied description
	UserDescription string `json:"udesc"`

	// Unique
	// Device-specific information
	Unique string `json:"unique"`

	// Vendor
	// Vendor name
	Vendor string `json:"vendor"`

	// Capacity
	// Capacity of the memory device
	Capacity string `json:"capacity"`
}

// NewMemoryDeviceInfo initializer for MemoryDeviceInfo struct
func NewMemoryDeviceInfo() MemoryDeviceInfo {
	obj := MemoryDeviceInfo{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// MemoryDeviceInfoWrapper JSON body wrapper for MemoryDeviceInfo
type MemoryDeviceInfoWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []MemoryDeviceInfo `json:"data"`
	} `json:"response"`
}

// MemoryDeviceStatus
// Memory Device Status Information
type MemoryDeviceStatus struct {
	// ConnectionType
	// Connection type for the device
	ConnectionType string `json:"conn_type"`

	// DeviceId
	// Unique identifier for the device
	DeviceId string `json:"dev_id"`

	// DeviceState
	// State of this device
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device type for this device
	DeviceType DeviceType `json:"device_type"`

	// PCIDeviceId
	// PCI Vendor-Unique Device identifier
	PCIDeviceId string `json:"did"`

	// GlobalId
	// Fabric global identifier
	GlobalId string `json:"fabr_gid"`

	// FabricId
	// Fabric identifier
	FabricId int32 `json:"fabr_id"`

	// FabricType
	// Fabric type
	FabricType string `json:"fabric_type"`

	// Flags
	// Hardware flags displayed as a hex string prefixed by 'ox'
	Flags string `json:"flags"`

	// Flags2
	// Additional hardware flags displayed as a hex string prefixed by 'ox'
	Flags2 string `json:"flags2"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// PCILaneCount
	// Number of PCI lanes for this device
	PCILaneCount int32 `json:"lanes"`

	// Location
	// Liqid coordinates for this component
	Location Coordinates `json:"location"`

	// Name
	// Name of this component
	Name string `json:"name"`

	// Owner
	// Liqid coordinates for the component directly above this in the containing fabric topology
	Owner Coordinates `json:"owner"`

	// PodId
	// Pod identifier - obsolete value which is generally always -1
	PodId int32 `json:"pod_id"`

	// PortGlobalId
	// Port Global Identifier
	PortGlobalId string `json:"port_gid"`

	// SledId
	// Obsolete value
	SledId string `json:"sled_id"`

	// SwitchGlobalId
	// Switch Global Identifier
	SwitchGlobalId string `json:"swit_gid"`

	// DeviceStatusType
	// Type of DeviceStatus entity
	DeviceStatusType string `json:"type"`

	// PCIVendorId
	// Unique PCI Vendor Identifier
	PCIVendorId string `json:"vid"`

	// CapacityMB
	// Capacity of the device expressed in MB
	CapacityMB int64 `json:"capacity(MB)"`

	// Unique
	// Internal value
	Unique string `json:"unique"`
}

// NewMemoryDeviceStatus initializer for MemoryDeviceStatus struct
func NewMemoryDeviceStatus() MemoryDeviceStatus {
	obj := MemoryDeviceStatus{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// MemoryDeviceStatusWrapper JSON body wrapper for MemoryDeviceStatus
type MemoryDeviceStatusWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []MemoryDeviceStatus `json:"data"`
	} `json:"response"`
}

// NameValuePair
// A simple tuple tying a value key or name to a value
type NameValuePair struct {
	// Name
	// The key or name associated with a value
	Name string `json:"name"`

	// ValueString
	// The value associated with the given key or name
	ValueString string `json:"valueString"`
}

// NewNameValuePair initializer for NameValuePair struct
func NewNameValuePair() NameValuePair {
	obj := NameValuePair{}
	return obj
}

// NameValuePairWrapper JSON body wrapper for NameValuePair
type NameValuePairWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []NameValuePair `json:"data"`
	} `json:"response"`
}

// NetworkDeviceInfo
// Contains non-status information regarding a network device
type NetworkDeviceInfo struct {
	// BusGeneration
	// Bus Generation
	BusGeneration string `json:"busgen"`

	// BusWidth
	// Bus Width
	BusWidth string `json:"buswidth"`

	// DeviceClass
	// Device class
	DeviceClass string `json:"class"`

	// ConnectionType
	// Connection Type
	ConnectionType string `json:"conn_type"`

	// DeviceIdentifier
	// Device Identifier
	DeviceIdentifier string `json:"dev_id"`

	// DeviceState
	// Device State
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device Type
	DeviceType DeviceType `json:"device_type"`

	// FabricGlobalId
	// Fabric global identifier
	FabricGlobalId string `json:"fabr_gid"`

	// FabricType
	// Fabric Type
	FabricType FabricType `json:"fabric_type"`

	// Family
	// Family
	Family string `json:"family"`

	// Flags
	// Flags
	Flags string `json:"flags"`

	// FirmwareRevision
	// Device firmware revision
	FirmwareRevision string `json:"fw_rev"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// Location
	// Liqid Coordinates for this device
	Location Coordinates `json:"location"`

	// Model
	// Model of this device
	Model string `json:"model"`

	// Name
	// Device name
	Name string `json:"name"`

	// Owner
	// Liqid Coordinates of the device above this one in the hierarchy
	Owner Coordinates `json:"owner"`

	// PartNumber
	// Device name
	PartNumber string `json:"part_num"`

	// PCIDeviceId
	// Vendor-unique device identifier expressed in hex as a '0x'-prefixed string
	PCIDeviceId string `json:"pci_device"`

	// PCIVendorId
	// PCI Vendor identifier expressed in hex as a '0x'-prefixed string
	PCIVendorId string `json:"pci_vendor"`

	// PodId
	// Obsolete value - should always be -1
	PodId int32 `json:"pod_id"`

	// SerialNumber
	// Device serial number
	SerialNumber string `json:"serial_num"`

	// SledId
	// Obsolete
	SledId string `json:"sled_id"`

	// UserDescription
	// User-supplied description
	UserDescription string `json:"udesc"`

	// Unique
	// Device-specific information
	Unique string `json:"unique"`

	// Vendor
	// Vendor name
	Vendor string `json:"vendor"`

	// LinkSpeed
	// Speed/bandwidth of the network link
	LinkSpeed string `json:"link_speed"`

	// DRAMSize
	// Dynamic RAM size
	DRAMSize string `json:"dram_sz"`

	// DRAMType
	// Dynamic RAM type
	DRAMType string `json:"dram_type"`
}

// NewNetworkDeviceInfo initializer for NetworkDeviceInfo struct
func NewNetworkDeviceInfo() NetworkDeviceInfo {
	obj := NetworkDeviceInfo{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// NetworkDeviceInfoWrapper JSON body wrapper for NetworkDeviceInfo
type NetworkDeviceInfoWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []NetworkDeviceInfo `json:"data"`
	} `json:"response"`
}

// NetworkDeviceStatus
// Network Device Status Information
type NetworkDeviceStatus struct {
	// ConnectionType
	// Connection type for the device
	ConnectionType string `json:"conn_type"`

	// DeviceId
	// Unique identifier for the device
	DeviceId string `json:"dev_id"`

	// DeviceState
	// State of this device
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device type for this device
	DeviceType DeviceType `json:"device_type"`

	// PCIDeviceId
	// PCI Vendor-Unique Device identifier
	PCIDeviceId string `json:"did"`

	// GlobalId
	// Fabric global identifier
	GlobalId string `json:"fabr_gid"`

	// FabricId
	// Fabric identifier
	FabricId int32 `json:"fabr_id"`

	// FabricType
	// Fabric type
	FabricType string `json:"fabric_type"`

	// Flags
	// Hardware flags displayed as a hex string prefixed by 'ox'
	Flags string `json:"flags"`

	// Flags2
	// Additional hardware flags displayed as a hex string prefixed by 'ox'
	Flags2 string `json:"flags2"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// PCILaneCount
	// Number of PCI lanes for this device
	PCILaneCount int32 `json:"lanes"`

	// Location
	// Liqid coordinates for this component
	Location Coordinates `json:"location"`

	// Name
	// Name of this component
	Name string `json:"name"`

	// Owner
	// Liqid coordinates for the component directly above this in the containing fabric topology
	Owner Coordinates `json:"owner"`

	// PodId
	// Pod identifier - obsolete value which is generally always -1
	PodId int32 `json:"pod_id"`

	// PortGlobalId
	// Port Global Identifier
	PortGlobalId string `json:"port_gid"`

	// SledId
	// Obsolete value
	SledId string `json:"sled_id"`

	// SwitchGlobalId
	// Switch Global Identifier
	SwitchGlobalId string `json:"swit_gid"`

	// DeviceStatusType
	// Type of DeviceStatus entity
	DeviceStatusType string `json:"type"`

	// PCIVendorId
	// Unique PCI Vendor Identifier
	PCIVendorId string `json:"vid"`

	// Unique
	// Internal value
	Unique string `json:"unique"`
}

// NewNetworkDeviceStatus initializer for NetworkDeviceStatus struct
func NewNetworkDeviceStatus() NetworkDeviceStatus {
	obj := NetworkDeviceStatus{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// NetworkDeviceStatusWrapper JSON body wrapper for NetworkDeviceStatus
type NetworkDeviceStatusWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []NetworkDeviceStatus `json:"data"`
	} `json:"response"`
}

// NetworkManagedCPU
// Describes the access information required to manage a CPU device (e.g., via IPMI)
type NetworkManagedCPU struct {
	// Credentials
	// Credentials necessary for managing the device
	Credentials Credentials `json:"credentials"`

	// IPAddress
	// IP Address or DNS host name of the manager for the managed device
	IPAddress string `json:"ip_address"`

	// PortNumber
	// Port number for managing the device
	PortNumber int32 `json:"port"`

	// ManagerType
	// Entity management type
	ManagerType ManageableType `json:"type"`

	// CPUName
	// CPU name
	CPUName string `json:"cpu_name"`
}

// NewNetworkManagedCPU initializer for NetworkManagedCPU struct
func NewNetworkManagedCPU() NetworkManagedCPU {
	obj := NetworkManagedCPU{}
	obj.Credentials = Credentials{}
	return obj
}

// NetworkManagedCPUWrapper JSON body wrapper for NetworkManagedCPU
type NetworkManagedCPUWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []NetworkManagedCPU `json:"data"`
	} `json:"response"`
}

// NetworkManagedEnclosure
// Describes the access information required to manage an enclosure
type NetworkManagedEnclosure struct {
	// Credentials
	// Credentials necessary for managing the device
	Credentials Credentials `json:"credentials"`

	// IPAddress
	// IP Address or DNS host name of the manager for the managed device
	IPAddress string `json:"ip_address"`

	// PortNumber
	// Port number for managing the device
	PortNumber int32 `json:"port"`

	// ManagerType
	// Entity management type
	ManagerType ManageableType `json:"type"`

	// EnclosureName
	// Enclosure name
	EnclosureName string `json:"enclosure_name"`
}

// NewNetworkManagedEnclosure initializer for NetworkManagedEnclosure struct
func NewNetworkManagedEnclosure() NetworkManagedEnclosure {
	obj := NetworkManagedEnclosure{}
	obj.Credentials = Credentials{}
	return obj
}

// NetworkManagedEnclosureWrapper JSON body wrapper for NetworkManagedEnclosure
type NetworkManagedEnclosureWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []NetworkManagedEnclosure `json:"data"`
	} `json:"response"`
}

// NodeStatus
// Status information regarding one particular node.
// A node should be thought of as a unique host or CPU.
// The use of the word 'node' does not imply any association with a clustered system.
type NodeStatus struct {
	// ConfigVersion
	// Configuration version
	ConfigVersion int32 `json:"cfg_vers"`

	// ConfigIdentifier
	// Configuration identifier
	ConfigIdentifier int32 `json:"cid"`

	// Comps
	// TODO
	Comps int32 `json:"comps"`

	// CurrentTime
	// Current time setting of the node
	CurrentTime string `json:"currtime"`

	// FabricId
	// Identifier of the containing fabric
	FabricId int32 `json:"fabr_id"`

	// Flags
	// Flag settings expressed as a hex value prefixed by '0x'
	Flags string `json:"flags"`

	// LinkCount
	// Number of links for this node
	LinkCount int32 `json:"links"`

	// Location
	// Liqid coordinates of this node
	Location Coordinates `json:"location"`

	// OperatingSystem
	// Operating system which is running on the node
	OperatingSystem OperatingSystemType `json:"os_type"`

	// RunState
	// Current running state of the node
	RunState RunType `json:"run"`

	// SoftwareVersion
	// Software version for the node
	SoftwareVersion int32 `json:"sw_vers"`

	// TargetCount
	// Number of targets
	TargetCount int32 `json:"targs"`

	// UpTime
	// Amount of time the system has been up
	UpTime string `json:"uptime"`
}

// NewNodeStatus initializer for NodeStatus struct
func NewNodeStatus() NodeStatus {
	obj := NodeStatus{}
	obj.Location = Coordinates{}
	return obj
}

// NodeStatusWrapper JSON body wrapper for NodeStatus
type NodeStatusWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []NodeStatus `json:"data"`
	} `json:"response"`
}

// Port
// Describes a switch port
type Port struct {
	// Type
	// Hardware type
	Type string `json:"type"`

	// Index
	// Internal index of this port
	Index string `json:"index"`

	// Flags
	// Hardware flags displayed as a hex string prefixed by 'ox'
	Flags string `json:"flags"`

	// State
	// Hardware state
	State string `json:"state"`

	// FabricGlobalId
	// Fabric-unique global identifier to which this entity is connected.
	FabricGlobalId string `json:"fabr_gid"`

	// PCILaneCount
	// Number of PCI lanes for this component
	PCILaneCount int32 `json:"lanes"`

	// ReceiverErrorCount
	// Running count of errors received for this component
	ReceiverErrorCount int32 `json:"rcv_errs"`

	// BadTLPCount
	// Running count of bad TLPs for this component
	BadTLPCount int32 `json:"bad_tlp"`

	// BadDLLPCount
	// Running count of bad DLLPs for this component
	BadDLLPCount int32 `json:"bad_dllp"`

	// ErrorCount
	// Running count of all errors for this component
	ErrorCount int32 `json:"errs"`

	// IngressBytes
	// Running count of incoming bytes
	IngressBytes int32 `json:"ibytes"`

	// EgressBytes
	// Running count of outgoing bytes
	EgressBytes int32 `json:"ebytes"`

	// Switches
	// Array of Switch objects describing the switches which are connected to this port
	Switches []Switch `json:"switches"`

	// CPU
	// A PortCPU struct describing the CPU which is connected to this port
	CPU PortCPU `json:"cpu"`

	// DeviceState
	// Describes this entity's current state
	DeviceState string `json:"device_state"`
}

// NewPort initializer for Port struct
func NewPort() Port {
	obj := Port{}
	obj.Switches = []Switch{}
	obj.CPU = PortCPU{}
	return obj
}

// PortWrapper JSON body wrapper for Port
type PortWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []Port `json:"data"`
	} `json:"response"`
}

// PortCPU
// Describes a CPU switch port
type PortCPU struct {
	// Name
	// Name of this entity
	Name string `json:"name"`

	// GlobalId
	// Global identifier for this entity
	GlobalId string `json:"gid"`

	// VendorId
	// Unique identifier of the hardware vendor expressed as a hex value prefixed with '0x'
	VendorId string `json:"vendorid"`

	// DeviceId
	// Vendor-unique identifier of the device expressed as a hex value prefixed with '0x'
	DeviceId string `json:"deviceid"`

	// BusWidth
	// Bus-width for the switch
	BusWidth string `json:"buswidth"`

	// Direction
	// TODO
	Direction string `json:"direction"`

	// Type
	// TODO
	Type string `json:"type"`
}

// NewPortCPU initializer for PortCPU struct
func NewPortCPU() PortCPU {
	obj := PortCPU{}
	return obj
}

// PortCPUWrapper JSON body wrapper for PortCPU
type PortCPUWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []PortCPU `json:"data"`
	} `json:"response"`
}

// PreDevice
// Describes a device before it is added to a group
type PreDevice struct {
	// DeviceType
	// Device type
	DeviceType DeviceType `json:"dev_type"`

	// FabricGlobalId
	// Fabric global id
	FabricGlobalId string `json:"fabr_gid"`

	// FabricId
	// Fabric identifier to which the device belongs
	FabricId int32 `json:"fabr_id"`

	// Flags
	// Flags for the device - a hex string prefixed by '0x'
	Flags string `json:"flags"`

	// GroupName
	// TODO
	GroupName string `json:"gname"`

	// GroupId
	// TODO
	GroupId int32 `json:"grp_id"`

	// Index
	// Internal index for this pre-device
	Index string `json:"index"`

	// PCILaneCount
	// Number of PCI lanes for this device
	PCILaneCount int32 `json:"lanes"`

	// MachineId
	// Machine identifier
	MachineId string `json:"mach_id"`

	// MachineName
	// Machine name
	MachineName string `json:"mname"`

	// DeviceName
	// Device name
	DeviceName string `json:"name"`

	// OwnerId
	// Owner identifier
	OwnerId string `json:"owner_id"`

	// PodId
	// Obsolete value - should always be -1
	PodId int32 `json:"pod_id"`

	// PortGlobalId
	// Device port global identifier
	PortGlobalId int32 `json:"port_gid"`

	// SwitchGlobalId
	// Device switch global identifier
	SwitchGlobalId int32 `json:"swit_gid"`
}

// NewPreDevice initializer for PreDevice struct
func NewPreDevice() PreDevice {
	obj := PreDevice{}
	obj.PodId = -1
	return obj
}

// PreDeviceWrapper JSON body wrapper for PreDevice
type PreDeviceWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []PreDevice `json:"data"`
	} `json:"response"`
}

// SSHStatus
// Describes the current state of SSH
type SSHStatus struct {
	// Active
	// Indicates whether SSH is active
	Active bool `json:"is-active"`

	// Enabled
	// Indicates whether SSH is enabled
	Enabled bool `json:"is-enabled"`
}

// NewSSHStatus initializer for SSHStatus struct
func NewSSHStatus() SSHStatus {
	obj := SSHStatus{}
	return obj
}

// SSHStatusWrapper JSON body wrapper for SSHStatus
type SSHStatusWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []SSHStatus `json:"data"`
	} `json:"response"`
}

// StorageDeviceInfo
// Contains non-status information regarding a storage device
type StorageDeviceInfo struct {
	// BusGeneration
	// Bus Generation
	BusGeneration string `json:"busgen"`

	// BusWidth
	// Bus Width
	BusWidth string `json:"buswidth"`

	// DeviceClass
	// Device class
	DeviceClass string `json:"class"`

	// ConnectionType
	// Connection Type
	ConnectionType string `json:"conn_type"`

	// DeviceIdentifier
	// Device Identifier
	DeviceIdentifier string `json:"dev_id"`

	// DeviceState
	// Device State
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device Type
	DeviceType DeviceType `json:"device_type"`

	// FabricGlobalId
	// Fabric global identifier
	FabricGlobalId string `json:"fabr_gid"`

	// FabricType
	// Fabric Type
	FabricType FabricType `json:"fabric_type"`

	// Family
	// Family
	Family string `json:"family"`

	// Flags
	// Flags
	Flags string `json:"flags"`

	// FirmwareRevision
	// Device firmware revision
	FirmwareRevision string `json:"fw_rev"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// Location
	// Liqid Coordinates for this device
	Location Coordinates `json:"location"`

	// Model
	// Model of this device
	Model string `json:"model"`

	// Name
	// Device name
	Name string `json:"name"`

	// Owner
	// Liqid Coordinates of the device above this one in the hierarchy
	Owner Coordinates `json:"owner"`

	// PartNumber
	// Device name
	PartNumber string `json:"part_num"`

	// PCIDeviceId
	// Vendor-unique device identifier expressed in hex as a '0x'-prefixed string
	PCIDeviceId string `json:"pci_device"`

	// PCIVendorId
	// PCI Vendor identifier expressed in hex as a '0x'-prefixed string
	PCIVendorId string `json:"pci_vendor"`

	// PodId
	// Obsolete value - should always be -1
	PodId int32 `json:"pod_id"`

	// SerialNumber
	// Device serial number
	SerialNumber string `json:"serial_num"`

	// SledId
	// Obsolete
	SledId string `json:"sled_id"`

	// UserDescription
	// User-supplied description
	UserDescription string `json:"udesc"`

	// Unique
	// Device-specific information
	Unique string `json:"unique"`

	// Vendor
	// Vendor name
	Vendor string `json:"vendor"`

	// Capacity
	// Capacity of the storage device
	Capacity string `json:"capacity"`
}

// NewStorageDeviceInfo initializer for StorageDeviceInfo struct
func NewStorageDeviceInfo() StorageDeviceInfo {
	obj := StorageDeviceInfo{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// StorageDeviceInfoWrapper JSON body wrapper for StorageDeviceInfo
type StorageDeviceInfoWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []StorageDeviceInfo `json:"data"`
	} `json:"response"`
}

// StorageDeviceStatus
// Storage Device Status Information
type StorageDeviceStatus struct {
	// ConnectionType
	// Connection type for the device
	ConnectionType string `json:"conn_type"`

	// DeviceId
	// Unique identifier for the device
	DeviceId string `json:"dev_id"`

	// DeviceState
	// State of this device
	DeviceState string `json:"device_state"`

	// DeviceType
	// Device type for this device
	DeviceType DeviceType `json:"device_type"`

	// PCIDeviceId
	// PCI Vendor-Unique Device identifier
	PCIDeviceId string `json:"did"`

	// GlobalId
	// Fabric global identifier
	GlobalId string `json:"fabr_gid"`

	// FabricId
	// Fabric identifier
	FabricId int32 `json:"fabr_id"`

	// FabricType
	// Fabric type
	FabricType string `json:"fabric_type"`

	// Flags
	// Hardware flags displayed as a hex string prefixed by 'ox'
	Flags string `json:"flags"`

	// Flags2
	// Additional hardware flags displayed as a hex string prefixed by 'ox'
	Flags2 string `json:"flags2"`

	// Index
	// Internal index of this device
	Index int32 `json:"index"`

	// PCILaneCount
	// Number of PCI lanes for this device
	PCILaneCount int32 `json:"lanes"`

	// Location
	// Liqid coordinates for this component
	Location Coordinates `json:"location"`

	// Name
	// Name of this component
	Name string `json:"name"`

	// Owner
	// Liqid coordinates for the component directly above this in the containing fabric topology
	Owner Coordinates `json:"owner"`

	// PodId
	// Pod identifier - obsolete value which is generally always -1
	PodId int32 `json:"pod_id"`

	// PortGlobalId
	// Port Global Identifier
	PortGlobalId string `json:"port_gid"`

	// SledId
	// Obsolete value
	SledId string `json:"sled_id"`

	// SwitchGlobalId
	// Switch Global Identifier
	SwitchGlobalId string `json:"swit_gid"`

	// DeviceStatusType
	// Type of DeviceStatus entity
	DeviceStatusType string `json:"type"`

	// PCIVendorId
	// Unique PCI Vendor Identifier
	PCIVendorId string `json:"vid"`

	// CapacityMB
	// Capacity of the device expressed in MB
	CapacityMB int64 `json:"capacity(MB)"`
}

// NewStorageDeviceStatus initializer for StorageDeviceStatus struct
func NewStorageDeviceStatus() StorageDeviceStatus {
	obj := StorageDeviceStatus{}
	obj.Location = Coordinates{}
	obj.Owner = Coordinates{}
	obj.PodId = -1
	return obj
}

// StorageDeviceStatusWrapper JSON body wrapper for StorageDeviceStatus
type StorageDeviceStatusWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []StorageDeviceStatus `json:"data"`
	} `json:"response"`
}

// Switch
// Details related to the PCI switch
type Switch struct {
	// Name
	// Switch name
	Name string `json:"name"`

	// GlobalId
	// Global identifier expressed as a hex value prefixed with '0x'
	GlobalId string `json:"gid"`

	// VendorId
	// Unique identifier of the hardware vendor expressed as a hex value prefixed with '0x'
	VendorId string `json:"vendorid"`

	// DeviceId
	// Vendor-unique identifier of the device expressed as a hex value prefixed with '0x'
	DeviceId string `json:"deviceid"`

	// BusWidth
	// Bus-width for the switch
	BusWidth string `json:"buswidth"`

	// Direction
	// TODO
	Direction string `json:"direction"`

	// Type
	// TODO
	Type string `json:"type"`

	// SwitchDevice
	// Additional details regarding the switch
	SwitchDevice SwitchDevice `json:"device"`
}

// NewSwitch initializer for Switch struct
func NewSwitch() Switch {
	obj := Switch{}
	obj.SwitchDevice = SwitchDevice{}
	return obj
}

// SwitchWrapper JSON body wrapper for Switch
type SwitchWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []Switch `json:"data"`
	} `json:"response"`
}

// SwitchDevice
// Additional details specific to the switch
type SwitchDevice struct {
	// Name
	// Switch name
	Name string `json:"name"`

	// GroupId
	// Unique identifier of the group expressed as a hex value prefixed with '0x'
	GroupId string `json:"gid"`

	// VendorId
	// Unique identifier of the hardware vendor expressed as a hex value prefixed with '0x'
	VendorId string `json:"vendorid"`

	// DeviceId
	// Vendor-unique identifier of the device expressed as a hex value prefixed with '0x'
	DeviceId string `json:"deviceid"`

	// BusWidth
	// Bus-width for the switch
	BusWidth string `json:"buswidth"`

	// Direction
	// TODO
	Direction string `json:"direction"`

	// Type
	// TODO
	Type string `json:"type"`
}

// NewSwitchDevice initializer for SwitchDevice struct
func NewSwitchDevice() SwitchDevice {
	obj := SwitchDevice{}
	return obj
}

// SwitchDeviceWrapper JSON body wrapper for SwitchDevice
type SwitchDeviceWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []SwitchDevice `json:"data"`
	} `json:"response"`
}

// Timestamp
// Contains a timestamp value
type Timestamp struct {
	// OldTimestamp
	// Old value with wrong size int and incorrectly-named JSON tag of 'epoch'
	OldTimestamp int32 `json:"epoch"`

	// Timestamp
	// A time and date based on the Unix epoch
	Timestamp int64 `json:"timestamp"`
}

// NewTimestamp initializer for Timestamp struct
func NewTimestamp() Timestamp {
	obj := Timestamp{}
	return obj
}

// TimestampWrapper JSON body wrapper for Timestamp
type TimestampWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []Timestamp `json:"data"`
	} `json:"response"`
}

// VersionInfo
// Describes the versions of the various software components which comprise the Director
type VersionInfo struct {
	// Branch
	// Development branch from which this component was built
	Branch string `json:"branch"`

	// ChangeSet
	// 
	ChangeSet string `json:"changeset"`

	// ShortChangeSet
	// 
	ShortChangeSet string `json:"changeset_short"`

	// Component
	// Name of the software component
	Component string `json:"component"`

	// Date
	// Date the component was built
	Date string `json:"date"`

	// ShortDate
	// Date the component was built
	ShortDate string `json:"date_short"`

	// Version
	// Component version string
	Version string `json:"version"`
}

// NewVersionInfo initializer for VersionInfo struct
func NewVersionInfo() VersionInfo {
	obj := VersionInfo{}
	return obj
}

// VersionInfoWrapper JSON body wrapper for VersionInfo
type VersionInfoWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []VersionInfo `json:"data"`
	} `json:"response"`
}

// FabricTopology
// An array of FabricItem entities
type FabricTopology []FabricItem

// FabricTopologyWrapper JSON body wrapper for FabricTopology
type FabricTopologyWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []FabricTopology `json:"data"`
	} `json:"response"`
}
