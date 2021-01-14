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

// NewWsMessagingOnDeleteMLModelParams creates a new WsMessagingOnDeleteMLModelParams object
// with the default values initialized.
func NewWsMessagingOnDeleteMLModelParams() *WsMessagingOnDeleteMLModelParams {
	var ()
	return &WsMessagingOnDeleteMLModelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingOnDeleteMLModelParamsWithTimeout creates a new WsMessagingOnDeleteMLModelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingOnDeleteMLModelParamsWithTimeout(timeout time.Duration) *WsMessagingOnDeleteMLModelParams {
	var ()
	return &WsMessagingOnDeleteMLModelParams{

		timeout: timeout,
	}
}

// NewWsMessagingOnDeleteMLModelParamsWithContext creates a new WsMessagingOnDeleteMLModelParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingOnDeleteMLModelParamsWithContext(ctx context.Context) *WsMessagingOnDeleteMLModelParams {
	var ()
	return &WsMessagingOnDeleteMLModelParams{

		Context: ctx,
	}
}

// NewWsMessagingOnDeleteMLModelParamsWithHTTPClient creates a new WsMessagingOnDeleteMLModelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingOnDeleteMLModelParamsWithHTTPClient(client *http.Client) *WsMessagingOnDeleteMLModelParams {
	var ()
	return &WsMessagingOnDeleteMLModelParams{
		HTTPClient: client,
	}
}

/*WsMessagingOnDeleteMLModelParams contains all the parameters to send to the API endpoint
for the ws messaging on delete m l model operation typically these are written to a http.Request
*/
type WsMessagingOnDeleteMLModelParams struct {

	/*Request*/
	Request *models.DeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging on delete m l model params
func (o *WsMessagingOnDeleteMLModelParams) WithTimeout(timeout time.Duration) *WsMessagingOnDeleteMLModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging on delete m l model params
func (o *WsMessagingOnDeleteMLModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging on delete m l model params
func (o *WsMessagingOnDeleteMLModelParams) WithContext(ctx context.Context) *WsMessagingOnDeleteMLModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging on delete m l model params
func (o *WsMessagingOnDeleteMLModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging on delete m l model params
func (o *WsMessagingOnDeleteMLModelParams) WithHTTPClient(client *http.Client) *WsMessagingOnDeleteMLModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging on delete m l model params
func (o *WsMessagingOnDeleteMLModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging on delete m l model params
func (o *WsMessagingOnDeleteMLModelParams) WithRequest(request *models.DeleteRequest) *WsMessagingOnDeleteMLModelParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging on delete m l model params
func (o *WsMessagingOnDeleteMLModelParams) SetRequest(request *models.DeleteRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingOnDeleteMLModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
