// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ApplicationContainers ApplicationContainers encapsulates the container names
// for a specific application on a specific edge.
// swagger:model ApplicationContainers
type ApplicationContainers struct {

	// application ID
	ApplicationID string `json:"applicationId,omitempty"`

	// container names
	ContainerNames []string `json:"containerNames"`

	// edge ID
	EdgeID string `json:"edgeId,omitempty"`
}

// Validate validates this application containers
func (m *ApplicationContainers) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationContainers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationContainers) UnmarshalBinary(b []byte) error {
	var res ApplicationContainers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}