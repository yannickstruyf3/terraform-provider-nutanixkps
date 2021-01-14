// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// JaegerInfo jaeger info
// swagger:model JaegerInfo
type JaegerInfo struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// integration
	Integration bool `json:"integration,omitempty"`

	// namespace selector
	NamespaceSelector bool `json:"namespaceSelector,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// white list istio system
	WhiteListIstioSystem []string `json:"whiteListIstioSystem"`
}

// Validate validates this jaeger info
func (m *JaegerInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JaegerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JaegerInfo) UnmarshalBinary(b []byte) error {
	var res JaegerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}