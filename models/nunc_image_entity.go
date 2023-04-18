// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NuncImageEntity nunc image entity
//
// swagger:model NuncImageEntity
type NuncImageEntity struct {

	// original url
	OriginalURL string `json:"original_url,omitempty"`

	// An object with keys that can be of type logo, favicon, cover_image, or open_graph_image
	VersionsUrls interface{} `json:"versions_urls,omitempty"`
}

// Validate validates this nunc image entity
func (m *NuncImageEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nunc image entity based on context it is used
func (m *NuncImageEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NuncImageEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NuncImageEntity) UnmarshalBinary(b []byte) error {
	var res NuncImageEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}