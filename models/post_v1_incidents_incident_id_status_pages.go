// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1IncidentsIncidentIDStatusPages Add a status page to an incident.
//
// swagger:model postV1IncidentsIncidentIdStatusPages
type PostV1IncidentsIncidentIDStatusPages struct {

	// integration id
	// Required: true
	IntegrationID *string `json:"integration_id"`

	// integration slug
	// Required: true
	IntegrationSlug *string `json:"integration_slug"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this post v1 incidents incident Id status pages
func (m *PostV1IncidentsIncidentIDStatusPages) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegrationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationSlug(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentsIncidentIDStatusPages) validateIntegrationID(formats strfmt.Registry) error {

	if err := validate.Required("integration_id", "body", m.IntegrationID); err != nil {
		return err
	}

	return nil
}

func (m *PostV1IncidentsIncidentIDStatusPages) validateIntegrationSlug(formats strfmt.Registry) error {

	if err := validate.Required("integration_slug", "body", m.IntegrationSlug); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 incidents incident Id status pages based on context it is used
func (m *PostV1IncidentsIncidentIDStatusPages) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentsIncidentIDStatusPages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentsIncidentIDStatusPages) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentsIncidentIDStatusPages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
