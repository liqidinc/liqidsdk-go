//
// Copyright (c) 2022 Liqid, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are not permitted without prior consent.
//

package liqidsdk

type ErrorMessage struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type StringArray []string
type StringArrayMap map[string][]string
type StringMap map[string]string

type BoolWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []bool         `json:"data"`
	} `json:"response"`
}

func NewBoolWrapper(value bool) BoolWrapper {
	body := BoolWrapper{}
	body.Response.Data = []bool{value}
	return body
}

type Int32Wrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []int32        `json:"data"`
	} `json:"response"`
}

func NewInt32Wrapper(value int32) Int32Wrapper {
	body := Int32Wrapper{}
	body.Response.Data = []int32{value}
	return body
}

type StringWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []string       `json:"data"`
	} `json:"response"`
}

type StringArrayWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []StringArray  `json:"data"`
	} `json:"response"`
}

type StringArrayMapWrapper struct {
	Response struct {
		Code   int              `json:"code"`
		Errors []ErrorMessage   `json:"errors"`
		Data   []StringArrayMap `json:"data"`
	} `json:"response"`
}

type StringMapWrapper struct {
	Response struct {
		Code   int            `json:"code"`
		Errors []ErrorMessage `json:"errors"`
		Data   []StringMap    `json:"data"`
	} `json:"response"`
}
