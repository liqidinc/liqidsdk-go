// File: ssh_gen.go
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

// Category: Ssh
// Includes functions which get and set SSH status

// GetSSHStatus Retrieves the current state of SSH
// Returns: A SSHStatus entity describing the state of SSH
func (client *LiqidClient) GetSSHStatus() (*SSHStatus, error) {
	var fn = "GetSSHStatus"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "ssh/enable"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper SSHStatusWrapper
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

// SetSSHStatus Sets the state of SSH
//
// Param active: Indicates whether SSH should be active
// Param enabled: Indicates whether SSH should be enabled
// Returns: A SSHStatus entity describing the state of SSH
func (client *LiqidClient) SetSSHStatus(active bool, enabled bool) (*SSHStatus, error) {
	var fn = "SetSSHStatus"
	client.traceLogger.Printf(LogFmtEnter2, fn, active, enabled)

	path := "ssh/enable"
	cliReq := NewClientRequest().SetMethod("POST").SetPartialPath(path)
	apiParam := NewSSHStatus()
	apiParam.Active = active
	apiParam.Enabled = enabled
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper SSHStatusWrapper
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
