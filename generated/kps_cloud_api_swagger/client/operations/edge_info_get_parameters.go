// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewEdgeInfoGetParams creates a new EdgeInfoGetParams object
// with the default values initialized.
func NewEdgeInfoGetParams() *EdgeInfoGetParams {
	var ()
	return &EdgeInfoGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeInfoGetParamsWithTimeout creates a new EdgeInfoGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEdgeInfoGetParamsWithTimeout(timeout time.Duration) *EdgeInfoGetParams {
	var ()
	return &EdgeInfoGetParams{

		timeout: timeout,
	}
}

// NewEdgeInfoGetParamsWithContext creates a new EdgeInfoGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewEdgeInfoGetParamsWithContext(ctx context.Context) *EdgeInfoGetParams {
	var ()
	return &EdgeInfoGetParams{

		Context: ctx,
	}
}

// NewEdgeInfoGetParamsWithHTTPClient creates a new EdgeInfoGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEdgeInfoGetParamsWithHTTPClient(client *http.Client) *EdgeInfoGetParams {
	var ()
	return &EdgeInfoGetParams{
		HTTPClient: client,
	}
}

/*EdgeInfoGetParams contains all the parameters to send to the API endpoint
for the edge info get operation typically these are written to a http.Request
*/
type EdgeInfoGetParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*EdgeID
	  ID for the edge

	*/
	EdgeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edge info get params
func (o *EdgeInfoGetParams) WithTimeout(timeout time.Duration) *EdgeInfoGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge info get params
func (o *EdgeInfoGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge info get params
func (o *EdgeInfoGetParams) WithContext(ctx context.Context) *EdgeInfoGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge info get params
func (o *EdgeInfoGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge info get params
func (o *EdgeInfoGetParams) WithHTTPClient(client *http.Client) *EdgeInfoGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge info get params
func (o *EdgeInfoGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the edge info get params
func (o *EdgeInfoGetParams) WithAuthorization(authorization string) *EdgeInfoGetParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the edge info get params
func (o *EdgeInfoGetParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithEdgeID adds the edgeID to the edge info get params
func (o *EdgeInfoGetParams) WithEdgeID(edgeID string) *EdgeInfoGetParams {
	o.SetEdgeID(edgeID)
	return o
}

// SetEdgeID adds the edgeId to the edge info get params
func (o *EdgeInfoGetParams) SetEdgeID(edgeID string) {
	o.EdgeID = edgeID
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeInfoGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param edgeId
	if err := r.SetPathParam("edgeId", o.EdgeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}