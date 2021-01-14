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

// ProjectGetScriptRuntimesReader is a Reader for the ProjectGetScriptRuntimes structure.
type ProjectGetScriptRuntimesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetScriptRuntimesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGetScriptRuntimesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectGetScriptRuntimesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetScriptRuntimesOK creates a ProjectGetScriptRuntimesOK with default headers values
func NewProjectGetScriptRuntimesOK() *ProjectGetScriptRuntimesOK {
	return &ProjectGetScriptRuntimesOK{}
}

/*ProjectGetScriptRuntimesOK handles this case with default header values.

Ok
*/
type ProjectGetScriptRuntimesOK struct {
	Payload []*models.ScriptRuntime
}

func (o *ProjectGetScriptRuntimesOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/scriptruntimes][%d] projectGetScriptRuntimesOK  %+v", 200, o.Payload)
}

func (o *ProjectGetScriptRuntimesOK) GetPayload() []*models.ScriptRuntime {
	return o.Payload
}

func (o *ProjectGetScriptRuntimesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetScriptRuntimesDefault creates a ProjectGetScriptRuntimesDefault with default headers values
func NewProjectGetScriptRuntimesDefault(code int) *ProjectGetScriptRuntimesDefault {
	return &ProjectGetScriptRuntimesDefault{
		_statusCode: code,
	}
}

/*ProjectGetScriptRuntimesDefault handles this case with default header values.

generic API error response
*/
type ProjectGetScriptRuntimesDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project get script runtimes default response
func (o *ProjectGetScriptRuntimesDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetScriptRuntimesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/scriptruntimes][%d] ProjectGetScriptRuntimes default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectGetScriptRuntimesDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectGetScriptRuntimesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
