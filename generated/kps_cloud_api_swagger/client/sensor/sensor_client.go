// Code generated by go-swagger; DO NOT EDIT.

package sensor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sensor API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sensor API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
EdgeGetSensorsV2 gets edge sensors by edge ID ntnx ignore

Retrieves all sensors for an edge by edge ID {edgeId}.
*/
func (a *Client) EdgeGetSensorsV2(params *EdgeGetSensorsV2Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeGetSensorsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeGetSensorsV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeGetSensorsV2",
		Method:             "GET",
		PathPattern:        "/v1.0/edges/{edgeId}/sensors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeGetSensorsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeGetSensorsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeGetSensorsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SensorCreateV2 creates sensor ntnx ignore

Creates a sensor.
*/
func (a *Client) SensorCreateV2(params *SensorCreateV2Params, authInfo runtime.ClientAuthInfoWriter) (*SensorCreateV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSensorCreateV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SensorCreateV2",
		Method:             "POST",
		PathPattern:        "/v1.0/sensors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SensorCreateV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SensorCreateV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SensorCreateV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SensorDeleteV2 deletes sensor ntnx ignore

Deletes the sensor with the given ID {id}.
*/
func (a *Client) SensorDeleteV2(params *SensorDeleteV2Params, authInfo runtime.ClientAuthInfoWriter) (*SensorDeleteV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSensorDeleteV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SensorDeleteV2",
		Method:             "DELETE",
		PathPattern:        "/v1.0/sensors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SensorDeleteV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SensorDeleteV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SensorDeleteV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SensorGetV2 gets sensor by ID ntnx ignore

Retrieves the sensor with the given ID {id}.
*/
func (a *Client) SensorGetV2(params *SensorGetV2Params, authInfo runtime.ClientAuthInfoWriter) (*SensorGetV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSensorGetV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SensorGetV2",
		Method:             "GET",
		PathPattern:        "/v1.0/sensors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SensorGetV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SensorGetV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SensorGetV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SensorListV2 gets sensors ntnx ignore

Retrieves all sensors for a tenant.
*/
func (a *Client) SensorListV2(params *SensorListV2Params, authInfo runtime.ClientAuthInfoWriter) (*SensorListV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSensorListV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SensorListV2",
		Method:             "GET",
		PathPattern:        "/v1.0/sensors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SensorListV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SensorListV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SensorListV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SensorUpdateV3 updates a sensor by ID ntnx ignore

Updates a sensor by ID {id}.
*/
func (a *Client) SensorUpdateV3(params *SensorUpdateV3Params, authInfo runtime.ClientAuthInfoWriter) (*SensorUpdateV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSensorUpdateV3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SensorUpdateV3",
		Method:             "PUT",
		PathPattern:        "/v1.0/sensors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SensorUpdateV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SensorUpdateV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SensorUpdateV3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}