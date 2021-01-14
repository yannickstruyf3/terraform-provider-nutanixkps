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

// DataIfcEndpoint DataIfcEndpoint is the endpoint within a given data Ifc
// swagger:model DataIfcEndpoint
type DataIfcEndpoint struct {

	// ID is the UUID of the data Src/Ifc
	// Required: true
	ID *string `json:"id"`

	// Name is the name of the field within the data Src/Ifc.
	// This defines invariant between data ifc and application.
	// If an application is associated w/ a given data ifc field name, that field
	// name cannot be removed from the data ifc until application is no longer the consumer of the
	// field.
	// Required: true
	Name *string `json:"name"`

	// Value is the name of the endpoint within the data Src/Ifc. This could be a topic name or bucket name, etc.
	// Artifacts of a data Ifc be affected by this.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this data ifc endpoint
func (m *DataIfcEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataIfcEndpoint) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DataIfcEndpoint) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DataIfcEndpoint) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataIfcEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataIfcEndpoint) UnmarshalBinary(b []byte) error {
	var res DataIfcEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
