// Code generated by go-swagger; DO NOT EDIT.

package data_driver_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// DataDriverConfigListReader is a Reader for the DataDriverConfigList structure.
type DataDriverConfigListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataDriverConfigListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDataDriverConfigListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDataDriverConfigListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataDriverConfigListOK creates a DataDriverConfigListOK with default headers values
func NewDataDriverConfigListOK() *DataDriverConfigListOK {
	return &DataDriverConfigListOK{}
}

/*DataDriverConfigListOK handles this case with default header values.

DataDriverConfigListResponse is a a data driver config listing response
*/
type DataDriverConfigListOK struct {
	Payload *models.DataDriverConfigListResponsePayload
}

func (o *DataDriverConfigListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/datadriverconfigs][%d] dataDriverConfigListOK  %+v", 200, o.Payload)
}

func (o *DataDriverConfigListOK) GetPayload() *models.DataDriverConfigListResponsePayload {
	return o.Payload
}

func (o *DataDriverConfigListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataDriverConfigListResponsePayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataDriverConfigListDefault creates a DataDriverConfigListDefault with default headers values
func NewDataDriverConfigListDefault(code int) *DataDriverConfigListDefault {
	return &DataDriverConfigListDefault{
		_statusCode: code,
	}
}

/*DataDriverConfigListDefault handles this case with default header values.

generic API error response
*/
type DataDriverConfigListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the data driver config list default response
func (o *DataDriverConfigListDefault) Code() int {
	return o._statusCode
}

func (o *DataDriverConfigListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/datadriverconfigs][%d] DataDriverConfigList default  %+v", o._statusCode, o.Payload)
}

func (o *DataDriverConfigListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *DataDriverConfigListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}