// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ThreeScaleHandler ThreeScaleHAndler represents the minimal info that a user needs to know from the UI to link a service with 3Scale site
// swagger:model ThreeScaleHandler
type ThreeScaleHandler struct {

	// access token
	AccessToken string `json:"accessToken,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// service Id
	ServiceID string `json:"serviceId,omitempty"`

	// system Url
	SystemURL string `json:"systemUrl,omitempty"`
}

// Validate validates this three scale handler
func (m *ThreeScaleHandler) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThreeScaleHandler) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThreeScaleHandler) UnmarshalBinary(b []byte) error {
	var res ThreeScaleHandler
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}