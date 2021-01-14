// Code generated by go-swagger; DO NOT EDIT.

package kiali

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// WorkloadMetricsReader is a Reader for the WorkloadMetrics structure.
type WorkloadMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkloadMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkloadMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWorkloadMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWorkloadMetricsOK creates a WorkloadMetricsOK with default headers values
func NewWorkloadMetricsOK() *WorkloadMetricsOK {
	return &WorkloadMetricsOK{}
}

/*WorkloadMetricsOK handles this case with default header values.

Metrics response model
*/
type WorkloadMetricsOK struct {
	Payload *models.KialiMetrics
}

func (o *WorkloadMetricsOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/workloads/{workload}/metrics][%d] workloadMetricsOK  %+v", 200, o.Payload)
}

func (o *WorkloadMetricsOK) GetPayload() *models.KialiMetrics {
	return o.Payload
}

func (o *WorkloadMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KialiMetrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkloadMetricsDefault creates a WorkloadMetricsDefault with default headers values
func NewWorkloadMetricsDefault(code int) *WorkloadMetricsDefault {
	return &WorkloadMetricsDefault{
		_statusCode: code,
	}
}

/*WorkloadMetricsDefault handles this case with default header values.

generic API error response
*/
type WorkloadMetricsDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the workload metrics default response
func (o *WorkloadMetricsDefault) Code() int {
	return o._statusCode
}

func (o *WorkloadMetricsDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/workloads/{workload}/metrics][%d] workloadMetrics default  %+v", o._statusCode, o.Payload)
}

func (o *WorkloadMetricsDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *WorkloadMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}