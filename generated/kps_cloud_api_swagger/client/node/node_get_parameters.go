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
)

// NewNodeGetParams creates a new NodeGetParams object
// with the default values initialized.
func NewNodeGetParams() *NodeGetParams {
	var ()
	return &NodeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodeGetParamsWithTimeout creates a new NodeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodeGetParamsWithTimeout(timeout time.Duration) *NodeGetParams {
	var ()
	return &NodeGetParams{

		timeout: timeout,
	}
}

// NewNodeGetParamsWithContext creates a new NodeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodeGetParamsWithContext(ctx context.Context) *NodeGetParams {
	var ()
	return &NodeGetParams{

		Context: ctx,
	}
}

// NewNodeGetParamsWithHTTPClient creates a new NodeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodeGetParamsWithHTTPClient(client *http.Client) *NodeGetParams {
	var ()
	return &NodeGetParams{
		HTTPClient: client,
	}
}

/*NodeGetParams contains all the parameters to send to the API endpoint
for the node get operation typically these are written to a http.Request
*/
type NodeGetParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*NodeID
	  ID for the node

	*/
	NodeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the node get params
func (o *NodeGetParams) WithTimeout(timeout time.Duration) *NodeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node get params
func (o *NodeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node get params
func (o *NodeGetParams) WithContext(ctx context.Context) *NodeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node get params
func (o *NodeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node get params
func (o *NodeGetParams) WithHTTPClient(client *http.Client) *NodeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node get params
func (o *NodeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the node get params
func (o *NodeGetParams) WithAuthorization(authorization string) *NodeGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the node get params
func (o *NodeGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithNodeID adds the nodeID to the node get params
func (o *NodeGetParams) WithNodeID(nodeID string) *NodeGetParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the node get params
func (o *NodeGetParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WriteToRequest writes these params to a swagger request
func (o *NodeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param nodeId
	if err := r.SetPathParam("nodeId", o.NodeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}