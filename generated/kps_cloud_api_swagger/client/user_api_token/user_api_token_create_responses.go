// Code generated by go-swagger; DO NOT EDIT.

package user_api_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// UserAPITokenCreateReader is a Reader for the UserAPITokenCreate structure.
type UserAPITokenCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserAPITokenCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserAPITokenCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserAPITokenCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserAPITokenCreateOK creates a UserAPITokenCreateOK with default headers values
func NewUserAPITokenCreateOK() *UserAPITokenCreateOK {
	return &UserAPITokenCreateOK{}
}

/*UserAPITokenCreateOK handles this case with default header values.

Ok
*/
type UserAPITokenCreateOK struct {
	Payload *models.UserAPITokenCreated
}

func (o *UserAPITokenCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/userapitokens][%d] userApiTokenCreateOK  %+v", 200, o.Payload)
}

func (o *UserAPITokenCreateOK) GetPayload() *models.UserAPITokenCreated {
	return o.Payload
}

func (o *UserAPITokenCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAPITokenCreated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAPITokenCreateDefault creates a UserAPITokenCreateDefault with default headers values
func NewUserAPITokenCreateDefault(code int) *UserAPITokenCreateDefault {
	return &UserAPITokenCreateDefault{
		_statusCode: code,
	}
}

/*UserAPITokenCreateDefault handles this case with default header values.

generic API error response
*/
type UserAPITokenCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the user Api token create default response
func (o *UserAPITokenCreateDefault) Code() int {
	return o._statusCode
}

func (o *UserAPITokenCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/userapitokens][%d] UserApiTokenCreate default  %+v", o._statusCode, o.Payload)
}

func (o *UserAPITokenCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *UserAPITokenCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
