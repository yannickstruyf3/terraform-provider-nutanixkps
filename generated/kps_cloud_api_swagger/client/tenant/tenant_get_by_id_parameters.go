// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTenantGetByIDParams creates a new TenantGetByIDParams object
// with the default values initialized.
func NewTenantGetByIDParams() *TenantGetByIDParams {
	var ()
	return &TenantGetByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTenantGetByIDParamsWithTimeout creates a new TenantGetByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTenantGetByIDParamsWithTimeout(timeout time.Duration) *TenantGetByIDParams {
	var ()
	return &TenantGetByIDParams{

		timeout: timeout,
	}
}

// NewTenantGetByIDParamsWithContext creates a new TenantGetByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewTenantGetByIDParamsWithContext(ctx context.Context) *TenantGetByIDParams {
	var ()
	return &TenantGetByIDParams{

		Context: ctx,
	}
}

// NewTenantGetByIDParamsWithHTTPClient creates a new TenantGetByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTenantGetByIDParamsWithHTTPClient(client *http.Client) *TenantGetByIDParams {
	var ()
	return &TenantGetByIDParams{
		HTTPClient: client,
	}
}

/*TenantGetByIDParams contains all the parameters to send to the API endpoint
for the tenant get by ID operation typically these are written to a http.Request
*/
type TenantGetByIDParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tenant get by ID params
func (o *TenantGetByIDParams) WithTimeout(timeout time.Duration) *TenantGetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenant get by ID params
func (o *TenantGetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenant get by ID params
func (o *TenantGetByIDParams) WithContext(ctx context.Context) *TenantGetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenant get by ID params
func (o *TenantGetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenant get by ID params
func (o *TenantGetByIDParams) WithHTTPClient(client *http.Client) *TenantGetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenant get by ID params
func (o *TenantGetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the tenant get by ID params
func (o *TenantGetByIDParams) WithAuthorization(authorization string) *TenantGetByIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the tenant get by ID params
func (o *TenantGetByIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the tenant get by ID params
func (o *TenantGetByIDParams) WithID(id string) *TenantGetByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenant get by ID params
func (o *TenantGetByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenantGetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}