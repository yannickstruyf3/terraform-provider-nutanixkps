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

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// NewDataPipelineUpdateParams creates a new DataPipelineUpdateParams object
// with the default values initialized.
func NewDataPipelineUpdateParams() *DataPipelineUpdateParams {
	var ()
	return &DataPipelineUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataPipelineUpdateParamsWithTimeout creates a new DataPipelineUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataPipelineUpdateParamsWithTimeout(timeout time.Duration) *DataPipelineUpdateParams {
	var ()
	return &DataPipelineUpdateParams{

		timeout: timeout,
	}
}

// NewDataPipelineUpdateParamsWithContext creates a new DataPipelineUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataPipelineUpdateParamsWithContext(ctx context.Context) *DataPipelineUpdateParams {
	var ()
	return &DataPipelineUpdateParams{

		Context: ctx,
	}
}

// NewDataPipelineUpdateParamsWithHTTPClient creates a new DataPipelineUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataPipelineUpdateParamsWithHTTPClient(client *http.Client) *DataPipelineUpdateParams {
	var ()
	return &DataPipelineUpdateParams{
		HTTPClient: client,
	}
}

/*DataPipelineUpdateParams contains all the parameters to send to the API endpoint
for the data pipeline update operation typically these are written to a http.Request
*/
type DataPipelineUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.DataPipeline
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data pipeline update params
func (o *DataPipelineUpdateParams) WithTimeout(timeout time.Duration) *DataPipelineUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data pipeline update params
func (o *DataPipelineUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data pipeline update params
func (o *DataPipelineUpdateParams) WithContext(ctx context.Context) *DataPipelineUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data pipeline update params
func (o *DataPipelineUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data pipeline update params
func (o *DataPipelineUpdateParams) WithHTTPClient(client *http.Client) *DataPipelineUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data pipeline update params
func (o *DataPipelineUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the data pipeline update params
func (o *DataPipelineUpdateParams) WithAuthorization(authorization string) *DataPipelineUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the data pipeline update params
func (o *DataPipelineUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the data pipeline update params
func (o *DataPipelineUpdateParams) WithBody(body *models.DataPipeline) *DataPipelineUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the data pipeline update params
func (o *DataPipelineUpdateParams) SetBody(body *models.DataPipeline) {
	o.Body = body
}

// WithID adds the id to the data pipeline update params
func (o *DataPipelineUpdateParams) WithID(id string) *DataPipelineUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the data pipeline update params
func (o *DataPipelineUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DataPipelineUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
