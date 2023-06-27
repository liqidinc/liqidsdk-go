// File: groupPool_gen.go
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

// Category: GroupPool
// Functions which manage edit mode for the group pool

// CancelGroupPoolEdit Cancels a pending group pool edit operation.
//                     All pending device attachments or detachments will be canceled.
//
// Param groupId: Identifier of the group for which edit mode is to be canceled
// Returns: A Group entity describing the group pool
func (client *LiqidClient) CancelGroupPoolEdit(groupId int32) (*GroupPool, error) {
	var fn = "CancelGroupPoolEdit"
	client.traceLogger.Printf(LogFmtEnter1, fn, groupId)

	path := "pool/cancel"
	cliReq := NewClientRequest().SetMethod("POST").SetPartialPath(path)
	apiParam := NewGroupPool()
	apiParam.GroupId = groupId
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
	apiParam.FabricId, presetErr = client.getCurrentFabricId()
	if presetErr != nil {
		return nil, presetErr
	}
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupPoolWrapper
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

// GroupPoolDone Completes a pending group pool edit operation.
//               All pending device attachments or detachments are finalized.
//
// Param groupId: Identifier of the group for which edit mode is to be completed
// Returns: A Group entity describing the group pool
func (client *LiqidClient) GroupPoolDone(groupId int32) (*GroupPool, error) {
	var fn = "GroupPoolDone"
	client.traceLogger.Printf(LogFmtEnter1, fn, groupId)

	path := "pool/done"
	cliReq := NewClientRequest().SetMethod("POST").SetPartialPath(path)
	apiParam := NewGroupPool()
	apiParam.GroupId = groupId
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
	apiParam.FabricId, presetErr = client.getCurrentFabricId()
	if presetErr != nil {
		return nil, presetErr
	}
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupPoolWrapper
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

// GroupPoolEdit Establishes edit mode for a particular group pool.
//               Must be invoked before adding devices to or removing devices from a group free pool.
//
// Param groupId: Identifier of the group for which edit mode is to be established
// Returns: A Group entity describing the group pool
func (client *LiqidClient) GroupPoolEdit(groupId int32) (*GroupPool, error) {
	var fn = "GroupPoolEdit"
	client.traceLogger.Printf(LogFmtEnter1, fn, groupId)

	path := "pool/edit"
	cliReq := NewClientRequest().SetMethod("POST").SetPartialPath(path)
	apiParam := NewGroupPool()
	apiParam.GroupId = groupId
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
	apiParam.FabricId, presetErr = client.getCurrentFabricId()
	if presetErr != nil {
		return nil, presetErr
	}
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupPoolWrapper
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
