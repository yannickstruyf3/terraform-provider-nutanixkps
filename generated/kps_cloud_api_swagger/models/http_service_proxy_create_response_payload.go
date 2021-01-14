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

// HTTPServiceProxyCreateResponsePayload HTTP service proxy create response payload
// swagger:model HTTPServiceProxyCreateResponsePayload
type HTTPServiceProxyCreateResponsePayload struct {

	// DNS URL of the service proxy endpoint
	// Valid only if setupDNS is set to true when creating the service proxy
	// Required: true
	DNSURL *string `json:"dnsURL"`

	// Expires at timestamp
	// Required: true
	// Format: date-time
	ExpiresAt *strfmt.DateTime `json:"expiresAt"`

	// ID of the entity
	// Required: true
	ID *string `json:"id"`

	// Password to login to the service when setupBasicAuth=true.
	Password string `json:"password,omitempty"`

	// URL of the service proxy endpoint
	// Required: true
	URL *string `json:"url"`

	// Username to login to the service when setupBasicAuth=true.
	Username string `json:"username,omitempty"`
}

// Validate validates this HTTP service proxy create response payload
func (m *HTTPServiceProxyCreateResponsePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HTTPServiceProxyCreateResponsePayload) validateDNSURL(formats strfmt.Registry) error {

	if err := validate.Required("dnsURL", "body", m.DNSURL); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxyCreateResponsePayload) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expiresAt", "body", m.ExpiresAt); err != nil {
		return err
	}

	if err := validate.FormatOf("expiresAt", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxyCreateResponsePayload) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxyCreateResponsePayload) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPServiceProxyCreateResponsePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPServiceProxyCreateResponsePayload) UnmarshalBinary(b []byte) error {
	var res HTTPServiceProxyCreateResponsePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
