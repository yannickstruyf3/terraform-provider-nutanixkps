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

// NewCloudCredsGetParams creates a new CloudCredsGetParams object
// with the default values initialized.
func NewCloudCredsGetParams() *CloudCredsGetParams {
	var ()
	return &CloudCredsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloudCredsGetParamsWithTimeout creates a new CloudCredsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloudCredsGetParamsWithTimeout(timeout time.Duration) *CloudCredsGetParams {
	var ()
	return &CloudCredsGetParams{

		timeout: timeout,
	}
}

// NewCloudCredsGetParamsWithContext creates a new CloudCredsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloudCredsGetParamsWithContext(ctx context.Context) *CloudCredsGetParams {
	var ()
	return &CloudCredsGetParams{

		Context: ctx,
	}
}

// NewCloudCredsGetParamsWithHTTPClient creates a new CloudCredsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloudCredsGetParamsWithHTTPClient(client *http.Client) *CloudCredsGetParams {
	var ()
	return &CloudCredsGetParams{
		HTTPClient: client,
	}
}

/*CloudCredsGetParams contains all the parameters to send to the API endpoint
for the cloud creds get operation typically these are written to a http.Request
*/
type CloudCredsGetParams struct {

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

// WithTimeout adds the timeout to the cloud creds get params
func (o *CloudCredsGetParams) WithTimeout(timeout time.Duration) *CloudCredsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud creds get params
func (o *CloudCredsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud creds get params
func (o *CloudCredsGetParams) WithContext(ctx context.Context) *CloudCredsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud creds get params
func (o *CloudCredsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud creds get params
func (o *CloudCredsGetParams) WithHTTPClient(client *http.Client) *CloudCredsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud creds get params
func (o *CloudCredsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the cloud creds get params
func (o *CloudCredsGetParams) WithAuthorization(authorization string) *CloudCredsGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the cloud creds get params
func (o *CloudCredsGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the cloud creds get params
func (o *CloudCredsGetParams) WithID(id string) *CloudCredsGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cloud creds get params
func (o *CloudCredsGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CloudCredsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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