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

// UserAPITokenGetReader is a Reader for the UserAPITokenGet structure.
type UserAPITokenGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserAPITokenGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserAPITokenGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserAPITokenGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserAPITokenGetOK creates a UserAPITokenGetOK with default headers values
func NewUserAPITokenGetOK() *UserAPITokenGetOK {
	return &UserAPITokenGetOK{}
}

/*UserAPITokenGetOK handles this case with default header values.

Ok
*/
type UserAPITokenGetOK struct {
	Payload []*models.UserAPIToken
}

func (o *UserAPITokenGetOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/userapitokens][%d] userApiTokenGetOK  %+v", 200, o.Payload)
}

func (o *UserAPITokenGetOK) GetPayload() []*models.UserAPIToken {
	return o.Payload
}

func (o *UserAPITokenGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAPITokenGetDefault creates a UserAPITokenGetDefault with default headers values
func NewUserAPITokenGetDefault(code int) *UserAPITokenGetDefault {
	return &UserAPITokenGetDefault{
		_statusCode: code,
	}
}

/*UserAPITokenGetDefault handles this case with default header values.

generic API error response
*/
type UserAPITokenGetDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the user Api token get default response
func (o *UserAPITokenGetDefault) Code() int {
	return o._statusCode
}

func (o *UserAPITokenGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/userapitokens][%d] UserApiTokenGet default  %+v", o._statusCode, o.Payload)
}

func (o *UserAPITokenGetDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *UserAPITokenGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
