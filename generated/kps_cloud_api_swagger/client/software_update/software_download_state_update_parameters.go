// Code generated by go-swagger; DO NOT EDIT.

package software_update

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

// NewSoftwareDownloadStateUpdateParams creates a new SoftwareDownloadStateUpdateParams object
// with the default values initialized.
func NewSoftwareDownloadStateUpdateParams() *SoftwareDownloadStateUpdateParams {
	var ()
	return &SoftwareDownloadStateUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwareDownloadStateUpdateParamsWithTimeout creates a new SoftwareDownloadStateUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSoftwareDownloadStateUpdateParamsWithTimeout(timeout time.Duration) *SoftwareDownloadStateUpdateParams {
	var ()
	return &SoftwareDownloadStateUpdateParams{

		timeout: timeout,
	}
}

// NewSoftwareDownloadStateUpdateParamsWithContext creates a new SoftwareDownloadStateUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSoftwareDownloadStateUpdateParamsWithContext(ctx context.Context) *SoftwareDownloadStateUpdateParams {
	var ()
	return &SoftwareDownloadStateUpdateParams{

		Context: ctx,
	}
}

// NewSoftwareDownloadStateUpdateParamsWithHTTPClient creates a new SoftwareDownloadStateUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSoftwareDownloadStateUpdateParamsWithHTTPClient(client *http.Client) *SoftwareDownloadStateUpdateParams {
	var ()
	return &SoftwareDownloadStateUpdateParams{
		HTTPClient: client,
	}
}

/*SoftwareDownloadStateUpdateParams contains all the parameters to send to the API endpoint
for the software download state update operation typically these are written to a http.Request
*/
type SoftwareDownloadStateUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*BatchID
	  ID for the batch

	*/
	BatchID string
	/*Body*/
	Body *models.SoftwareUpdateState

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) WithTimeout(timeout time.Duration) *SoftwareDownloadStateUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) WithContext(ctx context.Context) *SoftwareDownloadStateUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) WithHTTPClient(client *http.Client) *SoftwareDownloadStateUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) WithAuthorization(authorization string) *SoftwareDownloadStateUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBatchID adds the batchID to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) WithBatchID(batchID string) *SoftwareDownloadStateUpdateParams {
	o.SetBatchID(batchID)
	return o
}

// SetBatchID adds the batchId to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) SetBatchID(batchID string) {
	o.BatchID = batchID
}

// WithBody adds the body to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) WithBody(body *models.SoftwareUpdateState) *SoftwareDownloadStateUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the software download state update params
func (o *SoftwareDownloadStateUpdateParams) SetBody(body *models.SoftwareUpdateState) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwareDownloadStateUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param batchId
	if err := r.SetPathParam("batchId", o.BatchID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
