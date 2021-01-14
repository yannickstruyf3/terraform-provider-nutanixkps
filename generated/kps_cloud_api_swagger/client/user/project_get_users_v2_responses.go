// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ProjectGetUsersV2Reader is a Reader for the ProjectGetUsersV2 structure.
type ProjectGetUsersV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetUsersV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGetUsersV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectGetUsersV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetUsersV2OK creates a ProjectGetUsersV2OK with default headers values
func NewProjectGetUsersV2OK() *ProjectGetUsersV2OK {
	return &ProjectGetUsersV2OK{}
}

/*ProjectGetUsersV2OK handles this case with default header values.

Ok
*/
type ProjectGetUsersV2OK struct {
	Payload *models.UserListPayload
}

func (o *ProjectGetUsersV2OK) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/users][%d] projectGetUsersV2OK  %+v", 200, o.Payload)
}

func (o *ProjectGetUsersV2OK) GetPayload() *models.UserListPayload {
	return o.Payload
}

func (o *ProjectGetUsersV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetUsersV2Default creates a ProjectGetUsersV2Default with default headers values
func NewProjectGetUsersV2Default(code int) *ProjectGetUsersV2Default {
	return &ProjectGetUsersV2Default{
		_statusCode: code,
	}
}

/*ProjectGetUsersV2Default handles this case with default header values.

generic API error response
*/
type ProjectGetUsersV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project get users v2 default response
func (o *ProjectGetUsersV2Default) Code() int {
	return o._statusCode
}

func (o *ProjectGetUsersV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/users][%d] ProjectGetUsersV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectGetUsersV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectGetUsersV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}