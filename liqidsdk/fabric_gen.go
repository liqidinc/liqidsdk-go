// File: fabric_gen.go
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

// Category: Fabric
// These functions relate to the full configuration (fabric) which is managed by the currently-connected Director.

// ApplyFabricChanges applies changes to the fabric - currently one may only add a particular flag with a value
//
// Param flagName: The name of the flag to be added
// Param flagValue: The value for the flag to be added
// Param operation: Currently, only ADD is supported
// Returns: Describes the new toggle state
func (client *LiqidClient) ApplyFabricChanges(flagName string, flagValue string, operation FabricToggleCompositeOption) (*FabricToggleComposite, error) {
	var fn = "ApplyFabricChanges"
	client.traceLogger.Printf(LogFmtEnter3, fn, flagName, flagValue, operation)

	path := "fabric"
	cliReq := NewClientRequest().SetMethod("PUT").SetPartialPath(path)
	apiParam := NewFabricToggleComposite()
	apiParam.ControlFlag.Name = flagName
	apiParam.ControlFlag.ValueString = flagValue
	apiParam.Options = operation
	var presetErr error
	apiParam.Coordinates.Rack, presetErr = client.getCurrentRack()
	if presetErr != nil {
		return nil, presetErr
	}
	apiParam.Coordinates.Shelf, presetErr = client.getCurrentShelf()
	if presetErr != nil {
		return nil, presetErr
	}
	apiParam.Coordinates.Node, presetErr = client.getCurrentNode()
	if presetErr != nil {
		return nil, presetErr
	}
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper FabricToggleCompositeWrapper
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

// CancelEditFabric Cancels fabric edit mode - reverts all pending changes made to device connections
//
// Param machineId: Unique machine identifier
// Returns: A Machine entity describing the machine of interest
func (client *LiqidClient) CancelEditFabric(machineId int32) (*Machine, error) {
	var fn = "CancelEditFabric"
	client.traceLogger.Printf(LogFmtEnter1, fn, machineId)
	object, getErr := client.GetMachine(machineId)
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}

	path := "fabric/edit/cancel"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(object)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper MachineWrapper
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

// CancelReprogramFabric Cancels fabric reprogram mode
//
// Param machineId: Unique machine identifier
// Returns: A Machine entity describing the machine of interest
func (client *LiqidClient) CancelReprogramFabric(machineId int32) (*Machine, error) {
	var fn = "CancelReprogramFabric"
	client.traceLogger.Printf(LogFmtEnter1, fn, machineId)
	object, getErr := client.GetMachine(machineId)
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}

	path := "fabric/reprogram/cancel"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(object)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper MachineWrapper
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

// EditFabric enters fabric edit mode which allows the fabric to be electrically connected
//             - the system must be put into edit mode before a device is added to a machine
//
// Param machineId: Unique machine identifier
// Returns: A Machine entity describing the machine of interest
func (client *LiqidClient) EditFabric(machineId int32) (*Machine, error) {
	var fn = "EditFabric"
	client.traceLogger.Printf(LogFmtEnter1, fn, machineId)
	object, getErr := client.GetMachine(machineId)
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}

	path := "fabric/edit"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(object)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper MachineWrapper
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

// GetCurrentFabricId Retrieves the fabric id associated with the current default coordinates.
//                    This is not generally needed for SDK interactions, but may be useful for troubleshooting.
// Returns: The fabric id which is managed by the connected Director.
func (client *LiqidClient) GetCurrentFabricId() (int32, error) {
	var fn = "GetCurrentFabricId"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "fabric/id"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return 0, cliErr
	}

	var respWrapper Int32Wrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, 0, jsonErr)
		return 0, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, 0, fmtErr)
		return 0, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, 0, fmtErr)
		return 0, fmtErr
	}

	result := respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, result, nil)
	return result, nil
}

// GetFabricTopology Returns the current fabric topology
// Returns: A FabricTopology entity
func (client *LiqidClient) GetFabricTopology() (*FabricTopology, error) {
	var fn = "GetFabricTopology"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "fabric/topology"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper FabricTopologyWrapper
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

// GetFabricTypes Returns all known fabric types and associated id values
// Returns: An array of fabric types and identifiers
func (client *LiqidClient) GetFabricTypes() ([]NameValuePair, error) {
	var fn = "GetFabricTypes"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "fabric/types"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper NameValuePairWrapper
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

// ReprogramFabric Reprogram the fabric. This will result in devices associated with a machine being electrically connected to the machine.
//                 This step MUST be done in order for a device to be added to a machine.
//
// Param machineId: Unique machine identifier
// Returns: A Machine entity
func (client *LiqidClient) ReprogramFabric(machineId int32) (*Machine, error) {
	var fn = "ReprogramFabric"
	client.traceLogger.Printf(LogFmtEnter1, fn, machineId)
	object, getErr := client.GetMachine(machineId)
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}

	path := "fabric/reprogram"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliReq.SetJSONBody(object)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper MachineWrapper
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
