// Code generated by go-swagger; DO NOT EDIT.

package log_collector

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

// NewLogCollectorStopParams creates a new LogCollectorStopParams object
// with the default values initialized.
func NewLogCollectorStopParams() *LogCollectorStopParams {
	var ()
	return &LogCollectorStopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogCollectorStopParamsWithTimeout creates a new LogCollectorStopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogCollectorStopParamsWithTimeout(timeout time.Duration) *LogCollectorStopParams {
	var ()
	return &LogCollectorStopParams{

		timeout: timeout,
	}
}

// NewLogCollectorStopParamsWithContext creates a new LogCollectorStopParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogCollectorStopParamsWithContext(ctx context.Context) *LogCollectorStopParams {
	var ()
	return &LogCollectorStopParams{

		Context: ctx,
	}
}

// NewLogCollectorStopParamsWithHTTPClient creates a new LogCollectorStopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogCollectorStopParamsWithHTTPClient(client *http.Client) *LogCollectorStopParams {
	var ()
	return &LogCollectorStopParams{
		HTTPClient: client,
	}
}

/*LogCollectorStopParams contains all the parameters to send to the API endpoint
for the log collector stop operation typically these are written to a http.Request
*/
type LogCollectorStopParams struct {

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

// WithTimeout adds the timeout to the log collector stop params
func (o *LogCollectorStopParams) WithTimeout(timeout time.Duration) *LogCollectorStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the log collector stop params
func (o *LogCollectorStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the log collector stop params
func (o *LogCollectorStopParams) WithContext(ctx context.Context) *LogCollectorStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the log collector stop params
func (o *LogCollectorStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the log collector stop params
func (o *LogCollectorStopParams) WithHTTPClient(client *http.Client) *LogCollectorStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the log collector stop params
func (o *LogCollectorStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the log collector stop params
func (o *LogCollectorStopParams) WithAuthorization(authorization string) *LogCollectorStopParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the log collector stop params
func (o *LogCollectorStopParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the log collector stop params
func (o *LogCollectorStopParams) WithID(id string) *LogCollectorStopParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the log collector stop params
func (o *LogCollectorStopParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LogCollectorStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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