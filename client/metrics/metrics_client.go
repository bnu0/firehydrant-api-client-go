// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new metrics API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for metrics API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1MetricsIncidents(params *GetV1MetricsIncidentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsIncidentsOK, error)

	GetV1MetricsRetrospectives(params *GetV1MetricsRetrospectivesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsRetrospectivesOK, error)

	GetV1MetricsServicesServiceID(params *GetV1MetricsServicesServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsServicesServiceIDOK, error)

	GetV1MetricsUserInvolvements(params *GetV1MetricsUserInvolvementsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsUserInvolvementsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetV1MetricsIncidents Returns a report with time bucketed analytics data
*/
func (a *Client) GetV1MetricsIncidents(params *GetV1MetricsIncidentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsIncidentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsIncidentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsIncidents",
		Method:             "GET",
		PathPattern:        "/v1/metrics/incidents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsIncidentsReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsIncidentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsIncidents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1MetricsRetrospectives Returns a report with retrospective analytics data
*/
func (a *Client) GetV1MetricsRetrospectives(params *GetV1MetricsRetrospectivesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsRetrospectivesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsRetrospectivesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsRetrospectives",
		Method:             "GET",
		PathPattern:        "/v1/metrics/retrospectives",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsRetrospectivesReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsRetrospectivesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsRetrospectives: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1MetricsServicesServiceID Return metrics for a specific service
*/
func (a *Client) GetV1MetricsServicesServiceID(params *GetV1MetricsServicesServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsServicesServiceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsServicesServiceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsServicesServiceId",
		Method:             "GET",
		PathPattern:        "/v1/metrics/services/{service_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsServicesServiceIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsServicesServiceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsServicesServiceId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1MetricsUserInvolvements Returns a report with time bucketed analytics data
*/
func (a *Client) GetV1MetricsUserInvolvements(params *GetV1MetricsUserInvolvementsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsUserInvolvementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsUserInvolvementsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsUserInvolvements",
		Method:             "GET",
		PathPattern:        "/v1/metrics/user_involvements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsUserInvolvementsReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsUserInvolvementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsUserInvolvements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}