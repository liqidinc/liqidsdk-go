// File: manager_gen.go
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

// Category: Manager
// These functions manage the managed-entity descriptors which are used for discovering devices.

// GetDiscoveryTokens Returns all of the configured discovery tokens
// Returns: An array of DiscoveryToken entities
func (client *LiqidClient) GetDiscoveryTokens() ([]DiscoveryToken, error) {
	var fn = "GetDiscoveryTokens"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "manager/discovery"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper DiscoveryTokenWrapper
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

// GetDiscoveryTokensByType Returns a subset of the configured discovery tokens, specified by the TokenType parameter
//
// Param tokenType: The type of tokens requested
// Returns: An array of DiscoveryToken entities
func (client *LiqidClient) GetDiscoveryTokensByType(tokenType TokenType) ([]DiscoveryToken, error) {
	var fn = "GetDiscoveryTokensByType"
	client.traceLogger.Printf(LogFmtEnter1, fn, tokenType)

	path := "manager/discovery"
	path += "/" + fmt.Sprintf("%v", tokenType)
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper DiscoveryTokenWrapper
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

// GetNetworkIPMIManagedCPU Retrieves a particular network-IPMI-managed CPU given its name
// Returns: A NetworkManagedCPU entity
func (client *LiqidClient) GetNetworkIPMIManagedCPU(cpuName string) (*NetworkManagedCPU, error) {
	var fn = "GetNetworkIPMIManagedCPU"
	client.traceLogger.Printf(LogFmtEnter1, fn, cpuName)
	list, getErr := client.GetNetworkIPMIManagedCPUS()
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}
	for _, item := range list {
		if cpuName == item.CPUName {
			client.traceLogger.Printf(LogFmtReturn2, fn, item, nil)
			return &item, nil
		}
	}
	err := fmt.Errorf("cannot find requested item")
	client.traceLogger.Printf(LogFmtReturn2, fn, nil, err)
	return nil, err
}

// GetNetworkIPMIManagedCPUS Retrieves a list of management entries for IPMI-via-network CPUs
// Returns: An array of NetworkManagedCPU entities
func (client *LiqidClient) GetNetworkIPMIManagedCPUS() ([]NetworkManagedCPU, error) {
	var fn = "GetNetworkIPMIManagedCPUS"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "manager/network/ipmi/cpu"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper NetworkManagedCPUWrapper
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

// GetNetworkIPMIManagedEnclosure Retrieves a particular network-IPMI-managed enclosure given its name
// Returns: A NetworkManagedEnclosure entity
func (client *LiqidClient) GetNetworkIPMIManagedEnclosure(enclosureName string) (*NetworkManagedEnclosure, error) {
	var fn = "GetNetworkIPMIManagedEnclosure"
	client.traceLogger.Printf(LogFmtEnter1, fn, enclosureName)
	list, getErr := client.GetNetworkIPMIManagedEnclosures()
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}
	for _, item := range list {
		if enclosureName == item.EnclosureName {
			client.traceLogger.Printf(LogFmtReturn2, fn, item, nil)
			return &item, nil
		}
	}
	err := fmt.Errorf("cannot find requested item")
	client.traceLogger.Printf(LogFmtReturn2, fn, nil, err)
	return nil, err
}

// GetNetworkIPMIManagedEnclosures Retrieves a list of management entries for IPMI-via-network enclosures
// Returns: An array of NetworkManagedEnclosure entities
func (client *LiqidClient) GetNetworkIPMIManagedEnclosures() ([]NetworkManagedEnclosure, error) {
	var fn = "GetNetworkIPMIManagedEnclosures"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "manager/network/ipmi/enclosure"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper NetworkManagedEnclosureWrapper
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

// GetManagedEntities Reports all the managed entity entries used for discovering PCI devices
// Returns: Returns an array of ManagedEntity items describing the entries which are used for device discovery
func (client *LiqidClient) GetManagedEntities() ([]ManagedEntity, error) {
	var fn = "GetManagedEntities"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "manager/device"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper ManagedEntityWrapper
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

// GetManagedEntity Retrieves a particular managed entity
// Returns: A ManagedEntity entity describing the requested entity
func (client *LiqidClient) GetManagedEntity(pciVendorId string, pciDeviceId string) (*ManagedEntity, error) {
	var fn = "GetManagedEntity"
	client.traceLogger.Printf(LogFmtEnter2, fn, pciVendorId, pciDeviceId)
	list, getErr := client.GetManagedEntities()
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}
	for _, item := range list {
		if pciVendorId == item.PCIVendorId && pciDeviceId == item.PCIDeviceId {
			client.traceLogger.Printf(LogFmtReturn2, fn, item, nil)
			return &item, nil
		}
	}
	err := fmt.Errorf("cannot find requested item")
	client.traceLogger.Printf(LogFmtReturn2, fn, nil, err)
	return nil, err
}
