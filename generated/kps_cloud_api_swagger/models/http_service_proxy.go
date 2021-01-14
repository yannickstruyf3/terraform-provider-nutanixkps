// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HTTPServiceProxy HTTP service proxy
// swagger:model HTTPServiceProxy
type HTTPServiceProxy struct {

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// DNS URL of the service proxy endpoint
	// Valid only if setupDNS is set to true when creating the service proxy
	// Required: true
	DNSURL *string `json:"dnsURL"`

	// Duration of the http service proxy.
	// Required: true
	Duration *string `json:"duration"`

	// ntnx:ignore
	// Expires at timestamp - computed from createdAt, updatedAt time and duration
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expiresAt,omitempty"`

	// ntnx:ignore
	// Name of statefulset host this proxy is served.
	Hostname string `json:"hostname,omitempty"`

	// ntnx:ignore
	// TCP port on statefulset host this proxy is served.
	Hostport int64 `json:"hostport,omitempty"`

	// ID of the entity
	// Maximum character length is 64 for project, category, and runtime environment,
	// 36 for other entity types.
	ID string `json:"id,omitempty"`

	// HTTP service proxy name.
	// Unique within (tenant, service domain)
	// Required: true
	Name *string `json:"name"`

	// Password to login to the service when setupBasicAuth=true.
	Password string `json:"password,omitempty"`

	// ID of parent project, required when TYPE = PROJECT.
	ProjectID string `json:"projectId,omitempty"`

	// ntnx:ignore
	// Public Key. Used for session tracking. Required in Delete payload.
	PublicKey string `json:"publicKey,omitempty"`

	// Name of the http service.
	// Required: true
	ServiceName *string `json:"serviceName"`

	// Namespace of the http service, required when TYPE = SYSTEM
	ServiceNamespace string `json:"serviceNamespace,omitempty"`

	// Port of the http service.
	// Required: true
	ServicePort *int64 `json:"servicePort"`

	// ID of the service domain this entity belongs to
	// Required: true
	SvcDomainID *string `json:"svcDomainId"`

	// ntnx:ignore
	// Required: true
	TenantID *string `json:"tenantId"`

	// Service type for this http proxy.
	// Required: true
	// Enum: [SYSTEM PROJECT CUSTOM]
	Type *string `json:"type"`

	// URL of the service proxy endpoint
	// Required: true
	URL *string `json:"url"`

	// ntnx:ignore
	// timestamp feature supported by DB
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Username to login to the service when setupBasicAuth=true.
	Username string `json:"username,omitempty"`

	// ntnx:ignore
	// Version of entity, implemented using timestamp in nano seconds
	// This is set to float64 since JSON numbers are floating point
	// May lose precision due to truncation but should have milli-second precision
	Version float64 `json:"version,omitempty"`
}

// Validate validates this HTTP service proxy
func (m *HTTPServiceProxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvcDomainID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
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

func (m *HTTPServiceProxy) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxy) validateDNSURL(formats strfmt.Registry) error {

	if err := validate.Required("dnsURL", "body", m.DNSURL); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxy) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxy) validateExpiresAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expiresAt", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxy) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxy) validateServiceName(formats strfmt.Registry) error {

	if err := validate.Required("serviceName", "body", m.ServiceName); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxy) validateServicePort(formats strfmt.Registry) error {

	if err := validate.Required("servicePort", "body", m.ServicePort); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxy) validateSvcDomainID(formats strfmt.Registry) error {

	if err := validate.Required("svcDomainId", "body", m.SvcDomainID); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxy) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

var httpServiceProxyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SYSTEM","PROJECT","CUSTOM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpServiceProxyTypeTypePropEnum = append(httpServiceProxyTypeTypePropEnum, v)
	}
}

const (

	// HTTPServiceProxyTypeSYSTEM captures enum value "SYSTEM"
	HTTPServiceProxyTypeSYSTEM string = "SYSTEM"

	// HTTPServiceProxyTypePROJECT captures enum value "PROJECT"
	HTTPServiceProxyTypePROJECT string = "PROJECT"

	// HTTPServiceProxyTypeCUSTOM captures enum value "CUSTOM"
	HTTPServiceProxyTypeCUSTOM string = "CUSTOM"
)

// prop value enum
func (m *HTTPServiceProxy) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpServiceProxyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPServiceProxy) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxy) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *HTTPServiceProxy) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPServiceProxy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPServiceProxy) UnmarshalBinary(b []byte) error {
	var res HTTPServiceProxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}