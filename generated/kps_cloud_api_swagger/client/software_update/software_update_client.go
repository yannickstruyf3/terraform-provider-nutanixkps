// Code generated by go-swagger; DO NOT EDIT.

package software_update

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new software update API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for software update API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SoftwareDownloadBatchGet gets the batch in software download phase ntnx ignore

Retrieves progress details about a download operation as specified by its download batch ID.
*/
func (a *Client) SoftwareDownloadBatchGet(params *SoftwareDownloadBatchGetParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareDownloadBatchGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareDownloadBatchGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareDownloadBatchGet",
		Method:             "GET",
		PathPattern:        "/v1.0/softwareupdates/downloads/{batchId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareDownloadBatchGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareDownloadBatchGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareDownloadBatchGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareDownloadBatchList gets progress details for all download operations ntnx ignore

Retrieves progress details about each download operation for all in-progress or completed downloads.
*/
func (a *Client) SoftwareDownloadBatchList(params *SoftwareDownloadBatchListParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareDownloadBatchListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareDownloadBatchListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareDownloadBatchList",
		Method:             "GET",
		PathPattern:        "/v1.0/softwareupdates/downloads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareDownloadBatchListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareDownloadBatchListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareDownloadBatchListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareDownloadCreate starts a software download on the selected list of service domains ntnx ignore

Starts a software download on the selected list of service domains.
*/
func (a *Client) SoftwareDownloadCreate(params *SoftwareDownloadCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareDownloadCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareDownloadCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareDownloadCreate",
		Method:             "POST",
		PathPattern:        "/v1.0/softwareupdates/downloads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareDownloadCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareDownloadCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareDownloadCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareDownloadServiceDomainList gets all the service domains in the software download batch ntnx ignore

Retrieves all the service domains in the software download batch.
*/
func (a *Client) SoftwareDownloadServiceDomainList(params *SoftwareDownloadServiceDomainListParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareDownloadServiceDomainListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareDownloadServiceDomainListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareDownloadServiceDomainList",
		Method:             "GET",
		PathPattern:        "/v1.0/softwareupdates/downloads/{batchId}/servicedomains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareDownloadServiceDomainListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareDownloadServiceDomainListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareDownloadServiceDomainListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareDownloadStateUpdate updates the state of an existing software download ntnx ignore

Updates the state of an existing software download operation. The only reportable states are DOWNLOADING, DOWNLOAD_FAILED, DOWNLOAD_CANCELLED and DOWNLOADED.
*/
func (a *Client) SoftwareDownloadStateUpdate(params *SoftwareDownloadStateUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareDownloadStateUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareDownloadStateUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareDownloadStateUpdate",
		Method:             "PUT",
		PathPattern:        "/v1.0/softwareupdates/downloads/{batchId}/states",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareDownloadStateUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareDownloadStateUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareDownloadStateUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareDownloadUpdate updates an existing software download batch ntnx ignore

Updates an existing software download batch.
*/
func (a *Client) SoftwareDownloadUpdate(params *SoftwareDownloadUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareDownloadUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareDownloadUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareDownloadUpdate",
		Method:             "PUT",
		PathPattern:        "/v1.0/softwareupdates/downloads/{batchId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareDownloadUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareDownloadUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareDownloadUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareDownloadedServiceDomainList gets all service domains that have downloaded available software releases ntnx ignore

Retrieves all the service domains with the release downloaded.
*/
func (a *Client) SoftwareDownloadedServiceDomainList(params *SoftwareDownloadedServiceDomainListParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareDownloadedServiceDomainListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareDownloadedServiceDomainListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareDownloadedServiceDomainList",
		Method:             "GET",
		PathPattern:        "/v1.0/softwareupdates/releases/{release}/downloaded-servicedomains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareDownloadedServiceDomainListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareDownloadedServiceDomainListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareDownloadedServiceDomainListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareReleaseList gets all the available software releases ntnx ignore

Retrieves all the software releases available.
*/
func (a *Client) SoftwareReleaseList(params *SoftwareReleaseListParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareReleaseListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareReleaseListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareReleaseList",
		Method:             "GET",
		PathPattern:        "/v1.0/softwareupdates/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareReleaseListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareReleaseListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareReleaseListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareUpdateCredentialsCreate gets credentials to download software update files ntnx ignore

Retrieves credentials to download software update files.
*/
func (a *Client) SoftwareUpdateCredentialsCreate(params *SoftwareUpdateCredentialsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareUpdateCredentialsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareUpdateCredentialsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareUpdateCredentialsCreate",
		Method:             "POST",
		PathPattern:        "/v1.0/softwareupdates/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareUpdateCredentialsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareUpdateCredentialsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareUpdateCredentialsCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareUpdateServiceDomainList gets all the service domains and batches ntnx ignore

Retrieves all the service domains and batches.
*/
func (a *Client) SoftwareUpdateServiceDomainList(params *SoftwareUpdateServiceDomainListParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareUpdateServiceDomainListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareUpdateServiceDomainListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareUpdateServiceDomainList",
		Method:             "GET",
		PathPattern:        "/v1.0/softwareupdates/servicedomains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareUpdateServiceDomainListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareUpdateServiceDomainListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareUpdateServiceDomainListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareUpgradeBatchGet gets the batch in software upgrade phase ntnx ignore

Retrieves the batch in software upgrade phase.
*/
func (a *Client) SoftwareUpgradeBatchGet(params *SoftwareUpgradeBatchGetParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareUpgradeBatchGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareUpgradeBatchGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareUpgradeBatchGet",
		Method:             "GET",
		PathPattern:        "/v1.0/softwareupdates/upgrades/{batchId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareUpgradeBatchGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareUpgradeBatchGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareUpgradeBatchGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareUpgradeBatchList gets all the batches in software upgrade phase ntnx ignore

Retrieves all the batches in software upgrade phase.
*/
func (a *Client) SoftwareUpgradeBatchList(params *SoftwareUpgradeBatchListParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareUpgradeBatchListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareUpgradeBatchListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareUpgradeBatchList",
		Method:             "GET",
		PathPattern:        "/v1.0/softwareupdates/upgrades",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareUpgradeBatchListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareUpgradeBatchListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareUpgradeBatchListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareUpgradeCreate starts a software upgrade on the selected list of service domains ntnx ignore

Starts a software upgrade on the selected list of service domains.
*/
func (a *Client) SoftwareUpgradeCreate(params *SoftwareUpgradeCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareUpgradeCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareUpgradeCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareUpgradeCreate",
		Method:             "POST",
		PathPattern:        "/v1.0/softwareupdates/upgrades",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareUpgradeCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareUpgradeCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareUpgradeCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareUpgradeServiceDomainList gets all the service domains in the software upgrade batch ntnx ignore

Retrieves all the service domains in the software upgrade batch.
*/
func (a *Client) SoftwareUpgradeServiceDomainList(params *SoftwareUpgradeServiceDomainListParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareUpgradeServiceDomainListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareUpgradeServiceDomainListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareUpgradeServiceDomainList",
		Method:             "GET",
		PathPattern:        "/v1.0/softwareupdates/upgrades/{batchId}/servicedomains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareUpgradeServiceDomainListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareUpgradeServiceDomainListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareUpgradeServiceDomainListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareUpgradeStateUpdate updates the state of an existing software upgrade ntnx ignore

Updates the state of an existing software upgrade. The only reportable states are UPDATING, UPDATE_FAILED and UPDATED.
*/
func (a *Client) SoftwareUpgradeStateUpdate(params *SoftwareUpgradeStateUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareUpgradeStateUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareUpgradeStateUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareUpgradeStateUpdate",
		Method:             "PUT",
		PathPattern:        "/v1.0/softwareupdates/upgrades/{batchId}/states",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareUpgradeStateUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareUpgradeStateUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareUpgradeStateUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SoftwareUpgradeUpdate updates the state of an existing software upgrade ntnx ignore

Updates the state of an existing software upgrade.
*/
func (a *Client) SoftwareUpgradeUpdate(params *SoftwareUpgradeUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SoftwareUpgradeUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSoftwareUpgradeUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SoftwareUpgradeUpdate",
		Method:             "PUT",
		PathPattern:        "/v1.0/softwareupdates/upgrades/{batchId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SoftwareUpgradeUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SoftwareUpgradeUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SoftwareUpgradeUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}