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

// Node Node is the DB object and object model for nodes
//
// Node is a node in a service domain.
// swagger:model Node
type Node struct {

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Node description
	Description string `json:"description,omitempty"`

	// Node Gateway IP address
	// Required: true
	Gateway *string `json:"gateway"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// Node IP Address
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// is bootstrap master
	IsBootstrapMaster bool `json:"isBootstrapMaster,omitempty"`

	// Node name.
	// Maximum length edge name is determined by kubernetes.
	// Name length limited to 60 and contraints are defined here
	// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go
	// Required: true
	Name *string `json:"name"`

	// Node serial number
	// Required: true
	SerialNumber *string `json:"serialNumber"`

	// Node subnet mask
	// Required: true
	Subnet *string `json:"subnet"`

	// ID of the service domain this entity belongs to
	// Required: true
	SvcDomainID *string `json:"svcDomainId"`

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

	// role
	Role *NodeRole `json:"role,omitempty"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvcDomainID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateSerialNumber(formats strfmt.Registry) error {

	if err := validate.Required("serialNumber", "body", m.SerialNumber); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateSubnet(formats strfmt.Registry) error {

	if err := validate.Required("subnet", "body", m.Subnet); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateSvcDomainID(formats strfmt.Registry) error {

	if err := validate.Required("svcDomainId", "body", m.SvcDomainID); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
