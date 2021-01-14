// Code generated by go-swagger; DO NOT EDIT.

package user_props

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

// NewUserPropsUpdateV2Params creates a new UserPropsUpdateV2Params object
// with the default values initialized.
func NewUserPropsUpdateV2Params() *UserPropsUpdateV2Params {
	var ()
	return &UserPropsUpdateV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserPropsUpdateV2ParamsWithTimeout creates a new UserPropsUpdateV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserPropsUpdateV2ParamsWithTimeout(timeout time.Duration) *UserPropsUpdateV2Params {
	var ()
	return &UserPropsUpdateV2Params{

		timeout: timeout,
	}
}

// NewUserPropsUpdateV2ParamsWithContext creates a new UserPropsUpdateV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewUserPropsUpdateV2ParamsWithContext(ctx context.Context) *UserPropsUpdateV2Params {
	var ()
	return &UserPropsUpdateV2Params{

		Context: ctx,
	}
}

// NewUserPropsUpdateV2ParamsWithHTTPClient creates a new UserPropsUpdateV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserPropsUpdateV2ParamsWithHTTPClient(client *http.Client) *UserPropsUpdateV2Params {
	var ()
	return &UserPropsUpdateV2Params{
		HTTPClient: client,
	}
}

/*UserPropsUpdateV2Params contains all the parameters to send to the API endpoint
for the user props update v2 operation typically these are written to a http.Request
*/
type UserPropsUpdateV2Params struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.UserProps
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user props update v2 params
func (o *UserPropsUpdateV2Params) WithTimeout(timeout time.Duration) *UserPropsUpdateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user props update v2 params
func (o *UserPropsUpdateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user props update v2 params
func (o *UserPropsUpdateV2Params) WithContext(ctx context.Context) *UserPropsUpdateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user props update v2 params
func (o *UserPropsUpdateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user props update v2 params
func (o *UserPropsUpdateV2Params) WithHTTPClient(client *http.Client) *UserPropsUpdateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user props update v2 params
func (o *UserPropsUpdateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the user props update v2 params
func (o *UserPropsUpdateV2Params) WithAuthorization(authorization string) *UserPropsUpdateV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the user props update v2 params
func (o *UserPropsUpdateV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the user props update v2 params
func (o *UserPropsUpdateV2Params) WithBody(body *models.UserProps) *UserPropsUpdateV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user props update v2 params
func (o *UserPropsUpdateV2Params) SetBody(body *models.UserProps) {
	o.Body = body
}

// WithID adds the id to the user props update v2 params
func (o *UserPropsUpdateV2Params) WithID(id string) *UserPropsUpdateV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the user props update v2 params
func (o *UserPropsUpdateV2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserPropsUpdateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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