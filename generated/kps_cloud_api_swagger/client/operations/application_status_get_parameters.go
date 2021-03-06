// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewApplicationStatusGetParams creates a new ApplicationStatusGetParams object
// with the default values initialized.
func NewApplicationStatusGetParams() *ApplicationStatusGetParams {
	var ()
	return &ApplicationStatusGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationStatusGetParamsWithTimeout creates a new ApplicationStatusGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewApplicationStatusGetParamsWithTimeout(timeout time.Duration) *ApplicationStatusGetParams {
	var ()
	return &ApplicationStatusGetParams{

		timeout: timeout,
	}
}

// NewApplicationStatusGetParamsWithContext creates a new ApplicationStatusGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewApplicationStatusGetParamsWithContext(ctx context.Context) *ApplicationStatusGetParams {
	var ()
	return &ApplicationStatusGetParams{

		Context: ctx,
	}
}

// NewApplicationStatusGetParamsWithHTTPClient creates a new ApplicationStatusGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewApplicationStatusGetParamsWithHTTPClient(client *http.Client) *ApplicationStatusGetParams {
	var ()
	return &ApplicationStatusGetParams{
		HTTPClient: client,
	}
}

/*ApplicationStatusGetParams contains all the parameters to send to the API endpoint
for the application status get operation typically these are written to a http.Request
*/
type ApplicationStatusGetParams struct {

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

// WithTimeout adds the timeout to the application status get params
func (o *ApplicationStatusGetParams) WithTimeout(timeout time.Duration) *ApplicationStatusGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application status get params
func (o *ApplicationStatusGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application status get params
func (o *ApplicationStatusGetParams) WithContext(ctx context.Context) *ApplicationStatusGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application status get params
func (o *ApplicationStatusGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application status get params
func (o *ApplicationStatusGetParams) WithHTTPClient(client *http.Client) *ApplicationStatusGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application status get params
func (o *ApplicationStatusGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the application status get params
func (o *ApplicationStatusGetParams) WithAuthorization(authorization string) *ApplicationStatusGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the application status get params
func (o *ApplicationStatusGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the application status get params
func (o *ApplicationStatusGetParams) WithID(id string) *ApplicationStatusGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the application status get params
func (o *ApplicationStatusGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationStatusGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
