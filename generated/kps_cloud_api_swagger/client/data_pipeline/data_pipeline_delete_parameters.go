// Code generated by go-swagger; DO NOT EDIT.

package data_pipeline

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

// NewDataPipelineDeleteParams creates a new DataPipelineDeleteParams object
// with the default values initialized.
func NewDataPipelineDeleteParams() *DataPipelineDeleteParams {
	var ()
	return &DataPipelineDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataPipelineDeleteParamsWithTimeout creates a new DataPipelineDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataPipelineDeleteParamsWithTimeout(timeout time.Duration) *DataPipelineDeleteParams {
	var ()
	return &DataPipelineDeleteParams{

		timeout: timeout,
	}
}

// NewDataPipelineDeleteParamsWithContext creates a new DataPipelineDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataPipelineDeleteParamsWithContext(ctx context.Context) *DataPipelineDeleteParams {
	var ()
	return &DataPipelineDeleteParams{

		Context: ctx,
	}
}

// NewDataPipelineDeleteParamsWithHTTPClient creates a new DataPipelineDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataPipelineDeleteParamsWithHTTPClient(client *http.Client) *DataPipelineDeleteParams {
	var ()
	return &DataPipelineDeleteParams{
		HTTPClient: client,
	}
}

/*DataPipelineDeleteParams contains all the parameters to send to the API endpoint
for the data pipeline delete operation typically these are written to a http.Request
*/
type DataPipelineDeleteParams struct {

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

// WithTimeout adds the timeout to the data pipeline delete params
func (o *DataPipelineDeleteParams) WithTimeout(timeout time.Duration) *DataPipelineDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data pipeline delete params
func (o *DataPipelineDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data pipeline delete params
func (o *DataPipelineDeleteParams) WithContext(ctx context.Context) *DataPipelineDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data pipeline delete params
func (o *DataPipelineDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data pipeline delete params
func (o *DataPipelineDeleteParams) WithHTTPClient(client *http.Client) *DataPipelineDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data pipeline delete params
func (o *DataPipelineDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data pipeline delete params
func (o *DataPipelineDeleteParams) WithAuthorization(authorization string) *DataPipelineDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data pipeline delete params
func (o *DataPipelineDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the data pipeline delete params
func (o *DataPipelineDeleteParams) WithID(id string) *DataPipelineDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the data pipeline delete params
func (o *DataPipelineDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DataPipelineDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
