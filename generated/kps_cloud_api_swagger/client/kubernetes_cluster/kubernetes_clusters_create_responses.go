// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// KubernetesClustersCreateReader is a Reader for the KubernetesClustersCreate structure.
type KubernetesClustersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesClustersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesClustersCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKubernetesClustersCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKubernetesClustersCreateOK creates a KubernetesClustersCreateOK with default headers values
func NewKubernetesClustersCreateOK() *KubernetesClustersCreateOK {
	return &KubernetesClustersCreateOK{}
}

/*KubernetesClustersCreateOK handles this case with default header values.

Ok
*/
type KubernetesClustersCreateOK struct {
	Payload *models.CreateDocumentResponseV2
}

func (o *KubernetesClustersCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/kubernetesclusters][%d] kubernetesClustersCreateOK  %+v", 200, o.Payload)
}

func (o *KubernetesClustersCreateOK) GetPayload() *models.CreateDocumentResponseV2 {
	return o.Payload
}

func (o *KubernetesClustersCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesClustersCreateDefault creates a KubernetesClustersCreateDefault with default headers values
func NewKubernetesClustersCreateDefault(code int) *KubernetesClustersCreateDefault {
	return &KubernetesClustersCreateDefault{
		_statusCode: code,
	}
}

/*KubernetesClustersCreateDefault handles this case with default header values.

generic API error response
*/
type KubernetesClustersCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the kubernetes clusters create default response
func (o *KubernetesClustersCreateDefault) Code() int {
	return o._statusCode
}

func (o *KubernetesClustersCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/kubernetesclusters][%d] KubernetesClustersCreate default  %+v", o._statusCode, o.Payload)
}

func (o *KubernetesClustersCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *KubernetesClustersCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}