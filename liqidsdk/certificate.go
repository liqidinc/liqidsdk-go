//
// Copyright (c) 2019-2022 Liqid, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are not permitted without prior consent.
//

package liqidsdk

import (
	"fmt"
	"liqidsdk/zip"
	"net/http"
)

const (
	certificateBasePath = "certificate"
	CSRPath             = certificateBasePath + "/csr"
)

type CreateCSRRequestParameters struct {
	CommonName       string
	OrganizationUnit string
	Organization     string
	Locality         string
	StateOrProvince  string
	Country          string
}

// CreateCSRRequest asks the director to create CSR and KEY files
// Returns a first byte array containing the CSR content and a second byte array containing the KEY content.
func (client *LiqidClient) CreateCSRRequest(params *CreateCSRRequestParameters) ([]byte, []byte, error) {
	paramMap := make(map[string]string)
	if len(params.CommonName) > 0 {
		paramMap["commonName"] = params.CommonName
	}
	if len(params.OrganizationUnit) > 0 {
		paramMap["organizationUnit"] = params.OrganizationUnit
	}
	if len(params.Organization) > 0 {
		paramMap["organization"] = params.Organization
	}
	if len(params.Locality) > 0 {
		paramMap["locality"] = params.Locality
	}
	if len(params.StateOrProvince) > 0 {
		paramMap["state"] = params.StateOrProvince
	}
	if len(params.Country) > 0 {
		paramMap["country"] = params.Country
	}

	cliReq := NewClientRequest().SetMethod(http.MethodGet).SetPartialPath(CSRPath).SetParameters(paramMap)
	err := cliReq.Invoke(client)
	if err != nil {
		return nil, nil, err
	}

	container, err := zip.Unzip(cliReq.GetResponse())
	if err != nil {
		return nil, nil, err
	}

	var csrFile *zip.File
	var keyFile *zip.File

	for i := 0; i < len(container.Root.Content); i++ {
		file := container.Root.Content[i]
		if file.FileHeader.FileName == "liqid.csr" {
			csrFile = &file
		} else if file.FileHeader.FileName == "liqid.key" {
			keyFile = &file
		}
	}

	if csrFile == nil {
		return nil, nil, fmt.Errorf("no csr file returned")
	}
	if keyFile == nil {
		return nil, nil, fmt.Errorf("no key file returned")
	}

	return csrFile.Content, keyFile.Content, nil
}
