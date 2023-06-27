//
// Copyright (c) 2019-2022 Liqid, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are not permitted without prior consent.
//

package liqidsdk

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	apiBasePath       = "/liqid/api/v2/"
	defaultPort       = 8080
	defaultSecurePort = 8443
	defaultTimeout    = 15 //	seconds
	JSONContentType   = "application/json;charset=UTF-8"
)

const (
	LogErrorBit = 0x01
	LogWarnBit  = 0x02
	LogInfoBit  = 0x04
	LogDebugBit = 0x08
	LogTraceBit = 0x10
)

const (
	LogError = LogErrorBit
	LogWarn  = LogWarnBit | LogError
	LogInfo  = LogInfoBit | LogWarn
	LogDebug = LogDebugBit | LogInfo
	LogTrace = LogTraceBit | LogDebug
)

const (
	LogFmtEnter0  = "ENTER %s()"
	LogFmtEnter1  = "ENTER %s(%#v)"
	LogFmtEnter2  = "ENTER %s(%#v, %#v)"
	LogFmtEnter3  = "ENTER %s(%#v, %#v, %#v)"
	LogFmtEnter4  = "ENTER %s(%#v, %#v, %#v, %#v)"
	LogFmtEnter5  = "ENTER %s(%#v, %#v, %#v, %#v, %#v)"
	LogFmtEnter6  = "ENTER %s(%#v, %#v, %#v, %#v, %#v, %#v)"
	LogFmtEnter19 = "ENTER %s(%#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v)"
	LogFmtReturn0 = "%s RETURNING"
	LogFmtReturn1 = "%s RETURNING %#v"
	LogFmtReturn2 = "%s RETURNING %#v, %#v"
)

type AuthenticationMode int

var p2pSwitchFromFlag = map[bool]P2PType{
	true:  P2PTypeOn,
	false: P2PTypeOff,
}

// This interface allows unit tests to mock out the real http library
type iClient interface {
	Do(request *http.Request) (*http.Response, error)
}

// LiqidClient - this is the liqid client struct, which wraps a client struct.
//	SDK users will use this struct, obtained via NewLiqidClient (below)
type LiqidClient struct {
	useSecureHTTP      bool   //	true for HTTPS, false for HTTP
	host               string //	DNS name or IP address (do NOT include scheme/protocol)
	port               int
	basicAuthUsername  string
	basicAuthPassword  string
	authLabel          string
	authToken          string
	ignoreCertificates bool
	scheme             string //	auto-set to "http" or "https"
	timeout            int    //	in seconds
	httpClient         iClient

	logLevel      int
	debugLogger   *log.Logger
	errorLogger   *log.Logger
	infoLogger    *log.Logger
	traceLogger   *log.Logger
	warningLogger *log.Logger

	hasCurrentData     bool
	currentCoordinates Coordinates
	currentFabricId    int32
}

// NewLiqidClient creates a new LiqidClient struct - it must be separately initialized.
//	this convention allows us to require a secure flag and host string, defaulting everything else,
//	but allowing the caller to update the defaulted values before actual initialization.
func NewLiqidClient(secure bool, host string) *LiqidClient {
	result := LiqidClient{useSecureHTTP: secure, host: host}
	result.SetLogging(os.Stdout, log.Lshortfile, LogInfo)

	if secure {
		result.scheme = "https"
		result.port = defaultSecurePort
	} else {
		result.scheme = "http"
		result.port = defaultPort
	}

	result.basicAuthUsername = ""
	result.basicAuthPassword = ""
	result.authToken = ""
	result.ignoreCertificates = false
	result.timeout = defaultTimeout
	result.hasCurrentData = false

	return &result
}

// NewLiqidClientForTest sets up a LiqidClient with a mocked-out http.
func NewLiqidClientForTest(secure bool, host string, httpClient iClient) *LiqidClient {
	result := NewLiqidClient(secure, host)
	result.httpClient = httpClient
	return result
}

//	Logging -------------------------------------------------------------------

