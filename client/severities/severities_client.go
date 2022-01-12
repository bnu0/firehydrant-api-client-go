// Code generated by go-swagger; DO NOT EDIT.

package severities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new severities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for severities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1SeveritiesSeveritySlug(params *DeleteV1SeveritiesSeveritySlugParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1SeveritiesSeveritySlugOK, error)

	GetV1Severities(params *GetV1SeveritiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SeveritiesOK, error)

	GetV1SeveritiesSeveritySlug(params *GetV1SeveritiesSeveritySlugParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SeveritiesSeveritySlugOK, error)

	PatchV1SeveritiesSeveritySlug(params *PatchV1SeveritiesSeveritySlugParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1SeveritiesSeveritySlugOK, error)

	PostV1Severities(params *PostV1SeveritiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1SeveritiesCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteV1SeveritiesSeveritySlug deletes a specific severity

  Delete a specific severity
*/
func (a *Client) DeleteV1SeveritiesSeveritySlug(params *DeleteV1SeveritiesSeveritySlugParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1SeveritiesSeveritySlugOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1SeveritiesSeveritySlugParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1SeveritiesSeveritySlug",
		Method:             "DELETE",
		PathPattern:        "/v1/severities/{severity_slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1SeveritiesSeveritySlugReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1SeveritiesSeveritySlugOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1SeveritiesSeveritySlug: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1Severities lists severities

  Lists severities
*/
func (a *Client) GetV1Severities(params *GetV1SeveritiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SeveritiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SeveritiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1Severities",
		Method:             "GET",
		PathPattern:        "/v1/severities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SeveritiesReader{formats: a.formats},
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
	success, ok := result.(*GetV1SeveritiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1Severities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1SeveritiesSeveritySlug retrieves a specific severity

  Retrieve a specific severity
*/
func (a *Client) GetV1SeveritiesSeveritySlug(params *GetV1SeveritiesSeveritySlugParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SeveritiesSeveritySlugOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SeveritiesSeveritySlugParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SeveritiesSeveritySlug",
		Method:             "GET",
		PathPattern:        "/v1/severities/{severity_slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SeveritiesSeveritySlugReader{formats: a.formats},
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
	success, ok := result.(*GetV1SeveritiesSeveritySlugOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SeveritiesSeveritySlug: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchV1SeveritiesSeveritySlug updates a specific severity

  Update a specific severity
*/
func (a *Client) PatchV1SeveritiesSeveritySlug(params *PatchV1SeveritiesSeveritySlugParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1SeveritiesSeveritySlugOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1SeveritiesSeveritySlugParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1SeveritiesSeveritySlug",
		Method:             "PATCH",
		PathPattern:        "/v1/severities/{severity_slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1SeveritiesSeveritySlugReader{formats: a.formats},
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
	success, ok := result.(*PatchV1SeveritiesSeveritySlugOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1SeveritiesSeveritySlug: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostV1Severities creates severity

  Create a new severity
*/
func (a *Client) PostV1Severities(params *PostV1SeveritiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1SeveritiesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1SeveritiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1Severities",
		Method:             "POST",
		PathPattern:        "/v1/severities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1SeveritiesReader{formats: a.formats},
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
	success, ok := result.(*PostV1SeveritiesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1Severities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}