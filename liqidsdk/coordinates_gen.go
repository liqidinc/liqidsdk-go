// File: coordinates_gen.go
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

// Category: Coordinates
// These functions relate to internal API mechanisms, and are not particularly useful for the SDK client.
// They are presented here for completeness, and due to the possibility that one or another of them might eventually be needed.

// GetAvailableCoordinates Retrieves a list of entities indicating the known values which may be employed for identifying particular Liqid entities
//                         within a configuration. As a general rule, there will be only one such entry.
// Returns: Array of all known available Liqid coordinates to which REST API invocations may be directed
func (client *LiqidClient) GetAvailableCoordinates() ([]AvailableCoordinates, error) {
	var fn = "GetAvailableCoordinates"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "coordinates/available"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper AvailableCoordinatesWrapper
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

// GetDefaultCoordinates Retrieves the current default coordinates used for the various operations which are initiated by the REST API for the SDK.
// Returns: Liqid coordinates to which all REST API invocations are directed
func (client *LiqidClient) GetDefaultCoordinates() (*Coordinates, error) {
	var fn = "GetDefaultCoordinates"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "coordinates"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper CoordinatesWrapper
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

// SetDefaultCoordinates Sets the default coordinates used for subsequent operations initiated via the SDK.
//                       Specifying any set of values apart from what is available from GetAvailableCoordinates produces undefined behavior.
//
// Param rack: Rack component of the liqid coordinates for the REST API endpoint
//             Should always be zero.
// Param shelf: Shelf component of the liqid coordinates for the REST API endpoint
//              Should always be zero.
// Param node: Node component of the liqid coordinates for the REST API endpoint
// Returns: Liqid coordinates to which all future REST API invocations are directed
func (client *LiqidClient) SetDefaultCoordinates(rack *int32, shelf *int32, node int32) (*Coordinates, error) {
	var fn = "SetDefaultCoordinates"
	client.traceLogger.Printf(LogFmtEnter3, fn, rack, shelf, node)

	path := "coordinates"
	cliReq := NewClientRequest().SetMethod("POST").SetPartialPath(path)
	apiParam := NewCoordinates()
	if rack != nil {
		apiParam.Rack = *rack
	}
	if shelf != nil {
		apiParam.Shelf = *shelf
	}
	apiParam.Node = node
	cliReq.SetJSONBody(apiParam)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper CoordinatesWrapper
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
