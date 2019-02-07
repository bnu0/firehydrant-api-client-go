// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// NewPostV1ComponentsComponentIDEnvironmentsParams creates a new PostV1ComponentsComponentIDEnvironmentsParams object
// with the default values initialized.
func NewPostV1ComponentsComponentIDEnvironmentsParams() *PostV1ComponentsComponentIDEnvironmentsParams {
	var ()
	return &PostV1ComponentsComponentIDEnvironmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ComponentsComponentIDEnvironmentsParamsWithTimeout creates a new PostV1ComponentsComponentIDEnvironmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostV1ComponentsComponentIDEnvironmentsParamsWithTimeout(timeout time.Duration) *PostV1ComponentsComponentIDEnvironmentsParams {
	var ()
	return &PostV1ComponentsComponentIDEnvironmentsParams{

		timeout: timeout,
	}
}

// NewPostV1ComponentsComponentIDEnvironmentsParamsWithContext creates a new PostV1ComponentsComponentIDEnvironmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostV1ComponentsComponentIDEnvironmentsParamsWithContext(ctx context.Context) *PostV1ComponentsComponentIDEnvironmentsParams {
	var ()
	return &PostV1ComponentsComponentIDEnvironmentsParams{

		Context: ctx,
	}
}

// NewPostV1ComponentsComponentIDEnvironmentsParamsWithHTTPClient creates a new PostV1ComponentsComponentIDEnvironmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostV1ComponentsComponentIDEnvironmentsParamsWithHTTPClient(client *http.Client) *PostV1ComponentsComponentIDEnvironmentsParams {
	var ()
	return &PostV1ComponentsComponentIDEnvironmentsParams{
		HTTPClient: client,
	}
}

/*PostV1ComponentsComponentIDEnvironmentsParams contains all the parameters to send to the API endpoint
for the post v1 components component Id environments operation typically these are written to a http.Request
*/
type PostV1ComponentsComponentIDEnvironmentsParams struct {

	/*V1ComponentsComponentIDEnvironments*/
	V1ComponentsComponentIDEnvironments *models.PostV1ComponentsComponentIDEnvironments
	/*ComponentID
	  Component UUID

	*/
	ComponentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post v1 components component Id environments params
func (o *PostV1ComponentsComponentIDEnvironmentsParams) WithTimeout(timeout time.Duration) *PostV1ComponentsComponentIDEnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 components component Id environments params
func (o *PostV1ComponentsComponentIDEnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 components component Id environments params
func (o *PostV1ComponentsComponentIDEnvironmentsParams) WithContext(ctx context.Context) *PostV1ComponentsComponentIDEnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 components component Id environments params
func (o *PostV1ComponentsComponentIDEnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 components component Id environments params
func (o *PostV1ComponentsComponentIDEnvironmentsParams) WithHTTPClient(client *http.Client) *PostV1ComponentsComponentIDEnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 components component Id environments params
func (o *PostV1ComponentsComponentIDEnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1ComponentsComponentIDEnvironments adds the v1ComponentsComponentIDEnvironments to the post v1 components component Id environments params
func (o *PostV1ComponentsComponentIDEnvironmentsParams) WithV1ComponentsComponentIDEnvironments(v1ComponentsComponentIDEnvironments *models.PostV1ComponentsComponentIDEnvironments) *PostV1ComponentsComponentIDEnvironmentsParams {
	o.SetV1ComponentsComponentIDEnvironments(v1ComponentsComponentIDEnvironments)
	return o
}

// SetV1ComponentsComponentIDEnvironments adds the v1ComponentsComponentIdEnvironments to the post v1 components component Id environments params
func (o *PostV1ComponentsComponentIDEnvironmentsParams) SetV1ComponentsComponentIDEnvironments(v1ComponentsComponentIDEnvironments *models.PostV1ComponentsComponentIDEnvironments) {
	o.V1ComponentsComponentIDEnvironments = v1ComponentsComponentIDEnvironments
}

// WithComponentID adds the componentID to the post v1 components component Id environments params
func (o *PostV1ComponentsComponentIDEnvironmentsParams) WithComponentID(componentID string) *PostV1ComponentsComponentIDEnvironmentsParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the post v1 components component Id environments params
func (o *PostV1ComponentsComponentIDEnvironmentsParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ComponentsComponentIDEnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.V1ComponentsComponentIDEnvironments != nil {
		if err := r.SetBodyParam(o.V1ComponentsComponentIDEnvironments); err != nil {
			return err
		}
	}

	// path param component_id
	if err := r.SetPathParam("component_id", o.ComponentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}