// Code generated by go-swagger; DO NOT EDIT.

package log_collector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// LogCollectorStopReader is a Reader for the LogCollectorStop structure.
type LogCollectorStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogCollectorStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogCollectorStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogCollectorStopDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogCollectorStopOK creates a LogCollectorStopOK with default headers values
func NewLogCollectorStopOK() *LogCollectorStopOK {
	return &LogCollectorStopOK{}
}

/*LogCollectorStopOK handles this case with default header values.

Ok
*/
type LogCollectorStopOK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *LogCollectorStopOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/logs/collector/{id}/stop][%d] logCollectorStopOK  %+v", 200, o.Payload)
}

func (o *LogCollectorStopOK) GetPayload() *models.UpdateDocumentResponse {
	return o.Payload
}

func (o *LogCollectorStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogCollectorStopDefault creates a LogCollectorStopDefault with default headers values
func NewLogCollectorStopDefault(code int) *LogCollectorStopDefault {
	return &LogCollectorStopDefault{
		_statusCode: code,
	}
}

/*LogCollectorStopDefault handles this case with default header values.

generic API error response
*/
type LogCollectorStopDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the log collector stop default response
func (o *LogCollectorStopDefault) Code() int {
	return o._statusCode
}

func (o *LogCollectorStopDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/logs/collector/{id}/stop][%d] LogCollectorStop default  %+v", o._statusCode, o.Payload)
}

func (o *LogCollectorStopDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *LogCollectorStopDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}