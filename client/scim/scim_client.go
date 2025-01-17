// Code generated by go-swagger; DO NOT EDIT.

package scim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scim API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scim API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1ScimV2GroupsID(params *DeleteV1ScimV2GroupsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ScimV2GroupsIDNoContent, error)

	DeleteV1ScimV2UsersID(params *DeleteV1ScimV2UsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ScimV2UsersIDNoContent, error)

	GetV1ScimV2Groups(params *GetV1ScimV2GroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ScimV2GroupsOK, error)

	GetV1ScimV2GroupsID(params *GetV1ScimV2GroupsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ScimV2GroupsIDOK, error)

	GetV1ScimV2Users(params *GetV1ScimV2UsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ScimV2UsersOK, error)

	GetV1ScimV2UsersID(params *GetV1ScimV2UsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ScimV2UsersIDOK, error)

	PostV1ScimV2Groups(params *PostV1ScimV2GroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ScimV2GroupsCreated, error)

	PostV1ScimV2Users(params *PostV1ScimV2UsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ScimV2UsersCreated, error)

	PutV1ScimV2GroupsID(params *PutV1ScimV2GroupsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutV1ScimV2GroupsIDOK, error)

	PutV1ScimV2UsersID(params *PutV1ScimV2UsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutV1ScimV2UsersIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1ScimV2GroupsID deletes a team via the group protocol

SCIM endpoint to delete a Team (Colloquial for Group in the SCIM protocol).
*/
func (a *Client) DeleteV1ScimV2GroupsID(params *DeleteV1ScimV2GroupsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ScimV2GroupsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1ScimV2GroupsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1ScimV2GroupsId",
		Method:             "DELETE",
		PathPattern:        "/v1/scim/v2/Groups/{id}",
		ProducesMediaTypes: []string{"application/scim+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1ScimV2GroupsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteV1ScimV2GroupsIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1ScimV2GroupsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteV1ScimV2UsersID deletes a user using s c i m protocol

SCIM endpoint to delete a User. This endpoint will deactivate a confirmed OrganizationUser record in our system.
*/
func (a *Client) DeleteV1ScimV2UsersID(params *DeleteV1ScimV2UsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ScimV2UsersIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1ScimV2UsersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1ScimV2UsersId",
		Method:             "DELETE",
		PathPattern:        "/v1/scim/v2/Users/{id}",
		ProducesMediaTypes: []string{"application/scim+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1ScimV2UsersIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteV1ScimV2UsersIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1ScimV2UsersId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ScimV2Groups lists all teams via the group protocol

SCIM endpoint that lists all Teams (Colloquial for Group in the SCIM protocol)
*/
func (a *Client) GetV1ScimV2Groups(params *GetV1ScimV2GroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ScimV2GroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ScimV2GroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ScimV2Groups",
		Method:             "GET",
		PathPattern:        "/v1/scim/v2/Groups",
		ProducesMediaTypes: []string{"application/scim+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ScimV2GroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1ScimV2GroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ScimV2Groups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ScimV2GroupsID lists a team via the group protocol

SCIM endpoint that lists a Team (Colloquial for Group in the SCIM protocol)
*/
func (a *Client) GetV1ScimV2GroupsID(params *GetV1ScimV2GroupsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ScimV2GroupsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ScimV2GroupsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ScimV2GroupsId",
		Method:             "GET",
		PathPattern:        "/v1/scim/v2/Groups/{id}",
		ProducesMediaTypes: []string{"application/scim+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ScimV2GroupsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1ScimV2GroupsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ScimV2GroupsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ScimV2Users gets a list of users using s c i m protocol

SCIM endpoint that lists users. This endpoint will display a list of Users currently in the system.
*/
func (a *Client) GetV1ScimV2Users(params *GetV1ScimV2UsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ScimV2UsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ScimV2UsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ScimV2Users",
		Method:             "GET",
		PathPattern:        "/v1/scim/v2/Users",
		ProducesMediaTypes: []string{"application/scim+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ScimV2UsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1ScimV2UsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ScimV2Users: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ScimV2UsersID lists a user via the user protocol

SCIM endpoint that lists a User
*/
func (a *Client) GetV1ScimV2UsersID(params *GetV1ScimV2UsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ScimV2UsersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ScimV2UsersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ScimV2UsersId",
		Method:             "GET",
		PathPattern:        "/v1/scim/v2/Users/{id}",
		ProducesMediaTypes: []string{"application/scim+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ScimV2UsersIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1ScimV2UsersIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ScimV2UsersId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1ScimV2Groups creates a new team via the group protocol and assigns members to that team

SCIM endpoint to create a new Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role.
*/
func (a *Client) PostV1ScimV2Groups(params *PostV1ScimV2GroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ScimV2GroupsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ScimV2GroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1ScimV2Groups",
		Method:             "POST",
		PathPattern:        "/v1/scim/v2/Groups",
		ProducesMediaTypes: []string{"application/scim+json"},
		ConsumesMediaTypes: []string{"application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1ScimV2GroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1ScimV2GroupsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1ScimV2Groups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1ScimV2Users creates a new user using s c i m protocol

SCIM endpoint to create and provision a new User. This endpoint will provision the User, which allows them to accept their account throught their IDP or via the Forgot Password flow.
*/
func (a *Client) PostV1ScimV2Users(params *PostV1ScimV2UsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ScimV2UsersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ScimV2UsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1ScimV2Users",
		Method:             "POST",
		PathPattern:        "/v1/scim/v2/Users",
		ProducesMediaTypes: []string{"application/scim+json"},
		ConsumesMediaTypes: []string{"application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1ScimV2UsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1ScimV2UsersCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1ScimV2Users: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutV1ScimV2GroupsID updates a team via the group protocol and assigns members to that team

SCIM endpoint to update a Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role, any missing members will be removed from the team.
*/
func (a *Client) PutV1ScimV2GroupsID(params *PutV1ScimV2GroupsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutV1ScimV2GroupsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutV1ScimV2GroupsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putV1ScimV2GroupsId",
		Method:             "PUT",
		PathPattern:        "/v1/scim/v2/Groups/{id}",
		ProducesMediaTypes: []string{"application/scim+json"},
		ConsumesMediaTypes: []string{"application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutV1ScimV2GroupsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutV1ScimV2GroupsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putV1ScimV2GroupsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutV1ScimV2UsersID updates a user using s c i m protocol

PUT SCIM endpoint to update a User. This endpoint is used to replace a resource's attributes.
*/
func (a *Client) PutV1ScimV2UsersID(params *PutV1ScimV2UsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutV1ScimV2UsersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutV1ScimV2UsersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putV1ScimV2UsersId",
		Method:             "PUT",
		PathPattern:        "/v1/scim/v2/Users/{id}",
		ProducesMediaTypes: []string{"application/scim+json"},
		ConsumesMediaTypes: []string{"application/scim+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutV1ScimV2UsersIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutV1ScimV2UsersIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putV1ScimV2UsersId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
