// Code generated by go-swagger; DO NOT EDIT.

package service_instance

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

// NewServiceInstanceDeleteParams creates a new ServiceInstanceDeleteParams object
// with the default values initialized.
func NewServiceInstanceDeleteParams() *ServiceInstanceDeleteParams {
	var ()
	return &ServiceInstanceDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceInstanceDeleteParamsWithTimeout creates a new ServiceInstanceDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceInstanceDeleteParamsWithTimeout(timeout time.Duration) *ServiceInstanceDeleteParams {
	var ()
	return &ServiceInstanceDeleteParams{

		timeout: timeout,
	}
}

// NewServiceInstanceDeleteParamsWithContext creates a new ServiceInstanceDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceInstanceDeleteParamsWithContext(ctx context.Context) *ServiceInstanceDeleteParams {
	var ()
	return &ServiceInstanceDeleteParams{

		Context: ctx,
	}
}

// NewServiceInstanceDeleteParamsWithHTTPClient creates a new ServiceInstanceDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceInstanceDeleteParamsWithHTTPClient(client *http.Client) *ServiceInstanceDeleteParams {
	var ()
	return &ServiceInstanceDeleteParams{
		HTTPClient: client,
	}
}

/*ServiceInstanceDeleteParams contains all the parameters to send to the API endpoint
for the service instance delete operation typically these are written to a http.Request
*/
type ServiceInstanceDeleteParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*SvcInstanceID
	  Service Instance ID

	*/
	SvcInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service instance delete params
func (o *ServiceInstanceDeleteParams) WithTimeout(timeout time.Duration) *ServiceInstanceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service instance delete params
func (o *ServiceInstanceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service instance delete params
func (o *ServiceInstanceDeleteParams) WithContext(ctx context.Context) *ServiceInstanceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service instance delete params
func (o *ServiceInstanceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service instance delete params
func (o *ServiceInstanceDeleteParams) WithHTTPClient(client *http.Client) *ServiceInstanceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service instance delete params
func (o *ServiceInstanceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the service instance delete params
func (o *ServiceInstanceDeleteParams) WithAuthorization(authorization string) *ServiceInstanceDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the service instance delete params
func (o *ServiceInstanceDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithSvcInstanceID adds the svcInstanceID to the service instance delete params
func (o *ServiceInstanceDeleteParams) WithSvcInstanceID(svcInstanceID string) *ServiceInstanceDeleteParams {
	o.SetSvcInstanceID(svcInstanceID)
	return o
}

// SetSvcInstanceID adds the svcInstanceId to the service instance delete params
func (o *ServiceInstanceDeleteParams) SetSvcInstanceID(svcInstanceID string) {
	o.SvcInstanceID = svcInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceInstanceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param svcInstanceId
	if err := r.SetPathParam("svcInstanceId", o.SvcInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}