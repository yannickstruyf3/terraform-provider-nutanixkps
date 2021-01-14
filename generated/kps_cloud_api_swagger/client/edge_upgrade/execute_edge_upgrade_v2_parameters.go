// Code generated by go-swagger; DO NOT EDIT.

package edge_upgrade

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

// NewExecuteEdgeUpgradeV2Params creates a new ExecuteEdgeUpgradeV2Params object
// with the default values initialized.
func NewExecuteEdgeUpgradeV2Params() *ExecuteEdgeUpgradeV2Params {
	var ()
	return &ExecuteEdgeUpgradeV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteEdgeUpgradeV2ParamsWithTimeout creates a new ExecuteEdgeUpgradeV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewExecuteEdgeUpgradeV2ParamsWithTimeout(timeout time.Duration) *ExecuteEdgeUpgradeV2Params {
	var ()
	return &ExecuteEdgeUpgradeV2Params{

		timeout: timeout,
	}
}

// NewExecuteEdgeUpgradeV2ParamsWithContext creates a new ExecuteEdgeUpgradeV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewExecuteEdgeUpgradeV2ParamsWithContext(ctx context.Context) *ExecuteEdgeUpgradeV2Params {
	var ()
	return &ExecuteEdgeUpgradeV2Params{

		Context: ctx,
	}
}

// NewExecuteEdgeUpgradeV2ParamsWithHTTPClient creates a new ExecuteEdgeUpgradeV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExecuteEdgeUpgradeV2ParamsWithHTTPClient(client *http.Client) *ExecuteEdgeUpgradeV2Params {
	var ()
	return &ExecuteEdgeUpgradeV2Params{
		HTTPClient: client,
	}
}

/*ExecuteEdgeUpgradeV2Params contains all the parameters to send to the API endpoint
for the execute edge upgrade v2 operation typically these are written to a http.Request
*/
type ExecuteEdgeUpgradeV2Params struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body
	  This is an execute edge upgrade request description

	*/
	Body *models.ExecuteEdgeUpgrade

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the execute edge upgrade v2 params
func (o *ExecuteEdgeUpgradeV2Params) WithTimeout(timeout time.Duration) *ExecuteEdgeUpgradeV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute edge upgrade v2 params
func (o *ExecuteEdgeUpgradeV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute edge upgrade v2 params
func (o *ExecuteEdgeUpgradeV2Params) WithContext(ctx context.Context) *ExecuteEdgeUpgradeV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute edge upgrade v2 params
func (o *ExecuteEdgeUpgradeV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute edge upgrade v2 params
func (o *ExecuteEdgeUpgradeV2Params) WithHTTPClient(client *http.Client) *ExecuteEdgeUpgradeV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute edge upgrade v2 params
func (o *ExecuteEdgeUpgradeV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the execute edge upgrade v2 params
func (o *ExecuteEdgeUpgradeV2Params) WithAuthorization(authorization string) *ExecuteEdgeUpgradeV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the execute edge upgrade v2 params
func (o *ExecuteEdgeUpgradeV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the execute edge upgrade v2 params
func (o *ExecuteEdgeUpgradeV2Params) WithBody(body *models.ExecuteEdgeUpgrade) *ExecuteEdgeUpgradeV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the execute edge upgrade v2 params
func (o *ExecuteEdgeUpgradeV2Params) SetBody(body *models.ExecuteEdgeUpgrade) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteEdgeUpgradeV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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