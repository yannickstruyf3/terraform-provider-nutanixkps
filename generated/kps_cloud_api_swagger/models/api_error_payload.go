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

// APIErrorPayload The error message
// swagger:model APIErrorPayload
type APIErrorPayload struct {

	// Karbon Platform Services API error code
	// Required: true
	ErrorCode *int64 `json:"errorCode"`

	// Error message
	// Required: true
	Message *string `json:"message"`

	// HTTP status code for the response
	// Required: true
	StatusCode *int64 `json:"statusCode"`
}

// Validate validates this API error payload
func (m *APIErrorPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIErrorPayload) validateErrorCode(formats strfmt.Registry) error {

	if err := validate.Required("errorCode", "body", m.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (m *APIErrorPayload) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *APIErrorPayload) validateStatusCode(formats strfmt.Registry) error {

	if err := validate.Required("statusCode", "body", m.StatusCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIErrorPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIErrorPayload) UnmarshalBinary(b []byte) error {
	var res APIErrorPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
