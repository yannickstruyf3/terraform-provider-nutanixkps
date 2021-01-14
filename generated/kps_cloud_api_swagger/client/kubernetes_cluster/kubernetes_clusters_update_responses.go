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

// KubernetesClustersUpdateReader is a Reader for the KubernetesClustersUpdate structure.
type KubernetesClustersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesClustersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesClustersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKubernetesClustersUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKubernetesClustersUpdateOK creates a KubernetesClustersUpdateOK with default headers values
func NewKubernetesClustersUpdateOK() *KubernetesClustersUpdateOK {
	return &KubernetesClustersUpdateOK{}
}

/*KubernetesClustersUpdateOK handles this case with default header values.

Ok
*/
type KubernetesClustersUpdateOK struct {
	Payload *models.UpdateDocumentResponseV2
}

func (o *KubernetesClustersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1.0/kubernetesclusters/{id}][%d] kubernetesClustersUpdateOK  %+v", 200, o.Payload)
}

func (o *KubernetesClustersUpdateOK) GetPayload() *models.UpdateDocumentResponseV2 {
	return o.Payload
}

func (o *KubernetesClustersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesClustersUpdateDefault creates a KubernetesClustersUpdateDefault with default headers values
func NewKubernetesClustersUpdateDefault(code int) *KubernetesClustersUpdateDefault {
	return &KubernetesClustersUpdateDefault{
		_statusCode: code,
	}
}

/*KubernetesClustersUpdateDefault handles this case with default header values.

generic API error response
*/
type KubernetesClustersUpdateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the kubernetes clusters update default response
func (o *KubernetesClustersUpdateDefault) Code() int {
	return o._statusCode
}

func (o *KubernetesClustersUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1.0/kubernetesclusters/{id}][%d] KubernetesClustersUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *KubernetesClustersUpdateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *KubernetesClustersUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}