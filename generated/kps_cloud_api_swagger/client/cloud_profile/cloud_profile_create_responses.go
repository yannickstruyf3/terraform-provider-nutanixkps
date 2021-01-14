// Code generated by go-swagger; DO NOT EDIT.

package cloud_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// CloudProfileCreateReader is a Reader for the CloudProfileCreate structure.
type CloudProfileCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudProfileCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudProfileCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCloudProfileCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudProfileCreateOK creates a CloudProfileCreateOK with default headers values
func NewCloudProfileCreateOK() *CloudProfileCreateOK {
	return &CloudProfileCreateOK{}
}

/*CloudProfileCreateOK handles this case with default header values.

Ok
*/
type CloudProfileCreateOK struct {
	Payload *models.CreateDocumentResponseV2
}

func (o *CloudProfileCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/cloudprofiles][%d] cloudProfileCreateOK  %+v", 200, o.Payload)
}

func (o *CloudProfileCreateOK) GetPayload() *models.CreateDocumentResponseV2 {
	return o.Payload
}

func (o *CloudProfileCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudProfileCreateDefault creates a CloudProfileCreateDefault with default headers values
func NewCloudProfileCreateDefault(code int) *CloudProfileCreateDefault {
	return &CloudProfileCreateDefault{
		_statusCode: code,
	}
}

/*CloudProfileCreateDefault handles this case with default header values.

generic API error response
*/
type CloudProfileCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the cloud profile create default response
func (o *CloudProfileCreateDefault) Code() int {
	return o._statusCode
}

func (o *CloudProfileCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/cloudprofiles][%d] CloudProfileCreate default  %+v", o._statusCode, o.Payload)
}

func (o *CloudProfileCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *CloudProfileCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
