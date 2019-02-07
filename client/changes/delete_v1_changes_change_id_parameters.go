// Code generated by go-swagger; DO NOT EDIT.

package changes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteV1ChangesChangeIDParams creates a new DeleteV1ChangesChangeIDParams object
// with the default values initialized.
func NewDeleteV1ChangesChangeIDParams() *DeleteV1ChangesChangeIDParams {
	var ()
	return &DeleteV1ChangesChangeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ChangesChangeIDParamsWithTimeout creates a new DeleteV1ChangesChangeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteV1ChangesChangeIDParamsWithTimeout(timeout time.Duration) *DeleteV1ChangesChangeIDParams {
	var ()
	return &DeleteV1ChangesChangeIDParams{

		timeout: timeout,
	}
}

// NewDeleteV1ChangesChangeIDParamsWithContext creates a new DeleteV1ChangesChangeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteV1ChangesChangeIDParamsWithContext(ctx context.Context) *DeleteV1ChangesChangeIDParams {
	var ()
	return &DeleteV1ChangesChangeIDParams{

		Context: ctx,
	}
}

// NewDeleteV1ChangesChangeIDParamsWithHTTPClient creates a new DeleteV1ChangesChangeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteV1ChangesChangeIDParamsWithHTTPClient(client *http.Client) *DeleteV1ChangesChangeIDParams {
	var ()
	return &DeleteV1ChangesChangeIDParams{
		HTTPClient: client,
	}
}

/*DeleteV1ChangesChangeIDParams contains all the parameters to send to the API endpoint
for the delete v1 changes change Id operation typically these are written to a http.Request
*/
type DeleteV1ChangesChangeIDParams struct {

	/*ChangeID*/
	ChangeID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete v1 changes change Id params
func (o *DeleteV1ChangesChangeIDParams) WithTimeout(timeout time.Duration) *DeleteV1ChangesChangeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 changes change Id params
func (o *DeleteV1ChangesChangeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 changes change Id params
func (o *DeleteV1ChangesChangeIDParams) WithContext(ctx context.Context) *DeleteV1ChangesChangeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 changes change Id params
func (o *DeleteV1ChangesChangeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 changes change Id params
func (o *DeleteV1ChangesChangeIDParams) WithHTTPClient(client *http.Client) *DeleteV1ChangesChangeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 changes change Id params
func (o *DeleteV1ChangesChangeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeID adds the changeID to the delete v1 changes change Id params
func (o *DeleteV1ChangesChangeIDParams) WithChangeID(changeID int32) *DeleteV1ChangesChangeIDParams {
	o.SetChangeID(changeID)
	return o
}

// SetChangeID adds the changeId to the delete v1 changes change Id params
func (o *DeleteV1ChangesChangeIDParams) SetChangeID(changeID int32) {
	o.ChangeID = changeID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ChangesChangeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param change_id
	if err := r.SetPathParam("change_id", swag.FormatInt32(o.ChangeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}