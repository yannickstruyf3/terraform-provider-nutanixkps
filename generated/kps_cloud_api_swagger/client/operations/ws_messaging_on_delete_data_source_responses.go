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

// WsMessagingOnDeleteDataSourceReader is a Reader for the WsMessagingOnDeleteDataSource structure.
type WsMessagingOnDeleteDataSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WsMessagingOnDeleteDataSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWsMessagingOnDeleteDataSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWsMessagingOnDeleteDataSourceOK creates a WsMessagingOnDeleteDataSourceOK with default headers values
func NewWsMessagingOnDeleteDataSourceOK() *WsMessagingOnDeleteDataSourceOK {
	return &WsMessagingOnDeleteDataSourceOK{}
}

/*WsMessagingOnDeleteDataSourceOK handles this case with default header values.

Ok
*/
type WsMessagingOnDeleteDataSourceOK struct {
	Payload *models.ResponseBase
}

func (o *WsMessagingOnDeleteDataSourceOK) Error() string {
	return fmt.Sprintf("[POST /v1/wsdocs/onDeleteDataSource][%d] wsMessagingOnDeleteDataSourceOK  %+v", 200, o.Payload)
}

func (o *WsMessagingOnDeleteDataSourceOK) GetPayload() *models.ResponseBase {
	return o.Payload
}

func (o *WsMessagingOnDeleteDataSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseBase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}