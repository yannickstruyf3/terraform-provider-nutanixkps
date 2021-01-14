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

// FunctionListPayload payload for FunctionListResponse
// swagger:model FunctionListPayload
type FunctionListPayload struct {

	// list of functions
	// Required: true
	FunctionList []*Function `json:"result"`

	// Specify result order. Zero or more entries with format: &ltkey> [desc]
	// where orderByKeys lists allowed keys in each response.
	OrderBy string `json:"orderBy,omitempty"`

	// Keys that can be used in orderBy.
	OrderByKeys []string `json:"orderByKeys"`

	// 0-based index of the page to fetch results.
	// Required: true
	PageIndex *int64 `json:"pageIndex"`

	// Item count of each page.
	// Required: true
	PageSize *int64 `json:"pageSize"`

	// Count of all items matching the query.
	// Required: true
	TotalCount *int64 `json:"totalCount"`
}

// Validate validates this function list payload
func (m *FunctionListPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunctionList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FunctionListPayload) validateFunctionList(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.FunctionList); err != nil {
		return err
	}

	for i := 0; i < len(m.FunctionList); i++ {
		if swag.IsZero(m.FunctionList[i]) { // not required
			continue
		}

		if m.FunctionList[i] != nil {
			if err := m.FunctionList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FunctionListPayload) validatePageIndex(formats strfmt.Registry) error {

	if err := validate.Required("pageIndex", "body", m.PageIndex); err != nil {
		return err
	}

	return nil
}

func (m *FunctionListPayload) validatePageSize(formats strfmt.Registry) error {

	if err := validate.Required("pageSize", "body", m.PageSize); err != nil {
		return err
	}

	return nil
}

func (m *FunctionListPayload) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("totalCount", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FunctionListPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionListPayload) UnmarshalBinary(b []byte) error {
	var res FunctionListPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
