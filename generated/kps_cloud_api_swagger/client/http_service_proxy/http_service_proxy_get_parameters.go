// Code generated by go-swagger; DO NOT EDIT.

package http_service_proxy

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

// NewHTTPServiceProxyGetParams creates a new HTTPServiceProxyGetParams object
// with the default values initialized.
func NewHTTPServiceProxyGetParams() *HTTPServiceProxyGetParams {
	var ()
	return &HTTPServiceProxyGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHTTPServiceProxyGetParamsWithTimeout creates a new HTTPServiceProxyGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHTTPServiceProxyGetParamsWithTimeout(timeout time.Duration) *HTTPServiceProxyGetParams {
	var ()
	return &HTTPServiceProxyGetParams{

		timeout: timeout,
	}
}

// NewHTTPServiceProxyGetParamsWithContext creates a new HTTPServiceProxyGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewHTTPServiceProxyGetParamsWithContext(ctx context.Context) *HTTPServiceProxyGetParams {
	var ()
	return &HTTPServiceProxyGetParams{

		Context: ctx,
	}
}

// NewHTTPServiceProxyGetParamsWithHTTPClient creates a new HTTPServiceProxyGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHTTPServiceProxyGetParamsWithHTTPClient(client *http.Client) *HTTPServiceProxyGetParams {
	var ()
	return &HTTPServiceProxyGetParams{
		HTTPClient: client,
	}
}

/*HTTPServiceProxyGetParams contains all the parameters to send to the API endpoint
for the HTTP service proxy get operation typically these are written to a http.Request
*/
type HTTPServiceProxyGetParams struct {

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

// WithTimeout adds the timeout to the HTTP service proxy get params
func (o *HTTPServiceProxyGetParams) WithTimeout(timeout time.Duration) *HTTPServiceProxyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the HTTP service proxy get params
func (o *HTTPServiceProxyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the HTTP service proxy get params
func (o *HTTPServiceProxyGetParams) WithContext(ctx context.Context) *HTTPServiceProxyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the HTTP service proxy get params
func (o *HTTPServiceProxyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the HTTP service proxy get params
func (o *HTTPServiceProxyGetParams) WithHTTPClient(client *http.Client) *HTTPServiceProxyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the HTTP service proxy get params
func (o *HTTPServiceProxyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the HTTP service proxy get params
func (o *HTTPServiceProxyGetParams) WithAuthorization(authorization string) *HTTPServiceProxyGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the HTTP service proxy get params
func (o *HTTPServiceProxyGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the HTTP service proxy get params
func (o *HTTPServiceProxyGetParams) WithID(id string) *HTTPServiceProxyGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the HTTP service proxy get params
func (o *HTTPServiceProxyGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *HTTPServiceProxyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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