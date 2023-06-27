// File: machineDeviceRelator_gen.go
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

// Category: MachineDeviceRelator
// Functions which attach devices to and detach them from machines

// AddComputeDeviceToMachine Attaches a particular device to a machine
// Returns: A description of the relation being created
func (client *LiqidClient) AddComputeDeviceToMachine(groupId int32, deviceId string, machineId int32) (*GroupComputeDeviceRelator, error) {
	var fn = "AddComputeDeviceToMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineComputeDeviceRelator{}
	value0, err0 := client.GetGroupComputeDeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/compute"
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

// AddFPGADeviceToMachine Attaches a particular device to a machine
// Returns: A description of the relation being created
func (client *LiqidClient) AddFPGADeviceToMachine(groupId int32, deviceId string, machineId int32) (*GroupFPGADeviceRelator, error) {
	var fn = "AddFPGADeviceToMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineFPGADeviceRelator{}
	value0, err0 := client.GetGroupFPGADeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/fpga"
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

// AddGPUDeviceToMachine Attaches a particular device to a machine
// Returns: A description of the relation being created
func (client *LiqidClient) AddGPUDeviceToMachine(groupId int32, deviceId string, machineId int32) (*GroupGPUDeviceRelator, error) {
	var fn = "AddGPUDeviceToMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineGPUDeviceRelator{}
	value0, err0 := client.GetGroupGPUDeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/gpu"
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

// AddMemoryDeviceToMachine Attaches a particular device to a machine
// Returns: A description of the relation being created
func (client *LiqidClient) AddMemoryDeviceToMachine(groupId int32, deviceId string, machineId int32) (*GroupMemoryDeviceRelator, error) {
	var fn = "AddMemoryDeviceToMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineMemoryDeviceRelator{}
	value0, err0 := client.GetGroupMemoryDeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/memory"
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

// AddNetworkDeviceToMachine Attaches a particular device to a machine
// Returns: A description of the relation being created
func (client *LiqidClient) AddNetworkDeviceToMachine(groupId int32, deviceId string, machineId int32) (*GroupNetworkDeviceRelator, error) {
	var fn = "AddNetworkDeviceToMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineNetworkDeviceRelator{}
	value0, err0 := client.GetGroupNetworkDeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/network"
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

// AddStorageDeviceToMachine Attaches a particular device to a machine
// Returns: A description of the relation being created
func (client *LiqidClient) AddStorageDeviceToMachine(groupId int32, deviceId string, machineId int32) (*GroupStorageDeviceRelator, error) {
	var fn = "AddStorageDeviceToMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineStorageDeviceRelator{}
	value0, err0 := client.GetGroupStorageDeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/storage"
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

// RemoveComputeDeviceFromMachine Detaches a particular device from a machine, returning it to the group free pool
// Returns: A description of the relation being removed
func (client *LiqidClient) RemoveComputeDeviceFromMachine(groupId int32, deviceId string, machineId int32) (*GroupComputeDeviceRelator, error) {
	var fn = "RemoveComputeDeviceFromMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineComputeDeviceRelator{}
	value0, err0 := client.GetGroupComputeDeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/compute"
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

// RemoveFPGADeviceFromMachine Detaches a particular device from a machine, returning it to the group free pool
// Returns: A description of the relation being removed
func (client *LiqidClient) RemoveFPGADeviceFromMachine(groupId int32, deviceId string, machineId int32) (*GroupFPGADeviceRelator, error) {
	var fn = "RemoveFPGADeviceFromMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineFPGADeviceRelator{}
	value0, err0 := client.GetGroupFPGADeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/fpga"
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

// RemoveGPUDeviceFromMachine Detaches a particular device from a machine, returning it to the group free pool
// Returns: A description of the relation being removed
func (client *LiqidClient) RemoveGPUDeviceFromMachine(groupId int32, deviceId string, machineId int32) (*GroupGPUDeviceRelator, error) {
	var fn = "RemoveGPUDeviceFromMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineGPUDeviceRelator{}
	value0, err0 := client.GetGroupGPUDeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/gpu"
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

// RemoveMemoryDeviceFromMachine Detaches a particular device from a machine, returning it to the group free pool
// Returns: A description of the relation being removed
func (client *LiqidClient) RemoveMemoryDeviceFromMachine(groupId int32, deviceId string, machineId int32) (*GroupMemoryDeviceRelator, error) {
	var fn = "RemoveMemoryDeviceFromMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineMemoryDeviceRelator{}
	value0, err0 := client.GetGroupMemoryDeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/memory"
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

// RemoveNetworkDeviceFromMachine Detaches a particular device from a machine, returning it to the group free pool
// Returns: A description of the relation being removed
func (client *LiqidClient) RemoveNetworkDeviceFromMachine(groupId int32, deviceId string, machineId int32) (*GroupNetworkDeviceRelator, error) {
	var fn = "RemoveNetworkDeviceFromMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineNetworkDeviceRelator{}
	value0, err0 := client.GetGroupNetworkDeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/network"
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

// RemoveStorageDeviceFromMachine Detaches a particular device from a machine, returning it to the group free pool
// Returns: A description of the relation being removed
func (client *LiqidClient) RemoveStorageDeviceFromMachine(groupId int32, deviceId string, machineId int32) (*GroupStorageDeviceRelator, error) {
	var fn = "RemoveStorageDeviceFromMachine"
	client.traceLogger.Printf(LogFmtEnter3, fn, groupId, deviceId, machineId)
	var composite = MachineStorageDeviceRelator{}
	value0, err0 := client.GetGroupStorageDeviceRelator(groupId, deviceId)
	if err0 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err0)
		return nil, err0
	}
	composite.GroupDeviceRelator = *value0
	value1, err1 := client.GetMachine(machineId)
	if err1 != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, err1)
		return nil, err1
	}
	composite.Machine = *value1

	path := "relate/storage"
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
