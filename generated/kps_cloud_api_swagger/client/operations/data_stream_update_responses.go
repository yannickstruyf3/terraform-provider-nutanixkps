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

// DataStreamUpdateReader is a Reader for the DataStreamUpdate structure.
type DataStreamUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataStreamUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDataStreamUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDataStreamUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataStreamUpdateOK creates a DataStreamUpdateOK with default headers values
func NewDataStreamUpdateOK() *DataStreamUpdateOK {
	return &DataStreamUpdateOK{}
}

/*DataStreamUpdateOK handles this case with default header values.

Ok
*/
type DataStreamUpdateOK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *DataStreamUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/datastreams][%d] dataStreamUpdateOK  %+v", 200, o.Payload)
}

func (o *DataStreamUpdateOK) GetPayload() *models.UpdateDocumentResponse {
	return o.Payload
}

func (o *DataStreamUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataStreamUpdateDefault creates a DataStreamUpdateDefault with default headers values
func NewDataStreamUpdateDefault(code int) *DataStreamUpdateDefault {
	return &DataStreamUpdateDefault{
		_statusCode: code,
	}
}

/*DataStreamUpdateDefault handles this case with default header values.

generic API error response
*/
type DataStreamUpdateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the data stream update default response
func (o *DataStreamUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DataStreamUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/datastreams][%d] DataStreamUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *DataStreamUpdateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *DataStreamUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}