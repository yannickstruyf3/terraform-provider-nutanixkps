// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DataPipelineContainersBaseObject DataPipelineContainersBaseObject - dataPipelineId and edgeID for which the
// containers will listed.
// swagger:model DataPipelineContainersBaseObject
type DataPipelineContainersBaseObject struct {

	// data pipeline ID
	DataPipelineID string `json:"dataPipelineId,omitempty"`

	// edge ID
	EdgeID string `json:"edgeId,omitempty"`
}

// Validate validates this data pipeline containers base object
func (m *DataPipelineContainersBaseObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataPipelineContainersBaseObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataPipelineContainersBaseObject) UnmarshalBinary(b []byte) error {
	var res DataPipelineContainersBaseObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}