// Code generated by go-swagger; DO NOT EDIT.

package kiali

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

// NewNamespaceOverviewParams creates a new NamespaceOverviewParams object
// with the default values initialized.
func NewNamespaceOverviewParams() *NamespaceOverviewParams {
	var ()
	return &NamespaceOverviewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNamespaceOverviewParamsWithTimeout creates a new NamespaceOverviewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNamespaceOverviewParamsWithTimeout(timeout time.Duration) *NamespaceOverviewParams {
	var ()
	return &NamespaceOverviewParams{

		timeout: timeout,
	}
}

// NewNamespaceOverviewParamsWithContext creates a new NamespaceOverviewParams object
// with the default values initialized, and the ability to set a context for a request
func NewNamespaceOverviewParamsWithContext(ctx context.Context) *NamespaceOverviewParams {
	var ()
	return &NamespaceOverviewParams{

		Context: ctx,
	}
}

// NewNamespaceOverviewParamsWithHTTPClient creates a new NamespaceOverviewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNamespaceOverviewParamsWithHTTPClient(client *http.Client) *NamespaceOverviewParams {
	var ()
	return &NamespaceOverviewParams{
		HTTPClient: client,
	}
}

/*NamespaceOverviewParams contains all the parameters to send to the API endpoint
for the namespace overview operation typically these are written to a http.Request
*/
type NamespaceOverviewParams struct {

	/*Authorization
	  Format: Bearer &lt;token>, with &lt;token> from login API response.

	*/
	Authorization string
	/*Namespace
	  The namespace name.

	*/
	Namespace string
	/*ServiceDomain
	  ID of ServiceDomain to access.

	*/
	ServiceDomain string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the namespace overview params
func (o *NamespaceOverviewParams) WithTimeout(timeout time.Duration) *NamespaceOverviewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the namespace overview params
func (o *NamespaceOverviewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the namespace overview params
func (o *NamespaceOverviewParams) WithContext(ctx context.Context) *NamespaceOverviewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the namespace overview params
func (o *NamespaceOverviewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the namespace overview params
func (o *NamespaceOverviewParams) WithHTTPClient(client *http.Client) *NamespaceOverviewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the namespace overview params
func (o *NamespaceOverviewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the namespace overview params
func (o *NamespaceOverviewParams) WithAuthorization(authorization string) *NamespaceOverviewParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the namespace overview params
func (o *NamespaceOverviewParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithNamespace adds the namespace to the namespace overview params
func (o *NamespaceOverviewParams) WithNamespace(namespace string) *NamespaceOverviewParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the namespace overview params
func (o *NamespaceOverviewParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithServiceDomain adds the serviceDomain to the namespace overview params
func (o *NamespaceOverviewParams) WithServiceDomain(serviceDomain string) *NamespaceOverviewParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the namespace overview params
func (o *NamespaceOverviewParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WriteToRequest writes these params to a swagger request
func (o *NamespaceOverviewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param serviceDomain
	qrServiceDomain := o.ServiceDomain
	qServiceDomain := qrServiceDomain
	if qServiceDomain != "" {
		if err := r.SetQueryParam("serviceDomain", qServiceDomain); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}