// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// NodeOnboardedReader is a Reader for the NodeOnboarded structure.
type NodeOnboardedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeOnboardedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeOnboardedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodeOnboardedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeOnboardedOK creates a NodeOnboardedOK with default headers values
func NewNodeOnboardedOK() *NodeOnboardedOK {
	return &NodeOnboardedOK{}
}

/*NodeOnboardedOK handles this case with default header values.

Ok
*/
type NodeOnboardedOK struct {
	Payload *models.UpdateDocumentResponseV2
}

func (o *NodeOnboardedOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/nodeonboarded][%d] nodeOnboardedOK  %+v", 200, o.Payload)
}

func (o *NodeOnboardedOK) GetPayload() *models.UpdateDocumentResponseV2 {
	return o.Payload
}

func (o *NodeOnboardedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeOnboardedDefault creates a NodeOnboardedDefault with default headers values
func NewNodeOnboardedDefault(code int) *NodeOnboardedDefault {
	return &NodeOnboardedDefault{
		_statusCode: code,
	}
}

/*NodeOnboardedDefault handles this case with default header values.

generic API error response
*/
type NodeOnboardedDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the node onboarded default response
func (o *NodeOnboardedDefault) Code() int {
	return o._statusCode
}

func (o *NodeOnboardedDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/nodeonboarded][%d] NodeOnboarded default  %+v", o._statusCode, o.Payload)
}

func (o *NodeOnboardedDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *NodeOnboardedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}