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

// AppListItem AppListItem has the necessary information to display the console app list
// swagger:model AppListItem
type AppListItem struct {

	// Define if all Pods related to the Workloads of this app has an IstioSidecar deployed
	// Required: true
	IstioSidecar *bool `json:"istioSidecar"`

	// Labels for App
	Labels map[string]string `json:"labels,omitempty"`

	// Name of the application
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this app list item
func (m *AppListItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIstioSidecar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppListItem) validateIstioSidecar(formats strfmt.Registry) error {

	if err := validate.Required("istioSidecar", "body", m.IstioSidecar); err != nil {
		return err
	}

	return nil
}

func (m *AppListItem) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppListItem) UnmarshalBinary(b []byte) error {
	var res AppListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}