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

// ServiceDomainUpdateReader is a Reader for the ServiceDomainUpdate structure.
type ServiceDomainUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceDomainUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceDomainUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceDomainUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceDomainUpdateOK creates a ServiceDomainUpdateOK with default headers values
func NewServiceDomainUpdateOK() *ServiceDomainUpdateOK {
	return &ServiceDomainUpdateOK{}
}

/*ServiceDomainUpdateOK handles this case with default header values.

Ok
*/
type ServiceDomainUpdateOK struct {
	Payload *models.UpdateDocumentResponseV2
}

func (o *ServiceDomainUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1.0/servicedomains/{svcDomainId}][%d] serviceDomainUpdateOK  %+v", 200, o.Payload)
}

func (o *ServiceDomainUpdateOK) GetPayload() *models.UpdateDocumentResponseV2 {
	return o.Payload
}

func (o *ServiceDomainUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceDomainUpdateDefault creates a ServiceDomainUpdateDefault with default headers values
func NewServiceDomainUpdateDefault(code int) *ServiceDomainUpdateDefault {
	return &ServiceDomainUpdateDefault{
		_statusCode: code,
	}
}

/*ServiceDomainUpdateDefault handles this case with default header values.

generic API error response
*/
type ServiceDomainUpdateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service domain update default response
func (o *ServiceDomainUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ServiceDomainUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1.0/servicedomains/{svcDomainId}][%d] ServiceDomainUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceDomainUpdateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceDomainUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}