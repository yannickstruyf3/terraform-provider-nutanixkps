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

// NewWsMessagingOnCreateExecuteEdgeUpgradeParams creates a new WsMessagingOnCreateExecuteEdgeUpgradeParams object
// with the default values initialized.
func NewWsMessagingOnCreateExecuteEdgeUpgradeParams() *WsMessagingOnCreateExecuteEdgeUpgradeParams {
	var ()
	return &WsMessagingOnCreateExecuteEdgeUpgradeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingOnCreateExecuteEdgeUpgradeParamsWithTimeout creates a new WsMessagingOnCreateExecuteEdgeUpgradeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingOnCreateExecuteEdgeUpgradeParamsWithTimeout(timeout time.Duration) *WsMessagingOnCreateExecuteEdgeUpgradeParams {
	var ()
	return &WsMessagingOnCreateExecuteEdgeUpgradeParams{

		timeout: timeout,
	}
}

// NewWsMessagingOnCreateExecuteEdgeUpgradeParamsWithContext creates a new WsMessagingOnCreateExecuteEdgeUpgradeParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingOnCreateExecuteEdgeUpgradeParamsWithContext(ctx context.Context) *WsMessagingOnCreateExecuteEdgeUpgradeParams {
	var ()
	return &WsMessagingOnCreateExecuteEdgeUpgradeParams{

		Context: ctx,
	}
}

// NewWsMessagingOnCreateExecuteEdgeUpgradeParamsWithHTTPClient creates a new WsMessagingOnCreateExecuteEdgeUpgradeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingOnCreateExecuteEdgeUpgradeParamsWithHTTPClient(client *http.Client) *WsMessagingOnCreateExecuteEdgeUpgradeParams {
	var ()
	return &WsMessagingOnCreateExecuteEdgeUpgradeParams{
		HTTPClient: client,
	}
}

/*WsMessagingOnCreateExecuteEdgeUpgradeParams contains all the parameters to send to the API endpoint
for the ws messaging on create execute edge upgrade operation typically these are written to a http.Request
*/
type WsMessagingOnCreateExecuteEdgeUpgradeParams struct {

	/*Request*/
	Request *models.ObjectRequestBaseExecuteEdgeUpgrade

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging on create execute edge upgrade params
func (o *WsMessagingOnCreateExecuteEdgeUpgradeParams) WithTimeout(timeout time.Duration) *WsMessagingOnCreateExecuteEdgeUpgradeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging on create execute edge upgrade params
func (o *WsMessagingOnCreateExecuteEdgeUpgradeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging on create execute edge upgrade params
func (o *WsMessagingOnCreateExecuteEdgeUpgradeParams) WithContext(ctx context.Context) *WsMessagingOnCreateExecuteEdgeUpgradeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging on create execute edge upgrade params
func (o *WsMessagingOnCreateExecuteEdgeUpgradeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging on create execute edge upgrade params
func (o *WsMessagingOnCreateExecuteEdgeUpgradeParams) WithHTTPClient(client *http.Client) *WsMessagingOnCreateExecuteEdgeUpgradeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging on create execute edge upgrade params
func (o *WsMessagingOnCreateExecuteEdgeUpgradeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging on create execute edge upgrade params
func (o *WsMessagingOnCreateExecuteEdgeUpgradeParams) WithRequest(request *models.ObjectRequestBaseExecuteEdgeUpgrade) *WsMessagingOnCreateExecuteEdgeUpgradeParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging on create execute edge upgrade params
func (o *WsMessagingOnCreateExecuteEdgeUpgradeParams) SetRequest(request *models.ObjectRequestBaseExecuteEdgeUpgrade) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingOnCreateExecuteEdgeUpgradeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
