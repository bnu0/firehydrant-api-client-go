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
	"github.com/go-openapi/swag"

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1TeamsTeamIDParams creates a new PatchV1TeamsTeamIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1TeamsTeamIDParams() *PatchV1TeamsTeamIDParams {
	return &PatchV1TeamsTeamIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1TeamsTeamIDParamsWithTimeout creates a new PatchV1TeamsTeamIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1TeamsTeamIDParamsWithTimeout(timeout time.Duration) *PatchV1TeamsTeamIDParams {
	return &PatchV1TeamsTeamIDParams{
		timeout: timeout,
	}
}

// NewPatchV1TeamsTeamIDParamsWithContext creates a new PatchV1TeamsTeamIDParams object
// with the ability to set a context for a request.
func NewPatchV1TeamsTeamIDParamsWithContext(ctx context.Context) *PatchV1TeamsTeamIDParams {
	return &PatchV1TeamsTeamIDParams{
		Context: ctx,
	}
}

// NewPatchV1TeamsTeamIDParamsWithHTTPClient creates a new PatchV1TeamsTeamIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1TeamsTeamIDParamsWithHTTPClient(client *http.Client) *PatchV1TeamsTeamIDParams {
	return &PatchV1TeamsTeamIDParams{
		HTTPClient: client,
	}
}

/* PatchV1TeamsTeamIDParams contains all the parameters to send to the API endpoint
   for the patch v1 teams team Id operation.

   Typically these are written to a http.Request.
*/
type PatchV1TeamsTeamIDParams struct {

	// V1Teams.
	V1Teams *models.PatchV1Teams

	// TeamID.
	//
	// Format: int32
	TeamID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 teams team Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1TeamsTeamIDParams) WithDefaults() *PatchV1TeamsTeamIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 teams team Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1TeamsTeamIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 teams team Id params
func (o *PatchV1TeamsTeamIDParams) WithTimeout(timeout time.Duration) *PatchV1TeamsTeamIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 teams team Id params
func (o *PatchV1TeamsTeamIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 teams team Id params
func (o *PatchV1TeamsTeamIDParams) WithContext(ctx context.Context) *PatchV1TeamsTeamIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 teams team Id params
func (o *PatchV1TeamsTeamIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 teams team Id params
func (o *PatchV1TeamsTeamIDParams) WithHTTPClient(client *http.Client) *PatchV1TeamsTeamIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 teams team Id params
func (o *PatchV1TeamsTeamIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1Teams adds the v1Teams to the patch v1 teams team Id params
func (o *PatchV1TeamsTeamIDParams) WithV1Teams(v1Teams *models.PatchV1Teams) *PatchV1TeamsTeamIDParams {
	o.SetV1Teams(v1Teams)
	return o
}

// SetV1Teams adds the v1Teams to the patch v1 teams team Id params
func (o *PatchV1TeamsTeamIDParams) SetV1Teams(v1Teams *models.PatchV1Teams) {
	o.V1Teams = v1Teams
}

// WithTeamID adds the teamID to the patch v1 teams team Id params
func (o *PatchV1TeamsTeamIDParams) WithTeamID(teamID int32) *PatchV1TeamsTeamIDParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the patch v1 teams team Id params
func (o *PatchV1TeamsTeamIDParams) SetTeamID(teamID int32) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1TeamsTeamIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1Teams != nil {
		if err := r.SetBodyParam(o.V1Teams); err != nil {
			return err
		}
	}

	// path param team_id
	if err := r.SetPathParam("team_id", swag.FormatInt32(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}