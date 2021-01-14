// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserAPITokenCreatePayload user Api token create payload
// swagger:model UserApiTokenCreatePayload
type UserAPITokenCreatePayload struct {

	// Whether the token is active
	// Required: true
	Active *bool `json:"active"`
}

// Validate validates this user Api token create payload
func (m *UserAPITokenCreatePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAPITokenCreatePayload) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserAPITokenCreatePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAPITokenCreatePayload) UnmarshalBinary(b []byte) error {
	var res UserAPITokenCreatePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}