// File: groupDeviceRelator_gen.go
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

import (
	"fmt"
)

// Category: GroupDeviceRelator
// Functions which move devices into and out of group free pools

// AddComputeDeviceToGroup Moves a device from the system free pool to the group free pool for the indicated group
// Returns: A description of the relation being created
func (client *LiqidClient) AddComputeDeviceToGroup(deviceId string, groupId int32) (*GroupComputeDeviceRelator, error) {
	var fn = "AddComputeDeviceToGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupComputeDeviceRelator{}
	value0, err0 := client.GetComputeDeviceStatus(deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.DeviceStatus = *value0
	value1, err1 := client.GetGroup(groupId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Group = *value1

	path := "predevice/compute"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupComputeDeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// AddFPGADeviceToGroup Moves a device from the system free pool to the group free pool for the indicated group
// Returns: A description of the relation being created
func (client *LiqidClient) AddFPGADeviceToGroup(deviceId string, groupId int32) (*GroupFPGADeviceRelator, error) {
	var fn = "AddFPGADeviceToGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupFPGADeviceRelator{}
	value0, err0 := client.GetFPGADeviceStatus(deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.DeviceStatus = *value0
	value1, err1 := client.GetGroup(groupId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Group = *value1

	path := "predevice/fpga"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupFPGADeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// AddGPUDeviceToGroup Moves a device from the system free pool to the group free pool for the indicated group
// Returns: A description of the relation being created
func (client *LiqidClient) AddGPUDeviceToGroup(deviceId string, groupId int32) (*GroupGPUDeviceRelator, error) {
	var fn = "AddGPUDeviceToGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupGPUDeviceRelator{}
	value0, err0 := client.GetGPUDeviceStatus(deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.DeviceStatus = *value0
	value1, err1 := client.GetGroup(groupId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Group = *value1

	path := "predevice/gpu"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupGPUDeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// AddMemoryDeviceToGroup Moves a device from the system free pool to the group free pool for the indicated group
// Returns: A description of the relation being created
func (client *LiqidClient) AddMemoryDeviceToGroup(deviceId string, groupId int32) (*GroupMemoryDeviceRelator, error) {
	var fn = "AddMemoryDeviceToGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupMemoryDeviceRelator{}
	value0, err0 := client.GetGroup(groupId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.Group = *value0
	value1, err1 := client.GetMemoryDeviceStatus(deviceId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.DeviceStatus = *value1

	path := "predevice/mem"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupMemoryDeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// AddNetworkDeviceToGroup Moves a device from the system free pool to the group free pool for the indicated group
// Returns: A description of the relation being created
func (client *LiqidClient) AddNetworkDeviceToGroup(deviceId string, groupId int32) (*GroupNetworkDeviceRelator, error) {
	var fn = "AddNetworkDeviceToGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupNetworkDeviceRelator{}
	value0, err0 := client.GetGroup(groupId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.Group = *value0
	value1, err1 := client.GetNetworkDeviceStatus(deviceId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.DeviceStatus = *value1

	path := "predevice/network"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupNetworkDeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// AddStorageDeviceToGroup Moves a device from the system free pool to the group free pool for the indicated group
// Returns: A description of the relation being created
func (client *LiqidClient) AddStorageDeviceToGroup(deviceId string, groupId int32) (*GroupStorageDeviceRelator, error) {
	var fn = "AddStorageDeviceToGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupStorageDeviceRelator{}
	value0, err0 := client.GetGroup(groupId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.Group = *value0
	value1, err1 := client.GetStorageDeviceStatus(deviceId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.DeviceStatus = *value1

	path := "predevice/storage"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupStorageDeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// GetDevices Returns information regard all devices (or a subset thereof) for the system
//
// Param deviceType: Limits the device type of the devices to be queried
// Param groupId: Only return devices associated with the indicated group
// Param machineId: Only return devices associated with the indicated machine
// Returns: An array of PreDevice entities describing the various devices in the configuration
func (client *LiqidClient) GetDevices(deviceType *DeviceType, groupId *int32, machineId *int32) ([]PreDevice, error) {
	var fn = "GetDevices"
	client.traceLogger.Printf(LogFmtEnter3, fn, deviceType, groupId, machineId)

	path := "predevice"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	queryParams := make(map[string]string)
	if deviceType != nil {
		queryParams["dev_type"] = fmt.Sprintf("%v", *deviceType)
	}
	if groupId != nil {
		queryParams["grp_id"] = fmt.Sprintf("%v", *groupId)
	}
	if machineId != nil {
		queryParams["mach_id"] = fmt.Sprintf("%v", *machineId)
	}
	cliReq.SetParameters(queryParams)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper PreDeviceWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := respWrapper.Response.Data
	client.traceLogger.Printf(LogFmtReturn2, fn, result, nil)
	return result, nil
}

// GetGroupComputeDeviceRelator This function is required by the SDK architecture.
//                              It is not intended for use by the SDK client.
// Returns: A GroupComputeDeviceRelator entity
func (client *LiqidClient) GetGroupComputeDeviceRelator(groupId int32, deviceId string) (*GroupComputeDeviceRelator, error) {
	var fn = "GetGroupComputeDeviceRelator"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, deviceId)
	var composite = GroupComputeDeviceRelator{}
	value0, err0 := client.GetComputeDeviceStatus(deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.DeviceStatus = *value0
	value1, err1 := client.GetGroup(groupId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Group = *value1
    return &composite, nil
}

// GetGroupFPGADeviceRelator This function is required by the SDK architecture.
//                           It is not intended for use by the SDK client.
// Returns: A GroupFPGADeviceRelator entity
func (client *LiqidClient) GetGroupFPGADeviceRelator(groupId int32, deviceId string) (*GroupFPGADeviceRelator, error) {
	var fn = "GetGroupFPGADeviceRelator"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, deviceId)
	var composite = GroupFPGADeviceRelator{}
	value0, err0 := client.GetFPGADeviceStatus(deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.DeviceStatus = *value0
	value1, err1 := client.GetGroup(groupId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Group = *value1
    return &composite, nil
}

// GetGroupGPUDeviceRelator This function is required by the SDK architecture.
//                          It is not intended for use by the SDK client.
// Returns: A GroupGPUDeviceRelator entity
func (client *LiqidClient) GetGroupGPUDeviceRelator(groupId int32, deviceId string) (*GroupGPUDeviceRelator, error) {
	var fn = "GetGroupGPUDeviceRelator"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, deviceId)
	var composite = GroupGPUDeviceRelator{}
	value0, err0 := client.GetGPUDeviceStatus(deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.DeviceStatus = *value0
	value1, err1 := client.GetGroup(groupId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Group = *value1
    return &composite, nil
}

// GetGroupMemoryDeviceRelator This function is required by the SDK architecture.
//                             It is not intended for use by the SDK client.
// Returns: A GroupMemoryDeviceRelator entity
func (client *LiqidClient) GetGroupMemoryDeviceRelator(groupId int32, deviceId string) (*GroupMemoryDeviceRelator, error) {
	var fn = "GetGroupMemoryDeviceRelator"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, deviceId)
	var composite = GroupMemoryDeviceRelator{}
	value0, err0 := client.GetGroup(groupId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.Group = *value0
	value1, err1 := client.GetMemoryDeviceStatus(deviceId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.DeviceStatus = *value1
    return &composite, nil
}

// GetGroupNetworkDeviceRelator This function is required by the SDK architecture.
//                              It is not intended for use by the SDK client.
// Returns: A GroupNetworkDeviceRelator entity
func (client *LiqidClient) GetGroupNetworkDeviceRelator(groupId int32, deviceId string) (*GroupNetworkDeviceRelator, error) {
	var fn = "GetGroupNetworkDeviceRelator"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, deviceId)
	var composite = GroupNetworkDeviceRelator{}
	value0, err0 := client.GetGroup(groupId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.Group = *value0
	value1, err1 := client.GetNetworkDeviceStatus(deviceId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.DeviceStatus = *value1
    return &composite, nil
}

// GetGroupStorageDeviceRelator This function is required by the SDK architecture.
//                              It is not intended for use by the SDK client.
// Returns: A GroupStorageDeviceRelator entity
func (client *LiqidClient) GetGroupStorageDeviceRelator(groupId int32, deviceId string) (*GroupStorageDeviceRelator, error) {
	var fn = "GetGroupStorageDeviceRelator"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, deviceId)
	var composite = GroupStorageDeviceRelator{}
	value0, err0 := client.GetGroup(groupId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.Group = *value0
	value1, err1 := client.GetStorageDeviceStatus(deviceId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.DeviceStatus = *value1
    return &composite, nil
}

// RemoveComputeDeviceFromGroup Moves a device from the group free pool to the system free pool
// Returns: A description of the deleted relation
func (client *LiqidClient) RemoveComputeDeviceFromGroup(deviceId string, groupId int32) (*GroupComputeDeviceRelator, error) {
	var fn = "RemoveComputeDeviceFromGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupComputeDeviceRelator{}
	value0, err0 := client.GetComputeDeviceStatus(deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.DeviceStatus = *value0
	value1, err1 := client.GetGroup(groupId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Group = *value1

	path := "predevice/compute"
	cliReq := NewClientRequest()
	cliReq.SetMethod("DELETE")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupComputeDeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// RemoveFPGADeviceFromGroup Moves a device from the group free pool to the system free pool
// Returns: A description of the deleted relation
func (client *LiqidClient) RemoveFPGADeviceFromGroup(deviceId string, groupId int32) (*GroupFPGADeviceRelator, error) {
	var fn = "RemoveFPGADeviceFromGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupFPGADeviceRelator{}
	value0, err0 := client.GetFPGADeviceStatus(deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.DeviceStatus = *value0
	value1, err1 := client.GetGroup(groupId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Group = *value1

	path := "predevice/fpga"
	cliReq := NewClientRequest()
	cliReq.SetMethod("DELETE")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupFPGADeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// RemoveGPUDeviceFromGroup Moves a device from the group free pool to the system free pool
// Returns: A description of the deleted relation
func (client *LiqidClient) RemoveGPUDeviceFromGroup(deviceId string, groupId int32) (*GroupGPUDeviceRelator, error) {
	var fn = "RemoveGPUDeviceFromGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupGPUDeviceRelator{}
	value0, err0 := client.GetGPUDeviceStatus(deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.DeviceStatus = *value0
	value1, err1 := client.GetGroup(groupId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Group = *value1

	path := "predevice/gpu"
	cliReq := NewClientRequest()
	cliReq.SetMethod("DELETE")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupGPUDeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// RemoveMemoryDeviceFromGroup Moves a device from the group free pool to the system free pool
// Returns: A description of the deleted relation
func (client *LiqidClient) RemoveMemoryDeviceFromGroup(deviceId string, groupId int32) (*GroupMemoryDeviceRelator, error) {
	var fn = "RemoveMemoryDeviceFromGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupMemoryDeviceRelator{}
	value0, err0 := client.GetGroup(groupId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.Group = *value0
	value1, err1 := client.GetMemoryDeviceStatus(deviceId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.DeviceStatus = *value1

	path := "predevice/mem"
	cliReq := NewClientRequest()
	cliReq.SetMethod("DELETE")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupMemoryDeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// RemoveNetworkDeviceFromGroup Moves a device from the group free pool to the system free pool
// Returns: A description of the deleted relation
func (client *LiqidClient) RemoveNetworkDeviceFromGroup(deviceId string, groupId int32) (*GroupNetworkDeviceRelator, error) {
	var fn = "RemoveNetworkDeviceFromGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupNetworkDeviceRelator{}
	value0, err0 := client.GetGroup(groupId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.Group = *value0
	value1, err1 := client.GetNetworkDeviceStatus(deviceId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.DeviceStatus = *value1

	path := "predevice/network"
	cliReq := NewClientRequest()
	cliReq.SetMethod("DELETE")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupNetworkDeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}

// RemoveStorageDeviceFromGroup Moves a device from the group free pool to the system free pool
// Returns: A description of the deleted relation
func (client *LiqidClient) RemoveStorageDeviceFromGroup(deviceId string, groupId int32) (*GroupStorageDeviceRelator, error) {
	var fn = "RemoveStorageDeviceFromGroup"
	client.traceLogger.Printf(LogFmtEnter2, fn, deviceId, groupId)
	var composite = GroupStorageDeviceRelator{}
	value0, err0 := client.GetGroup(groupId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.Group = *value0
	value1, err1 := client.GetStorageDeviceStatus(deviceId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.DeviceStatus = *value1

	path := "predevice/storage"
	cliReq := NewClientRequest()
	cliReq.SetMethod("DELETE")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(composite)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupStorageDeviceRelatorWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, jsonErr)
		return nil, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, fmtErr)
		return nil, fmtErr
	}

	result := &respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, *result, nil)
	return result, nil
}
