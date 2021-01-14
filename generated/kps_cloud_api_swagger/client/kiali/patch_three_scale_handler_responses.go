// Code generated by go-swagger; DO NOT EDIT.

package kiali

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// PatchThreeScaleHandlerReader is a Reader for the PatchThreeScaleHandler structure.
type PatchThreeScaleHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchThreeScaleHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchThreeScaleHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPatchThreeScaleHandlerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchThreeScaleHandlerOK creates a PatchThreeScaleHandlerOK with default headers values
func NewPatchThreeScaleHandlerOK() *PatchThreeScaleHandlerOK {
	return &PatchThreeScaleHandlerOK{}
}

/*PatchThreeScaleHandlerOK handles this case with default header values.

List of ThreeScale handlers created from Kiali to be used in the adapter's configuration
*/
type PatchThreeScaleHandlerOK struct {
	Payload models.ThreeScaleHandlers
}

func (o *PatchThreeScaleHandlerOK) Error() string {
	return fmt.Sprintf("[PATCH /v1.0/kiali/threescale/handlers/{threescaleHandlerName}][%d] patchThreeScaleHandlerOK  %+v", 200, o.Payload)
}

func (o *PatchThreeScaleHandlerOK) GetPayload() models.ThreeScaleHandlers {
	return o.Payload
}

func (o *PatchThreeScaleHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchThreeScaleHandlerDefault creates a PatchThreeScaleHandlerDefault with default headers values
func NewPatchThreeScaleHandlerDefault(code int) *PatchThreeScaleHandlerDefault {
	return &PatchThreeScaleHandlerDefault{
		_statusCode: code,
	}
}

/*PatchThreeScaleHandlerDefault handles this case with default header values.

generic API error response
*/
type PatchThreeScaleHandlerDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the patch three scale handler default response
func (o *PatchThreeScaleHandlerDefault) Code() int {
	return o._statusCode
}

func (o *PatchThreeScaleHandlerDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1.0/kiali/threescale/handlers/{threescaleHandlerName}][%d] patchThreeScaleHandler default  %+v", o._statusCode, o.Payload)
}

func (o *PatchThreeScaleHandlerDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *PatchThreeScaleHandlerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}