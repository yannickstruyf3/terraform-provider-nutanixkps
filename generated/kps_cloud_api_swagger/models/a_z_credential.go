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

// AZCredential AZCredential - Azure credentials.
// swagger:model AZCredential
type AZCredential struct {

	// Azure storage account name and access key.
	// When you create a storage account, Azure generates 2 access keys. Provide the primary access key here.
	// Required: true
	StorageAccountName *string `json:"storageAccountName"`

	// storage key
	// Required: true
	StorageKey *string `json:"storageKey"`
}

// Validate validates this a z credential
func (m *AZCredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorageAccountName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AZCredential) validateStorageAccountName(formats strfmt.Registry) error {

	if err := validate.Required("storageAccountName", "body", m.StorageAccountName); err != nil {
		return err
	}

	return nil
}

func (m *AZCredential) validateStorageKey(formats strfmt.Registry) error {

	if err := validate.Required("storageKey", "body", m.StorageKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AZCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AZCredential) UnmarshalBinary(b []byte) error {
	var res AZCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}