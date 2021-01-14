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

// NewWsMessagingOnCreateProjectServiceParams creates a new WsMessagingOnCreateProjectServiceParams object
// with the default values initialized.
func NewWsMessagingOnCreateProjectServiceParams() *WsMessagingOnCreateProjectServiceParams {
	var ()
	return &WsMessagingOnCreateProjectServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWsMessagingOnCreateProjectServiceParamsWithTimeout creates a new WsMessagingOnCreateProjectServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWsMessagingOnCreateProjectServiceParamsWithTimeout(timeout time.Duration) *WsMessagingOnCreateProjectServiceParams {
	var ()
	return &WsMessagingOnCreateProjectServiceParams{

		timeout: timeout,
	}
}

// NewWsMessagingOnCreateProjectServiceParamsWithContext creates a new WsMessagingOnCreateProjectServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewWsMessagingOnCreateProjectServiceParamsWithContext(ctx context.Context) *WsMessagingOnCreateProjectServiceParams {
	var ()
	return &WsMessagingOnCreateProjectServiceParams{

		Context: ctx,
	}
}

// NewWsMessagingOnCreateProjectServiceParamsWithHTTPClient creates a new WsMessagingOnCreateProjectServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWsMessagingOnCreateProjectServiceParamsWithHTTPClient(client *http.Client) *WsMessagingOnCreateProjectServiceParams {
	var ()
	return &WsMessagingOnCreateProjectServiceParams{
		HTTPClient: client,
	}
}

/*WsMessagingOnCreateProjectServiceParams contains all the parameters to send to the API endpoint
for the ws messaging on create project service operation typically these are written to a http.Request
*/
type WsMessagingOnCreateProjectServiceParams struct {

	/*Request*/
	Request *models.ObjectRequestBaseProjectService

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ws messaging on create project service params
func (o *WsMessagingOnCreateProjectServiceParams) WithTimeout(timeout time.Duration) *WsMessagingOnCreateProjectServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ws messaging on create project service params
func (o *WsMessagingOnCreateProjectServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ws messaging on create project service params
func (o *WsMessagingOnCreateProjectServiceParams) WithContext(ctx context.Context) *WsMessagingOnCreateProjectServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ws messaging on create project service params
func (o *WsMessagingOnCreateProjectServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ws messaging on create project service params
func (o *WsMessagingOnCreateProjectServiceParams) WithHTTPClient(client *http.Client) *WsMessagingOnCreateProjectServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ws messaging on create project service params
func (o *WsMessagingOnCreateProjectServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the ws messaging on create project service params
func (o *WsMessagingOnCreateProjectServiceParams) WithRequest(request *models.ObjectRequestBaseProjectService) *WsMessagingOnCreateProjectServiceParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the ws messaging on create project service params
func (o *WsMessagingOnCreateProjectServiceParams) SetRequest(request *models.ObjectRequestBaseProjectService) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *WsMessagingOnCreateProjectServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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