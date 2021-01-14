// Code generated by go-swagger; DO NOT EDIT.

package service_class

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ServiceClassListReader is a Reader for the ServiceClassList structure.
type ServiceClassListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceClassListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceClassListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceClassListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceClassListOK creates a ServiceClassListOK with default headers values
func NewServiceClassListOK() *ServiceClassListOK {
	return &ServiceClassListOK{}
}

/*ServiceClassListOK handles this case with default header values.

Ok
*/
type ServiceClassListOK struct {
	Payload *models.ServiceClassListPayload
}

func (o *ServiceClassListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/serviceclasses][%d] serviceClassListOK  %+v", 200, o.Payload)
}

func (o *ServiceClassListOK) GetPayload() *models.ServiceClassListPayload {
	return o.Payload
}

func (o *ServiceClassListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceClassListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceClassListDefault creates a ServiceClassListDefault with default headers values
func NewServiceClassListDefault(code int) *ServiceClassListDefault {
	return &ServiceClassListDefault{
		_statusCode: code,
	}
}

/*ServiceClassListDefault handles this case with default header values.

generic API error response
*/
type ServiceClassListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service class list default response
func (o *ServiceClassListDefault) Code() int {
	return o._statusCode
}

func (o *ServiceClassListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/serviceclasses][%d] ServiceClassList default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceClassListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceClassListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
