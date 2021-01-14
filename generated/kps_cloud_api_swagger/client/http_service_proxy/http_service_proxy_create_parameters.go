// Code generated by go-swagger; DO NOT EDIT.

package http_service_proxy

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

// NewHTTPServiceProxyCreateParams creates a new HTTPServiceProxyCreateParams object
// with the default values initialized.
func NewHTTPServiceProxyCreateParams() *HTTPServiceProxyCreateParams {
	var ()
	return &HTTPServiceProxyCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHTTPServiceProxyCreateParamsWithTimeout creates a new HTTPServiceProxyCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHTTPServiceProxyCreateParamsWithTimeout(timeout time.Duration) *HTTPServiceProxyCreateParams {
	var ()
	return &HTTPServiceProxyCreateParams{

		timeout: timeout,
	}
}

// NewHTTPServiceProxyCreateParamsWithContext creates a new HTTPServiceProxyCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewHTTPServiceProxyCreateParamsWithContext(ctx context.Context) *HTTPServiceProxyCreateParams {
	var ()
	return &HTTPServiceProxyCreateParams{

		Context: ctx,
	}
}

// NewHTTPServiceProxyCreateParamsWithHTTPClient creates a new HTTPServiceProxyCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHTTPServiceProxyCreateParamsWithHTTPClient(client *http.Client) *HTTPServiceProxyCreateParams {
	var ()
	return &HTTPServiceProxyCreateParams{
		HTTPClient: client,
	}
}

/*HTTPServiceProxyCreateParams contains all the parameters to send to the API endpoint
for the HTTP service proxy create operation typically these are written to a http.Request
*/
type HTTPServiceProxyCreateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.HTTPServiceProxyCreateParamPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the HTTP service proxy create params
func (o *HTTPServiceProxyCreateParams) WithTimeout(timeout time.Duration) *HTTPServiceProxyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the HTTP service proxy create params
func (o *HTTPServiceProxyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the HTTP service proxy create params
func (o *HTTPServiceProxyCreateParams) WithContext(ctx context.Context) *HTTPServiceProxyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the HTTP service proxy create params
func (o *HTTPServiceProxyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the HTTP service proxy create params
func (o *HTTPServiceProxyCreateParams) WithHTTPClient(client *http.Client) *HTTPServiceProxyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the HTTP service proxy create params
func (o *HTTPServiceProxyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the HTTP service proxy create params
func (o *HTTPServiceProxyCreateParams) WithAuthorization(authorization string) *HTTPServiceProxyCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the HTTP service proxy create params
func (o *HTTPServiceProxyCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the HTTP service proxy create params
func (o *HTTPServiceProxyCreateParams) WithBody(body *models.HTTPServiceProxyCreateParamPayload) *HTTPServiceProxyCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the HTTP service proxy create params
func (o *HTTPServiceProxyCreateParams) SetBody(body *models.HTTPServiceProxyCreateParamPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *HTTPServiceProxyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
