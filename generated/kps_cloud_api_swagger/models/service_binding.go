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

// ServiceBinding ServiceBinding holds the Service Binding information
// swagger:model ServiceBinding
type ServiceBinding struct {

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// Minimum version of the Service Domain supporting this Service Class
	// Required: true
	MinSvcDomainVersion *string `json:"minSvcDomainVersion"`

	// name
	// Required: true
	Name *string `json:"name"`

	// parameters
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// svc class ID
	SvcClassID string `json:"svcClassId,omitempty"`

	// svc class name
	SvcClassName string `json:"svcClassName,omitempty"`

	// Version of the Service Class type
	// Required: true
	SvcVersion *string `json:"svcVersion"`

	// ntnx:ignore
	// Required: true
	TenantID *string `json:"tenantId"`

	// Type of the Service Class e.g Kafka
	// Required: true
	Type *string `json:"type"`

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// ntnx:ignore
	// Version of entity, implemented using timestamp in nano seconds
	// This is set to float64 since JSON numbers are floating point
	// May lose precision due to truncation but should have milli-second precision
	Version float64 `json:"version,omitempty"`

	// bind resource
	BindResource *ServiceBindingResource `json:"bindResource,omitempty"`

	// scope
	// Required: true
	Scope ServiceClassScopeType `json:"scope"`
}

// Validate validates this service binding
func (m *ServiceBinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinSvcDomainVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvcVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBindResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceBinding) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceBinding) validateMinSvcDomainVersion(formats strfmt.Registry) error {

	if err := validate.Required("minSvcDomainVersion", "body", m.MinSvcDomainVersion); err != nil {
		return err
	}

	return nil
}

func (m *ServiceBinding) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ServiceBinding) validateSvcVersion(formats strfmt.Registry) error {

	if err := validate.Required("svcVersion", "body", m.SvcVersion); err != nil {
		return err
	}

	return nil
}

func (m *ServiceBinding) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceBinding) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ServiceBinding) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceBinding) validateBindResource(formats strfmt.Registry) error {

	if swag.IsZero(m.BindResource) { // not required
		return nil
	}

	if m.BindResource != nil {
		if err := m.BindResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bindResource")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceBinding) validateScope(formats strfmt.Registry) error {

	if err := m.Scope.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceBinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceBinding) UnmarshalBinary(b []byte) error {
	var res ServiceBinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}