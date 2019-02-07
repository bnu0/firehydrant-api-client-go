// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

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

// NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams creates a new DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams object
// with the default values initialized.
func NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams() *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams {
	var ()
	return &DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParamsWithTimeout creates a new DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParamsWithTimeout(timeout time.Duration) *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams {
	var ()
	return &DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams{

		timeout: timeout,
	}
}

// NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParamsWithContext creates a new DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParamsWithContext(ctx context.Context) *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams {
	var ()
	return &DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams{

		Context: ctx,
	}
}

// NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParamsWithHTTPClient creates a new DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParamsWithHTTPClient(client *http.Client) *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams {
	var ()
	return &DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams{
		HTTPClient: client,
	}
}

/*DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams contains all the parameters to send to the API endpoint
for the delete v1 post mortems reports report Id events from incident incident event Id operation typically these are written to a http.Request
*/
type DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams struct {

	/*IncidentEventID*/
	IncidentEventID string
	/*ReportID*/
	ReportID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete v1 post mortems reports report Id events from incident incident event Id params
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) WithTimeout(timeout time.Duration) *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 post mortems reports report Id events from incident incident event Id params
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 post mortems reports report Id events from incident incident event Id params
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) WithContext(ctx context.Context) *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 post mortems reports report Id events from incident incident event Id params
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 post mortems reports report Id events from incident incident event Id params
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) WithHTTPClient(client *http.Client) *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 post mortems reports report Id events from incident incident event Id params
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentEventID adds the incidentEventID to the delete v1 post mortems reports report Id events from incident incident event Id params
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) WithIncidentEventID(incidentEventID string) *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams {
	o.SetIncidentEventID(incidentEventID)
	return o
}

// SetIncidentEventID adds the incidentEventId to the delete v1 post mortems reports report Id events from incident incident event Id params
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) SetIncidentEventID(incidentEventID string) {
	o.IncidentEventID = incidentEventID
}

// WithReportID adds the reportID to the delete v1 post mortems reports report Id events from incident incident event Id params
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) WithReportID(reportID int32) *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the delete v1 post mortems reports report Id events from incident incident event Id params
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) SetReportID(reportID int32) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_event_id
	if err := r.SetPathParam("incident_event_id", o.IncidentEventID); err != nil {
		return err
	}

	// path param report_id
	if err := r.SetPathParam("report_id", swag.FormatInt32(o.ReportID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}