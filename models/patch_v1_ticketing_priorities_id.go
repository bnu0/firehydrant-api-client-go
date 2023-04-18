// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchV1TicketingPrioritiesID Update a single ticketing priority's attributes
//
// swagger:model patchV1TicketingPrioritiesId
type PatchV1TicketingPrioritiesID struct {

	// name
	Name string `json:"name,omitempty"`

	// The position that this priority should take in your list of priorities. Priorities should be ordered from highest to lowest, with the highest priority at 0. If a position isn't specified, the new priority will be added to the end of the list; if another priority already exists at the specified position, this priority will shift that priority and all priorities down the list.
	Position int32 `json:"position,omitempty"`
}

// Validate validates this patch v1 ticketing priorities Id
func (m *PatchV1TicketingPrioritiesID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch v1 ticketing priorities Id based on context it is used
func (m *PatchV1TicketingPrioritiesID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1TicketingPrioritiesID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1TicketingPrioritiesID) UnmarshalBinary(b []byte) error {
	var res PatchV1TicketingPrioritiesID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
