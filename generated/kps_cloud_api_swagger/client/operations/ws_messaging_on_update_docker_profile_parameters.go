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

// NewWsMessagingOnUpdateDockerProfileParams creates a new WsMessagingOnUpdateDockerProfileParams object
// with the default values initialized.
func NewWsMessagingOnUpdateDockerProfileParams() *WsMessagingOnUpdateDockerProfileParams {
	var ()
	return &WsMessagingOnUpdateDockerProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingOnUpdateDockerProfileParamsWithTimeout creates a new WsMessagingOnUpdateDockerProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingOnUpdateDockerProfileParamsWithTimeout(timeout time.Duration) *WsMessagingOnUpdateDockerProfileParams {
	var ()
	return &WsMessagingOnUpdateDockerProfileParams{

		timeout: timeout,
	}
}

// NewWsMessagingOnUpdateDockerProfileParamsWithContext creates a new WsMessagingOnUpdateDockerProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingOnUpdateDockerProfileParamsWithContext(ctx context.Context) *WsMessagingOnUpdateDockerProfileParams {
	var ()
	return &WsMessagingOnUpdateDockerProfileParams{

		Context: ctx,
	}
}

// NewWsMessagingOnUpdateDockerProfileParamsWithHTTPClient creates a new WsMessagingOnUpdateDockerProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingOnUpdateDockerProfileParamsWithHTTPClient(client *http.Client) *WsMessagingOnUpdateDockerProfileParams {
	var ()
	return &WsMessagingOnUpdateDockerProfileParams{
		HTTPClient: client,
	}
}

/*WsMessagingOnUpdateDockerProfileParams contains all the parameters to send to the API endpoint
for the ws messaging on update docker profile operation typically these are written to a http.Request
*/
type WsMessagingOnUpdateDockerProfileParams struct {

	/*Request*/
	Request *models.ObjectRequestBaseDockerProfile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging on update docker profile params
func (o *WsMessagingOnUpdateDockerProfileParams) WithTimeout(timeout time.Duration) *WsMessagingOnUpdateDockerProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging on update docker profile params
func (o *WsMessagingOnUpdateDockerProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging on update docker profile params
func (o *WsMessagingOnUpdateDockerProfileParams) WithContext(ctx context.Context) *WsMessagingOnUpdateDockerProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging on update docker profile params
func (o *WsMessagingOnUpdateDockerProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging on update docker profile params
func (o *WsMessagingOnUpdateDockerProfileParams) WithHTTPClient(client *http.Client) *WsMessagingOnUpdateDockerProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging on update docker profile params
func (o *WsMessagingOnUpdateDockerProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging on update docker profile params
func (o *WsMessagingOnUpdateDockerProfileParams) WithRequest(request *models.ObjectRequestBaseDockerProfile) *WsMessagingOnUpdateDockerProfileParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging on update docker profile params
func (o *WsMessagingOnUpdateDockerProfileParams) SetRequest(request *models.ObjectRequestBaseDockerProfile) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingOnUpdateDockerProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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