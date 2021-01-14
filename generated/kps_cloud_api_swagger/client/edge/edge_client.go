// Code generated by go-swagger; DO NOT EDIT.

package edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new edge API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
EdgeCreateV2 creates edge ntnx ignore

Create an edge.
*/
func (a *Client) EdgeCreateV2(params *EdgeCreateV2Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeCreateV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeCreateV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeCreateV2",
		Method:             "POST",
		PathPattern:        "/v1.0/edges",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeCreateV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeCreateV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeCreateV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeDeleteV2 deletes edge by ID ntnx ignore

Deletes the edge with the given ID  {edgeId}.
*/
func (a *Client) EdgeDeleteV2(params *EdgeDeleteV2Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeDeleteV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeDeleteV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeDeleteV2",
		Method:             "DELETE",
		PathPattern:        "/v1.0/edges/{edgeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeDeleteV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeDeleteV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeDeleteV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeGetV2 gets edge by ID ntnx ignore

Retrieves the edge with the given ID {edgeId}.
*/
func (a *Client) EdgeGetV2(params *EdgeGetV2Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeGetV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeGetV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeGetV2",
		Method:             "GET",
		PathPattern:        "/v1.0/edges/{edgeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeGetV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeGetV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeGetV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeListV2 gets edges ntnx ignore

Retrieves all edges for a tenant.
*/
func (a *Client) EdgeListV2(params *EdgeListV2Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeListV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeListV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeListV2",
		Method:             "GET",
		PathPattern:        "/v1.0/edges",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeListV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeListV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeListV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeUpdateV3 updates edge by its ID ntnx ignore

Updates an edge by its ID {id}.
*/
func (a *Client) EdgeUpdateV3(params *EdgeUpdateV3Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeUpdateV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeUpdateV3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeUpdateV3",
		Method:             "PUT",
		PathPattern:        "/v1.0/edges/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeUpdateV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeUpdateV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeUpdateV3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProjectGetEdgesV2 gets project edges by ID ntnx ignore

Retrieves all edges for a project by project ID {projectId}.
*/
func (a *Client) ProjectGetEdgesV2(params *ProjectGetEdgesV2Params, authInfo runtime.ClientAuthInfoWriter) (*ProjectGetEdgesV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectGetEdgesV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectGetEdgesV2",
		Method:             "GET",
		PathPattern:        "/v1.0/projects/{projectId}/edges",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProjectGetEdgesV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectGetEdgesV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectGetEdgesV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
