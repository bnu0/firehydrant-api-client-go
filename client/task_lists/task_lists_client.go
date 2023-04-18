// Code generated by go-swagger; DO NOT EDIT.

package task_lists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new task lists API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for task lists API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1TaskListsTaskListID(params *DeleteV1TaskListsTaskListIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1TaskListsTaskListIDOK, error)

	GetV1TaskLists(params *GetV1TaskListsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1TaskListsOK, error)

	GetV1TaskListsTaskListID(params *GetV1TaskListsTaskListIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1TaskListsTaskListIDOK, error)

	PatchV1TaskListsTaskListID(params *PatchV1TaskListsTaskListIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1TaskListsTaskListIDOK, error)

	PostV1TaskLists(params *PostV1TaskListsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1TaskListsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1TaskListsTaskListID deletes a task list

Delete a task list
*/
func (a *Client) DeleteV1TaskListsTaskListID(params *DeleteV1TaskListsTaskListIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1TaskListsTaskListIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1TaskListsTaskListIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1TaskListsTaskListId",
		Method:             "DELETE",
		PathPattern:        "/v1/task_lists/{task_list_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1TaskListsTaskListIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1TaskListsTaskListIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1TaskListsTaskListId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1TaskLists lists all task lists

Lists all task lists for your organization
*/
func (a *Client) GetV1TaskLists(params *GetV1TaskListsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1TaskListsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TaskListsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1TaskLists",
		Method:             "GET",
		PathPattern:        "/v1/task_lists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1TaskListsReader{formats: a.formats},
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
	success, ok := result.(*GetV1TaskListsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1TaskLists: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1TaskListsTaskListID retrieves a task list

Retrieves a single task list by ID
*/
func (a *Client) GetV1TaskListsTaskListID(params *GetV1TaskListsTaskListIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1TaskListsTaskListIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TaskListsTaskListIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1TaskListsTaskListId",
		Method:             "GET",
		PathPattern:        "/v1/task_lists/{task_list_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1TaskListsTaskListIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1TaskListsTaskListIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1TaskListsTaskListId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1TaskListsTaskListID updates a task list

Updates a task list's attributes and task list items
*/
func (a *Client) PatchV1TaskListsTaskListID(params *PatchV1TaskListsTaskListIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1TaskListsTaskListIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1TaskListsTaskListIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1TaskListsTaskListId",
		Method:             "PATCH",
		PathPattern:        "/v1/task_lists/{task_list_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1TaskListsTaskListIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1TaskListsTaskListIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1TaskListsTaskListId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1TaskLists creates a task list

Creates a new task list
*/
func (a *Client) PostV1TaskLists(params *PostV1TaskListsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1TaskListsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1TaskListsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1TaskLists",
		Method:             "POST",
		PathPattern:        "/v1/task_lists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1TaskListsReader{formats: a.formats},
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
	success, ok := result.(*PostV1TaskListsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1TaskLists: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}