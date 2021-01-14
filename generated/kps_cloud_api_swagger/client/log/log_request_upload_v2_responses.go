// Code generated by go-swagger; DO NOT EDIT.

package log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// LogRequestUploadV2Reader is a Reader for the LogRequestUploadV2 structure.
type LogRequestUploadV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogRequestUploadV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogRequestUploadV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogRequestUploadV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogRequestUploadV2OK creates a LogRequestUploadV2OK with default headers values
func NewLogRequestUploadV2OK() *LogRequestUploadV2OK {
	return &LogRequestUploadV2OK{}
}

/*LogRequestUploadV2OK handles this case with default header values.

Ok
*/
type LogRequestUploadV2OK struct {
	Payload []*models.LogUploadPayload
}

func (o *LogRequestUploadV2OK) Error() string {
	return fmt.Sprintf("[POST /v1.0/logs/requestupload][%d] logRequestUploadV2OK  %+v", 200, o.Payload)
}

func (o *LogRequestUploadV2OK) GetPayload() []*models.LogUploadPayload {
	return o.Payload
}

func (o *LogRequestUploadV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogRequestUploadV2Default creates a LogRequestUploadV2Default with default headers values
func NewLogRequestUploadV2Default(code int) *LogRequestUploadV2Default {
	return &LogRequestUploadV2Default{
		_statusCode: code,
	}
}

/*LogRequestUploadV2Default handles this case with default header values.

generic API error response
*/
type LogRequestUploadV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the log request upload v2 default response
func (o *LogRequestUploadV2Default) Code() int {
	return o._statusCode
}

func (o *LogRequestUploadV2Default) Error() string {
	return fmt.Sprintf("[POST /v1.0/logs/requestupload][%d] LogRequestUploadV2 default  %+v", o._statusCode, o.Payload)
}

func (o *LogRequestUploadV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *LogRequestUploadV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}