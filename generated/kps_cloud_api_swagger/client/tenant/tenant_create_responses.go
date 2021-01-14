// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// TenantCreateReader is a Reader for the TenantCreate structure.
type TenantCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenantCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenantCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenantCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenantCreateOK creates a TenantCreateOK with default headers values
func NewTenantCreateOK() *TenantCreateOK {
	return &TenantCreateOK{}
}

/*TenantCreateOK handles this case with default header values.

Ok
*/
type TenantCreateOK struct {
	Payload *models.CreateDocumentResponseV2
}

func (o *TenantCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/tenants][%d] tenantCreateOK  %+v", 200, o.Payload)
}

func (o *TenantCreateOK) GetPayload() *models.CreateDocumentResponseV2 {
	return o.Payload
}

func (o *TenantCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantCreateDefault creates a TenantCreateDefault with default headers values
func NewTenantCreateDefault(code int) *TenantCreateDefault {
	return &TenantCreateDefault{
		_statusCode: code,
	}
}

/*TenantCreateDefault handles this case with default header values.

generic API error response
*/
type TenantCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the tenant create default response
func (o *TenantCreateDefault) Code() int {
	return o._statusCode
}

func (o *TenantCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/tenants][%d] TenantCreate default  %+v", o._statusCode, o.Payload)
}

func (o *TenantCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *TenantCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}