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

// EdgeInfoListReader is a Reader for the EdgeInfoList structure.
type EdgeInfoListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeInfoListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeInfoListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEdgeInfoListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeInfoListOK creates a EdgeInfoListOK with default headers values
func NewEdgeInfoListOK() *EdgeInfoListOK {
	return &EdgeInfoListOK{}
}

/*EdgeInfoListOK handles this case with default header values.

Ok
*/
type EdgeInfoListOK struct {
	Payload []*models.EdgeUsageInfo
}

func (o *EdgeInfoListOK) Error() string {
	return fmt.Sprintf("[GET /v1/edgesInfo][%d] edgeInfoListOK  %+v", 200, o.Payload)
}

func (o *EdgeInfoListOK) GetPayload() []*models.EdgeUsageInfo {
	return o.Payload
}

func (o *EdgeInfoListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeInfoListDefault creates a EdgeInfoListDefault with default headers values
func NewEdgeInfoListDefault(code int) *EdgeInfoListDefault {
	return &EdgeInfoListDefault{
		_statusCode: code,
	}
}

/*EdgeInfoListDefault handles this case with default header values.

generic API error response
*/
type EdgeInfoListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the edge info list default response
func (o *EdgeInfoListDefault) Code() int {
	return o._statusCode
}

func (o *EdgeInfoListDefault) Error() string {
	return fmt.Sprintf("[GET /v1/edgesInfo][%d] EdgeInfoList default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeInfoListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *EdgeInfoListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}