// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RelatedChangeEventEntity Update a change event
// swagger:model RelatedChangeEventEntity
type RelatedChangeEventEntity struct {

	// change event
	ChangeEvent *ChangeEventEntity `json:"change_event,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by
	CreatedBy *AuthorEntity `json:"created_by,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// incident id
	IncidentID string `json:"incident_id,omitempty"`

	// type
	// Enum: [caused fixed suspect dismissed]
	Type string `json:"type,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// The reason why this change event is related to this incident
	Why string `json:"why,omitempty"`
}

// Validate validates this related change event entity
func (m *RelatedChangeEventEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangeEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelatedChangeEventEntity) validateChangeEvent(formats strfmt.Registry) error {

	if swag.IsZero(m.ChangeEvent) { // not required
		return nil
	}

	if m.ChangeEvent != nil {
		if err := m.ChangeEvent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("change_event")
			}
			return err
		}
	}

	return nil
}

func (m *RelatedChangeEventEntity) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelatedChangeEventEntity) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

var relatedChangeEventEntityTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["caused","fixed","suspect","dismissed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		relatedChangeEventEntityTypeTypePropEnum = append(relatedChangeEventEntityTypeTypePropEnum, v)
	}
}

const (

	// RelatedChangeEventEntityTypeCaused captures enum value "caused"
	RelatedChangeEventEntityTypeCaused string = "caused"

	// RelatedChangeEventEntityTypeFixed captures enum value "fixed"
	RelatedChangeEventEntityTypeFixed string = "fixed"

	// RelatedChangeEventEntityTypeSuspect captures enum value "suspect"
	RelatedChangeEventEntityTypeSuspect string = "suspect"

	// RelatedChangeEventEntityTypeDismissed captures enum value "dismissed"
	RelatedChangeEventEntityTypeDismissed string = "dismissed"
)

// prop value enum
func (m *RelatedChangeEventEntity) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, relatedChangeEventEntityTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RelatedChangeEventEntity) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *RelatedChangeEventEntity) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelatedChangeEventEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelatedChangeEventEntity) UnmarshalBinary(b []byte) error {
	var res RelatedChangeEventEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}