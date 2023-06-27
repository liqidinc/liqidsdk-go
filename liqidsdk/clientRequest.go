package liqidsdk

import (
	"bytes"
	"encoding/json"
)

type ClientRequest struct {
	httpMethod  string
	parameters  map[string]string
	partialPath string
	body        []byte
	jsonBody    interface{}
	response    []byte
	contentType string
}

func NewClientRequest() *ClientRequest {
	return &ClientRequest{}
}

func (req *ClientRequest) SetMethod(method string) *ClientRequest {
	req.httpMethod = method
	return req
}

func (req *ClientRequest) SetParameters(parameters map[string]string) *ClientRequest {
	req.parameters = parameters
	return req
}

func (req *ClientRequest) SetPartialPath(partialPath string) *ClientRequest {
	req.partialPath = partialPath
	return req
}

func (req *ClientRequest) SetBody(body []byte) *ClientRequest {
	req.body = body
	req.contentType = "plain/text"
	return req
}

func (req *ClientRequest) SetContentType(contentType string) *ClientRequest {
	req.contentType = contentType
	return req
}

func (req *ClientRequest) SetJSONBody(jsonRequest interface{}) *ClientRequest {
	req.jsonBody = jsonRequest
	req.contentType = "application/json"
	return req
}

func (req *ClientRequest) Invoke(client *LiqidClient) error {
	if req.jsonBody != nil {
		buffer := new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(&req.jsonBody)
		if err != nil {
			return err
		}
		req.body = buffer.Bytes()
	}

	resp, err := (*client).do(req.httpMethod, req.partialPath, req.parameters, req.body, req.contentType)
	if err != nil {
		return err
	}

	req.response = resp
	return err
}

func (req *ClientRequest) GetResponse() []byte {
	return req.response
}

func (req *ClientRequest) GetJSONResponse(jsonResponse interface{}) error {
	if req.response != nil && len(req.response) > 0 {
		return json.NewDecoder(bytes.NewReader(req.response)).Decode(jsonResponse)
	}

	return nil
}
