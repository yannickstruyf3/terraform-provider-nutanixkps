// Code generated by go-swagger; DO NOT EDIT.

package cloud_profile

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

// NewCloudProfileGetParams creates a new CloudProfileGetParams object
// with the default values initialized.
func NewCloudProfileGetParams() *CloudProfileGetParams {
	var ()
	return &CloudProfileGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloudProfileGetParamsWithTimeout creates a new CloudProfileGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloudProfileGetParamsWithTimeout(timeout time.Duration) *CloudProfileGetParams {
	var ()
	return &CloudProfileGetParams{

		timeout: timeout,
	}
}

// NewCloudProfileGetParamsWithContext creates a new CloudProfileGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloudProfileGetParamsWithContext(ctx context.Context) *CloudProfileGetParams {
	var ()
	return &CloudProfileGetParams{

		Context: ctx,
	}
}

// NewCloudProfileGetParamsWithHTTPClient creates a new CloudProfileGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloudProfileGetParamsWithHTTPClient(client *http.Client) *CloudProfileGetParams {
	var ()
	return &CloudProfileGetParams{
		HTTPClient: client,
	}
}

/*CloudProfileGetParams contains all the parameters to send to the API endpoint
for the cloud profile get operation typically these are written to a http.Request
*/
type CloudProfileGetParams struct {

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

// WithTimeout adds the timeout to the cloud profile get params
func (o *CloudProfileGetParams) WithTimeout(timeout time.Duration) *CloudProfileGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud profile get params
func (o *CloudProfileGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud profile get params
func (o *CloudProfileGetParams) WithContext(ctx context.Context) *CloudProfileGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud profile get params
func (o *CloudProfileGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud profile get params
func (o *CloudProfileGetParams) WithHTTPClient(client *http.Client) *CloudProfileGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud profile get params
func (o *CloudProfileGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the cloud profile get params
func (o *CloudProfileGetParams) WithAuthorization(authorization string) *CloudProfileGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the cloud profile get params
func (o *CloudProfileGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the cloud profile get params
func (o *CloudProfileGetParams) WithID(id string) *CloudProfileGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cloud profile get params
func (o *CloudProfileGetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CloudProfileGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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