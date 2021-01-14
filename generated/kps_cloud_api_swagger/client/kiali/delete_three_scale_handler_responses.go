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

// DeleteThreeScaleHandlerReader is a Reader for the DeleteThreeScaleHandler structure.
type DeleteThreeScaleHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteThreeScaleHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteThreeScaleHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteThreeScaleHandlerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteThreeScaleHandlerOK creates a DeleteThreeScaleHandlerOK with default headers values
func NewDeleteThreeScaleHandlerOK() *DeleteThreeScaleHandlerOK {
	return &DeleteThreeScaleHandlerOK{}
}

/*DeleteThreeScaleHandlerOK handles this case with default header values.

List of ThreeScale handlers created from Kiali to be used in the adapter's configuration
*/
type DeleteThreeScaleHandlerOK struct {
	Payload models.ThreeScaleHandlers
}

func (o *DeleteThreeScaleHandlerOK) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/kiali/threescale/handlers/{threescaleHandlerName}][%d] deleteThreeScaleHandlerOK  %+v", 200, o.Payload)
}

func (o *DeleteThreeScaleHandlerOK) GetPayload() models.ThreeScaleHandlers {
	return o.Payload
}

func (o *DeleteThreeScaleHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteThreeScaleHandlerDefault creates a DeleteThreeScaleHandlerDefault with default headers values
func NewDeleteThreeScaleHandlerDefault(code int) *DeleteThreeScaleHandlerDefault {
	return &DeleteThreeScaleHandlerDefault{
		_statusCode: code,
	}
}

/*DeleteThreeScaleHandlerDefault handles this case with default header values.

generic API error response
*/
type DeleteThreeScaleHandlerDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the delete three scale handler default response
func (o *DeleteThreeScaleHandlerDefault) Code() int {
	return o._statusCode
}

func (o *DeleteThreeScaleHandlerDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/kiali/threescale/handlers/{threescaleHandlerName}][%d] deleteThreeScaleHandler default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteThreeScaleHandlerDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *DeleteThreeScaleHandlerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
