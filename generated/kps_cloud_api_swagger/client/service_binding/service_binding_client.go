// Code generated by go-swagger; DO NOT EDIT.

package service_binding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new service binding API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service binding API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ServiceBindingCreate creates a service binding

Create a Service Binding
*/
func (a *Client) ServiceBindingCreate(params *ServiceBindingCreateParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBindingCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceBindingCreate",
		Method:             "POST",
		PathPattern:        "/v1.0/servicebindings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceBindingCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceBindingCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ServiceBindingDelete deletes a service binding

Delete a Service Binding
*/
func (a *Client) ServiceBindingDelete(params *ServiceBindingDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBindingDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceBindingDelete",
		Method:             "DELETE",
		PathPattern:        "/v1.0/servicebindings/{svcBindingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceBindingDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceBindingDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ServiceBindingGet gets a service binding

Get a Service Binding
*/
func (a *Client) ServiceBindingGet(params *ServiceBindingGetParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBindingGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceBindingGet",
		Method:             "GET",
		PathPattern:        "/v1.0/servicebindings/{svcBindingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceBindingGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceBindingGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ServiceBindingList lists service bindings

List Service Bindings
*/
func (a *Client) ServiceBindingList(params *ServiceBindingListParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBindingListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceBindingList",
		Method:             "GET",
		PathPattern:        "/v1.0/servicebindings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceBindingListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceBindingListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ServiceBindingStatusList gets the status of service binding

Get the status of Service Binding on Service Domains
*/
func (a *Client) ServiceBindingStatusList(params *ServiceBindingStatusListParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBindingStatusListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingStatusListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ServiceBindingStatusList",
		Method:             "GET",
		PathPattern:        "/v1.0/servicebindings/{svcBindingId}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingStatusListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceBindingStatusListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceBindingStatusListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}