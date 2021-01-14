// Code generated by go-swagger; DO NOT EDIT.

package data_driver_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new data driver config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data driver config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DataDriverConfigCreate creates a data driver config parameters ntnx ignore

Create a data driver config parameters.
*/
func (a *Client) DataDriverConfigCreate(params *DataDriverConfigCreateParams, authInfo runtime.ClientAuthInfoWriter) (*DataDriverConfigCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataDriverConfigCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DataDriverConfigCreate",
		Method:             "POST",
		PathPattern:        "/v1.0/datadriverconfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataDriverConfigCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataDriverConfigCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataDriverConfigCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataDriverConfigDelete deletes a specific data driver config parameters ntnx ignore

Delete a data driver config parameters with a given ID {id}.
*/
func (a *Client) DataDriverConfigDelete(params *DataDriverConfigDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*DataDriverConfigDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataDriverConfigDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DataDriverConfigDelete",
		Method:             "DELETE",
		PathPattern:        "/v1.0/datadriverconfigs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataDriverConfigDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataDriverConfigDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataDriverConfigDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataDriverConfigGet gets a data driver config parameters by ID ntnx ignore

Get a data driver config parameters according to its given ID {id}.
*/
func (a *Client) DataDriverConfigGet(params *DataDriverConfigGetParams, authInfo runtime.ClientAuthInfoWriter) (*DataDriverConfigGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataDriverConfigGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DataDriverConfigGet",
		Method:             "GET",
		PathPattern:        "/v1.0/datadriverconfigs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataDriverConfigGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataDriverConfigGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataDriverConfigGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataDriverConfigList gets a data driver config parameters by ID ntnx ignore

Get a data driver config parameters according to its given ID {id}.
*/
func (a *Client) DataDriverConfigList(params *DataDriverConfigListParams, authInfo runtime.ClientAuthInfoWriter) (*DataDriverConfigListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataDriverConfigListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DataDriverConfigList",
		Method:             "GET",
		PathPattern:        "/v1.0/datadriverconfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataDriverConfigListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataDriverConfigListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataDriverConfigListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataDriverConfigUpdate updates a data driver config parameters ntnx ignore

Update a data driver config parameters.
*/
func (a *Client) DataDriverConfigUpdate(params *DataDriverConfigUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*DataDriverConfigUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataDriverConfigUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DataDriverConfigUpdate",
		Method:             "PUT",
		PathPattern:        "/v1.0/datadriverconfigs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataDriverConfigUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataDriverConfigUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataDriverConfigUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}