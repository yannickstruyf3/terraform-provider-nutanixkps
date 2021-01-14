// Code generated by go-swagger; DO NOT EDIT.

package runtime_environment

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

// NewRuntimeEnvironmentUpdateParams creates a new RuntimeEnvironmentUpdateParams object
// with the default values initialized.
func NewRuntimeEnvironmentUpdateParams() *RuntimeEnvironmentUpdateParams {
	var ()
	return &RuntimeEnvironmentUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRuntimeEnvironmentUpdateParamsWithTimeout creates a new RuntimeEnvironmentUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRuntimeEnvironmentUpdateParamsWithTimeout(timeout time.Duration) *RuntimeEnvironmentUpdateParams {
	var ()
	return &RuntimeEnvironmentUpdateParams{

		timeout: timeout,
	}
}

// NewRuntimeEnvironmentUpdateParamsWithContext creates a new RuntimeEnvironmentUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewRuntimeEnvironmentUpdateParamsWithContext(ctx context.Context) *RuntimeEnvironmentUpdateParams {
	var ()
	return &RuntimeEnvironmentUpdateParams{

		Context: ctx,
	}
}

// NewRuntimeEnvironmentUpdateParamsWithHTTPClient creates a new RuntimeEnvironmentUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRuntimeEnvironmentUpdateParamsWithHTTPClient(client *http.Client) *RuntimeEnvironmentUpdateParams {
	var ()
	return &RuntimeEnvironmentUpdateParams{
		HTTPClient: client,
	}
}

/*RuntimeEnvironmentUpdateParams contains all the parameters to send to the API endpoint
for the runtime environment update operation typically these are written to a http.Request
*/
type RuntimeEnvironmentUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.RuntimeEnvironment
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) WithTimeout(timeout time.Duration) *RuntimeEnvironmentUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) WithContext(ctx context.Context) *RuntimeEnvironmentUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) WithHTTPClient(client *http.Client) *RuntimeEnvironmentUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) WithAuthorization(authorization string) *RuntimeEnvironmentUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) WithBody(body *models.RuntimeEnvironment) *RuntimeEnvironmentUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) SetBody(body *models.RuntimeEnvironment) {
	o.Body = body
}

// WithID adds the id to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) WithID(id string) *RuntimeEnvironmentUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the runtime environment update params
func (o *RuntimeEnvironmentUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RuntimeEnvironmentUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}