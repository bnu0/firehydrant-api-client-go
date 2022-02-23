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

// PostV1Priorities Create a new priority
//
// swagger:model postV1Priorities
type PostV1Priorities struct {

	// description
	Description string `json:"description,omitempty"`

	// slug
	// Required: true
	Slug *string `json:"slug"`
}

// Validate validates this post v1 priorities
func (m *PostV1Priorities) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1Priorities) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 priorities based on context it is used
func (m *PostV1Priorities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1Priorities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1Priorities) UnmarshalBinary(b []byte) error {
	var res PostV1Priorities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
