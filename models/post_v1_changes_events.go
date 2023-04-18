// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1ChangesEvents Create a change event
//
// swagger:model postV1ChangesEvents
type PostV1ChangesEvents struct {

	// JSON objects representing attachments, see attachments documentation for the schema
	Attachments []*PostV1ChangesEventsAttachmentsItems0 `json:"attachments"`

	// Array of additional authors to add to the change event, the creating actor will automatically be added as an author
	Authors []*PostV1ChangesEventsAuthorsItems0 `json:"authors"`

	// If provided and valid, the event will be linked to all changes that have the same identities. Identity *values* must be unique.
	ChangeIdentities []*PostV1ChangesEventsChangeIdentitiesItems0 `json:"change_identities"`

	// An array of change IDs
	Changes []string `json:"changes"`

	// description
	Description string `json:"description,omitempty"`

	// ends at
	// Format: date-time
	EndsAt strfmt.DateTime `json:"ends_at,omitempty"`

	// An array of environment IDs
	Environments []string `json:"environments"`

	// The ID of a change event as assigned by an external provider
	ExternalID string `json:"external_id,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// An array of service IDs
	Services []string `json:"services"`

	// starts at
	// Format: date-time
	StartsAt strfmt.DateTime `json:"starts_at,omitempty"`

	// summary
	// Required: true
	Summary *string `json:"summary"`
}

// Validate validates this post v1 changes events
func (m *PostV1ChangesEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangeIdentities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ChangesEvents) validateAttachments(formats strfmt.Registry) error {
	if swag.IsZero(m.Attachments) { // not required
		return nil
	}

	for i := 0; i < len(m.Attachments); i++ {
		if swag.IsZero(m.Attachments[i]) { // not required
			continue
		}

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1ChangesEvents) validateAuthors(formats strfmt.Registry) error {
	if swag.IsZero(m.Authors) { // not required
		return nil
	}

	for i := 0; i < len(m.Authors); i++ {
		if swag.IsZero(m.Authors[i]) { // not required
			continue
		}

		if m.Authors[i] != nil {
			if err := m.Authors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("authors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("authors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1ChangesEvents) validateChangeIdentities(formats strfmt.Registry) error {
	if swag.IsZero(m.ChangeIdentities) { // not required
		return nil
	}

	for i := 0; i < len(m.ChangeIdentities); i++ {
		if swag.IsZero(m.ChangeIdentities[i]) { // not required
			continue
		}

		if m.ChangeIdentities[i] != nil {
			if err := m.ChangeIdentities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("change_identities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("change_identities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1ChangesEvents) validateEndsAt(formats strfmt.Registry) error {
	if swag.IsZero(m.EndsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("ends_at", "body", "date-time", m.EndsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ChangesEvents) validateStartsAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("starts_at", "body", "date-time", m.StartsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ChangesEvents) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("summary", "body", m.Summary); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post v1 changes events based on the context it is used
func (m *PostV1ChangesEvents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChangeIdentities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ChangesEvents) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attachments); i++ {

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1ChangesEvents) contextValidateAuthors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Authors); i++ {

		if m.Authors[i] != nil {
			if err := m.Authors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("authors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("authors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1ChangesEvents) contextValidateChangeIdentities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChangeIdentities); i++ {

		if m.ChangeIdentities[i] != nil {
			if err := m.ChangeIdentities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("change_identities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("change_identities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ChangesEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ChangesEvents) UnmarshalBinary(b []byte) error {
	var res PostV1ChangesEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1ChangesEventsAttachmentsItems0 post v1 changes events attachments items0
//
// swagger:model PostV1ChangesEventsAttachmentsItems0
type PostV1ChangesEventsAttachmentsItems0 struct {

	// type
	// Required: true
	// Enum: [link]
	Type *string `json:"type"`
}

// Validate validates this post v1 changes events attachments items0
func (m *PostV1ChangesEventsAttachmentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postV1ChangesEventsAttachmentsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["link"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postV1ChangesEventsAttachmentsItems0TypeTypePropEnum = append(postV1ChangesEventsAttachmentsItems0TypeTypePropEnum, v)
	}
}

const (

	// PostV1ChangesEventsAttachmentsItems0TypeLink captures enum value "link"
	PostV1ChangesEventsAttachmentsItems0TypeLink string = "link"
)

// prop value enum
func (m *PostV1ChangesEventsAttachmentsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postV1ChangesEventsAttachmentsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostV1ChangesEventsAttachmentsItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 changes events attachments items0 based on context it is used
func (m *PostV1ChangesEventsAttachmentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ChangesEventsAttachmentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ChangesEventsAttachmentsItems0) UnmarshalBinary(b []byte) error {
	var res PostV1ChangesEventsAttachmentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1ChangesEventsAuthorsItems0 post v1 changes events authors items0
//
// swagger:model PostV1ChangesEventsAuthorsItems0
type PostV1ChangesEventsAuthorsItems0 struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// source
	// Required: true
	Source *string `json:"source"`

	// source id
	// Required: true
	SourceID *string `json:"source_id"`
}

// Validate validates this post v1 changes events authors items0
func (m *PostV1ChangesEventsAuthorsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ChangesEventsAuthorsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ChangesEventsAuthorsItems0) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ChangesEventsAuthorsItems0) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("source_id", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 changes events authors items0 based on context it is used
func (m *PostV1ChangesEventsAuthorsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ChangesEventsAuthorsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ChangesEventsAuthorsItems0) UnmarshalBinary(b []byte) error {
	var res PostV1ChangesEventsAuthorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1ChangesEventsChangeIdentitiesItems0 post v1 changes events change identities items0
//
// swagger:model PostV1ChangesEventsChangeIdentitiesItems0
type PostV1ChangesEventsChangeIdentitiesItems0 struct {

	// type
	// Required: true
	Type *string `json:"type"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this post v1 changes events change identities items0
func (m *PostV1ChangesEventsChangeIdentitiesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ChangesEventsChangeIdentitiesItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ChangesEventsChangeIdentitiesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 changes events change identities items0 based on context it is used
func (m *PostV1ChangesEventsChangeIdentitiesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ChangesEventsChangeIdentitiesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ChangesEventsChangeIdentitiesItems0) UnmarshalBinary(b []byte) error {
	var res PostV1ChangesEventsChangeIdentitiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
