// File: messageDigest_gen.go
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

// Category: MessageDigest
// These functions relate to secure authentication of REST API function calls.
// They are not of use to the SDK clients, as the SDK provides wrapper methods for these functions.
// See commentary on authentication models and functions related to logging in and logging out.

// CreateMessageDigest Creates a new digest associated with a given label.
//                     The digest will be returned to the client, and will not be otherwise exposed by the Director.
//                     This digest should be used for authenticating subsequent REST API invocations.
//                     At the end of the client session, this digest should be deleted by invoking DeleteMessageDigest.
//
// Param label: The label to be associated with a newly-created digest
// Returns: Contains information regarding the created digest.
func (client *LiqidClient) CreateMessageDigest(label string) (*DigestInfo, error) {
	var fn = "CreateMessageDigest"
	client.traceLogger.Printf(LogFmtEnter1, fn, label)

	path := "digest"
	path += "/" + fmt.Sprintf("%v", label)
	cliReq := NewClientRequest()
	cliReq.SetMethod("POST")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper DigestInfoWrapper
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

// GetMessageDigestLabels Retrieves all existing message digest labels.
// Returns: These are labels which can be used for login/logout authentication.
//          The labels are NOT authentication tokens.
func (client *LiqidClient) GetMessageDigestLabels() ([]string, error) {
	var fn = "GetMessageDigestLabels"
	client.traceLogger.Printf(LogFmtEnter0, fn)

	path := "digest"
	cliReq := NewClientRequest()
	cliReq.SetMethod("GET")
	cliReq.SetPartialPath(path)
	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, nil, cliErr)
		return nil, cliErr
	}

	var respWrapper StringWrapper
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
