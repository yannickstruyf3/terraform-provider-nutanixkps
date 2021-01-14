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

// ServiceDomainGetEffectiveProfileReader is a Reader for the ServiceDomainGetEffectiveProfile structure.
type ServiceDomainGetEffectiveProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceDomainGetEffectiveProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceDomainGetEffectiveProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceDomainGetEffectiveProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceDomainGetEffectiveProfileOK creates a ServiceDomainGetEffectiveProfileOK with default headers values
func NewServiceDomainGetEffectiveProfileOK() *ServiceDomainGetEffectiveProfileOK {
	return &ServiceDomainGetEffectiveProfileOK{}
}

/*ServiceDomainGetEffectiveProfileOK handles this case with default header values.

Ok
*/
type ServiceDomainGetEffectiveProfileOK struct {
	Payload *models.ServiceDomainProfile
}

func (o *ServiceDomainGetEffectiveProfileOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicedomains/{svcDomainId}/effectiveprofile][%d] serviceDomainGetEffectiveProfileOK  %+v", 200, o.Payload)
}

func (o *ServiceDomainGetEffectiveProfileOK) GetPayload() *models.ServiceDomainProfile {
	return o.Payload
}

func (o *ServiceDomainGetEffectiveProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDomainProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceDomainGetEffectiveProfileDefault creates a ServiceDomainGetEffectiveProfileDefault with default headers values
func NewServiceDomainGetEffectiveProfileDefault(code int) *ServiceDomainGetEffectiveProfileDefault {
	return &ServiceDomainGetEffectiveProfileDefault{
		_statusCode: code,
	}
}

/*ServiceDomainGetEffectiveProfileDefault handles this case with default header values.

generic API error response
*/
type ServiceDomainGetEffectiveProfileDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service domain get effective profile default response
func (o *ServiceDomainGetEffectiveProfileDefault) Code() int {
	return o._statusCode
}

func (o *ServiceDomainGetEffectiveProfileDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicedomains/{svcDomainId}/effectiveprofile][%d] ServiceDomainGetEffectiveProfile default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceDomainGetEffectiveProfileDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceDomainGetEffectiveProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
