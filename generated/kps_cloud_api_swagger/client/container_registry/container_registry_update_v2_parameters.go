// Code generated by go-swagger; DO NOT EDIT.

package container_registry

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

// NewContainerRegistryUpdateV2Params creates a new ContainerRegistryUpdateV2Params object
// with the default values initialized.
func NewContainerRegistryUpdateV2Params() *ContainerRegistryUpdateV2Params {
	var ()
	return &ContainerRegistryUpdateV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewContainerRegistryUpdateV2ParamsWithTimeout creates a new ContainerRegistryUpdateV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewContainerRegistryUpdateV2ParamsWithTimeout(timeout time.Duration) *ContainerRegistryUpdateV2Params {
	var ()
	return &ContainerRegistryUpdateV2Params{

		timeout: timeout,
	}
}

// NewContainerRegistryUpdateV2ParamsWithContext creates a new ContainerRegistryUpdateV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewContainerRegistryUpdateV2ParamsWithContext(ctx context.Context) *ContainerRegistryUpdateV2Params {
	var ()
	return &ContainerRegistryUpdateV2Params{

		Context: ctx,
	}
}

// NewContainerRegistryUpdateV2ParamsWithHTTPClient creates a new ContainerRegistryUpdateV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContainerRegistryUpdateV2ParamsWithHTTPClient(client *http.Client) *ContainerRegistryUpdateV2Params {
	var ()
	return &ContainerRegistryUpdateV2Params{
		HTTPClient: client,
	}
}

/*ContainerRegistryUpdateV2Params contains all the parameters to send to the API endpoint
for the container registry update v2 operation typically these are written to a http.Request
*/
type ContainerRegistryUpdateV2Params struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.ContainerRegistryV2
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) WithTimeout(timeout time.Duration) *ContainerRegistryUpdateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) WithContext(ctx context.Context) *ContainerRegistryUpdateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) WithHTTPClient(client *http.Client) *ContainerRegistryUpdateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) WithAuthorization(authorization string) *ContainerRegistryUpdateV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) WithBody(body *models.ContainerRegistryV2) *ContainerRegistryUpdateV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) SetBody(body *models.ContainerRegistryV2) {
	o.Body = body
}

// WithID adds the id to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) WithID(id string) *ContainerRegistryUpdateV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the container registry update v2 params
func (o *ContainerRegistryUpdateV2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerRegistryUpdateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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