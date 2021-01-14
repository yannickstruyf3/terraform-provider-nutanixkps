// Code generated by go-swagger; DO NOT EDIT.

package kiali

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWorkloadMetricsParams creates a new WorkloadMetricsParams object
// with the default values initialized.
func NewWorkloadMetricsParams() *WorkloadMetricsParams {
	var (
		avgDefault             = bool(true)
		directionDefault       = string("outbound")
		durationDefault        = int64(1800)
		rateFuncDefault        = string("rate")
		rateIntervalDefault    = string("1m")
		reporterDefault        = string("source")
		requestProtocolDefault = string("all protocols")
		stepDefault            = int64(15)
	)
	return &WorkloadMetricsParams{
		Avg:             &avgDefault,
		Direction:       &directionDefault,
		Duration:        &durationDefault,
		RateFunc:        &rateFuncDefault,
		RateInterval:    &rateIntervalDefault,
		Reporter:        &reporterDefault,
		RequestProtocol: &requestProtocolDefault,
		Step:            &stepDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkloadMetricsParamsWithTimeout creates a new WorkloadMetricsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkloadMetricsParamsWithTimeout(timeout time.Duration) *WorkloadMetricsParams {
	var (
		avgDefault             = bool(true)
		directionDefault       = string("outbound")
		durationDefault        = int64(1800)
		rateFuncDefault        = string("rate")
		rateIntervalDefault    = string("1m")
		reporterDefault        = string("source")
		requestProtocolDefault = string("all protocols")
		stepDefault            = int64(15)
	)
	return &WorkloadMetricsParams{
		Avg:             &avgDefault,
		Direction:       &directionDefault,
		Duration:        &durationDefault,
		RateFunc:        &rateFuncDefault,
		RateInterval:    &rateIntervalDefault,
		Reporter:        &reporterDefault,
		RequestProtocol: &requestProtocolDefault,
		Step:            &stepDefault,

		timeout: timeout,
	}
}

// NewWorkloadMetricsParamsWithContext creates a new WorkloadMetricsParams object
// with the default values initialized, and the ability to set a context for a request
func NewWorkloadMetricsParamsWithContext(ctx context.Context) *WorkloadMetricsParams {
	var (
		avgDefault             = bool(true)
		directionDefault       = string("outbound")
		durationDefault        = int64(1800)
		rateFuncDefault        = string("rate")
		rateIntervalDefault    = string("1m")
		reporterDefault        = string("source")
		requestProtocolDefault = string("all protocols")
		stepDefault            = int64(15)
	)
	return &WorkloadMetricsParams{
		Avg:             &avgDefault,
		Direction:       &directionDefault,
		Duration:        &durationDefault,
		RateFunc:        &rateFuncDefault,
		RateInterval:    &rateIntervalDefault,
		Reporter:        &reporterDefault,
		RequestProtocol: &requestProtocolDefault,
		Step:            &stepDefault,

		Context: ctx,
	}
}

// NewWorkloadMetricsParamsWithHTTPClient creates a new WorkloadMetricsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWorkloadMetricsParamsWithHTTPClient(client *http.Client) *WorkloadMetricsParams {
	var (
		avgDefault             = bool(true)
		directionDefault       = string("outbound")
		durationDefault        = int64(1800)
		rateFuncDefault        = string("rate")
		rateIntervalDefault    = string("1m")
		reporterDefault        = string("source")
		requestProtocolDefault = string("all protocols")
		stepDefault            = int64(15)
	)
	return &WorkloadMetricsParams{
		Avg:             &avgDefault,
		Direction:       &directionDefault,
		Duration:        &durationDefault,
		RateFunc:        &rateFuncDefault,
		RateInterval:    &rateIntervalDefault,
		Reporter:        &reporterDefault,
		RequestProtocol: &requestProtocolDefault,
		Step:            &stepDefault,
		HTTPClient:      client,
	}
}

/*WorkloadMetricsParams contains all the parameters to send to the API endpoint
for the workload metrics operation typically these are written to a http.Request
*/
type WorkloadMetricsParams struct {

	/*Authorization
	  Format: Bearer &lt;token>, with &lt;token> from login API response.

	*/
	Authorization string
	/*Avg
	  Flag for fetching histogram average. Default is true.

	*/
	Avg *bool
	/*ByLabels
	  List of labels to use for grouping metrics (via Prometheus 'by' clause).

	*/
	ByLabels []string
	/*Direction
	  Traffic direction: 'inbound' or 'outbound'.

	*/
	Direction *string
	/*Duration
	  Duration of the query period, in seconds.

	*/
	Duration *int64
	/*Filters
	  List of metrics to fetch. Fetch all metrics when empty. List entries are Kiali internal metric names.

	*/
	Filters []string
	/*Namespace
	  The namespace name.

	*/
	Namespace string
	/*Quantiles
	  List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99].

	*/
	Quantiles []string
	/*RateFunc
	  Prometheus function used to calculate rate: 'rate' or 'irate'.

	*/
	RateFunc *string
	/*RateInterval
	  Interval used for rate and histogram calculation.

	*/
	RateInterval *string
	/*Reporter
	  Istio telemetry reporter: 'source' or 'destination'.

	*/
	Reporter *string
	/*RequestProtocol
	  Desired request protocol for the telemetry: For example, 'http' or 'grpc'.

	*/
	RequestProtocol *string
	/*ServiceDomain
	  ID of ServiceDomain to access.

	*/
	ServiceDomain string
	/*Step
	  Step between [graph] datapoints, in seconds.

	*/
	Step *int64
	/*Version
	  Filters metrics by the specified version.

	*/
	Version *string
	/*Workload
	  The workload name.

	*/
	Workload string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the workload metrics params
func (o *WorkloadMetricsParams) WithTimeout(timeout time.Duration) *WorkloadMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workload metrics params
func (o *WorkloadMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workload metrics params
func (o *WorkloadMetricsParams) WithContext(ctx context.Context) *WorkloadMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workload metrics params
func (o *WorkloadMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workload metrics params
func (o *WorkloadMetricsParams) WithHTTPClient(client *http.Client) *WorkloadMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workload metrics params
func (o *WorkloadMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the workload metrics params
func (o *WorkloadMetricsParams) WithAuthorization(authorization string) *WorkloadMetricsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the workload metrics params
func (o *WorkloadMetricsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAvg adds the avg to the workload metrics params
func (o *WorkloadMetricsParams) WithAvg(avg *bool) *WorkloadMetricsParams {
	o.SetAvg(avg)
	return o
}

// SetAvg adds the avg to the workload metrics params
func (o *WorkloadMetricsParams) SetAvg(avg *bool) {
	o.Avg = avg
}

// WithByLabels adds the byLabels to the workload metrics params
func (o *WorkloadMetricsParams) WithByLabels(byLabels []string) *WorkloadMetricsParams {
	o.SetByLabels(byLabels)
	return o
}

// SetByLabels adds the byLabels to the workload metrics params
func (o *WorkloadMetricsParams) SetByLabels(byLabels []string) {
	o.ByLabels = byLabels
}

// WithDirection adds the direction to the workload metrics params
func (o *WorkloadMetricsParams) WithDirection(direction *string) *WorkloadMetricsParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the workload metrics params
func (o *WorkloadMetricsParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithDuration adds the duration to the workload metrics params
func (o *WorkloadMetricsParams) WithDuration(duration *int64) *WorkloadMetricsParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the workload metrics params
func (o *WorkloadMetricsParams) SetDuration(duration *int64) {
	o.Duration = duration
}

// WithFilters adds the filters to the workload metrics params
func (o *WorkloadMetricsParams) WithFilters(filters []string) *WorkloadMetricsParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the workload metrics params
func (o *WorkloadMetricsParams) SetFilters(filters []string) {
	o.Filters = filters
}

// WithNamespace adds the namespace to the workload metrics params
func (o *WorkloadMetricsParams) WithNamespace(namespace string) *WorkloadMetricsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the workload metrics params
func (o *WorkloadMetricsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithQuantiles adds the quantiles to the workload metrics params
func (o *WorkloadMetricsParams) WithQuantiles(quantiles []string) *WorkloadMetricsParams {
	o.SetQuantiles(quantiles)
	return o
}

// SetQuantiles adds the quantiles to the workload metrics params
func (o *WorkloadMetricsParams) SetQuantiles(quantiles []string) {
	o.Quantiles = quantiles
}

// WithRateFunc adds the rateFunc to the workload metrics params
func (o *WorkloadMetricsParams) WithRateFunc(rateFunc *string) *WorkloadMetricsParams {
	o.SetRateFunc(rateFunc)
	return o
}

// SetRateFunc adds the rateFunc to the workload metrics params
func (o *WorkloadMetricsParams) SetRateFunc(rateFunc *string) {
	o.RateFunc = rateFunc
}

// WithRateInterval adds the rateInterval to the workload metrics params
func (o *WorkloadMetricsParams) WithRateInterval(rateInterval *string) *WorkloadMetricsParams {
	o.SetRateInterval(rateInterval)
	return o
}

// SetRateInterval adds the rateInterval to the workload metrics params
func (o *WorkloadMetricsParams) SetRateInterval(rateInterval *string) {
	o.RateInterval = rateInterval
}

// WithReporter adds the reporter to the workload metrics params
func (o *WorkloadMetricsParams) WithReporter(reporter *string) *WorkloadMetricsParams {
	o.SetReporter(reporter)
	return o
}

// SetReporter adds the reporter to the workload metrics params
func (o *WorkloadMetricsParams) SetReporter(reporter *string) {
	o.Reporter = reporter
}

// WithRequestProtocol adds the requestProtocol to the workload metrics params
func (o *WorkloadMetricsParams) WithRequestProtocol(requestProtocol *string) *WorkloadMetricsParams {
	o.SetRequestProtocol(requestProtocol)
	return o
}

// SetRequestProtocol adds the requestProtocol to the workload metrics params
func (o *WorkloadMetricsParams) SetRequestProtocol(requestProtocol *string) {
	o.RequestProtocol = requestProtocol
}

// WithServiceDomain adds the serviceDomain to the workload metrics params
func (o *WorkloadMetricsParams) WithServiceDomain(serviceDomain string) *WorkloadMetricsParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the workload metrics params
func (o *WorkloadMetricsParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WithStep adds the step to the workload metrics params
func (o *WorkloadMetricsParams) WithStep(step *int64) *WorkloadMetricsParams {
	o.SetStep(step)
	return o
}

// SetStep adds the step to the workload metrics params
func (o *WorkloadMetricsParams) SetStep(step *int64) {
	o.Step = step
}

// WithVersion adds the version to the workload metrics params
func (o *WorkloadMetricsParams) WithVersion(version *string) *WorkloadMetricsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the workload metrics params
func (o *WorkloadMetricsParams) SetVersion(version *string) {
	o.Version = version
}

// WithWorkload adds the workload to the workload metrics params
func (o *WorkloadMetricsParams) WithWorkload(workload string) *WorkloadMetricsParams {
	o.SetWorkload(workload)
	return o
}

// SetWorkload adds the workload to the workload metrics params
func (o *WorkloadMetricsParams) SetWorkload(workload string) {
	o.Workload = workload
}

// WriteToRequest writes these params to a swagger request
func (o *WorkloadMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Avg != nil {

		// query param avg
		var qrAvg bool
		if o.Avg != nil {
			qrAvg = *o.Avg
		}
		qAvg := swag.FormatBool(qrAvg)
		if qAvg != "" {
			if err := r.SetQueryParam("avg", qAvg); err != nil {
				return err
			}
		}

	}

	valuesByLabels := o.ByLabels

	joinedByLabels := swag.JoinByFormat(valuesByLabels, "")
	// query array param byLabels[]
	if err := r.SetQueryParam("byLabels[]", joinedByLabels...); err != nil {
		return err
	}

	if o.Direction != nil {

		// query param direction
		var qrDirection string
		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {
			if err := r.SetQueryParam("direction", qDirection); err != nil {
				return err
			}
		}

	}

	if o.Duration != nil {

		// query param duration
		var qrDuration int64
		if o.Duration != nil {
			qrDuration = *o.Duration
		}
		qDuration := swag.FormatInt64(qrDuration)
		if qDuration != "" {
			if err := r.SetQueryParam("duration", qDuration); err != nil {
				return err
			}
		}

	}

	valuesFilters := o.Filters

	joinedFilters := swag.JoinByFormat(valuesFilters, "")
	// query array param filters[]
	if err := r.SetQueryParam("filters[]", joinedFilters...); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	valuesQuantiles := o.Quantiles

	joinedQuantiles := swag.JoinByFormat(valuesQuantiles, "")
	// query array param quantiles[]
	if err := r.SetQueryParam("quantiles[]", joinedQuantiles...); err != nil {
		return err
	}

	if o.RateFunc != nil {

		// query param rateFunc
		var qrRateFunc string
		if o.RateFunc != nil {
			qrRateFunc = *o.RateFunc
		}
		qRateFunc := qrRateFunc
		if qRateFunc != "" {
			if err := r.SetQueryParam("rateFunc", qRateFunc); err != nil {
				return err
			}
		}

	}

	if o.RateInterval != nil {

		// query param rateInterval
		var qrRateInterval string
		if o.RateInterval != nil {
			qrRateInterval = *o.RateInterval
		}
		qRateInterval := qrRateInterval
		if qRateInterval != "" {
			if err := r.SetQueryParam("rateInterval", qRateInterval); err != nil {
				return err
			}
		}

	}

	if o.Reporter != nil {

		// query param reporter
		var qrReporter string
		if o.Reporter != nil {
			qrReporter = *o.Reporter
		}
		qReporter := qrReporter
		if qReporter != "" {
			if err := r.SetQueryParam("reporter", qReporter); err != nil {
				return err
			}
		}

	}

	if o.RequestProtocol != nil {

		// query param requestProtocol
		var qrRequestProtocol string
		if o.RequestProtocol != nil {
			qrRequestProtocol = *o.RequestProtocol
		}
		qRequestProtocol := qrRequestProtocol
		if qRequestProtocol != "" {
			if err := r.SetQueryParam("requestProtocol", qRequestProtocol); err != nil {
				return err
			}
		}

	}

	// query param serviceDomain
	qrServiceDomain := o.ServiceDomain
	qServiceDomain := qrServiceDomain
	if qServiceDomain != "" {
		if err := r.SetQueryParam("serviceDomain", qServiceDomain); err != nil {
			return err
		}
	}

	if o.Step != nil {

		// query param step
		var qrStep int64
		if o.Step != nil {
			qrStep = *o.Step
		}
		qStep := swag.FormatInt64(qrStep)
		if qStep != "" {
			if err := r.SetQueryParam("step", qStep); err != nil {
				return err
			}
		}

	}

	if o.Version != nil {

		// query param version
		var qrVersion string
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	// path param workload
	if err := r.SetPathParam("workload", o.Workload); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}