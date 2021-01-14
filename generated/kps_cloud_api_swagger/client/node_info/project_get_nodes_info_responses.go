// Code generated by go-swagger; DO NOT EDIT.

package node_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ProjectGetNodesInfoReader is a Reader for the ProjectGetNodesInfo structure.
type ProjectGetNodesInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetNodesInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGetNodesInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectGetNodesInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetNodesInfoOK creates a ProjectGetNodesInfoOK with default headers values
func NewProjectGetNodesInfoOK() *ProjectGetNodesInfoOK {
	return &ProjectGetNodesInfoOK{}
}

/*ProjectGetNodesInfoOK handles this case with default header values.

Ok
*/
type ProjectGetNodesInfoOK struct {
	Payload *models.NodeInfoListPayload
}

func (o *ProjectGetNodesInfoOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/nodesinfo][%d] projectGetNodesInfoOK  %+v", 200, o.Payload)
}

func (o *ProjectGetNodesInfoOK) GetPayload() *models.NodeInfoListPayload {
	return o.Payload
}

func (o *ProjectGetNodesInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeInfoListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetNodesInfoDefault creates a ProjectGetNodesInfoDefault with default headers values
func NewProjectGetNodesInfoDefault(code int) *ProjectGetNodesInfoDefault {
	return &ProjectGetNodesInfoDefault{
		_statusCode: code,
	}
}

/*ProjectGetNodesInfoDefault handles this case with default header values.

generic API error response
*/
type ProjectGetNodesInfoDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project get nodes info default response
func (o *ProjectGetNodesInfoDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetNodesInfoDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/nodesinfo][%d] ProjectGetNodesInfo default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectGetNodesInfoDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectGetNodesInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}