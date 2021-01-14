// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUserDeleteV2Params creates a new UserDeleteV2Params object
// with the default values initialized.
func NewUserDeleteV2Params() *UserDeleteV2Params {
	var ()
	return &UserDeleteV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserDeleteV2ParamsWithTimeout creates a new UserDeleteV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserDeleteV2ParamsWithTimeout(timeout time.Duration) *UserDeleteV2Params {
	var ()
	return &UserDeleteV2Params{

		timeout: timeout,
	}
}

// NewUserDeleteV2ParamsWithContext creates a new UserDeleteV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewUserDeleteV2ParamsWithContext(ctx context.Context) *UserDeleteV2Params {
	var ()
	return &UserDeleteV2Params{

		Context: ctx,
	}
}

// NewUserDeleteV2ParamsWithHTTPClient creates a new UserDeleteV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserDeleteV2ParamsWithHTTPClient(client *http.Client) *UserDeleteV2Params {
	var ()
	return &UserDeleteV2Params{
		HTTPClient: client,
	}
}

/*UserDeleteV2Params contains all the parameters to send to the API endpoint
for the user delete v2 operation typically these are written to a http.Request
*/
type UserDeleteV2Params struct {

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

// WithTimeout adds the timeout to the user delete v2 params
func (o *UserDeleteV2Params) WithTimeout(timeout time.Duration) *UserDeleteV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user delete v2 params
func (o *UserDeleteV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user delete v2 params
func (o *UserDeleteV2Params) WithContext(ctx context.Context) *UserDeleteV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user delete v2 params
func (o *UserDeleteV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user delete v2 params
func (o *UserDeleteV2Params) WithHTTPClient(client *http.Client) *UserDeleteV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user delete v2 params
func (o *UserDeleteV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the user delete v2 params
func (o *UserDeleteV2Params) WithAuthorization(authorization string) *UserDeleteV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the user delete v2 params
func (o *UserDeleteV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the user delete v2 params
func (o *UserDeleteV2Params) WithID(id string) *UserDeleteV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the user delete v2 params
func (o *UserDeleteV2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserDeleteV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
