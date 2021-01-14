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

// ProjectGetDockerProfilesReader is a Reader for the ProjectGetDockerProfiles structure.
type ProjectGetDockerProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetDockerProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGetDockerProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectGetDockerProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetDockerProfilesOK creates a ProjectGetDockerProfilesOK with default headers values
func NewProjectGetDockerProfilesOK() *ProjectGetDockerProfilesOK {
	return &ProjectGetDockerProfilesOK{}
}

/*ProjectGetDockerProfilesOK handles this case with default header values.

Ok
*/
type ProjectGetDockerProfilesOK struct {
	Payload []*models.DockerProfile
}

func (o *ProjectGetDockerProfilesOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/dockerprofiles][%d] projectGetDockerProfilesOK  %+v", 200, o.Payload)
}

func (o *ProjectGetDockerProfilesOK) GetPayload() []*models.DockerProfile {
	return o.Payload
}

func (o *ProjectGetDockerProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetDockerProfilesDefault creates a ProjectGetDockerProfilesDefault with default headers values
func NewProjectGetDockerProfilesDefault(code int) *ProjectGetDockerProfilesDefault {
	return &ProjectGetDockerProfilesDefault{
		_statusCode: code,
	}
}

/*ProjectGetDockerProfilesDefault handles this case with default header values.

generic API error response
*/
type ProjectGetDockerProfilesDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project get docker profiles default response
func (o *ProjectGetDockerProfilesDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetDockerProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/dockerprofiles][%d] ProjectGetDockerProfiles default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectGetDockerProfilesDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectGetDockerProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}