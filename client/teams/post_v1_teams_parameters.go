// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// NewPostV1TeamsParams creates a new PostV1TeamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1TeamsParams() *PostV1TeamsParams {
	return &PostV1TeamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1TeamsParamsWithTimeout creates a new PostV1TeamsParams object
// with the ability to set a timeout on a request.
func NewPostV1TeamsParamsWithTimeout(timeout time.Duration) *PostV1TeamsParams {
	return &PostV1TeamsParams{
		timeout: timeout,
	}
}

// NewPostV1TeamsParamsWithContext creates a new PostV1TeamsParams object
// with the ability to set a context for a request.
func NewPostV1TeamsParamsWithContext(ctx context.Context) *PostV1TeamsParams {
	return &PostV1TeamsParams{
		Context: ctx,
	}
}

// NewPostV1TeamsParamsWithHTTPClient creates a new PostV1TeamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1TeamsParamsWithHTTPClient(client *http.Client) *PostV1TeamsParams {
	return &PostV1TeamsParams{
		HTTPClient: client,
	}
}

/* PostV1TeamsParams contains all the parameters to send to the API endpoint
   for the post v1 teams operation.

   Typically these are written to a http.Request.
*/
type PostV1TeamsParams struct {

	// V1Teams.
	V1Teams *models.PostV1Teams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1TeamsParams) WithDefaults() *PostV1TeamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1TeamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 teams params
func (o *PostV1TeamsParams) WithTimeout(timeout time.Duration) *PostV1TeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 teams params
func (o *PostV1TeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 teams params
func (o *PostV1TeamsParams) WithContext(ctx context.Context) *PostV1TeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 teams params
func (o *PostV1TeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 teams params
func (o *PostV1TeamsParams) WithHTTPClient(client *http.Client) *PostV1TeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 teams params
func (o *PostV1TeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1Teams adds the v1Teams to the post v1 teams params
func (o *PostV1TeamsParams) WithV1Teams(v1Teams *models.PostV1Teams) *PostV1TeamsParams {
	o.SetV1Teams(v1Teams)
	return o
}

// SetV1Teams adds the v1Teams to the post v1 teams params
func (o *PostV1TeamsParams) SetV1Teams(v1Teams *models.PostV1Teams) {
	o.V1Teams = v1Teams
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1TeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1Teams != nil {
		if err := r.SetBodyParam(o.V1Teams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}