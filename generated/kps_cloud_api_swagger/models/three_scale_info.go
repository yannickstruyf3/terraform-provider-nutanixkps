// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ThreeScaleInfo ThreeScaleInfo shows if 3scale adapter is enabled in cluster and if user has permissions on adapter's configuration
// swagger:model ThreeScaleInfo
type ThreeScaleInfo struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// permissions
	Permissions *ResourcePermissions `json:"permissions,omitempty"`
}

// Validate validates this three scale info
func (m *ThreeScaleInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThreeScaleInfo) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThreeScaleInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThreeScaleInfo) UnmarshalBinary(b []byte) error {
	var res ThreeScaleInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}