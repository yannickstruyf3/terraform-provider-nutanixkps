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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSoftwareUpdateServiceDomainListParams creates a new SoftwareUpdateServiceDomainListParams object
// with the default values initialized.
func NewSoftwareUpdateServiceDomainListParams() *SoftwareUpdateServiceDomainListParams {
	var ()
	return &SoftwareUpdateServiceDomainListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwareUpdateServiceDomainListParamsWithTimeout creates a new SoftwareUpdateServiceDomainListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSoftwareUpdateServiceDomainListParamsWithTimeout(timeout time.Duration) *SoftwareUpdateServiceDomainListParams {
	var ()
	return &SoftwareUpdateServiceDomainListParams{

		timeout: timeout,
	}
}

// NewSoftwareUpdateServiceDomainListParamsWithContext creates a new SoftwareUpdateServiceDomainListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSoftwareUpdateServiceDomainListParamsWithContext(ctx context.Context) *SoftwareUpdateServiceDomainListParams {
	var ()
	return &SoftwareUpdateServiceDomainListParams{

		Context: ctx,
	}
}

// NewSoftwareUpdateServiceDomainListParamsWithHTTPClient creates a new SoftwareUpdateServiceDomainListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSoftwareUpdateServiceDomainListParamsWithHTTPClient(client *http.Client) *SoftwareUpdateServiceDomainListParams {
	var ()
	return &SoftwareUpdateServiceDomainListParams{
		HTTPClient: client,
	}
}

/*SoftwareUpdateServiceDomainListParams contains all the parameters to send to the API endpoint
for the software update service domain list operation typically these are written to a http.Request
*/
type SoftwareUpdateServiceDomainListParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*IsLatestBatch
	  Fetch only latest batches

	*/
	IsLatestBatch *bool
	/*SvcDomainID
	  Service Domain ID

	*/
	SvcDomainID *string
	/*Type
	  Software DOWNLOAD/UPGRADE

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) WithTimeout(timeout time.Duration) *SoftwareUpdateServiceDomainListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) WithContext(ctx context.Context) *SoftwareUpdateServiceDomainListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) WithHTTPClient(client *http.Client) *SoftwareUpdateServiceDomainListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) WithAuthorization(authorization string) *SoftwareUpdateServiceDomainListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithIsLatestBatch adds the isLatestBatch to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) WithIsLatestBatch(isLatestBatch *bool) *SoftwareUpdateServiceDomainListParams {
	o.SetIsLatestBatch(isLatestBatch)
	return o
}

// SetIsLatestBatch adds the isLatestBatch to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) SetIsLatestBatch(isLatestBatch *bool) {
	o.IsLatestBatch = isLatestBatch
}

// WithSvcDomainID adds the svcDomainID to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) WithSvcDomainID(svcDomainID *string) *SoftwareUpdateServiceDomainListParams {
	o.SetSvcDomainID(svcDomainID)
	return o
}

// SetSvcDomainID adds the svcDomainId to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) SetSvcDomainID(svcDomainID *string) {
	o.SvcDomainID = svcDomainID
}

// WithType adds the typeVar to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) WithType(typeVar *string) *SoftwareUpdateServiceDomainListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the software update service domain list params
func (o *SoftwareUpdateServiceDomainListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwareUpdateServiceDomainListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.IsLatestBatch != nil {

		// query param isLatestBatch
		var qrIsLatestBatch bool
		if o.IsLatestBatch != nil {
			qrIsLatestBatch = *o.IsLatestBatch
		}
		qIsLatestBatch := swag.FormatBool(qrIsLatestBatch)
		if qIsLatestBatch != "" {
			if err := r.SetQueryParam("isLatestBatch", qIsLatestBatch); err != nil {
				return err
			}
		}

	}

	if o.SvcDomainID != nil {

		// query param svcDomainId
		var qrSvcDomainID string
		if o.SvcDomainID != nil {
			qrSvcDomainID = *o.SvcDomainID
		}
		qSvcDomainID := qrSvcDomainID
		if qSvcDomainID != "" {
			if err := r.SetQueryParam("svcDomainId", qSvcDomainID); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}