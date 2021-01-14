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

// SetupSSHTunnelingReader is a Reader for the SetupSSHTunneling structure.
type SetupSSHTunnelingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetupSSHTunnelingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetupSSHTunnelingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetupSSHTunnelingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetupSSHTunnelingOK creates a SetupSSHTunnelingOK with default headers values
func NewSetupSSHTunnelingOK() *SetupSSHTunnelingOK {
	return &SetupSSHTunnelingOK{}
}

/*SetupSSHTunnelingOK handles this case with default header values.

Ok
*/
type SetupSSHTunnelingOK struct {
	Payload *models.WstunPayload
}

func (o *SetupSSHTunnelingOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/setupsshtunneling][%d] setupSshTunnelingOK  %+v", 200, o.Payload)
}

func (o *SetupSSHTunnelingOK) GetPayload() *models.WstunPayload {
	return o.Payload
}

func (o *SetupSSHTunnelingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WstunPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetupSSHTunnelingDefault creates a SetupSSHTunnelingDefault with default headers values
func NewSetupSSHTunnelingDefault(code int) *SetupSSHTunnelingDefault {
	return &SetupSSHTunnelingDefault{
		_statusCode: code,
	}
}

/*SetupSSHTunnelingDefault handles this case with default header values.

generic API error response
*/
type SetupSSHTunnelingDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the setup SSH tunneling default response
func (o *SetupSSHTunnelingDefault) Code() int {
	return o._statusCode
}

func (o *SetupSSHTunnelingDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/setupsshtunneling][%d] SetupSSHTunneling default  %+v", o._statusCode, o.Payload)
}

func (o *SetupSSHTunnelingDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *SetupSSHTunnelingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}