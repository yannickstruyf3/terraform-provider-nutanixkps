// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// UserDeleteReader is a Reader for the UserDelete structure.
type UserDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserDeleteOK creates a UserDeleteOK with default headers values
func NewUserDeleteOK() *UserDeleteOK {
	return &UserDeleteOK{}
}

/*UserDeleteOK handles this case with default header values.

Ok
*/
type UserDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *UserDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/{id}][%d] userDeleteOK  %+v", 200, o.Payload)
}

func (o *UserDeleteOK) GetPayload() *models.DeleteDocumentResponse {
	return o.Payload
}

func (o *UserDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserDeleteDefault creates a UserDeleteDefault with default headers values
func NewUserDeleteDefault(code int) *UserDeleteDefault {
	return &UserDeleteDefault{
		_statusCode: code,
	}
}

/*UserDeleteDefault handles this case with default header values.

generic API error response
*/
type UserDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the user delete default response
func (o *UserDeleteDefault) Code() int {
	return o._statusCode
}

func (o *UserDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/{id}][%d] UserDelete default  %+v", o._statusCode, o.Payload)
}

func (o *UserDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *UserDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
