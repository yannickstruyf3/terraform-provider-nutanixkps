// Code generated by go-swagger; DO NOT EDIT.

package user_api_token

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

// NewUserAPITokenDeleteParams creates a new UserAPITokenDeleteParams object
// with the default values initialized.
func NewUserAPITokenDeleteParams() *UserAPITokenDeleteParams {
	var ()
	return &UserAPITokenDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserAPITokenDeleteParamsWithTimeout creates a new UserAPITokenDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserAPITokenDeleteParamsWithTimeout(timeout time.Duration) *UserAPITokenDeleteParams {
	var ()
	return &UserAPITokenDeleteParams{

		timeout: timeout,
	}
}

// NewUserAPITokenDeleteParamsWithContext creates a new UserAPITokenDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserAPITokenDeleteParamsWithContext(ctx context.Context) *UserAPITokenDeleteParams {
	var ()
	return &UserAPITokenDeleteParams{

		Context: ctx,
	}
}

// NewUserAPITokenDeleteParamsWithHTTPClient creates a new UserAPITokenDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserAPITokenDeleteParamsWithHTTPClient(client *http.Client) *UserAPITokenDeleteParams {
	var ()
	return &UserAPITokenDeleteParams{
		HTTPClient: client,
	}
}

/*UserAPITokenDeleteParams contains all the parameters to send to the API endpoint
for the user Api token delete operation typically these are written to a http.Request
*/
type UserAPITokenDeleteParams struct {

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

// WithTimeout adds the timeout to the user Api token delete params
func (o *UserAPITokenDeleteParams) WithTimeout(timeout time.Duration) *UserAPITokenDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user Api token delete params
func (o *UserAPITokenDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user Api token delete params
func (o *UserAPITokenDeleteParams) WithContext(ctx context.Context) *UserAPITokenDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user Api token delete params
func (o *UserAPITokenDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user Api token delete params
func (o *UserAPITokenDeleteParams) WithHTTPClient(client *http.Client) *UserAPITokenDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user Api token delete params
func (o *UserAPITokenDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the user Api token delete params
func (o *UserAPITokenDeleteParams) WithAuthorization(authorization string) *UserAPITokenDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the user Api token delete params
func (o *UserAPITokenDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the user Api token delete params
func (o *UserAPITokenDeleteParams) WithID(id string) *UserAPITokenDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user Api token delete params
func (o *UserAPITokenDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserAPITokenDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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