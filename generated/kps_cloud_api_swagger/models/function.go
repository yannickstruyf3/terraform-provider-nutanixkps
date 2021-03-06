// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Function Function is object model for function
//
// Function represent lambdas:
// functions or transformations that can be applied
// to Data Pipelines.
// Functions are tenant-wide objects and the same function
// may be run within an edge, across all edges of a tenant
// or on tenant data in the cloud.
// swagger:model Function
type Function struct {

	// ntnx:ignore
	//
	// Whether this is a built-in runtime
	//
	// This should be required, but is not marked as such due to backward compatibility.
	// Required: true
	Builtin *bool `json:"builtin"`

	// The source code for the function script.
	// Required: true
	Code *string `json:"code"`

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Provide a description for your function code/script.
	Description string `json:"description,omitempty"`

	// Runtime environment for the function code/script.
	// Required: true
	Environment *string `json:"environment"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// Programming language for the function code/script.
	// Supported languages are python and javascript
	// Required: true
	Language *string `json:"language"`

	// Function name.
	// Required: true
	Name *string `json:"name"`

	// Array of script parameters.
	// Required: true
	Params []*ScriptParam `json:"params"`

	// ID of parent project, required for custom (non-builtin) scripts.
	ProjectID string `json:"projectId,omitempty"`

	// ID of the ScriptRuntime to use to run this script
	RuntimeID string `json:"runtimeId,omitempty"`

	// Docker image tag of the ScriptRuntime to use to run this script.
	// If missing or empty, then backend should treat it as "latest"
	RuntimeTag string `json:"runtimeTag,omitempty"`

	// ntnx:ignore
	// Required: true
	TenantID *string `json:"tenantId"`

	// Type of function code/script: Transformation or Function.
	// Transformation takes a data stream as input
	// and produces a different data stream as output.
	// Function takes a data stream as input
	// but has no constraint on output.
	// Required: true
	// Enum: [Transformation Function]
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
}

// Validate validates this function
func (m *Function) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuiltin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParams(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Function) validateBuiltin(formats strfmt.Registry) error {

	if err := validate.Required("builtin", "body", m.Builtin); err != nil {
		return err
	}

	return nil
}

func (m *Function) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *Function) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Function) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *Function) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *Function) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Function) validateParams(formats strfmt.Registry) error {

	if err := validate.Required("params", "body", m.Params); err != nil {
		return err
	}

	for i := 0; i < len(m.Params); i++ {
		if swag.IsZero(m.Params[i]) { // not required
			continue
		}

		if m.Params[i] != nil {
			if err := m.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Function) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

var functionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Transformation","Function"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		functionTypeTypePropEnum = append(functionTypeTypePropEnum, v)
	}
}

const (

	// FunctionTypeTransformation captures enum value "Transformation"
	FunctionTypeTransformation string = "Transformation"

	// FunctionTypeFunction captures enum value "Function"
	FunctionTypeFunction string = "Function"
)

// prop value enum
func (m *Function) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, functionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Function) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Function) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Function) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Function) UnmarshalBinary(b []byte) error {
	var res Function
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
