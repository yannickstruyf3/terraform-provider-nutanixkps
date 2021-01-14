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

// ContainerRegistryUpdateReader is a Reader for the ContainerRegistryUpdate structure.
type ContainerRegistryUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRegistryUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerRegistryUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContainerRegistryUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContainerRegistryUpdateOK creates a ContainerRegistryUpdateOK with default headers values
func NewContainerRegistryUpdateOK() *ContainerRegistryUpdateOK {
	return &ContainerRegistryUpdateOK{}
}

/*ContainerRegistryUpdateOK handles this case with default header values.

Ok
*/
type ContainerRegistryUpdateOK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *ContainerRegistryUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/containerregistries/{id}][%d] containerRegistryUpdateOK  %+v", 200, o.Payload)
}

func (o *ContainerRegistryUpdateOK) GetPayload() *models.UpdateDocumentResponse {
	return o.Payload
}

func (o *ContainerRegistryUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRegistryUpdateDefault creates a ContainerRegistryUpdateDefault with default headers values
func NewContainerRegistryUpdateDefault(code int) *ContainerRegistryUpdateDefault {
	return &ContainerRegistryUpdateDefault{
		_statusCode: code,
	}
}

/*ContainerRegistryUpdateDefault handles this case with default header values.

generic API error response
*/
type ContainerRegistryUpdateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the container registry update default response
func (o *ContainerRegistryUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ContainerRegistryUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/containerregistries/{id}][%d] ContainerRegistryUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *ContainerRegistryUpdateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ContainerRegistryUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
