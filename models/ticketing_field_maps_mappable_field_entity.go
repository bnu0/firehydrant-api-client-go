// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TicketingFieldMapsMappableFieldEntity Ticketing_FieldMaps_MappableFieldEntity model
//
// swagger:model Ticketing_FieldMaps_MappableFieldEntity
type TicketingFieldMapsMappableFieldEntity struct {

	// The allowed values of the field
	AllowedValues string `json:"allowed_values,omitempty"`

	// The human-readable name of the field
	Label string `json:"label,omitempty"`

	// If the field is required to be mapped
	Required string `json:"required,omitempty"`

	// The allowed type of the field
	Type string `json:"type,omitempty"`

	// The ID of the field
	Value string `json:"value,omitempty"`
}

// Validate validates this ticketing field maps mappable field entity
func (m *TicketingFieldMapsMappableFieldEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticketing field maps mappable field entity based on context it is used
func (m *TicketingFieldMapsMappableFieldEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TicketingFieldMapsMappableFieldEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TicketingFieldMapsMappableFieldEntity) UnmarshalBinary(b []byte) error {
	var res TicketingFieldMapsMappableFieldEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}