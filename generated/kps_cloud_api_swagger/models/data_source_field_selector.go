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

// DataSourceFieldSelector DataSourceFieldSelector - data source field selector struct
//
// A DataSourceFieldSelector specifies a chosen category value
// and a specified scope (that is, which fields) where this value applies.
// The user annotates each data source field with one or more
// CategoryInfo objects.
// Categories enables the user to specify the data pipeline input.
// The list of categories specified is checked against each data source
// to determine if a field in the DataSource is
// included in the input of the data pipeline.
// swagger:model DataSourceFieldSelector
type DataSourceFieldSelector struct {

	// The category ID.
	// Required: true
	ID *string `json:"id"`

	// Field name(s) applicable to this CategoryInfo.
	// The special value '\_\_ALL\_\_' indicates that CategoryInfo is applicable to
	// all fields in this data source.
	// Required: true
	Scope []string `json:"scope"`

	// An allowed value to choose for the category.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this data source field selector
func (m *DataSourceFieldSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
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

func (m *DataSourceFieldSelector) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DataSourceFieldSelector) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *DataSourceFieldSelector) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSourceFieldSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSourceFieldSelector) UnmarshalBinary(b []byte) error {
	var res DataSourceFieldSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
