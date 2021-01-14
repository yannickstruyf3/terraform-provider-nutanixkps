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

// NewUserGetV2Params creates a new UserGetV2Params object
// with the default values initialized.
func NewUserGetV2Params() *UserGetV2Params {
	var ()
	return &UserGetV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserGetV2ParamsWithTimeout creates a new UserGetV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserGetV2ParamsWithTimeout(timeout time.Duration) *UserGetV2Params {
	var ()
	return &UserGetV2Params{

		timeout: timeout,
	}
}

// NewUserGetV2ParamsWithContext creates a new UserGetV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewUserGetV2ParamsWithContext(ctx context.Context) *UserGetV2Params {
	var ()
	return &UserGetV2Params{

		Context: ctx,
	}
}

// NewUserGetV2ParamsWithHTTPClient creates a new UserGetV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserGetV2ParamsWithHTTPClient(client *http.Client) *UserGetV2Params {
	var ()
	return &UserGetV2Params{
		HTTPClient: client,
	}
}

/*UserGetV2Params contains all the parameters to send to the API endpoint
for the user get v2 operation typically these are written to a http.Request
*/
type UserGetV2Params struct {

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

// WithTimeout adds the timeout to the user get v2 params
func (o *UserGetV2Params) WithTimeout(timeout time.Duration) *UserGetV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user get v2 params
func (o *UserGetV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user get v2 params
func (o *UserGetV2Params) WithContext(ctx context.Context) *UserGetV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user get v2 params
func (o *UserGetV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user get v2 params
func (o *UserGetV2Params) WithHTTPClient(client *http.Client) *UserGetV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user get v2 params
func (o *UserGetV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the user get v2 params
func (o *UserGetV2Params) WithAuthorization(authorization string) *UserGetV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the user get v2 params
func (o *UserGetV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the user get v2 params
func (o *UserGetV2Params) WithID(id string) *UserGetV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the user get v2 params
func (o *UserGetV2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserGetV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
