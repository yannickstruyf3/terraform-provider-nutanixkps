// Code generated by go-swagger; DO NOT EDIT.

package service_binding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ServiceBindingListReader is a Reader for the ServiceBindingList structure.
type ServiceBindingListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBindingListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBindingListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceBindingListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceBindingListOK creates a ServiceBindingListOK with default headers values
func NewServiceBindingListOK() *ServiceBindingListOK {
	return &ServiceBindingListOK{}
}

/*ServiceBindingListOK handles this case with default header values.

Ok
*/
type ServiceBindingListOK struct {
	Payload *models.ServiceBindingListPayload
}

func (o *ServiceBindingListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicebindings][%d] serviceBindingListOK  %+v", 200, o.Payload)
}

func (o *ServiceBindingListOK) GetPayload() *models.ServiceBindingListPayload {
	return o.Payload
}

func (o *ServiceBindingListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceBindingListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingListDefault creates a ServiceBindingListDefault with default headers values
func NewServiceBindingListDefault(code int) *ServiceBindingListDefault {
	return &ServiceBindingListDefault{
		_statusCode: code,
	}
}

/*ServiceBindingListDefault handles this case with default header values.

generic API error response
*/
type ServiceBindingListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service binding list default response
func (o *ServiceBindingListDefault) Code() int {
	return o._statusCode
}

func (o *ServiceBindingListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicebindings][%d] ServiceBindingList default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceBindingListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceBindingListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
