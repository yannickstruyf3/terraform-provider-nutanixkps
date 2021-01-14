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

// SoftwareUpdateState SoftwareUpdateState is the model for updating state called by service domain (cluster)
// swagger:model SoftwareUpdateState
type SoftwareUpdateState struct {

	// batch ID
	// Required: true
	BatchID *string `json:"batchId"`

	// Created timestamp
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// ETA in mins
	ETA int64 `json:"eta,omitempty"`

	// Failure reason if any
	FailureReason string `json:"failureReason,omitempty"`

	// Progress in percentage
	Progress int64 `json:"progress,omitempty"`

	// release
	// Required: true
	Release *string `json:"release"`

	// State updated timestamp
	// Format: date-time
	StateUpdatedAt strfmt.DateTime `json:"stateUpdatedAt,omitempty"`

	// svc domain ID
	// Required: true
	SvcDomainID *string `json:"svcDomainId"`

	// Record updated timestamp
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// state
	State SoftwareUpdateStateType `json:"state,omitempty"`
}

// Validate validates this software update state
func (m *SoftwareUpdateState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBatchID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvcDomainID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareUpdateState) validateBatchID(formats strfmt.Registry) error {

	if err := validate.Required("batchId", "body", m.BatchID); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareUpdateState) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareUpdateState) validateRelease(formats strfmt.Registry) error {

	if err := validate.Required("release", "body", m.Release); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareUpdateState) validateStateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StateUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("stateUpdatedAt", "body", "date-time", m.StateUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareUpdateState) validateSvcDomainID(formats strfmt.Registry) error {

	if err := validate.Required("svcDomainId", "body", m.SvcDomainID); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareUpdateState) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareUpdateState) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareUpdateState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareUpdateState) UnmarshalBinary(b []byte) error {
	var res SoftwareUpdateState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