//	SetLogging allows the SDK client to change the current logging parameters.
//	may be invoked prior to Initialize()
func (client *LiqidClient) SetLogging(file io.Writer, flags int, level int) {
	discard := log.New(ioutil.Discard, "", 0)
	if level&LogDebug != 0 {
		client.debugLogger = log.New(file, "DEBUG:", flags|log.Lshortfile)
	} else {
		client.debugLogger = discard
	}

	if level&LogError != 0 {
		client.errorLogger = log.New(file, "ERROR:", flags|log.Lshortfile)
	} else {
		client.errorLogger = discard
	}

	if level&LogInfo != 0 {
		client.infoLogger = log.New(file, "INFO: ", flags|log.Lshortfile)
	} else {
		client.infoLogger = discard
	}

	if level&LogTrace != 0 {
		client.traceLogger = log.New(file, "TRACE:", flags|log.Lshortfile)
	} else {
		client.traceLogger = discard
	}

	if level&LogTrace != 0 {
		client.warningLogger = log.New(file, "WARN: ", flags|log.Lshortfile)
	} else {
		client.warningLogger = discard
	}

	client.logLevel = level
}

// Initialize sets up the httpClient struct for use.
//	It should NOT be invoked by unit tests, as it would override the mock http client which the
//	unit unittest should have put in place after invoking NewLiqidClientForTest.
func (client *LiqidClient) Initialize() error {
	var fn = "Initialize"
	client.traceLogger.Printf(LogFmtEnter0, fn)
	hc := http.Client{}
	hc.Timeout = time.Second * time.Duration(client.timeout)
	if client.ignoreCertificates {
		hc.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}
	client.httpClient = &hc
	client.traceLogger.Printf(LogFmtReturn0, fn)
	return nil
}

//	API Authentication --------------------------------------------------------

//	ClearCredentials clears out the base-64 encoded authentication string previously set
func (client *LiqidClient) ClearCredentials() error {
	var fn = "ClearCredentials"
	client.traceLogger.Printf(LogFmtEnter0, fn)
	var err error

	if client.HasCredentials() {
		client.basicAuthUsername = ""
		client.basicAuthPassword = ""
	} else {
		err = fmt.Errorf("credentials were not established")
	}

	client.traceLogger.Printf(LogFmtReturn1, fn, err)
	return err
}

//	SetCredentials establishes the base-64 encoded basic authentication string to be sent
//	in the headers of future HTTP requests.
func (client *LiqidClient) SetCredentials(username string, password string) error {
	var fn = "SetCredentials"
	client.traceLogger.Printf(LogFmtEnter2, fn, username, "********")
	var err error

	if client.HasCredentials() {
		err = fmt.Errorf("client already has credentials established")
		goto end
	} else if client.IsLoggedIn() {
		err = fmt.Errorf("client is already logged in")
		goto end
	}

	client.basicAuthUsername = username
	client.basicAuthPassword = password

end:
	client.traceLogger.Printf(LogFmtReturn1, fn, err)
	return err
}

func (client *LiqidClient) HasCredentials() bool {
	return len(client.basicAuthUsername) > 0
}

//	Login establishes a token with the director, for the given application.
//	This token is subsequently used for authentication on all future http requests.
func (client *LiqidClient) Login(label string, username string, password string) error {
	var fn = "Login"
	client.traceLogger.Printf(LogFmtEnter3, fn, label, username, "********")
	var err error
	var digestInfo *DigestInfo

	if client.HasCredentials() {
		err = fmt.Errorf("client already has credentials established")
		goto end
	} else if client.IsLoggedIn() {
		err = fmt.Errorf("client is already logged in")
		goto end
	} else if label == "slurm" {
		//	Protect permanent slurm token
		err = fmt.Errorf("permission denied")
		goto end
	}

	//	Use basic auth just long enough to generate a digest / label / whatever
	err = client.SetCredentials(username, password)
	if err != nil {
		goto end
	}

	//	Remove any left-over token for this label
	_, _ = client.DeleteMessageDigest(label)

	//	Create a new digest token
	digestInfo, err = client.CreateMessageDigest(label)
	_ = client.ClearCredentials()
	if err != nil {
		goto end
	}

	client.authToken = digestInfo.DigestKey
	client.authLabel = label

end:
	client.traceLogger.Printf(LogFmtReturn1, fn, err)
	return err
}

//	Logout deletes the bearer token from the director, and clears it from this client.
func (client *LiqidClient) Logout() error {
	var fn = "Logout"
	client.traceLogger.Printf(LogFmtEnter0, fn)
	var err error

	if client.IsLoggedIn() {
		_, err = client.DeleteMessageDigest(client.authLabel)
		if err != nil {
			goto end
		}

		client.authToken = ""
		client.authLabel = ""
		client.traceLogger.Printf("Logout RETURNING nil")
	} else {
		err = fmt.Errorf("client was not logged-in")
	}

end:
	client.traceLogger.Printf(LogFmtReturn1, fn, err)
	return err
}

