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

// InsertAuditLogV2Reader is a Reader for the InsertAuditLogV2 structure.
type InsertAuditLogV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InsertAuditLogV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInsertAuditLogV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInsertAuditLogV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInsertAuditLogV2OK creates a InsertAuditLogV2OK with default headers values
func NewInsertAuditLogV2OK() *InsertAuditLogV2OK {
	return &InsertAuditLogV2OK{}
}

/*InsertAuditLogV2OK handles this case with default header values.

Ok
*/
type InsertAuditLogV2OK struct {
	Payload []*models.AuditLogV2
}

func (o *InsertAuditLogV2OK) Error() string {
	return fmt.Sprintf("[PUT /v1.0/auditlogsV2][%d] insertAuditLogV2OK  %+v", 200, o.Payload)
}

func (o *InsertAuditLogV2OK) GetPayload() []*models.AuditLogV2 {
	return o.Payload
}

func (o *InsertAuditLogV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInsertAuditLogV2Default creates a InsertAuditLogV2Default with default headers values
func NewInsertAuditLogV2Default(code int) *InsertAuditLogV2Default {
	return &InsertAuditLogV2Default{
		_statusCode: code,
	}
}

/*InsertAuditLogV2Default handles this case with default header values.

generic API error response
*/
type InsertAuditLogV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the insert audit log v2 default response
func (o *InsertAuditLogV2Default) Code() int {
	return o._statusCode
}

func (o *InsertAuditLogV2Default) Error() string {
	return fmt.Sprintf("[PUT /v1.0/auditlogsV2][%d] InsertAuditLogV2 default  %+v", o._statusCode, o.Payload)
}

func (o *InsertAuditLogV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *InsertAuditLogV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}