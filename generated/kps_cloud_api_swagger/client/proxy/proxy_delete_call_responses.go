// Code generated by go-swagger; DO NOT EDIT.

package proxy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ProxyDeleteCallReader is a Reader for the ProxyDeleteCall structure.
type ProxyDeleteCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProxyDeleteCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProxyDeleteCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProxyDeleteCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProxyDeleteCallOK creates a ProxyDeleteCallOK with default headers values
func NewProxyDeleteCallOK() *ProxyDeleteCallOK {
	return &ProxyDeleteCallOK{}
}

/*ProxyDeleteCallOK handles this case with default header values.

Ok
*/
type ProxyDeleteCallOK struct {
	Payload models.ProxyResponsePayload
}

func (o *ProxyDeleteCallOK) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/proxy/*path][%d] proxyDeleteCallOK  %+v", 200, o.Payload)
}

func (o *ProxyDeleteCallOK) GetPayload() models.ProxyResponsePayload {
	return o.Payload
}

func (o *ProxyDeleteCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProxyDeleteCallDefault creates a ProxyDeleteCallDefault with default headers values
func NewProxyDeleteCallDefault(code int) *ProxyDeleteCallDefault {
	return &ProxyDeleteCallDefault{
		_statusCode: code,
	}
}

/*ProxyDeleteCallDefault handles this case with default header values.

generic API error response
*/
type ProxyDeleteCallDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the proxy delete call default response
func (o *ProxyDeleteCallDefault) Code() int {
	return o._statusCode
}

func (o *ProxyDeleteCallDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/proxy/*path][%d] ProxyDeleteCall default  %+v", o._statusCode, o.Payload)
}

func (o *ProxyDeleteCallDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProxyDeleteCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