func (client *LiqidClient) IsLoggedIn() bool {
	return len(client.authToken) > 0
}

func (client *LiqidClient) getFullPath(partialPath string) string {
	result := fmt.Sprintf("%s://%s:%d/%s/%s", client.scheme, client.host, client.port, apiBasePath, partialPath)
	return result
}

//	API invocation ------------------------------------------------------------

func (client *LiqidClient) URIFromPartialPathAndParameters(partialPath string, parameters map[string]string) string {
	uri := fmt.Sprintf("%s://%s:%v%s%s", client.scheme, client.host, client.port, apiBasePath, partialPath)
	if parameters != nil {
		prefix := "?"
		for key, value := range parameters {
			uri += fmt.Sprintf("%s%s=%s", prefix, key, value)
			prefix = "&"
		}
	}

	return uri
}

func (client *LiqidClient) do(method string, partialPath string, parameters map[string]string, body []byte, contentType string) ([]byte, error) {
	if client.httpClient == nil {
		err := fmt.Errorf("liqid client was not initialized")
		return nil, err
	}

	var reader io.Reader = nil
	if body != nil && len(body) > 0 {
		reader = bytes.NewReader(body)
	}

	uri := client.URIFromPartialPathAndParameters(partialPath, parameters)
	req, err := http.NewRequest(method, uri, reader)
	if err != nil {
		return nil, err
	}

	if len(contentType) > 0 {
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Content-Length", "0")
		req.Header.Set("accept", contentType)
	}

	if client.IsLoggedIn() {
		req.Header.Add("X-API-Key", client.authToken)
	} else if client.HasCredentials() {
		req.SetBasicAuth(client.basicAuthUsername, client.basicAuthPassword)
	}

	client.debugLogger.Printf("===>%v:%v", method, uri)
	for k, v := range req.Header {
		client.debugLogger.Printf("   >%q: %q", k, v)
	}
	if body != nil {
		client.debugLogger.Printf("   >%v", string(body))
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		if Body != nil {
			err := Body.Close()
			if err != nil {
				client.errorLogger.Printf("Caught %#v", err)
			}
		}
	}(response.Body)

	client.debugLogger.Printf("<===%v:%v", response.StatusCode, response.Status)

	buf := new(bytes.Buffer)
	if response.Body != nil {
		_, err = buf.ReadFrom(response.Body)
		if err != nil {
			return nil, err
		}
	}

	client.debugLogger.Printf("<===%v", buf.String())
	if response.StatusCode < 200 || response.StatusCode > 299 {
		err = fmt.Errorf("received http response %v:%s", response.StatusCode, response.Status)
	}

	return buf.Bytes(), err
}

// FormatError is a convenience function for converting the messy JSON code and ErrorMessage structs into
//	a single readable string for logging purposes
func FormatError(code int, errors []ErrorMessage) string {
	msg := fmt.Sprintf("json outer code=%d errors=", code)
	for _, e := range errors {
		msg += fmt.Sprintf("[%s:%d:%s]", e.Type, e.Code, e.Message)
	}
	return msg
}

//	grab current values --------------------------------------------------------

func (client *LiqidClient) getCurrentValues() error {
	if !client.hasCurrentData {
		coordinates, err := client.GetDefaultCoordinates()
		if err != nil {
			return err
		}
		client.currentCoordinates = *coordinates

		client.currentFabricId, err = client.GetCurrentFabricId()
		if err != nil {
			return err
		}

		client.hasCurrentData = true
	}

	return nil
}

func (client *LiqidClient) getCurrentRack() (int32, error) {
	return client.currentCoordinates.Rack, client.getCurrentValues()
}

func (client *LiqidClient) getCurrentShelf() (int32, error) {
	return client.currentCoordinates.Shelf, client.getCurrentValues()
}

func (client *LiqidClient) getCurrentNode() (int32, error) {
	return client.currentCoordinates.Node, client.getCurrentValues()
}

func (client *LiqidClient) getCurrentFabricId() (int32, error) {
	return client.currentFabricId, client.getCurrentValues()
}

//	hard-coded SDK functions ---------------------------------------------------

