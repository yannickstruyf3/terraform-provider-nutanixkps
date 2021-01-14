// Code generated by go-swagger; DO NOT EDIT.

package edge_inventory_delta

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new edge inventory delta API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge inventory delta API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetEdgeInventoryDelta gets the edge inventory delta ntnx ignore

Retrieves the edge inventory delta: changes for any entity associated with the edge.
Entities are projects, applications, data pipelines, functions, and so on.
*/
func (a *Client) GetEdgeInventoryDelta(params *GetEdgeInventoryDeltaParams, authInfo runtime.ClientAuthInfoWriter) (*GetEdgeInventoryDeltaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEdgeInventoryDeltaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEdgeInventoryDelta",
		Method:             "POST",
		PathPattern:        "/v1.0/edgeinventorydelta",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEdgeInventoryDeltaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEdgeInventoryDeltaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetEdgeInventoryDeltaDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}