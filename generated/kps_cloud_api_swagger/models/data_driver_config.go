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

// DataDriverConfig DataDriverConfig is object model for data driver instance's dynamic configuration
//
// A dynamic instance config represents a logical Data Source/Sink integration's dynamic configuration.
// swagger:model DataDriverConfig
type DataDriverConfig struct {

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// data driver instance ID
	// Required: true
	DataDriverInstanceID *string `json:"dataDriverInstanceID"`

	// description
	Description string `json:"description,omitempty"`

	// Edges listed according to ID where the data driver config is deployed.
	// Only relevant if the parent project EdgeSelectorType value is set to Explicit.
	EdgeIDs []string `json:"edgeIds"`

	// Select edges according to CategoryInfo.
	// Only relevant if the parent project EdgeSelectorType value is set to Category.
	EdgeSelectors []*CategoryInfo `json:"edgeSelectors"`

	// Edges to be excluded from the data driver config deployment.
	ExcludeEdgeIDs []string `json:"excludeEdgeIds"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// A list of Category labels for this data driver config.
	Labels []*CategoryInfo `json:"labels"`

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

	// parameters
	Parameters DataDriverParametersValues `json:"parameters,omitempty"`
}

// Validate validates this data driver config
func (m *DataDriverConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataDriverInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeSelectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
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

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataDriverConfig) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverConfig) validateDataDriverInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("dataDriverInstanceID", "body", m.DataDriverInstanceID); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverConfig) validateEdgeSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeSelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.EdgeSelectors); i++ {
		if swag.IsZero(m.EdgeSelectors[i]) { // not required
			continue
		}

		if m.EdgeSelectors[i] != nil {
			if err := m.EdgeSelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeSelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataDriverConfig) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataDriverConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverConfig) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverConfig) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DataDriverConfig) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if err := m.Parameters.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("parameters")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataDriverConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataDriverConfig) UnmarshalBinary(b []byte) error {
	var res DataDriverConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}