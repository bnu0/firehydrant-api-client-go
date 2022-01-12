// Code generated by go-swagger; DO NOT EDIT.

package severities

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
	"github.com/go-openapi/swag"
)

// NewGetV1SeveritiesParams creates a new GetV1SeveritiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SeveritiesParams() *GetV1SeveritiesParams {
	return &GetV1SeveritiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SeveritiesParamsWithTimeout creates a new GetV1SeveritiesParams object
// with the ability to set a timeout on a request.
func NewGetV1SeveritiesParamsWithTimeout(timeout time.Duration) *GetV1SeveritiesParams {
	return &GetV1SeveritiesParams{
		timeout: timeout,
	}
}

// NewGetV1SeveritiesParamsWithContext creates a new GetV1SeveritiesParams object
// with the ability to set a context for a request.
func NewGetV1SeveritiesParamsWithContext(ctx context.Context) *GetV1SeveritiesParams {
	return &GetV1SeveritiesParams{
		Context: ctx,
	}
}

// NewGetV1SeveritiesParamsWithHTTPClient creates a new GetV1SeveritiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SeveritiesParamsWithHTTPClient(client *http.Client) *GetV1SeveritiesParams {
	return &GetV1SeveritiesParams{
		HTTPClient: client,
	}
}

/* GetV1SeveritiesParams contains all the parameters to send to the API endpoint
   for the get v1 severities operation.

   Typically these are written to a http.Request.
*/
type GetV1SeveritiesParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 severities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SeveritiesParams) WithDefaults() *GetV1SeveritiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 severities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SeveritiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 severities params
func (o *GetV1SeveritiesParams) WithTimeout(timeout time.Duration) *GetV1SeveritiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 severities params
func (o *GetV1SeveritiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 severities params
func (o *GetV1SeveritiesParams) WithContext(ctx context.Context) *GetV1SeveritiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 severities params
func (o *GetV1SeveritiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 severities params
func (o *GetV1SeveritiesParams) WithHTTPClient(client *http.Client) *GetV1SeveritiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 severities params
func (o *GetV1SeveritiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 severities params
func (o *GetV1SeveritiesParams) WithPage(page *int32) *GetV1SeveritiesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 severities params
func (o *GetV1SeveritiesParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 severities params
func (o *GetV1SeveritiesParams) WithPerPage(perPage *int32) *GetV1SeveritiesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 severities params
func (o *GetV1SeveritiesParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SeveritiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}