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

// GetThreeScaleInfoReader is a Reader for the GetThreeScaleInfo structure.
type GetThreeScaleInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetThreeScaleInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetThreeScaleInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetThreeScaleInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetThreeScaleInfoOK creates a GetThreeScaleInfoOK with default headers values
func NewGetThreeScaleInfoOK() *GetThreeScaleInfoOK {
	return &GetThreeScaleInfoOK{}
}

/*GetThreeScaleInfoOK handles this case with default header values.

Return if ThreeScale adapter is enabled in Istio and if user has permissions to write adapter's configuration
*/
type GetThreeScaleInfoOK struct {
	Payload *models.ThreeScaleInfo
}

func (o *GetThreeScaleInfoOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/threescale][%d] getThreeScaleInfoOK  %+v", 200, o.Payload)
}

func (o *GetThreeScaleInfoOK) GetPayload() *models.ThreeScaleInfo {
	return o.Payload
}

func (o *GetThreeScaleInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThreeScaleInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetThreeScaleInfoDefault creates a GetThreeScaleInfoDefault with default headers values
func NewGetThreeScaleInfoDefault(code int) *GetThreeScaleInfoDefault {
	return &GetThreeScaleInfoDefault{
		_statusCode: code,
	}
}

/*GetThreeScaleInfoDefault handles this case with default header values.

generic API error response
*/
type GetThreeScaleInfoDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the get three scale info default response
func (o *GetThreeScaleInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetThreeScaleInfoDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/threescale][%d] getThreeScaleInfo default  %+v", o._statusCode, o.Payload)
}

func (o *GetThreeScaleInfoDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *GetThreeScaleInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}