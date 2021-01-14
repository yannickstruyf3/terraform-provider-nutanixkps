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

// NodeOnboardInfo NodeOnboardInfo is object that relays post onboard info.
// swagger:model NodeOnboardInfo
type NodeOnboardInfo struct {

	// node ID
	// Required: true
	NodeID *string `json:"id"`

	// node version
	// Required: true
	NodeVersion *string `json:"nodeVersion"`

	// SSH public key
	// Required: true
	SSHPublicKey *string `json:"sshPublicKey"`
}

// Validate validates this node onboard info
func (m *NodeOnboardInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHPublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeOnboardInfo) validateNodeID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.NodeID); err != nil {
		return err
	}

	return nil
}

func (m *NodeOnboardInfo) validateNodeVersion(formats strfmt.Registry) error {

	if err := validate.Required("nodeVersion", "body", m.NodeVersion); err != nil {
		return err
	}

	return nil
}

func (m *NodeOnboardInfo) validateSSHPublicKey(formats strfmt.Registry) error {

	if err := validate.Required("sshPublicKey", "body", m.SSHPublicKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeOnboardInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeOnboardInfo) UnmarshalBinary(b []byte) error {
	var res NodeOnboardInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
