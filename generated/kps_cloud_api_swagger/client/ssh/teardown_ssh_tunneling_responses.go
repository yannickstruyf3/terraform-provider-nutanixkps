// Code generated by go-swagger; DO NOT EDIT.

package ssh

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// TeardownSSHTunnelingReader is a Reader for the TeardownSSHTunneling structure.
type TeardownSSHTunnelingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeardownSSHTunnelingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeardownSSHTunnelingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTeardownSSHTunnelingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTeardownSSHTunnelingOK creates a TeardownSSHTunnelingOK with default headers values
func NewTeardownSSHTunnelingOK() *TeardownSSHTunnelingOK {
	return &TeardownSSHTunnelingOK{}
}

/*TeardownSSHTunnelingOK handles this case with default header values.

Ok
*/
type TeardownSSHTunnelingOK struct {
}

func (o *TeardownSSHTunnelingOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/teardownsshtunneling][%d] teardownSshTunnelingOK ", 200)
}

func (o *TeardownSSHTunnelingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTeardownSSHTunnelingDefault creates a TeardownSSHTunnelingDefault with default headers values
func NewTeardownSSHTunnelingDefault(code int) *TeardownSSHTunnelingDefault {
	return &TeardownSSHTunnelingDefault{
		_statusCode: code,
	}
}

/*TeardownSSHTunnelingDefault handles this case with default header values.

generic API error response
*/
type TeardownSSHTunnelingDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the teardown SSH tunneling default response
func (o *TeardownSSHTunnelingDefault) Code() int {
	return o._statusCode
}

func (o *TeardownSSHTunnelingDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/teardownsshtunneling][%d] TeardownSSHTunneling default  %+v", o._statusCode, o.Payload)
}

func (o *TeardownSSHTunnelingDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *TeardownSSHTunnelingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}