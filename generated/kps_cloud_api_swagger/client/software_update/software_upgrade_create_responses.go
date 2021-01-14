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

// SoftwareUpgradeCreateReader is a Reader for the SoftwareUpgradeCreate structure.
type SoftwareUpgradeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwareUpgradeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSoftwareUpgradeCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwareUpgradeCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwareUpgradeCreateOK creates a SoftwareUpgradeCreateOK with default headers values
func NewSoftwareUpgradeCreateOK() *SoftwareUpgradeCreateOK {
	return &SoftwareUpgradeCreateOK{}
}

/*SoftwareUpgradeCreateOK handles this case with default header values.

Ok
*/
type SoftwareUpgradeCreateOK struct {
	Payload *models.CreateDocumentResponseV2
}

func (o *SoftwareUpgradeCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/softwareupdates/upgrades][%d] softwareUpgradeCreateOK  %+v", 200, o.Payload)
}

func (o *SoftwareUpgradeCreateOK) GetPayload() *models.CreateDocumentResponseV2 {
	return o.Payload
}

func (o *SoftwareUpgradeCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareUpgradeCreateDefault creates a SoftwareUpgradeCreateDefault with default headers values
func NewSoftwareUpgradeCreateDefault(code int) *SoftwareUpgradeCreateDefault {
	return &SoftwareUpgradeCreateDefault{
		_statusCode: code,
	}
}

/*SoftwareUpgradeCreateDefault handles this case with default header values.

generic API error response
*/
type SoftwareUpgradeCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the software upgrade create default response
func (o *SoftwareUpgradeCreateDefault) Code() int {
	return o._statusCode
}

func (o *SoftwareUpgradeCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/softwareupdates/upgrades][%d] SoftwareUpgradeCreate default  %+v", o._statusCode, o.Payload)
}

func (o *SoftwareUpgradeCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *SoftwareUpgradeCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}