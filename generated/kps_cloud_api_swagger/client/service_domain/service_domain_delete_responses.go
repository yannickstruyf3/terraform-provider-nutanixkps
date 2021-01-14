// Code generated by go-swagger; DO NOT EDIT.

package service_domain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ServiceDomainDeleteReader is a Reader for the ServiceDomainDelete structure.
type ServiceDomainDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceDomainDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceDomainDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceDomainDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceDomainDeleteOK creates a ServiceDomainDeleteOK with default headers values
func NewServiceDomainDeleteOK() *ServiceDomainDeleteOK {
	return &ServiceDomainDeleteOK{}
}

/*ServiceDomainDeleteOK handles this case with default header values.

Ok
*/
type ServiceDomainDeleteOK struct {
	Payload *models.DeleteDocumentResponseV2
}

func (o *ServiceDomainDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/servicedomains/{svcDomainId}][%d] serviceDomainDeleteOK  %+v", 200, o.Payload)
}

func (o *ServiceDomainDeleteOK) GetPayload() *models.DeleteDocumentResponseV2 {
	return o.Payload
}

func (o *ServiceDomainDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceDomainDeleteDefault creates a ServiceDomainDeleteDefault with default headers values
func NewServiceDomainDeleteDefault(code int) *ServiceDomainDeleteDefault {
	return &ServiceDomainDeleteDefault{
		_statusCode: code,
	}
}

/*ServiceDomainDeleteDefault handles this case with default header values.

generic API error response
*/
type ServiceDomainDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service domain delete default response
func (o *ServiceDomainDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ServiceDomainDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/servicedomains/{svcDomainId}][%d] ServiceDomainDelete default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceDomainDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceDomainDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}