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

// NewUserAPITokenListParams creates a new UserAPITokenListParams object
// with the default values initialized.
func NewUserAPITokenListParams() *UserAPITokenListParams {
	var ()
	return &UserAPITokenListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserAPITokenListParamsWithTimeout creates a new UserAPITokenListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserAPITokenListParamsWithTimeout(timeout time.Duration) *UserAPITokenListParams {
	var ()
	return &UserAPITokenListParams{

		timeout: timeout,
	}
}

// NewUserAPITokenListParamsWithContext creates a new UserAPITokenListParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserAPITokenListParamsWithContext(ctx context.Context) *UserAPITokenListParams {
	var ()
	return &UserAPITokenListParams{

		Context: ctx,
	}
}

// NewUserAPITokenListParamsWithHTTPClient creates a new UserAPITokenListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserAPITokenListParamsWithHTTPClient(client *http.Client) *UserAPITokenListParams {
	var ()
	return &UserAPITokenListParams{
		HTTPClient: client,
	}
}

/*UserAPITokenListParams contains all the parameters to send to the API endpoint
for the user Api token list operation typically these are written to a http.Request
*/
type UserAPITokenListParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user Api token list params
func (o *UserAPITokenListParams) WithTimeout(timeout time.Duration) *UserAPITokenListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user Api token list params
func (o *UserAPITokenListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user Api token list params
func (o *UserAPITokenListParams) WithContext(ctx context.Context) *UserAPITokenListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user Api token list params
func (o *UserAPITokenListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user Api token list params
func (o *UserAPITokenListParams) WithHTTPClient(client *http.Client) *UserAPITokenListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user Api token list params
func (o *UserAPITokenListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the user Api token list params
func (o *UserAPITokenListParams) WithAuthorization(authorization string) *UserAPITokenListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the user Api token list params
func (o *UserAPITokenListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *UserAPITokenListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
