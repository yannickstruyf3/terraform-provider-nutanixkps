// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Trace Trace is a list of spans
// swagger:model Trace
type Trace struct {

	// processes
	Processes map[string]Process `json:"processes,omitempty"`

	// spans
	Spans []*Span `json:"spans"`

	// warnings
	Warnings []string `json:"warnings"`

	// trace ID
	TraceID TraceID `json:"traceID,omitempty"`
}

// Validate validates this trace
func (m *Trace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcesses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Trace) validateProcesses(formats strfmt.Registry) error {

	if swag.IsZero(m.Processes) { // not required
		return nil
	}

	for k := range m.Processes {

		if err := validate.Required("processes"+"."+k, "body", m.Processes[k]); err != nil {
			return err
		}
		if val, ok := m.Processes[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Trace) validateSpans(formats strfmt.Registry) error {

	if swag.IsZero(m.Spans) { // not required
		return nil
	}

	for i := 0; i < len(m.Spans); i++ {
		if swag.IsZero(m.Spans[i]) { // not required
			continue
		}

		if m.Spans[i] != nil {
			if err := m.Spans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Trace) validateTraceID(formats strfmt.Registry) error {

	if swag.IsZero(m.TraceID) { // not required
		return nil
	}

	if err := m.TraceID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("traceID")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Trace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Trace) UnmarshalBinary(b []byte) error {
	var res Trace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
