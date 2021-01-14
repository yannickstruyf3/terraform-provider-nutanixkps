// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// LogUploadCompleteReader is a Reader for the LogUploadComplete structure.
type LogUploadCompleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogUploadCompleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewLogUploadCompleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewLogUploadCompleteDefault creates a LogUploadCompleteDefault with default headers values
func NewLogUploadCompleteDefault(code int) *LogUploadCompleteDefault {
	return &LogUploadCompleteDefault{
		_statusCode: code,
	}
}

/*LogUploadCompleteDefault handles this case with default header values.

generic API error response
*/
type LogUploadCompleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the log upload complete default response
func (o *LogUploadCompleteDefault) Code() int {
	return o._statusCode
}

func (o *LogUploadCompleteDefault) Error() string {
	return fmt.Sprintf("[POST /v1/logs/uploadComplete][%d] LogUploadComplete default  %+v", o._statusCode, o.Payload)
}

func (o *LogUploadCompleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *LogUploadCompleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}