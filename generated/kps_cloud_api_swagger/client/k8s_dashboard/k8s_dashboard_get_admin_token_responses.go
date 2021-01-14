// Code generated by go-swagger; DO NOT EDIT.

package k8s_dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// K8sDashboardGetAdminTokenReader is a Reader for the K8sDashboardGetAdminToken structure.
type K8sDashboardGetAdminTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *K8sDashboardGetAdminTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewK8sDashboardGetAdminTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewK8sDashboardGetAdminTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewK8sDashboardGetAdminTokenOK creates a K8sDashboardGetAdminTokenOK with default headers values
func NewK8sDashboardGetAdminTokenOK() *K8sDashboardGetAdminTokenOK {
	return &K8sDashboardGetAdminTokenOK{}
}

/*K8sDashboardGetAdminTokenOK handles this case with default header values.

Ok
*/
type K8sDashboardGetAdminTokenOK struct {
	Payload *models.K8sDashboardTokenResponsePayload
}

func (o *K8sDashboardGetAdminTokenOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/k8sdashboard/{svcDomainId}/adminToken][%d] k8sDashboardGetAdminTokenOK  %+v", 200, o.Payload)
}

func (o *K8sDashboardGetAdminTokenOK) GetPayload() *models.K8sDashboardTokenResponsePayload {
	return o.Payload
}

func (o *K8sDashboardGetAdminTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8sDashboardTokenResponsePayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewK8sDashboardGetAdminTokenDefault creates a K8sDashboardGetAdminTokenDefault with default headers values
func NewK8sDashboardGetAdminTokenDefault(code int) *K8sDashboardGetAdminTokenDefault {
	return &K8sDashboardGetAdminTokenDefault{
		_statusCode: code,
	}
}

/*K8sDashboardGetAdminTokenDefault handles this case with default header values.

generic API error response
*/
type K8sDashboardGetAdminTokenDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the k8s dashboard get admin token default response
func (o *K8sDashboardGetAdminTokenDefault) Code() int {
	return o._statusCode
}

func (o *K8sDashboardGetAdminTokenDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/k8sdashboard/{svcDomainId}/adminToken][%d] K8sDashboardGetAdminToken default  %+v", o._statusCode, o.Payload)
}

func (o *K8sDashboardGetAdminTokenDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *K8sDashboardGetAdminTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}