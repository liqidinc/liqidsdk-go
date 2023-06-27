// File: deviceStatus_gen.go
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

// Category: DeviceStatus
// Functions which retrieve the status of various devices

// GetAllDevicesStatus Retrieves status for all devices
// Returns: An array of DeviceStatus entities describing the status of all devices in the system
func (client *LiqidClient) GetAllDevicesStatus() ([]DeviceStatus, error) {
	var fn = "GetAllDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper DeviceStatusWrapper
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

// GetComputeDeviceStatus Retrieves status for a particular compute device
// Returns: A ComputeDeviceStatus entity describing the status of the requested device
func (client *LiqidClient) GetComputeDeviceStatus(deviceId string) (*ComputeDeviceStatus, error) {
	var fn = "GetComputeDeviceStatus"
	client.traceLogger.Printf(LogFmtEnter1, fn, deviceId)
	list, getErr := client.GetComputeDevicesStatus()
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}
	for _, item := range list {
		if deviceId == item.DeviceId {
			client.traceLogger.Printf(LogFmtReturn2, fn, item, nil)
			return &item, nil
		}
	}
	err := fmt.Errorf("cannot find requested item")
	client.traceLogger.Printf(LogFmtReturn2, fn, nil, err)
	return nil, err
}

// GetComputeDevicesStatus Retrieves status for compute devices
// Returns: An array of ComputeDeviceStatus entities describing the status of compute devices in the system
func (client *LiqidClient) GetComputeDevicesStatus() ([]ComputeDeviceStatus, error) {
	var fn = "GetComputeDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/compute"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	queryParams := make(map[string]string)
	queryParams["filter"] = "false"
	cliReq.SetParameters(queryParams)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper ComputeDeviceStatusWrapper
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

// GetComputeDevicesStatusWithMultiplePortsStatus Retrieves status for compute devices
// Returns: An array of ComputeDeviceStatus entities describing the status of compute devices in the system
func (client *LiqidClient) GetComputeDevicesStatusWithMultiplePortsStatus() ([]ComputeDeviceStatus, error) {
	var fn = "GetComputeDevicesStatusWithMultiplePortsStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/compute/parents"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper ComputeDeviceStatusWrapper
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

// GetFPGADeviceStatus Retrieves status for a particular FPGA device
// Returns: A FPGADeviceStatus entity describing the status of the requested device
func (client *LiqidClient) GetFPGADeviceStatus(deviceId string) (*FPGADeviceStatus, error) {
	var fn = "GetFPGADeviceStatus"
	client.traceLogger.Printf(LogFmtEnter1, fn, deviceId)
	list, getErr := client.GetFPGADevicesStatus()
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}
	for _, item := range list {
		if deviceId == item.DeviceId {
			client.traceLogger.Printf(LogFmtReturn2, fn, item, nil)
			return &item, nil
		}
	}
	err := fmt.Errorf("cannot find requested item")
	client.traceLogger.Printf(LogFmtReturn2, fn, nil, err)
	return nil, err
}

// GetFPGADevicesStatus Retrieves status for FPGA devices
// Returns: An array of FPGADeviceStatus entities describing the status of FPGA devices in the system
func (client *LiqidClient) GetFPGADevicesStatus() ([]FPGADeviceStatus, error) {
	var fn = "GetFPGADevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/fpga"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	queryParams := make(map[string]string)
	queryParams["filter"] = "false"
	cliReq.SetParameters(queryParams)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper FPGADeviceStatusWrapper
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

// GetFreeComputeDevicesStatus Retrieves status for compute devices which are not assigned to a group or machine
// Returns: An array of ComputeDeviceStatus entities describing the status of compute devices in the system
func (client *LiqidClient) GetFreeComputeDevicesStatus() ([]ComputeDeviceStatus, error) {
	var fn = "GetFreeComputeDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/compute"
	cliReq := NewClientRequest().SetMethod("GET").SetPartialPath(path)
	apiParam := NewComputeDeviceStatus()
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper ComputeDeviceStatusWrapper
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

// GetFreeFPGADevicesStatus Retrieves status for compute FPGA which are not assigned to a group or machine
// Returns: An array of FPGADeviceStatus entities describing the status of FPGA devices in the system
func (client *LiqidClient) GetFreeFPGADevicesStatus() ([]FPGADeviceStatus, error) {
	var fn = "GetFreeFPGADevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/fpga"
	cliReq := NewClientRequest().SetMethod("GET").SetPartialPath(path)
	apiParam := NewFPGADeviceStatus()
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper FPGADeviceStatusWrapper
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

// GetFreeGPUDevicesStatus Retrieves status for compute GPU which are not assigned to a group or machine
// Returns: An array of GPUDeviceStatus entities describing the status of GPU devices in the system
func (client *LiqidClient) GetFreeGPUDevicesStatus() ([]GPUDeviceStatus, error) {
	var fn = "GetFreeGPUDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/gpu"
	cliReq := NewClientRequest().SetMethod("GET").SetPartialPath(path)
	apiParam := NewGPUDeviceStatus()
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GPUDeviceStatusWrapper
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

// GetFreeMemoryDevicesStatus Retrieves status for memory devices which are not assigned to a group or machine
// Returns: An array of MemoryDeviceStatus entities describing the status of memory devices in the system
func (client *LiqidClient) GetFreeMemoryDevicesStatus() ([]MemoryDeviceStatus, error) {
	var fn = "GetFreeMemoryDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/mem"
	cliReq := NewClientRequest().SetMethod("GET").SetPartialPath(path)
	apiParam := NewMemoryDeviceStatus()
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper MemoryDeviceStatusWrapper
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

// GetFreeNetworkDevicesStatus Retrieves status for network devices which are not assigned to a group or machine
// Returns: An array of NetworkDeviceStatus entities describing the status of network devices in the system
func (client *LiqidClient) GetFreeNetworkDevicesStatus() ([]NetworkDeviceStatus, error) {
	var fn = "GetFreeNetworkDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/network"
	cliReq := NewClientRequest().SetMethod("GET").SetPartialPath(path)
	apiParam := NewNetworkDeviceStatus()
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper NetworkDeviceStatusWrapper
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

// GetFreeStorageDevicesStatus Retrieves status for storage devices which are not assigned to a group or machine
// Returns: An array of StorageDeviceStatus entities describing the status of storage devices in the system
func (client *LiqidClient) GetFreeStorageDevicesStatus() ([]StorageDeviceStatus, error) {
	var fn = "GetFreeStorageDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/storage"
	cliReq := NewClientRequest().SetMethod("GET").SetPartialPath(path)
	apiParam := NewStorageDeviceStatus()
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper StorageDeviceStatusWrapper
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

// GetGPUDeviceStatus Retrieves status for a particular GPU device
// Returns: A GPUDeviceStatus entity describing the status of the requested device
func (client *LiqidClient) GetGPUDeviceStatus(deviceId string) (*GPUDeviceStatus, error) {
	var fn = "GetGPUDeviceStatus"
	client.traceLogger.Printf(LogFmtEnter1, fn, deviceId)
	list, getErr := client.GetGPUDevicesStatus()
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}
	for _, item := range list {
		if deviceId == item.DeviceId {
			client.traceLogger.Printf(LogFmtReturn2, fn, item, nil)
			return &item, nil
		}
	}
	err := fmt.Errorf("cannot find requested item")
	client.traceLogger.Printf(LogFmtReturn2, fn, nil, err)
	return nil, err
}

// GetGPUDevicesStatus Retrieves status for GPU devices
// Returns: An array of GPUDeviceStatus entities describing the status of GPU devices in the system
func (client *LiqidClient) GetGPUDevicesStatus() ([]GPUDeviceStatus, error) {
	var fn = "GetGPUDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/gpu"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	queryParams := make(map[string]string)
	queryParams["filter"] = "false"
	cliReq.SetParameters(queryParams)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GPUDeviceStatusWrapper
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

// GetMemoryDeviceStatus Retrieves status for a particular memory device
// Returns: A MemoryDeviceStatus entity describing the status of the requested device
func (client *LiqidClient) GetMemoryDeviceStatus(deviceId string) (*MemoryDeviceStatus, error) {
	var fn = "GetMemoryDeviceStatus"
	client.traceLogger.Printf(LogFmtEnter1, fn, deviceId)
	list, getErr := client.GetMemoryDevicesStatus()
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}
	for _, item := range list {
		if deviceId == item.DeviceId {
			client.traceLogger.Printf(LogFmtReturn2, fn, item, nil)
			return &item, nil
		}
	}
	err := fmt.Errorf("cannot find requested item")
	client.traceLogger.Printf(LogFmtReturn2, fn, nil, err)
	return nil, err
}

// GetMemoryDevicesStatus Retrieves status for memory devices
// Returns: An array of MemoryDeviceStatus entities describing the status of memory devices in the system
func (client *LiqidClient) GetMemoryDevicesStatus() ([]MemoryDeviceStatus, error) {
	var fn = "GetMemoryDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/mem"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	queryParams := make(map[string]string)
	queryParams["filter"] = "false"
	cliReq.SetParameters(queryParams)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper MemoryDeviceStatusWrapper
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

// GetNetworkDeviceStatus Retrieves status for a particular network device
// Returns: A NetworkDeviceStatus entity describing the status of the requested device
func (client *LiqidClient) GetNetworkDeviceStatus(deviceId string) (*NetworkDeviceStatus, error) {
	var fn = "GetNetworkDeviceStatus"
	client.traceLogger.Printf(LogFmtEnter1, fn, deviceId)
	list, getErr := client.GetNetworkDevicesStatus()
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}
	for _, item := range list {
		if deviceId == item.DeviceId {
			client.traceLogger.Printf(LogFmtReturn2, fn, item, nil)
			return &item, nil
		}
	}
	err := fmt.Errorf("cannot find requested item")
	client.traceLogger.Printf(LogFmtReturn2, fn, nil, err)
	return nil, err
}

// GetNetworkDevicesStatus Retrieves status for network devices
// Returns: An array of NetworkDeviceStatus entities describing the status of network devices in the system
func (client *LiqidClient) GetNetworkDevicesStatus() ([]NetworkDeviceStatus, error) {
	var fn = "GetNetworkDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/network"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	queryParams := make(map[string]string)
	queryParams["filter"] = "false"
	cliReq.SetParameters(queryParams)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper NetworkDeviceStatusWrapper
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

// GetStorageDeviceStatus Retrieves status for a particular storage device
// Returns: A StorageDeviceStatus entity describing the status of the requested device
func (client *LiqidClient) GetStorageDeviceStatus(deviceId string) (*StorageDeviceStatus, error) {
	var fn = "GetStorageDeviceStatus"
	client.traceLogger.Printf(LogFmtEnter1, fn, deviceId)
	list, getErr := client.GetStorageDevicesStatus()
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}
	for _, item := range list {
		if deviceId == item.DeviceId {
			client.traceLogger.Printf(LogFmtReturn2, fn, item, nil)
			return &item, nil
		}
	}
	err := fmt.Errorf("cannot find requested item")
	client.traceLogger.Printf(LogFmtReturn2, fn, nil, err)
	return nil, err
}

// GetStorageDevicesStatus Retrieves status for all storage devices
// Returns: An array of StorageDeviceStatus entities describing the status of storage devices in the system
func (client *LiqidClient) GetStorageDevicesStatus() ([]StorageDeviceStatus, error) {
	var fn = "GetStorageDevicesStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "status/storage"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	queryParams := make(map[string]string)
	queryParams["filter"] = "false"
	cliReq.SetParameters(queryParams)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper StorageDeviceStatusWrapper
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
