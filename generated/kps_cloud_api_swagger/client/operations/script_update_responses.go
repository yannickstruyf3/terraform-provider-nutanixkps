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

// ScriptUpdateReader is a Reader for the ScriptUpdate structure.
type ScriptUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScriptUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScriptUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScriptUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScriptUpdateOK creates a ScriptUpdateOK with default headers values
func NewScriptUpdateOK() *ScriptUpdateOK {
	return &ScriptUpdateOK{}
}

/*ScriptUpdateOK handles this case with default header values.

Ok
*/
type ScriptUpdateOK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *ScriptUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/scripts][%d] scriptUpdateOK  %+v", 200, o.Payload)
}

func (o *ScriptUpdateOK) GetPayload() *models.UpdateDocumentResponse {
	return o.Payload
}

func (o *ScriptUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScriptUpdateDefault creates a ScriptUpdateDefault with default headers values
func NewScriptUpdateDefault(code int) *ScriptUpdateDefault {
	return &ScriptUpdateDefault{
		_statusCode: code,
	}
}

/*ScriptUpdateDefault handles this case with default header values.

generic API error response
*/
type ScriptUpdateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the script update default response
func (o *ScriptUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ScriptUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/scripts][%d] ScriptUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *ScriptUpdateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ScriptUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}