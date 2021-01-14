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

// LogUploadReader is a Reader for the LogUpload structure.
type LogUploadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogUploadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewLogUploadDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewLogUploadDefault creates a LogUploadDefault with default headers values
func NewLogUploadDefault(code int) *LogUploadDefault {
	return &LogUploadDefault{
		_statusCode: code,
	}
}

/*LogUploadDefault handles this case with default header values.

generic API error response
*/
type LogUploadDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the log upload default response
func (o *LogUploadDefault) Code() int {
	return o._statusCode
}

func (o *LogUploadDefault) Error() string {
	return fmt.Sprintf("[POST /v1/logs/upload][%d] LogUpload default  %+v", o._statusCode, o.Payload)
}

func (o *LogUploadDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *LogUploadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}