// File: powerManagement_gen.go
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

// Category: PowerManagement
// Manages power states for a node

// RebootNode Reboots a particular node
//
// Param groupId: Unique identifier of the group which contains the node of interest
// Param machineId: Unique machine identifier containing the node of interest
// Returns: A Machine entity describing the node
func (client *LiqidClient) RebootNode(groupId int32, machineId int32) (*Machine, error) {
	var fn = "RebootNode"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, machineId)

	path := "power/reboot"
	path += "/" + fmt.Sprintf("%v", groupId)
	path += "/" + fmt.Sprintf("%v", machineId)
	cliReq := NewClientRequest()
	cliReq.SetMethod("PUT")
	cliReq.SetPartialPath(path)
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

// RestartNode Restarts a particular node
//
// Param groupId: Unique identifier of the group which contains the node of interest
// Param machineId: Unique machine identifier containing the node of interest
// Returns: A Machine entity describing the node
func (client *LiqidClient) RestartNode(groupId int32, machineId int32) (*Machine, error) {
	var fn = "RestartNode"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, machineId)

	path := "power/restart"
	path += "/" + fmt.Sprintf("%v", groupId)
	path += "/" + fmt.Sprintf("%v", machineId)
	cliReq := NewClientRequest()
	cliReq.SetMethod("PUT")
	cliReq.SetPartialPath(path)
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

// ShutdownNode Shuts down a particular node
//
// Param groupId: Unique identifier of the group which contains the node of interest
// Param machineId: Unique machine identifier containing the node of interest
// Returns: A Machine entity describing the node
func (client *LiqidClient) ShutdownNode(groupId int32, machineId int32) (*Machine, error) {
	var fn = "ShutdownNode"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, machineId)

	path := "power/shutdown"
	path += "/" + fmt.Sprintf("%v", groupId)
	path += "/" + fmt.Sprintf("%v", machineId)
	cliReq := NewClientRequest()
	cliReq.SetMethod("PUT")
	cliReq.SetPartialPath(path)
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

// StartNode Starts a particular node
//
// Param groupId: Unique identifier of the group which contains the node of interest
// Param machineId: Unique machine identifier containing the node of interest
// Returns: A Machine entity describing the node
func (client *LiqidClient) StartNode(groupId int32, machineId int32) (*Machine, error) {
	var fn = "StartNode"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, machineId)

	path := "power/start"
	path += "/" + fmt.Sprintf("%v", groupId)
	path += "/" + fmt.Sprintf("%v", machineId)
	cliReq := NewClientRequest()
	cliReq.SetMethod("PUT")
	cliReq.SetPartialPath(path)
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
