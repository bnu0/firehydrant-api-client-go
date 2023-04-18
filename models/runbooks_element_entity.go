// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunbooksElementEntity runbooks element entity
//
// swagger:model Runbooks_ElementEntity
type RunbooksElementEntity struct {

	// dynamic select
	DynamicSelect *RunbooksElementDynamicSelectEntity `json:"dynamic_select,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// input
	Input *RunbooksElementInputEntity `json:"input,omitempty"`

	// markdown
	Markdown *RunbooksElementMarkdownEntity `json:"markdown,omitempty"`

	// plain text
	PlainText *RunbooksElementMarkdownEntity `json:"plain_text,omitempty"`

	// textarea
	Textarea *RunbooksElementTextareaEntity `json:"textarea,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this runbooks element entity
func (m *RunbooksElementEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDynamicSelect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkdown(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlainText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTextarea(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunbooksElementEntity) validateDynamicSelect(formats strfmt.Registry) error {
	if swag.IsZero(m.DynamicSelect) { // not required
		return nil
	}

	if m.DynamicSelect != nil {
		if err := m.DynamicSelect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamic_select")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamic_select")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksElementEntity) validateInput(formats strfmt.Registry) error {
	if swag.IsZero(m.Input) { // not required
		return nil
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksElementEntity) validateMarkdown(formats strfmt.Registry) error {
	if swag.IsZero(m.Markdown) { // not required
		return nil
	}

	if m.Markdown != nil {
		if err := m.Markdown.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("markdown")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("markdown")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksElementEntity) validatePlainText(formats strfmt.Registry) error {
	if swag.IsZero(m.PlainText) { // not required
		return nil
	}

	if m.PlainText != nil {
		if err := m.PlainText.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plain_text")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plain_text")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksElementEntity) validateTextarea(formats strfmt.Registry) error {
	if swag.IsZero(m.Textarea) { // not required
		return nil
	}

	if m.Textarea != nil {
		if err := m.Textarea.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("textarea")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("textarea")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this runbooks element entity based on the context it is used
func (m *RunbooksElementEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDynamicSelect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMarkdown(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlainText(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTextarea(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunbooksElementEntity) contextValidateDynamicSelect(ctx context.Context, formats strfmt.Registry) error {

	if m.DynamicSelect != nil {
		if err := m.DynamicSelect.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamic_select")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamic_select")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksElementEntity) contextValidateInput(ctx context.Context, formats strfmt.Registry) error {

	if m.Input != nil {
		if err := m.Input.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksElementEntity) contextValidateMarkdown(ctx context.Context, formats strfmt.Registry) error {

	if m.Markdown != nil {
		if err := m.Markdown.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("markdown")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("markdown")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksElementEntity) contextValidatePlainText(ctx context.Context, formats strfmt.Registry) error {

	if m.PlainText != nil {
		if err := m.PlainText.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plain_text")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plain_text")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksElementEntity) contextValidateTextarea(ctx context.Context, formats strfmt.Registry) error {

	if m.Textarea != nil {
		if err := m.Textarea.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("textarea")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("textarea")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunbooksElementEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunbooksElementEntity) UnmarshalBinary(b []byte) error {
	var res RunbooksElementEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}