// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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

// NewPatchV1WebhooksWebhookIDParams creates a new PatchV1WebhooksWebhookIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1WebhooksWebhookIDParams() *PatchV1WebhooksWebhookIDParams {
	return &PatchV1WebhooksWebhookIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1WebhooksWebhookIDParamsWithTimeout creates a new PatchV1WebhooksWebhookIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1WebhooksWebhookIDParamsWithTimeout(timeout time.Duration) *PatchV1WebhooksWebhookIDParams {
	return &PatchV1WebhooksWebhookIDParams{
		timeout: timeout,
	}
}

// NewPatchV1WebhooksWebhookIDParamsWithContext creates a new PatchV1WebhooksWebhookIDParams object
// with the ability to set a context for a request.
func NewPatchV1WebhooksWebhookIDParamsWithContext(ctx context.Context) *PatchV1WebhooksWebhookIDParams {
	return &PatchV1WebhooksWebhookIDParams{
		Context: ctx,
	}
}

// NewPatchV1WebhooksWebhookIDParamsWithHTTPClient creates a new PatchV1WebhooksWebhookIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1WebhooksWebhookIDParamsWithHTTPClient(client *http.Client) *PatchV1WebhooksWebhookIDParams {
	return &PatchV1WebhooksWebhookIDParams{
		HTTPClient: client,
	}
}

/* PatchV1WebhooksWebhookIDParams contains all the parameters to send to the API endpoint
   for the patch v1 webhooks webhook Id operation.

   Typically these are written to a http.Request.
*/
type PatchV1WebhooksWebhookIDParams struct {

	// V1Webhooks.
	V1Webhooks *models.PatchV1Webhooks

	/* WebhookID.

	   ID of a webhook
	*/
	WebhookID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 webhooks webhook Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1WebhooksWebhookIDParams) WithDefaults() *PatchV1WebhooksWebhookIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 webhooks webhook Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1WebhooksWebhookIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 webhooks webhook Id params
func (o *PatchV1WebhooksWebhookIDParams) WithTimeout(timeout time.Duration) *PatchV1WebhooksWebhookIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 webhooks webhook Id params
func (o *PatchV1WebhooksWebhookIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 webhooks webhook Id params
func (o *PatchV1WebhooksWebhookIDParams) WithContext(ctx context.Context) *PatchV1WebhooksWebhookIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 webhooks webhook Id params
func (o *PatchV1WebhooksWebhookIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 webhooks webhook Id params
func (o *PatchV1WebhooksWebhookIDParams) WithHTTPClient(client *http.Client) *PatchV1WebhooksWebhookIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 webhooks webhook Id params
func (o *PatchV1WebhooksWebhookIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1Webhooks adds the v1Webhooks to the patch v1 webhooks webhook Id params
func (o *PatchV1WebhooksWebhookIDParams) WithV1Webhooks(v1Webhooks *models.PatchV1Webhooks) *PatchV1WebhooksWebhookIDParams {
	o.SetV1Webhooks(v1Webhooks)
	return o
}

// SetV1Webhooks adds the v1Webhooks to the patch v1 webhooks webhook Id params
func (o *PatchV1WebhooksWebhookIDParams) SetV1Webhooks(v1Webhooks *models.PatchV1Webhooks) {
	o.V1Webhooks = v1Webhooks
}

// WithWebhookID adds the webhookID to the patch v1 webhooks webhook Id params
func (o *PatchV1WebhooksWebhookIDParams) WithWebhookID(webhookID string) *PatchV1WebhooksWebhookIDParams {
	o.SetWebhookID(webhookID)
	return o
}

// SetWebhookID adds the webhookId to the patch v1 webhooks webhook Id params
func (o *PatchV1WebhooksWebhookIDParams) SetWebhookID(webhookID string) {
	o.WebhookID = webhookID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1WebhooksWebhookIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1Webhooks != nil {
		if err := r.SetBodyParam(o.V1Webhooks); err != nil {
			return err
		}
	}

	// path param webhook_id
	if err := r.SetPathParam("webhook_id", o.WebhookID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}