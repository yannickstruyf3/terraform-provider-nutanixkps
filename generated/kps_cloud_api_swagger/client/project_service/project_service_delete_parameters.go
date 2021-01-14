// Code generated by go-swagger; DO NOT EDIT.

package project_service

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

// NewProjectServiceDeleteParams creates a new ProjectServiceDeleteParams object
// with the default values initialized.
func NewProjectServiceDeleteParams() *ProjectServiceDeleteParams {
	var ()
	return &ProjectServiceDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProjectServiceDeleteParamsWithTimeout creates a new ProjectServiceDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProjectServiceDeleteParamsWithTimeout(timeout time.Duration) *ProjectServiceDeleteParams {
	var ()
	return &ProjectServiceDeleteParams{

		timeout: timeout,
	}
}

// NewProjectServiceDeleteParamsWithContext creates a new ProjectServiceDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewProjectServiceDeleteParamsWithContext(ctx context.Context) *ProjectServiceDeleteParams {
	var ()
	return &ProjectServiceDeleteParams{

		Context: ctx,
	}
}

// NewProjectServiceDeleteParamsWithHTTPClient creates a new ProjectServiceDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProjectServiceDeleteParamsWithHTTPClient(client *http.Client) *ProjectServiceDeleteParams {
	var ()
	return &ProjectServiceDeleteParams{
		HTTPClient: client,
	}
}

/*ProjectServiceDeleteParams contains all the parameters to send to the API endpoint
for the project service delete operation typically these are written to a http.Request
*/
type ProjectServiceDeleteParams struct {

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

// WithTimeout adds the timeout to the project service delete params
func (o *ProjectServiceDeleteParams) WithTimeout(timeout time.Duration) *ProjectServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project service delete params
func (o *ProjectServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project service delete params
func (o *ProjectServiceDeleteParams) WithContext(ctx context.Context) *ProjectServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project service delete params
func (o *ProjectServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project service delete params
func (o *ProjectServiceDeleteParams) WithHTTPClient(client *http.Client) *ProjectServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project service delete params
func (o *ProjectServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the project service delete params
func (o *ProjectServiceDeleteParams) WithAuthorization(authorization string) *ProjectServiceDeleteParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the project service delete params
func (o *ProjectServiceDeleteParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the project service delete params
func (o *ProjectServiceDeleteParams) WithID(id string) *ProjectServiceDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the project service delete params
func (o *ProjectServiceDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
