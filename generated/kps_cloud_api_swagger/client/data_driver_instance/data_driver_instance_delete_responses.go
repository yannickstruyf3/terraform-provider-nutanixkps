// Code generated by go-swagger; DO NOT EDIT.

package data_driver_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// DataDriverInstanceDeleteReader is a Reader for the DataDriverInstanceDelete structure.
type DataDriverInstanceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataDriverInstanceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDataDriverInstanceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDataDriverInstanceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataDriverInstanceDeleteOK creates a DataDriverInstanceDeleteOK with default headers values
func NewDataDriverInstanceDeleteOK() *DataDriverInstanceDeleteOK {
	return &DataDriverInstanceDeleteOK{}
}

/*DataDriverInstanceDeleteOK handles this case with default header values.

Ok
*/
type DataDriverInstanceDeleteOK struct {
	Payload *models.DeleteDocumentResponseV2
}

func (o *DataDriverInstanceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/datadriverinstances/{id}][%d] dataDriverInstanceDeleteOK  %+v", 200, o.Payload)
}

func (o *DataDriverInstanceDeleteOK) GetPayload() *models.DeleteDocumentResponseV2 {
	return o.Payload
}

func (o *DataDriverInstanceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataDriverInstanceDeleteDefault creates a DataDriverInstanceDeleteDefault with default headers values
func NewDataDriverInstanceDeleteDefault(code int) *DataDriverInstanceDeleteDefault {
	return &DataDriverInstanceDeleteDefault{
		_statusCode: code,
	}
}

/*DataDriverInstanceDeleteDefault handles this case with default header values.

generic API error response
*/
type DataDriverInstanceDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the data driver instance delete default response
func (o *DataDriverInstanceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DataDriverInstanceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/datadriverinstances/{id}][%d] DataDriverInstanceDelete default  %+v", o._statusCode, o.Payload)
}

func (o *DataDriverInstanceDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *DataDriverInstanceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}