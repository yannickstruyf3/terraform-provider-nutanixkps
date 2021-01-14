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

// NodeListReader is a Reader for the NodeList structure.
type NodeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodeListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeListOK creates a NodeListOK with default headers values
func NewNodeListOK() *NodeListOK {
	return &NodeListOK{}
}

/*NodeListOK handles this case with default header values.

Ok
*/
type NodeListOK struct {
	Payload *models.NodeListPayload
}

func (o *NodeListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/nodes][%d] nodeListOK  %+v", 200, o.Payload)
}

func (o *NodeListOK) GetPayload() *models.NodeListPayload {
	return o.Payload
}

func (o *NodeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeListDefault creates a NodeListDefault with default headers values
func NewNodeListDefault(code int) *NodeListDefault {
	return &NodeListDefault{
		_statusCode: code,
	}
}

/*NodeListDefault handles this case with default header values.

generic API error response
*/
type NodeListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the node list default response
func (o *NodeListDefault) Code() int {
	return o._statusCode
}

func (o *NodeListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/nodes][%d] NodeList default  %+v", o._statusCode, o.Payload)
}

func (o *NodeListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *NodeListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}