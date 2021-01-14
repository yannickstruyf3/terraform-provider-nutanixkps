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

// UserUpdateV2Reader is a Reader for the UserUpdateV2 structure.
type UserUpdateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserUpdateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserUpdateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserUpdateV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserUpdateV2OK creates a UserUpdateV2OK with default headers values
func NewUserUpdateV2OK() *UserUpdateV2OK {
	return &UserUpdateV2OK{}
}

/*UserUpdateV2OK handles this case with default header values.

Ok
*/
type UserUpdateV2OK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *UserUpdateV2OK) Error() string {
	return fmt.Sprintf("[PUT /v1/users/{id}][%d] userUpdateV2OK  %+v", 200, o.Payload)
}

func (o *UserUpdateV2OK) GetPayload() *models.UpdateDocumentResponse {
	return o.Payload
}

func (o *UserUpdateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserUpdateV2Default creates a UserUpdateV2Default with default headers values
func NewUserUpdateV2Default(code int) *UserUpdateV2Default {
	return &UserUpdateV2Default{
		_statusCode: code,
	}
}

/*UserUpdateV2Default handles this case with default header values.

generic API error response
*/
type UserUpdateV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the user update v2 default response
func (o *UserUpdateV2Default) Code() int {
	return o._statusCode
}

func (o *UserUpdateV2Default) Error() string {
	return fmt.Sprintf("[PUT /v1/users/{id}][%d] UserUpdateV2 default  %+v", o._statusCode, o.Payload)
}

func (o *UserUpdateV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *UserUpdateV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}