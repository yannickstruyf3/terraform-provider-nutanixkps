// Code generated by go-swagger; DO NOT EDIT.

package auditlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// AuditLogGetV2Reader is a Reader for the AuditLogGetV2 structure.
type AuditLogGetV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuditLogGetV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuditLogGetV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuditLogGetV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuditLogGetV2OK creates a AuditLogGetV2OK with default headers values
func NewAuditLogGetV2OK() *AuditLogGetV2OK {
	return &AuditLogGetV2OK{}
}

/*AuditLogGetV2OK handles this case with default header values.

Ok
*/
type AuditLogGetV2OK struct {
	Payload []*models.AuditLog
}

func (o *AuditLogGetV2OK) Error() string {
	return fmt.Sprintf("[GET /v1.0/auditlogs/{id}][%d] auditLogGetV2OK  %+v", 200, o.Payload)
}

func (o *AuditLogGetV2OK) GetPayload() []*models.AuditLog {
	return o.Payload
}

func (o *AuditLogGetV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuditLogGetV2Default creates a AuditLogGetV2Default with default headers values
func NewAuditLogGetV2Default(code int) *AuditLogGetV2Default {
	return &AuditLogGetV2Default{
		_statusCode: code,
	}
}

/*AuditLogGetV2Default handles this case with default header values.

generic API error response
*/
type AuditLogGetV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the audit log get v2 default response
func (o *AuditLogGetV2Default) Code() int {
	return o._statusCode
}

func (o *AuditLogGetV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/auditlogs/{id}][%d] AuditLogGetV2 default  %+v", o._statusCode, o.Payload)
}

func (o *AuditLogGetV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *AuditLogGetV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
