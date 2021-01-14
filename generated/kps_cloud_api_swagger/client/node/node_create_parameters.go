// Code generated by go-swagger; DO NOT EDIT.

package node

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

// NewNodeCreateParams creates a new NodeCreateParams object
// with the default values initialized.
func NewNodeCreateParams() *NodeCreateParams {
	var ()
	return &NodeCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodeCreateParamsWithTimeout creates a new NodeCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodeCreateParamsWithTimeout(timeout time.Duration) *NodeCreateParams {
	var ()
	return &NodeCreateParams{

		timeout: timeout,
	}
}

// NewNodeCreateParamsWithContext creates a new NodeCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodeCreateParamsWithContext(ctx context.Context) *NodeCreateParams {
	var ()
	return &NodeCreateParams{

		Context: ctx,
	}
}

// NewNodeCreateParamsWithHTTPClient creates a new NodeCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodeCreateParamsWithHTTPClient(client *http.Client) *NodeCreateParams {
	var ()
	return &NodeCreateParams{
		HTTPClient: client,
	}
}

/*NodeCreateParams contains all the parameters to send to the API endpoint
for the node create operation typically these are written to a http.Request
*/
type NodeCreateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body
	  Parameters and values used when creating a node

	*/
	Body *models.Node

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the node create params
func (o *NodeCreateParams) WithTimeout(timeout time.Duration) *NodeCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node create params
func (o *NodeCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node create params
func (o *NodeCreateParams) WithContext(ctx context.Context) *NodeCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node create params
func (o *NodeCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node create params
func (o *NodeCreateParams) WithHTTPClient(client *http.Client) *NodeCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node create params
func (o *NodeCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the node create params
func (o *NodeCreateParams) WithAuthorization(authorization string) *NodeCreateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the node create params
func (o *NodeCreateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the node create params
func (o *NodeCreateParams) WithBody(body *models.Node) *NodeCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the node create params
func (o *NodeCreateParams) SetBody(body *models.Node) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *NodeCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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