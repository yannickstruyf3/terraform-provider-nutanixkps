// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewWsMessagingOnCreateScriptRuntimeParams creates a new WsMessagingOnCreateScriptRuntimeParams object
// with the default values initialized.
func NewWsMessagingOnCreateScriptRuntimeParams() *WsMessagingOnCreateScriptRuntimeParams {
	var ()
	return &WsMessagingOnCreateScriptRuntimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingOnCreateScriptRuntimeParamsWithTimeout creates a new WsMessagingOnCreateScriptRuntimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingOnCreateScriptRuntimeParamsWithTimeout(timeout time.Duration) *WsMessagingOnCreateScriptRuntimeParams {
	var ()
	return &WsMessagingOnCreateScriptRuntimeParams{

		timeout: timeout,
	}
}

// NewWsMessagingOnCreateScriptRuntimeParamsWithContext creates a new WsMessagingOnCreateScriptRuntimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingOnCreateScriptRuntimeParamsWithContext(ctx context.Context) *WsMessagingOnCreateScriptRuntimeParams {
	var ()
	return &WsMessagingOnCreateScriptRuntimeParams{

		Context: ctx,
	}
}

// NewWsMessagingOnCreateScriptRuntimeParamsWithHTTPClient creates a new WsMessagingOnCreateScriptRuntimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingOnCreateScriptRuntimeParamsWithHTTPClient(client *http.Client) *WsMessagingOnCreateScriptRuntimeParams {
	var ()
	return &WsMessagingOnCreateScriptRuntimeParams{
		HTTPClient: client,
	}
}

/*WsMessagingOnCreateScriptRuntimeParams contains all the parameters to send to the API endpoint
for the ws messaging on create script runtime operation typically these are written to a http.Request
*/
type WsMessagingOnCreateScriptRuntimeParams struct {

	/*Request*/
	Request *models.ObjectRequestBaseScriptRuntime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging on create script runtime params
func (o *WsMessagingOnCreateScriptRuntimeParams) WithTimeout(timeout time.Duration) *WsMessagingOnCreateScriptRuntimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging on create script runtime params
func (o *WsMessagingOnCreateScriptRuntimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging on create script runtime params
func (o *WsMessagingOnCreateScriptRuntimeParams) WithContext(ctx context.Context) *WsMessagingOnCreateScriptRuntimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging on create script runtime params
func (o *WsMessagingOnCreateScriptRuntimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging on create script runtime params
func (o *WsMessagingOnCreateScriptRuntimeParams) WithHTTPClient(client *http.Client) *WsMessagingOnCreateScriptRuntimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging on create script runtime params
func (o *WsMessagingOnCreateScriptRuntimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging on create script runtime params
func (o *WsMessagingOnCreateScriptRuntimeParams) WithRequest(request *models.ObjectRequestBaseScriptRuntime) *WsMessagingOnCreateScriptRuntimeParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging on create script runtime params
func (o *WsMessagingOnCreateScriptRuntimeParams) SetRequest(request *models.ObjectRequestBaseScriptRuntime) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingOnCreateScriptRuntimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
