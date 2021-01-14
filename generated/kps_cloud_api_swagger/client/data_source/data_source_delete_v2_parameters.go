// Code generated by go-swagger; DO NOT EDIT.

package data_source

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

// NewDataSourceDeleteV2Params creates a new DataSourceDeleteV2Params object
// with the default values initialized.
func NewDataSourceDeleteV2Params() *DataSourceDeleteV2Params {
	var ()
	return &DataSourceDeleteV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataSourceDeleteV2ParamsWithTimeout creates a new DataSourceDeleteV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataSourceDeleteV2ParamsWithTimeout(timeout time.Duration) *DataSourceDeleteV2Params {
	var ()
	return &DataSourceDeleteV2Params{

		timeout: timeout,
	}
}

// NewDataSourceDeleteV2ParamsWithContext creates a new DataSourceDeleteV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDataSourceDeleteV2ParamsWithContext(ctx context.Context) *DataSourceDeleteV2Params {
	var ()
	return &DataSourceDeleteV2Params{

		Context: ctx,
	}
}

// NewDataSourceDeleteV2ParamsWithHTTPClient creates a new DataSourceDeleteV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataSourceDeleteV2ParamsWithHTTPClient(client *http.Client) *DataSourceDeleteV2Params {
	var ()
	return &DataSourceDeleteV2Params{
		HTTPClient: client,
	}
}

/*DataSourceDeleteV2Params contains all the parameters to send to the API endpoint
for the data source delete v2 operation typically these are written to a http.Request
*/
type DataSourceDeleteV2Params struct {

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

// WithTimeout adds the timeout to the data source delete v2 params
func (o *DataSourceDeleteV2Params) WithTimeout(timeout time.Duration) *DataSourceDeleteV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data source delete v2 params
func (o *DataSourceDeleteV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data source delete v2 params
func (o *DataSourceDeleteV2Params) WithContext(ctx context.Context) *DataSourceDeleteV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data source delete v2 params
func (o *DataSourceDeleteV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data source delete v2 params
func (o *DataSourceDeleteV2Params) WithHTTPClient(client *http.Client) *DataSourceDeleteV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data source delete v2 params
func (o *DataSourceDeleteV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data source delete v2 params
func (o *DataSourceDeleteV2Params) WithAuthorization(authorization string) *DataSourceDeleteV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data source delete v2 params
func (o *DataSourceDeleteV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the data source delete v2 params
func (o *DataSourceDeleteV2Params) WithID(id string) *DataSourceDeleteV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the data source delete v2 params
func (o *DataSourceDeleteV2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DataSourceDeleteV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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