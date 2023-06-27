// File: group_gen.go
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

// Category: Group
// Includes functions related to the configuration of a group.

// ClearGroups Deletes all groups and all machines within those groups.
//             Restores all devices to the system free pool.
//             This is effectively a soft configuration reset which does not rediscover devices.
// Returns: Returns true if successful, throws an error otherwise.
func (client *LiqidClient) ClearGroups() (bool, error) {
	var fn = "ClearGroups"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "group/clear"
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return false, cliErr
	}

	var respWrapper BoolWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, false, jsonErr)
		return false, jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, false, fmtErr)
		return false, fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, false, fmtErr)
		return false, fmtErr
	}

	result := respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, result, nil)
	return result, nil
}

// CreateGroupWithId Creates a new group within the current fabric.
//
// Param groupId: Unique identifier to be associated with the group
// Param groupName: Unique name of the group
// Returns: A Group entity which describes the newly-created group
//
// Deprecated: SDK Clients should use CreateGroup instead, which does not require that the client specify a group id.
func (client *LiqidClient) CreateGroupWithId(groupId int32, groupName string) (*Group, error) {
	var fn = "CreateGroupWithId"
	client.traceLogger.Printf(LogFmtEnter2, fn, groupId, groupName)

	path := "group"
	cliReq := NewClientRequest().SetMethod("POST").SetPartialPath(path)
	apiParam := NewGroup()
	apiParam.GroupId = groupId
	apiParam.GroupName = groupName
	var presetErr error
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

	var respWrapper GroupWrapper
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

// DeleteGroup Deletes a configured group along with all the machines in the group.
//             Returns all devices related to the group or to the machines in the group to the system free pool.
//
// Param groupId: Unique identifier of the group to be deleted
// Returns: Returns a Group item describing the deleted group
func (client *LiqidClient) DeleteGroup(groupId int32) (*Group, error) {
	var fn = "DeleteGroup"
	client.traceLogger.Printf(LogFmtEnter1, fn, groupId)

	path := "group"
	path += "/" + fmt.Sprintf("%v", groupId)
	cliReq := NewClientRequest()
	cliReq.SetMethod("DELETE")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupWrapper
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

// GetGroup Retrieves a particular group given its identifier
// Returns: A Group entity describing the indicated group
func (client *LiqidClient) GetGroup(groupId int32) (*Group, error) {
	var fn = "GetGroup"
	client.traceLogger.Printf(LogFmtEnter1, fn, groupId)
	list, getErr := client.GetGroups()
	if getErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, getErr)
		return nil, getErr
	}
	for _, item := range list {
		if groupId == item.GroupId {
			client.traceLogger.Printf(LogFmtReturn2, fn, item, nil)
			return &item, nil
		}
	}
	err := fmt.Errorf("cannot find requested item")
	client.traceLogger.Printf(LogFmtReturn2, fn, nil, err)
	return nil, err
}

// GetGroups Retrieves all configured groups
// Returns: An array of Group entities describing the configured groups
func (client *LiqidClient) GetGroups() ([]Group, error) {
	var fn = "GetGroups"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "group"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupWrapper
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

// GetGroupDetails Retrieves details regarding a particular configured group
//
// Param groupId: Unique identifier of the group of interest
// Returns: A GroupDetails entity containing details for the indicated group
func (client *LiqidClient) GetGroupDetails(groupId int32) (*GroupDetails, error) {
	var fn = "GetGroupDetails"
	client.traceLogger.Printf(LogFmtEnter1, fn, groupId)

	path := "group/details"
	path += "/" + fmt.Sprintf("%v", groupId)
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper GroupDetailsWrapper
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

// GetGroupIdByName Retrieves the unique identifier of a configured group given its name
//
// Param groupName: Name of the group of interest
// Returns: Returns the unique identifier of the indicated group
func (client *LiqidClient) GetGroupIdByName(groupName string) (int32, error) {
	var fn = "GetGroupIdByName"
	client.traceLogger.Printf(LogFmtEnter1, fn, groupName)

	path := "group/mapper"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	queryParams := make(map[string]string)
	queryParams["group-name"] = fmt.Sprintf("%v", groupName)
	cliReq.SetParameters(queryParams)
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

// GetNextGroupId Retrieves the next sequential unused group identifier
// Returns: The next sequential unused group identifier
//
// Deprecated: Use of this function in conjunction with CreateGroupWithId opens a potential race condition.
//             This problem exists internally in the Director, and is reflected both in the REST API and in this SDK.
//             A future version of the Director will provide a means of creating a group whereby the group id is internally generated.
//             For this reason, SDK clients are encouraged to use the CreateGroup function which wraps the Get/Create mechanism in one SDK function call.
//             For the time being, this does not correct the race condition; however, it protects SDK clients from the eventual removal of CreateGroupWithId and GetNextGroupId.
func (client *LiqidClient) GetNextGroupId() (int32, error) {
	var fn = "GetNextGroupId"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "group/nextid"
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
