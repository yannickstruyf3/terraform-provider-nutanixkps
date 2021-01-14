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

// NewCloudProfileDeleteParams creates a new CloudProfileDeleteParams object
// with the default values initialized.
func NewCloudProfileDeleteParams() *CloudProfileDeleteParams {
	var ()
	return &CloudProfileDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloudProfileDeleteParamsWithTimeout creates a new CloudProfileDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloudProfileDeleteParamsWithTimeout(timeout time.Duration) *CloudProfileDeleteParams {
	var ()
	return &CloudProfileDeleteParams{

		timeout: timeout,
	}
}

// NewCloudProfileDeleteParamsWithContext creates a new CloudProfileDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloudProfileDeleteParamsWithContext(ctx context.Context) *CloudProfileDeleteParams {
	var ()
	return &CloudProfileDeleteParams{

		Context: ctx,
	}
}

// NewCloudProfileDeleteParamsWithHTTPClient creates a new CloudProfileDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloudProfileDeleteParamsWithHTTPClient(client *http.Client) *CloudProfileDeleteParams {
	var ()
	return &CloudProfileDeleteParams{
		HTTPClient: client,
	}
}

/*CloudProfileDeleteParams contains all the parameters to send to the API endpoint
for the cloud profile delete operation typically these are written to a http.Request
*/
type CloudProfileDeleteParams struct {

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

// WithTimeout adds the timeout to the cloud profile delete params
func (o *CloudProfileDeleteParams) WithTimeout(timeout time.Duration) *CloudProfileDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud profile delete params
func (o *CloudProfileDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud profile delete params
func (o *CloudProfileDeleteParams) WithContext(ctx context.Context) *CloudProfileDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud profile delete params
func (o *CloudProfileDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud profile delete params
func (o *CloudProfileDeleteParams) WithHTTPClient(client *http.Client) *CloudProfileDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud profile delete params
func (o *CloudProfileDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the cloud profile delete params
func (o *CloudProfileDeleteParams) WithAuthorization(authorization string) *CloudProfileDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the cloud profile delete params
func (o *CloudProfileDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the cloud profile delete params
func (o *CloudProfileDeleteParams) WithID(id string) *CloudProfileDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cloud profile delete params
func (o *CloudProfileDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CloudProfileDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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