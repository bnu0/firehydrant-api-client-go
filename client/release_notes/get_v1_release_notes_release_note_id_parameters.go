// Code generated by go-swagger; DO NOT EDIT.

package release_notes

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
)

// NewGetV1ReleaseNotesReleaseNoteIDParams creates a new GetV1ReleaseNotesReleaseNoteIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ReleaseNotesReleaseNoteIDParams() *GetV1ReleaseNotesReleaseNoteIDParams {
	return &GetV1ReleaseNotesReleaseNoteIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ReleaseNotesReleaseNoteIDParamsWithTimeout creates a new GetV1ReleaseNotesReleaseNoteIDParams object
// with the ability to set a timeout on a request.
func NewGetV1ReleaseNotesReleaseNoteIDParamsWithTimeout(timeout time.Duration) *GetV1ReleaseNotesReleaseNoteIDParams {
	return &GetV1ReleaseNotesReleaseNoteIDParams{
		timeout: timeout,
	}
}

// NewGetV1ReleaseNotesReleaseNoteIDParamsWithContext creates a new GetV1ReleaseNotesReleaseNoteIDParams object
// with the ability to set a context for a request.
func NewGetV1ReleaseNotesReleaseNoteIDParamsWithContext(ctx context.Context) *GetV1ReleaseNotesReleaseNoteIDParams {
	return &GetV1ReleaseNotesReleaseNoteIDParams{
		Context: ctx,
	}
}

// NewGetV1ReleaseNotesReleaseNoteIDParamsWithHTTPClient creates a new GetV1ReleaseNotesReleaseNoteIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ReleaseNotesReleaseNoteIDParamsWithHTTPClient(client *http.Client) *GetV1ReleaseNotesReleaseNoteIDParams {
	return &GetV1ReleaseNotesReleaseNoteIDParams{
		HTTPClient: client,
	}
}

/*
GetV1ReleaseNotesReleaseNoteIDParams contains all the parameters to send to the API endpoint

	for the get v1 release notes release note Id operation.

	Typically these are written to a http.Request.
*/
type GetV1ReleaseNotesReleaseNoteIDParams struct {

	// ReleaseNoteID.
	ReleaseNoteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 release notes release note Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ReleaseNotesReleaseNoteIDParams) WithDefaults() *GetV1ReleaseNotesReleaseNoteIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 release notes release note Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ReleaseNotesReleaseNoteIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 release notes release note Id params
func (o *GetV1ReleaseNotesReleaseNoteIDParams) WithTimeout(timeout time.Duration) *GetV1ReleaseNotesReleaseNoteIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 release notes release note Id params
func (o *GetV1ReleaseNotesReleaseNoteIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 release notes release note Id params
func (o *GetV1ReleaseNotesReleaseNoteIDParams) WithContext(ctx context.Context) *GetV1ReleaseNotesReleaseNoteIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 release notes release note Id params
func (o *GetV1ReleaseNotesReleaseNoteIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 release notes release note Id params
func (o *GetV1ReleaseNotesReleaseNoteIDParams) WithHTTPClient(client *http.Client) *GetV1ReleaseNotesReleaseNoteIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 release notes release note Id params
func (o *GetV1ReleaseNotesReleaseNoteIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReleaseNoteID adds the releaseNoteID to the get v1 release notes release note Id params
func (o *GetV1ReleaseNotesReleaseNoteIDParams) WithReleaseNoteID(releaseNoteID string) *GetV1ReleaseNotesReleaseNoteIDParams {
	o.SetReleaseNoteID(releaseNoteID)
	return o
}

// SetReleaseNoteID adds the releaseNoteId to the get v1 release notes release note Id params
func (o *GetV1ReleaseNotesReleaseNoteIDParams) SetReleaseNoteID(releaseNoteID string) {
	o.ReleaseNoteID = releaseNoteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ReleaseNotesReleaseNoteIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param release_note_id
	if err := r.SetPathParam("release_note_id", o.ReleaseNoteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
