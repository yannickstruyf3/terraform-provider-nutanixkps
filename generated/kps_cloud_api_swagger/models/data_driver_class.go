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

// DataDriverClass DataDriverClass is object model for data driver class
//
// A data driver class represents a logical IoT Data Source/Sink integration.
// swagger:model DataDriverClass
type DataDriverClass struct {

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Externa lversion of a data driver.
	// It is possible to have multiple data drivers with the same name, but different versions.
	// Required: true
	DataDriverVersion *string `json:"driverVersion"`

	// description
	Description string `json:"description,omitempty"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// min svc domain version
	MinSvcDomainVersion string `json:"minSvcDomainVersion,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// ntnx:ignore
	// Required: true
	TenantID *string `json:"tenantId"`

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// ntnx:ignore
	// Version of entity, implemented using timestamp in nano seconds
	// This is set to float64 since JSON numbers are floating point
	// May lose precision due to truncation but should have milli-second precision
	Version float64 `json:"version,omitempty"`

	// The YAML content for the application.
	// Required: true
	YamlData *string `json:"yamlData"`

	// dynamic parameter schema
	DynamicParameterSchema DataDriverParametersSchema `json:"dynamicParameterSchema,omitempty"`

	// static parameter schema
	StaticParameterSchema DataDriverParametersSchema `json:"staticParameterSchema,omitempty"`

	// type
	// Required: true
	Type DataDriverClassType `json:"type"`
}

// Validate validates this data driver class
func (m *DataDriverClass) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataDriverVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYamlData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDynamicParameterSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticParameterSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataDriverClass) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverClass) validateDataDriverVersion(formats strfmt.Registry) error {

	if err := validate.Required("driverVersion", "body", m.DataDriverVersion); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverClass) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverClass) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverClass) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverClass) validateYamlData(formats strfmt.Registry) error {

	if err := validate.Required("yamlData", "body", m.YamlData); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverClass) validateDynamicParameterSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.DynamicParameterSchema) { // not required
		return nil
	}

	if err := m.DynamicParameterSchema.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dynamicParameterSchema")
		}
		return err
	}

	return nil
}

func (m *DataDriverClass) validateStaticParameterSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.StaticParameterSchema) { // not required
		return nil
	}

	if err := m.StaticParameterSchema.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("staticParameterSchema")
		}
		return err
	}

	return nil
}

func (m *DataDriverClass) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataDriverClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataDriverClass) UnmarshalBinary(b []byte) error {
	var res DataDriverClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
