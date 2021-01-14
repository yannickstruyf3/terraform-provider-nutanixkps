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

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// NewSensorUpdateParams creates a new SensorUpdateParams object
// with the default values initialized.
func NewSensorUpdateParams() *SensorUpdateParams {
	var ()
	return &SensorUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSensorUpdateParamsWithTimeout creates a new SensorUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSensorUpdateParamsWithTimeout(timeout time.Duration) *SensorUpdateParams {
	var ()
	return &SensorUpdateParams{

		timeout: timeout,
	}
}

// NewSensorUpdateParamsWithContext creates a new SensorUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSensorUpdateParamsWithContext(ctx context.Context) *SensorUpdateParams {
	var ()
	return &SensorUpdateParams{

		Context: ctx,
	}
}

// NewSensorUpdateParamsWithHTTPClient creates a new SensorUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSensorUpdateParamsWithHTTPClient(client *http.Client) *SensorUpdateParams {
	var ()
	return &SensorUpdateParams{
		HTTPClient: client,
	}
}

/*SensorUpdateParams contains all the parameters to send to the API endpoint
for the sensor update operation typically these are written to a http.Request
*/
type SensorUpdateParams struct {

	/*Authorization
	  Format: Bearer <token>, with <token> from login API response.

	*/
	Authorization string
	/*Body*/
	Body *models.Sensor

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sensor update params
func (o *SensorUpdateParams) WithTimeout(timeout time.Duration) *SensorUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sensor update params
func (o *SensorUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sensor update params
func (o *SensorUpdateParams) WithContext(ctx context.Context) *SensorUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sensor update params
func (o *SensorUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sensor update params
func (o *SensorUpdateParams) WithHTTPClient(client *http.Client) *SensorUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sensor update params
func (o *SensorUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the sensor update params
func (o *SensorUpdateParams) WithAuthorization(authorization string) *SensorUpdateParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the sensor update params
func (o *SensorUpdateParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the sensor update params
func (o *SensorUpdateParams) WithBody(body *models.Sensor) *SensorUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the sensor update params
func (o *SensorUpdateParams) SetBody(body *models.Sensor) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SensorUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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