// Code generated by go-swagger; DO NOT EDIT.

package helm

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

// NewHelmApplicationCreateParams creates a new HelmApplicationCreateParams object
// with the default values initialized.
func NewHelmApplicationCreateParams() *HelmApplicationCreateParams {
	var ()
	return &HelmApplicationCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHelmApplicationCreateParamsWithTimeout creates a new HelmApplicationCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHelmApplicationCreateParamsWithTimeout(timeout time.Duration) *HelmApplicationCreateParams {
	var ()
	return &HelmApplicationCreateParams{

		timeout: timeout,
	}
}

// NewHelmApplicationCreateParamsWithContext creates a new HelmApplicationCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewHelmApplicationCreateParamsWithContext(ctx context.Context) *HelmApplicationCreateParams {
	var ()
	return &HelmApplicationCreateParams{

		Context: ctx,
	}
}

// NewHelmApplicationCreateParamsWithHTTPClient creates a new HelmApplicationCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHelmApplicationCreateParamsWithHTTPClient(client *http.Client) *HelmApplicationCreateParams {
	var ()
	return &HelmApplicationCreateParams{
		HTTPClient: client,
	}
}

/*HelmApplicationCreateParams contains all the parameters to send to the API endpoint
for the helm application create operation typically these are written to a http.Request
*/
type HelmApplicationCreateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Application*/
	Application string
	/*Chart*/
	Chart runtime.NamedReadCloser
	/*Values*/
	Values runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the helm application create params
func (o *HelmApplicationCreateParams) WithTimeout(timeout time.Duration) *HelmApplicationCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the helm application create params
func (o *HelmApplicationCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the helm application create params
func (o *HelmApplicationCreateParams) WithContext(ctx context.Context) *HelmApplicationCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the helm application create params
func (o *HelmApplicationCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the helm application create params
func (o *HelmApplicationCreateParams) WithHTTPClient(client *http.Client) *HelmApplicationCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the helm application create params
func (o *HelmApplicationCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the helm application create params
func (o *HelmApplicationCreateParams) WithAuthorization(authorization string) *HelmApplicationCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the helm application create params
func (o *HelmApplicationCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithApplication adds the application to the helm application create params
func (o *HelmApplicationCreateParams) WithApplication(application string) *HelmApplicationCreateParams {
	o.SetApplication(application)
	return o
}

// SetApplication adds the application to the helm application create params
func (o *HelmApplicationCreateParams) SetApplication(application string) {
	o.Application = application
}

// WithChart adds the chart to the helm application create params
func (o *HelmApplicationCreateParams) WithChart(chart runtime.NamedReadCloser) *HelmApplicationCreateParams {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the helm application create params
func (o *HelmApplicationCreateParams) SetChart(chart runtime.NamedReadCloser) {
	o.Chart = chart
}

// WithValues adds the values to the helm application create params
func (o *HelmApplicationCreateParams) WithValues(values runtime.NamedReadCloser) *HelmApplicationCreateParams {
	o.SetValues(values)
	return o
}

// SetValues adds the values to the helm application create params
func (o *HelmApplicationCreateParams) SetValues(values runtime.NamedReadCloser) {
	o.Values = values
}

// WriteToRequest writes these params to a swagger request
func (o *HelmApplicationCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// form param application
	frApplication := o.Application
	fApplication := frApplication
	if fApplication != "" {
		if err := r.SetFormParam("application", fApplication); err != nil {
			return err
		}
	}

	// form file param chart
	if err := r.SetFileParam("chart", o.Chart); err != nil {
		return err
	}

	if o.Values != nil {

		if o.Values != nil {

			// form file param values
			if err := r.SetFileParam("values", o.Values); err != nil {
				return err
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
