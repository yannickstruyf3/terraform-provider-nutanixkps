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

// ContainerRegistryDeleteReader is a Reader for the ContainerRegistryDelete structure.
type ContainerRegistryDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRegistryDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerRegistryDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContainerRegistryDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContainerRegistryDeleteOK creates a ContainerRegistryDeleteOK with default headers values
func NewContainerRegistryDeleteOK() *ContainerRegistryDeleteOK {
	return &ContainerRegistryDeleteOK{}
}

/*ContainerRegistryDeleteOK handles this case with default header values.

Ok
*/
type ContainerRegistryDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *ContainerRegistryDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/containerregistries/{id}][%d] containerRegistryDeleteOK  %+v", 200, o.Payload)
}

func (o *ContainerRegistryDeleteOK) GetPayload() *models.DeleteDocumentResponse {
	return o.Payload
}

func (o *ContainerRegistryDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRegistryDeleteDefault creates a ContainerRegistryDeleteDefault with default headers values
func NewContainerRegistryDeleteDefault(code int) *ContainerRegistryDeleteDefault {
	return &ContainerRegistryDeleteDefault{
		_statusCode: code,
	}
}

/*ContainerRegistryDeleteDefault handles this case with default header values.

generic API error response
*/
type ContainerRegistryDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the container registry delete default response
func (o *ContainerRegistryDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ContainerRegistryDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/containerregistries/{id}][%d] ContainerRegistryDelete default  %+v", o._statusCode, o.Payload)
}

func (o *ContainerRegistryDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ContainerRegistryDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
