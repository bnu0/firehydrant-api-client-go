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

// MetricsMetricsEntity Metrics_MetricsEntity model
//
// swagger:model Metrics_MetricsEntity
type MetricsMetricsEntity struct {

	// The size of returned buckets. Can be one of: day, week, month, or all_time.
	BucketSize int32 `json:"bucket_size,omitempty"`

	// buckets
	Buckets []interface{} `json:"buckets"`

	// The field by which the metrics are grouped. Can be one of: total, severity, priority, functionality, service, environment, or user.
	By string `json:"by,omitempty"`

	// display information
	DisplayInformation interface{} `json:"display_information,omitempty"`

	// keys
	Keys []string `json:"keys"`

	// sort
	Sort *MetricsMetricsEntitySortEntity `json:"sort,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this metrics metrics entity
func (m *MetricsMetricsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsMetricsEntity) validateSort(formats strfmt.Registry) error {
	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	if m.Sort != nil {
		if err := m.Sort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sort")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrics metrics entity based on the context it is used
func (m *MetricsMetricsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsMetricsEntity) contextValidateSort(ctx context.Context, formats strfmt.Registry) error {

	if m.Sort != nil {
		if err := m.Sort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sort")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsMetricsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsMetricsEntity) UnmarshalBinary(b []byte) error {
	var res MetricsMetricsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
