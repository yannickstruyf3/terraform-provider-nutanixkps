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

// KubernetesClusterHandlePayload KubernetesClusterHandlePayload payload for Kubernetes Cluster handle call
// swagger:model KubernetesClusterHandlePayload
type KubernetesClusterHandlePayload struct {

	// tenant ID
	// Required: true
	TenantID *string `json:"tenantId"`

	// token
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this kubernetes cluster handle payload
func (m *KubernetesClusterHandlePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesClusterHandlePayload) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesClusterHandlePayload) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesClusterHandlePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesClusterHandlePayload) UnmarshalBinary(b []byte) error {
	var res KubernetesClusterHandlePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}