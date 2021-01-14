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

// NodeDeleteReader is a Reader for the NodeDelete structure.
type NodeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodeDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeDeleteOK creates a NodeDeleteOK with default headers values
func NewNodeDeleteOK() *NodeDeleteOK {
	return &NodeDeleteOK{}
}

/*NodeDeleteOK handles this case with default header values.

Ok
*/
type NodeDeleteOK struct {
	Payload *models.DeleteDocumentResponseV2
}

func (o *NodeDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/nodes/{nodeId}][%d] nodeDeleteOK  %+v", 200, o.Payload)
}

func (o *NodeDeleteOK) GetPayload() *models.DeleteDocumentResponseV2 {
	return o.Payload
}

func (o *NodeDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeDeleteDefault creates a NodeDeleteDefault with default headers values
func NewNodeDeleteDefault(code int) *NodeDeleteDefault {
	return &NodeDeleteDefault{
		_statusCode: code,
	}
}

/*NodeDeleteDefault handles this case with default header values.

generic API error response
*/
type NodeDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the node delete default response
func (o *NodeDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NodeDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/nodes/{nodeId}][%d] NodeDelete default  %+v", o._statusCode, o.Payload)
}

func (o *NodeDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *NodeDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
