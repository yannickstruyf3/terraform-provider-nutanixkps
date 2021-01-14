// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Iter8TrafficControl iter8 traffic control
// swagger:model Iter8TrafficControl
type Iter8TrafficControl struct {

	// algorithm
	Algorithm string `json:"algorithm,omitempty"`

	// interval
	Interval string `json:"interval,omitempty"`

	// max iterations
	MaxIterations int64 `json:"maxIterations,omitempty"`

	// max traffic percentage
	MaxTrafficPercentage float64 `json:"maxTrafficPercentage,omitempty"`

	// traffic step size
	TrafficStepSize float64 `json:"trafficStepSize,omitempty"`
}

// Validate validates this iter8 traffic control
func (m *Iter8TrafficControl) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Iter8TrafficControl) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Iter8TrafficControl) UnmarshalBinary(b []byte) error {
	var res Iter8TrafficControl
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}