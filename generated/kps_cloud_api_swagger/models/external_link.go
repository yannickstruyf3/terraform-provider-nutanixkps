// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExternalLink ExternalLink provides links to external dashboards (e.g. to Grafana)
// swagger:model ExternalLink
type ExternalLink struct {

	// name
	Name string `json:"name,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// variables
	Variables *MonitoringDashboardExternalLinkVariables `json:"variables,omitempty"`
}

// Validate validates this external link
func (m *ExternalLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalLink) validateVariables(formats strfmt.Registry) error {

	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	if m.Variables != nil {
		if err := m.Variables.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variables")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalLink) UnmarshalBinary(b []byte) error {
	var res ExternalLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
