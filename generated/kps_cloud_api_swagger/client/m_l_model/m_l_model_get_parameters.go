// Code generated by go-swagger; DO NOT EDIT.

package m_l_model

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

// NewMLModelGetParams creates a new MLModelGetParams object
// with the default values initialized.
func NewMLModelGetParams() *MLModelGetParams {
	var ()
	return &MLModelGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMLModelGetParamsWithTimeout creates a new MLModelGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMLModelGetParamsWithTimeout(timeout time.Duration) *MLModelGetParams {
	var ()
	return &MLModelGetParams{

		timeout: timeout,
	}
}

// NewMLModelGetParamsWithContext creates a new MLModelGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMLModelGetParamsWithContext(ctx context.Context) *MLModelGetParams {
	var ()
	return &MLModelGetParams{

		Context: ctx,
	}
}

// NewMLModelGetParamsWithHTTPClient creates a new MLModelGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMLModelGetParamsWithHTTPClient(client *http.Client) *MLModelGetParams {
	var ()
	return &MLModelGetParams{
		HTTPClient: client,
	}
}

/*MLModelGetParams contains all the parameters to send to the API endpoint
for the m l model get operation typically these are written to a http.Request
*/
type MLModelGetParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from the login API response.

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

// WithTimeout adds the timeout to the m l model get params
func (o *MLModelGetParams) WithTimeout(timeout time.Duration) *MLModelGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the m l model get params
func (o *MLModelGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the m l model get params
func (o *MLModelGetParams) WithContext(ctx context.Context) *MLModelGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the m l model get params
func (o *MLModelGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the m l model get params
func (o *MLModelGetParams) WithHTTPClient(client *http.Client) *MLModelGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the m l model get params
func (o *MLModelGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the m l model get params
func (o *MLModelGetParams) WithAuthorization(authorization string) *MLModelGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the m l model get params
func (o *MLModelGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the m l model get params
func (o *MLModelGetParams) WithID(id string) *MLModelGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the m l model get params
func (o *MLModelGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MLModelGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
