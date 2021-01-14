// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewShortLoginTokenV1Params creates a new ShortLoginTokenV1Params object
// with the default values initialized.
func NewShortLoginTokenV1Params() *ShortLoginTokenV1Params {
	var ()
	return &ShortLoginTokenV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewShortLoginTokenV1ParamsWithTimeout creates a new ShortLoginTokenV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewShortLoginTokenV1ParamsWithTimeout(timeout time.Duration) *ShortLoginTokenV1Params {
	var ()
	return &ShortLoginTokenV1Params{

		timeout: timeout,
	}
}

// NewShortLoginTokenV1ParamsWithContext creates a new ShortLoginTokenV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewShortLoginTokenV1ParamsWithContext(ctx context.Context) *ShortLoginTokenV1Params {
	var ()
	return &ShortLoginTokenV1Params{

		Context: ctx,
	}
}

// NewShortLoginTokenV1ParamsWithHTTPClient creates a new ShortLoginTokenV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShortLoginTokenV1ParamsWithHTTPClient(client *http.Client) *ShortLoginTokenV1Params {
	var ()
	return &ShortLoginTokenV1Params{
		HTTPClient: client,
	}
}

/*ShortLoginTokenV1Params contains all the parameters to send to the API endpoint
for the short login token v1 operation typically these are written to a http.Request
*/
type ShortLoginTokenV1Params struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the short login token v1 params
func (o *ShortLoginTokenV1Params) WithTimeout(timeout time.Duration) *ShortLoginTokenV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the short login token v1 params
func (o *ShortLoginTokenV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the short login token v1 params
func (o *ShortLoginTokenV1Params) WithContext(ctx context.Context) *ShortLoginTokenV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the short login token v1 params
func (o *ShortLoginTokenV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the short login token v1 params
func (o *ShortLoginTokenV1Params) WithHTTPClient(client *http.Client) *ShortLoginTokenV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the short login token v1 params
func (o *ShortLoginTokenV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the short login token v1 params
func (o *ShortLoginTokenV1Params) WithAuthorization(authorization string) *ShortLoginTokenV1Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the short login token v1 params
func (o *ShortLoginTokenV1Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *ShortLoginTokenV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}