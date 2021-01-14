// Code generated by go-swagger; DO NOT EDIT.

package user_public_key

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

// NewUserPublicKeyUpdateParams creates a new UserPublicKeyUpdateParams object
// with the default values initialized.
func NewUserPublicKeyUpdateParams() *UserPublicKeyUpdateParams {
	var ()
	return &UserPublicKeyUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserPublicKeyUpdateParamsWithTimeout creates a new UserPublicKeyUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserPublicKeyUpdateParamsWithTimeout(timeout time.Duration) *UserPublicKeyUpdateParams {
	var ()
	return &UserPublicKeyUpdateParams{

		timeout: timeout,
	}
}

// NewUserPublicKeyUpdateParamsWithContext creates a new UserPublicKeyUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserPublicKeyUpdateParamsWithContext(ctx context.Context) *UserPublicKeyUpdateParams {
	var ()
	return &UserPublicKeyUpdateParams{

		Context: ctx,
	}
}

// NewUserPublicKeyUpdateParamsWithHTTPClient creates a new UserPublicKeyUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserPublicKeyUpdateParamsWithHTTPClient(client *http.Client) *UserPublicKeyUpdateParams {
	var ()
	return &UserPublicKeyUpdateParams{
		HTTPClient: client,
	}
}

/*UserPublicKeyUpdateParams contains all the parameters to send to the API endpoint
for the user public key update operation typically these are written to a http.Request
*/
type UserPublicKeyUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.UserPublicKeyUpdatePayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user public key update params
func (o *UserPublicKeyUpdateParams) WithTimeout(timeout time.Duration) *UserPublicKeyUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user public key update params
func (o *UserPublicKeyUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user public key update params
func (o *UserPublicKeyUpdateParams) WithContext(ctx context.Context) *UserPublicKeyUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user public key update params
func (o *UserPublicKeyUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user public key update params
func (o *UserPublicKeyUpdateParams) WithHTTPClient(client *http.Client) *UserPublicKeyUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user public key update params
func (o *UserPublicKeyUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the user public key update params
func (o *UserPublicKeyUpdateParams) WithAuthorization(authorization string) *UserPublicKeyUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the user public key update params
func (o *UserPublicKeyUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the user public key update params
func (o *UserPublicKeyUpdateParams) WithBody(body *models.UserPublicKeyUpdatePayload) *UserPublicKeyUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user public key update params
func (o *UserPublicKeyUpdateParams) SetBody(body *models.UserPublicKeyUpdatePayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UserPublicKeyUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
