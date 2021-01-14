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
)

// NewContainerRegistryDeleteV2Params creates a new ContainerRegistryDeleteV2Params object
// with the default values initialized.
func NewContainerRegistryDeleteV2Params() *ContainerRegistryDeleteV2Params {
	var ()
	return &ContainerRegistryDeleteV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewContainerRegistryDeleteV2ParamsWithTimeout creates a new ContainerRegistryDeleteV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewContainerRegistryDeleteV2ParamsWithTimeout(timeout time.Duration) *ContainerRegistryDeleteV2Params {
	var ()
	return &ContainerRegistryDeleteV2Params{

		timeout: timeout,
	}
}

// NewContainerRegistryDeleteV2ParamsWithContext creates a new ContainerRegistryDeleteV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewContainerRegistryDeleteV2ParamsWithContext(ctx context.Context) *ContainerRegistryDeleteV2Params {
	var ()
	return &ContainerRegistryDeleteV2Params{

		Context: ctx,
	}
}

// NewContainerRegistryDeleteV2ParamsWithHTTPClient creates a new ContainerRegistryDeleteV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContainerRegistryDeleteV2ParamsWithHTTPClient(client *http.Client) *ContainerRegistryDeleteV2Params {
	var ()
	return &ContainerRegistryDeleteV2Params{
		HTTPClient: client,
	}
}

/*ContainerRegistryDeleteV2Params contains all the parameters to send to the API endpoint
for the container registry delete v2 operation typically these are written to a http.Request
*/
type ContainerRegistryDeleteV2Params struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*ID
	  ID of the entity

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the container registry delete v2 params
func (o *ContainerRegistryDeleteV2Params) WithTimeout(timeout time.Duration) *ContainerRegistryDeleteV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container registry delete v2 params
func (o *ContainerRegistryDeleteV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container registry delete v2 params
func (o *ContainerRegistryDeleteV2Params) WithContext(ctx context.Context) *ContainerRegistryDeleteV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container registry delete v2 params
func (o *ContainerRegistryDeleteV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container registry delete v2 params
func (o *ContainerRegistryDeleteV2Params) WithHTTPClient(client *http.Client) *ContainerRegistryDeleteV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container registry delete v2 params
func (o *ContainerRegistryDeleteV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the container registry delete v2 params
func (o *ContainerRegistryDeleteV2Params) WithAuthorization(authorization string) *ContainerRegistryDeleteV2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the container registry delete v2 params
func (o *ContainerRegistryDeleteV2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the container registry delete v2 params
func (o *ContainerRegistryDeleteV2Params) WithID(id string) *ContainerRegistryDeleteV2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the container registry delete v2 params
func (o *ContainerRegistryDeleteV2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerRegistryDeleteV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
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
