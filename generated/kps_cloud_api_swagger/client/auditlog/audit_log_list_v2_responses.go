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

// AuditLogListV2Reader is a Reader for the AuditLogListV2 structure.
type AuditLogListV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuditLogListV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuditLogListV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuditLogListV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuditLogListV2OK creates a AuditLogListV2OK with default headers values
func NewAuditLogListV2OK() *AuditLogListV2OK {
	return &AuditLogListV2OK{}
}

/*AuditLogListV2OK handles this case with default header values.

Ok
*/
type AuditLogListV2OK struct {
	Payload *models.AuditLogListResponsePayload
}

func (o *AuditLogListV2OK) Error() string {
	return fmt.Sprintf("[GET /v1.0/auditlogs][%d] auditLogListV2OK  %+v", 200, o.Payload)
}

func (o *AuditLogListV2OK) GetPayload() *models.AuditLogListResponsePayload {
	return o.Payload
}

func (o *AuditLogListV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditLogListResponsePayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuditLogListV2Default creates a AuditLogListV2Default with default headers values
func NewAuditLogListV2Default(code int) *AuditLogListV2Default {
	return &AuditLogListV2Default{
		_statusCode: code,
	}
}

/*AuditLogListV2Default handles this case with default header values.

generic API error response
*/
type AuditLogListV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the audit log list v2 default response
func (o *AuditLogListV2Default) Code() int {
	return o._statusCode
}

func (o *AuditLogListV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/auditlogs][%d] AuditLogListV2 default  %+v", o._statusCode, o.Payload)
}

func (o *AuditLogListV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *AuditLogListV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
