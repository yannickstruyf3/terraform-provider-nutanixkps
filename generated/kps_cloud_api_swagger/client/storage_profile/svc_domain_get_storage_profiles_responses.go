// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// SvcDomainGetStorageProfilesReader is a Reader for the SvcDomainGetStorageProfiles structure.
type SvcDomainGetStorageProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvcDomainGetStorageProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvcDomainGetStorageProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvcDomainGetStorageProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvcDomainGetStorageProfilesOK creates a SvcDomainGetStorageProfilesOK with default headers values
func NewSvcDomainGetStorageProfilesOK() *SvcDomainGetStorageProfilesOK {
	return &SvcDomainGetStorageProfilesOK{}
}

/*SvcDomainGetStorageProfilesOK handles this case with default header values.

Ok
*/
type SvcDomainGetStorageProfilesOK struct {
	Payload *models.StorageProfileListResponsePayload
}

func (o *SvcDomainGetStorageProfilesOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicedomains/{svcDomainId}/storageprofiles][%d] svcDomainGetStorageProfilesOK  %+v", 200, o.Payload)
}

func (o *SvcDomainGetStorageProfilesOK) GetPayload() *models.StorageProfileListResponsePayload {
	return o.Payload
}

func (o *SvcDomainGetStorageProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageProfileListResponsePayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvcDomainGetStorageProfilesDefault creates a SvcDomainGetStorageProfilesDefault with default headers values
func NewSvcDomainGetStorageProfilesDefault(code int) *SvcDomainGetStorageProfilesDefault {
	return &SvcDomainGetStorageProfilesDefault{
		_statusCode: code,
	}
}

/*SvcDomainGetStorageProfilesDefault handles this case with default header values.

generic API error response
*/
type SvcDomainGetStorageProfilesDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the svc domain get storage profiles default response
func (o *SvcDomainGetStorageProfilesDefault) Code() int {
	return o._statusCode
}

func (o *SvcDomainGetStorageProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicedomains/{svcDomainId}/storageprofiles][%d] SvcDomainGetStorageProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *SvcDomainGetStorageProfilesDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *SvcDomainGetStorageProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}