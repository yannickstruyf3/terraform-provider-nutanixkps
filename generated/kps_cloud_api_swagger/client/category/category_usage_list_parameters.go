// Code generated by go-swagger; DO NOT EDIT.

package category

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

// NewCategoryUsageListParams creates a new CategoryUsageListParams object
// with the default values initialized.
func NewCategoryUsageListParams() *CategoryUsageListParams {
	var ()
	return &CategoryUsageListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCategoryUsageListParamsWithTimeout creates a new CategoryUsageListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCategoryUsageListParamsWithTimeout(timeout time.Duration) *CategoryUsageListParams {
	var ()
	return &CategoryUsageListParams{

		timeout: timeout,
	}
}

// NewCategoryUsageListParamsWithContext creates a new CategoryUsageListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCategoryUsageListParamsWithContext(ctx context.Context) *CategoryUsageListParams {
	var ()
	return &CategoryUsageListParams{

		Context: ctx,
	}
}

// NewCategoryUsageListParamsWithHTTPClient creates a new CategoryUsageListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCategoryUsageListParamsWithHTTPClient(client *http.Client) *CategoryUsageListParams {
	var ()
	return &CategoryUsageListParams{
		HTTPClient: client,
	}
}

/*CategoryUsageListParams contains all the parameters to send to the API endpoint
for the category usage list operation typically these are written to a http.Request
*/
type CategoryUsageListParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the category usage list params
func (o *CategoryUsageListParams) WithTimeout(timeout time.Duration) *CategoryUsageListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the category usage list params
func (o *CategoryUsageListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the category usage list params
func (o *CategoryUsageListParams) WithContext(ctx context.Context) *CategoryUsageListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the category usage list params
func (o *CategoryUsageListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the category usage list params
func (o *CategoryUsageListParams) WithHTTPClient(client *http.Client) *CategoryUsageListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the category usage list params
func (o *CategoryUsageListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the category usage list params
func (o *CategoryUsageListParams) WithAuthorization(authorization string) *CategoryUsageListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the category usage list params
func (o *CategoryUsageListParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *CategoryUsageListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
