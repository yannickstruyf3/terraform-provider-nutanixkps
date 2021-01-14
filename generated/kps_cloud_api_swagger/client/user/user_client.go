// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
IsEmailAvailable checks if the given email is available for create user ntnx ignore

Checks if the given email is available for create user.
*/
func (a *Client) IsEmailAvailable(params *IsEmailAvailableParams, authInfo runtime.ClientAuthInfoWriter) (*IsEmailAvailableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsEmailAvailableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "IsEmailAvailable",
		Method:             "GET",
		PathPattern:        "/v1.0/isemailavailable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IsEmailAvailableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IsEmailAvailableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IsEmailAvailableDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProjectGetUsersV2 gets project users ntnx ignore

Retrievesall users for a project  by project ID {projectId}.
*/
func (a *Client) ProjectGetUsersV2(params *ProjectGetUsersV2Params, authInfo runtime.ClientAuthInfoWriter) (*ProjectGetUsersV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectGetUsersV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectGetUsersV2",
		Method:             "GET",
		PathPattern:        "/v1.0/projects/{projectId}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProjectGetUsersV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectGetUsersV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectGetUsersV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UserCreateV2 creates user ntnx ignore

Creates a user.
*/
func (a *Client) UserCreateV2(params *UserCreateV2Params, authInfo runtime.ClientAuthInfoWriter) (*UserCreateV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCreateV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserCreateV2",
		Method:             "POST",
		PathPattern:        "/v1.0/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserCreateV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserCreateV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UserCreateV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UserDeleteV2 deletes user by ID ntnx ignore

Deletes the user with the given ID {id}.
*/
func (a *Client) UserDeleteV2(params *UserDeleteV2Params, authInfo runtime.ClientAuthInfoWriter) (*UserDeleteV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserDeleteV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserDeleteV2",
		Method:             "DELETE",
		PathPattern:        "/v1.0/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserDeleteV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserDeleteV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UserDeleteV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UserGetV2 gets user by ID ntnx ignore

Retrieves a user with the given id {id}.
*/
func (a *Client) UserGetV2(params *UserGetV2Params, authInfo runtime.ClientAuthInfoWriter) (*UserGetV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGetV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserGetV2",
		Method:             "GET",
		PathPattern:        "/v1.0/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserGetV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserGetV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UserGetV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UserListV2 gets users ntnx ignore

Retrieves all users for a tenant.
*/
func (a *Client) UserListV2(params *UserListV2Params, authInfo runtime.ClientAuthInfoWriter) (*UserListV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserListV2",
		Method:             "GET",
		PathPattern:        "/v1.0/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserListV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserListV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UserListV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UserUpdateV3 updates user with a given ID ntnx ignore

Updates a user with a given ID {id}.
*/
func (a *Client) UserUpdateV3(params *UserUpdateV3Params, authInfo runtime.ClientAuthInfoWriter) (*UserUpdateV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserUpdateV3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserUpdateV3",
		Method:             "PUT",
		PathPattern:        "/v1.0/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserUpdateV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserUpdateV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UserUpdateV3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}