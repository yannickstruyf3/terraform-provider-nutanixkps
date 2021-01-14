// Code generated by go-swagger; DO NOT EDIT.

package software_update

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// SoftwareDownloadBatchListReader is a Reader for the SoftwareDownloadBatchList structure.
type SoftwareDownloadBatchListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwareDownloadBatchListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSoftwareDownloadBatchListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwareDownloadBatchListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwareDownloadBatchListOK creates a SoftwareDownloadBatchListOK with default headers values
func NewSoftwareDownloadBatchListOK() *SoftwareDownloadBatchListOK {
	return &SoftwareDownloadBatchListOK{}
}

/*SoftwareDownloadBatchListOK handles this case with default header values.

Ok
*/
type SoftwareDownloadBatchListOK struct {
	Payload *models.SoftwareUpdateBatchListPayload
}

func (o *SoftwareDownloadBatchListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/softwareupdates/downloads][%d] softwareDownloadBatchListOK  %+v", 200, o.Payload)
}

func (o *SoftwareDownloadBatchListOK) GetPayload() *models.SoftwareUpdateBatchListPayload {
	return o.Payload
}

func (o *SoftwareDownloadBatchListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwareUpdateBatchListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareDownloadBatchListDefault creates a SoftwareDownloadBatchListDefault with default headers values
func NewSoftwareDownloadBatchListDefault(code int) *SoftwareDownloadBatchListDefault {
	return &SoftwareDownloadBatchListDefault{
		_statusCode: code,
	}
}

/*SoftwareDownloadBatchListDefault handles this case with default header values.

generic API error response
*/
type SoftwareDownloadBatchListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the software download batch list default response
func (o *SoftwareDownloadBatchListDefault) Code() int {
	return o._statusCode
}

func (o *SoftwareDownloadBatchListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/softwareupdates/downloads][%d] SoftwareDownloadBatchList default  %+v", o._statusCode, o.Payload)
}

func (o *SoftwareDownloadBatchListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *SoftwareDownloadBatchListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}