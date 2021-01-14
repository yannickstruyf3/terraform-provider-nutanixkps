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

// NewWsMessagingOnUpdateDataSourceParams creates a new WsMessagingOnUpdateDataSourceParams object
// with the default values initialized.
func NewWsMessagingOnUpdateDataSourceParams() *WsMessagingOnUpdateDataSourceParams {
	var ()
	return &WsMessagingOnUpdateDataSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingOnUpdateDataSourceParamsWithTimeout creates a new WsMessagingOnUpdateDataSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingOnUpdateDataSourceParamsWithTimeout(timeout time.Duration) *WsMessagingOnUpdateDataSourceParams {
	var ()
	return &WsMessagingOnUpdateDataSourceParams{

		timeout: timeout,
	}
}

// NewWsMessagingOnUpdateDataSourceParamsWithContext creates a new WsMessagingOnUpdateDataSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingOnUpdateDataSourceParamsWithContext(ctx context.Context) *WsMessagingOnUpdateDataSourceParams {
	var ()
	return &WsMessagingOnUpdateDataSourceParams{

		Context: ctx,
	}
}

// NewWsMessagingOnUpdateDataSourceParamsWithHTTPClient creates a new WsMessagingOnUpdateDataSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingOnUpdateDataSourceParamsWithHTTPClient(client *http.Client) *WsMessagingOnUpdateDataSourceParams {
	var ()
	return &WsMessagingOnUpdateDataSourceParams{
		HTTPClient: client,
	}
}

/*WsMessagingOnUpdateDataSourceParams contains all the parameters to send to the API endpoint
for the ws messaging on update data source operation typically these are written to a http.Request
*/
type WsMessagingOnUpdateDataSourceParams struct {

	/*Request*/
	Request *models.ObjectRequestBaseDataSource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging on update data source params
func (o *WsMessagingOnUpdateDataSourceParams) WithTimeout(timeout time.Duration) *WsMessagingOnUpdateDataSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging on update data source params
func (o *WsMessagingOnUpdateDataSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging on update data source params
func (o *WsMessagingOnUpdateDataSourceParams) WithContext(ctx context.Context) *WsMessagingOnUpdateDataSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging on update data source params
func (o *WsMessagingOnUpdateDataSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging on update data source params
func (o *WsMessagingOnUpdateDataSourceParams) WithHTTPClient(client *http.Client) *WsMessagingOnUpdateDataSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging on update data source params
func (o *WsMessagingOnUpdateDataSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging on update data source params
func (o *WsMessagingOnUpdateDataSourceParams) WithRequest(request *models.ObjectRequestBaseDataSource) *WsMessagingOnUpdateDataSourceParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging on update data source params
func (o *WsMessagingOnUpdateDataSourceParams) SetRequest(request *models.ObjectRequestBaseDataSource) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingOnUpdateDataSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
