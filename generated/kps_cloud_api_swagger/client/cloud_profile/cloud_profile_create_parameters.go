// Code generated by go-swagger; DO NOT EDIT.

package cloud_profile

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

// NewCloudProfileCreateParams creates a new CloudProfileCreateParams object
// with the default values initialized.
func NewCloudProfileCreateParams() *CloudProfileCreateParams {
	var ()
	return &CloudProfileCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloudProfileCreateParamsWithTimeout creates a new CloudProfileCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloudProfileCreateParamsWithTimeout(timeout time.Duration) *CloudProfileCreateParams {
	var ()
	return &CloudProfileCreateParams{

		timeout: timeout,
	}
}

// NewCloudProfileCreateParamsWithContext creates a new CloudProfileCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloudProfileCreateParamsWithContext(ctx context.Context) *CloudProfileCreateParams {
	var ()
	return &CloudProfileCreateParams{

		Context: ctx,
	}
}

// NewCloudProfileCreateParamsWithHTTPClient creates a new CloudProfileCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloudProfileCreateParamsWithHTTPClient(client *http.Client) *CloudProfileCreateParams {
	var ()
	return &CloudProfileCreateParams{
		HTTPClient: client,
	}
}

/*CloudProfileCreateParams contains all the parameters to send to the API endpoint
for the cloud profile create operation typically these are written to a http.Request
*/
type CloudProfileCreateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body
	  Description for the cloud profile.

	*/
	Body *models.CloudProfile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cloud profile create params
func (o *CloudProfileCreateParams) WithTimeout(timeout time.Duration) *CloudProfileCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud profile create params
func (o *CloudProfileCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud profile create params
func (o *CloudProfileCreateParams) WithContext(ctx context.Context) *CloudProfileCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud profile create params
func (o *CloudProfileCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud profile create params
func (o *CloudProfileCreateParams) WithHTTPClient(client *http.Client) *CloudProfileCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud profile create params
func (o *CloudProfileCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the cloud profile create params
func (o *CloudProfileCreateParams) WithAuthorization(authorization string) *CloudProfileCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the cloud profile create params
func (o *CloudProfileCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the cloud profile create params
func (o *CloudProfileCreateParams) WithBody(body *models.CloudProfile) *CloudProfileCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cloud profile create params
func (o *CloudProfileCreateParams) SetBody(body *models.CloudProfile) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CloudProfileCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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