// CreateGroup Creates a new group.
//  groupName (req): unique name for the group
//  returns: description of the newly-created group
func (client *LiqidClient) CreateGroup(groupName string) (*Group, error) {
	var fn = "CreateGroup"
	client.traceLogger.Printf(LogFmtEnter1, fn, groupName)
	var group *Group

	var err error
	var groupId int32

	groupId, err = client.GetNextGroupId()
	if err != nil {
		goto end
	}

	group, err = client.CreateGroupWithId(groupId, groupName)

end:
	client.traceLogger.Printf(LogFmtReturn2, fn, groupId, err)
	return group, err
}

// DeleteMessageDigest
// Deletes a previously-created message digest.
// Param label: The label associated with the digest which is to be deleted
// Returns: Returns the label associated with the deleted message digest
func (client *LiqidClient) DeleteMessageDigest(label string) (string, error) {
	var fn = "DeleteMessageDigest"
	client.traceLogger.Printf(LogFmtEnter1, fn, label)

	path := "digest"
	path += "/" + fmt.Sprintf("%v", label)
	cliReq := NewClientRequest().SetMethod("DELETE").SetPartialPath(path)

	cliErr := cliReq.Invoke(client)
	if cliErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, "", cliErr)
		return "", cliErr
	}

	var respWrapper StringWrapper
	jsonErr := cliReq.GetJSONResponse(&respWrapper)
	if jsonErr != nil {
		client.traceLogger.Printf(LogFmtReturn2, fn, "", jsonErr)
		return "", jsonErr
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		fmtErr := fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		client.traceLogger.Printf(LogFmtReturn2, fn, "", fmtErr)
		return "", fmtErr
	}

	if len(respWrapper.Response.Data) != 1 {
		fmtErr := fmt.Errorf("malformed data encountered - we should have exactly one data entity")
		client.traceLogger.Printf(LogFmtReturn2, fn, "", fmtErr)
		return "", fmtErr
	}

	result := respWrapper.Response.Data[0]
	client.traceLogger.Printf(LogFmtReturn2, fn, result, nil)
	return result, nil
}

func (client *LiqidClient) GetNextMachineId() (int32, error) {
	var fn = "GetNextMachineId"
	client.traceLogger.Printf(LogFmtEnter0)
	var result int32
	var val int
	var err error
	var respWrapper StringWrapper

	path := "machine/nextid"
	cliReq := NewClientRequest().SetMethod(http.MethodGet).SetPartialPath(path)

	if err = cliReq.Invoke(client); err != nil {
		goto end
	}

	err = cliReq.GetJSONResponse(&respWrapper)
	if err != nil {
		goto end
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		err = fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		goto end
	}

	val, err = strconv.Atoi(respWrapper.Response.Data[0])
	result = int32(val)

end:
	client.traceLogger.Printf(LogFmtReturn2, fn, result, err)
	return result, err
}

func (client *LiqidClient) EnableP2PForMachine(machineId int32, enabledFlag bool) error {
	var fn = "EnableP2PForMachine"
	client.traceLogger.Printf(LogFmtEnter2, machineId, enabledFlag)
	var err error
	var machine Machine
	reqWrapper := MachineWrapper{}
	var respWrapper MachineWrapper

	//	We have to read the entire Machine struct from the controller first,
	//	update the flag, then rewrite it.
	cliReq := NewClientRequest().SetMethod(http.MethodGet).SetPartialPath("machine")
	queryParams := make(map[string]string)
	queryParams["mach_id"] = string(machineId)
	cliReq.SetParameters(queryParams)

	if err = cliReq.Invoke(client); err != nil {
		goto end
	}

	if err = cliReq.GetJSONResponse(&respWrapper); err != nil {
		goto end
	}

	if (respWrapper.Response.Code != 0) || (respWrapper.Response.Errors != nil) {
		err = fmt.Errorf("%v", FormatError(respWrapper.Response.Code, respWrapper.Response.Errors))
		goto end
	}

	machine = respWrapper.Response.Data[0]
	machine.P2PEnabled = p2pSwitchFromFlag[enabledFlag]
	reqWrapper.Response.Data = []Machine{machine}

	err = NewClientRequest().SetMethod(http.MethodPost).SetPartialPath("machine").Invoke(client)

end:
	client.traceLogger.Printf(LogFmtReturn1, fn, err)
	return err
}
