// Code generated by go-swagger; DO NOT EDIT.

package software_update

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// SoftwareUpdateServiceDomainListReader is a Reader for the SoftwareUpdateServiceDomainList structure.
type SoftwareUpdateServiceDomainListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwareUpdateServiceDomainListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSoftwareUpdateServiceDomainListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwareUpdateServiceDomainListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwareUpdateServiceDomainListOK creates a SoftwareUpdateServiceDomainListOK with default headers values
func NewSoftwareUpdateServiceDomainListOK() *SoftwareUpdateServiceDomainListOK {
	return &SoftwareUpdateServiceDomainListOK{}
}

/*SoftwareUpdateServiceDomainListOK handles this case with default header values.

Ok
*/
type SoftwareUpdateServiceDomainListOK struct {
	Payload *models.SoftwareUpdateServiceDomainListPayload
}

func (o *SoftwareUpdateServiceDomainListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/softwareupdates/servicedomains][%d] softwareUpdateServiceDomainListOK  %+v", 200, o.Payload)
}

func (o *SoftwareUpdateServiceDomainListOK) GetPayload() *models.SoftwareUpdateServiceDomainListPayload {
	return o.Payload
}

func (o *SoftwareUpdateServiceDomainListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwareUpdateServiceDomainListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareUpdateServiceDomainListDefault creates a SoftwareUpdateServiceDomainListDefault with default headers values
func NewSoftwareUpdateServiceDomainListDefault(code int) *SoftwareUpdateServiceDomainListDefault {
	return &SoftwareUpdateServiceDomainListDefault{
		_statusCode: code,
	}
}

/*SoftwareUpdateServiceDomainListDefault handles this case with default header values.

generic API error response
*/
type SoftwareUpdateServiceDomainListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the software update service domain list default response
func (o *SoftwareUpdateServiceDomainListDefault) Code() int {
	return o._statusCode
}

func (o *SoftwareUpdateServiceDomainListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/softwareupdates/servicedomains][%d] SoftwareUpdateServiceDomainList default  %+v", o._statusCode, o.Payload)
}

func (o *SoftwareUpdateServiceDomainListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *SoftwareUpdateServiceDomainListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
