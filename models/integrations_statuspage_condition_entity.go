// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IntegrationsStatuspageConditionEntity integrations statuspage condition entity
//
// swagger:model Integrations_Statuspage_ConditionEntity
type IntegrationsStatuspageConditionEntity struct {

	// condition id
	ConditionID string `json:"condition_id,omitempty"`

	// condition name
	ConditionName string `json:"condition_name,omitempty"`

	// statuspageio condition
	// Enum: [operational major_outage minor_outage degraded_performance]
	StatuspageioCondition string `json:"statuspageio_condition,omitempty"`
}

// Validate validates this integrations statuspage condition entity
func (m *IntegrationsStatuspageConditionEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatuspageioCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var integrationsStatuspageConditionEntityTypeStatuspageioConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["operational","major_outage","minor_outage","degraded_performance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		integrationsStatuspageConditionEntityTypeStatuspageioConditionPropEnum = append(integrationsStatuspageConditionEntityTypeStatuspageioConditionPropEnum, v)
	}
}

const (

	// IntegrationsStatuspageConditionEntityStatuspageioConditionOperational captures enum value "operational"
	IntegrationsStatuspageConditionEntityStatuspageioConditionOperational string = "operational"

	// IntegrationsStatuspageConditionEntityStatuspageioConditionMajorOutage captures enum value "major_outage"
	IntegrationsStatuspageConditionEntityStatuspageioConditionMajorOutage string = "major_outage"

	// IntegrationsStatuspageConditionEntityStatuspageioConditionMinorOutage captures enum value "minor_outage"
	IntegrationsStatuspageConditionEntityStatuspageioConditionMinorOutage string = "minor_outage"

	// IntegrationsStatuspageConditionEntityStatuspageioConditionDegradedPerformance captures enum value "degraded_performance"
	IntegrationsStatuspageConditionEntityStatuspageioConditionDegradedPerformance string = "degraded_performance"
)

// prop value enum
func (m *IntegrationsStatuspageConditionEntity) validateStatuspageioConditionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, integrationsStatuspageConditionEntityTypeStatuspageioConditionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IntegrationsStatuspageConditionEntity) validateStatuspageioCondition(formats strfmt.Registry) error {
	if swag.IsZero(m.StatuspageioCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatuspageioConditionEnum("statuspageio_condition", "body", m.StatuspageioCondition); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this integrations statuspage condition entity based on context it is used
func (m *IntegrationsStatuspageConditionEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsStatuspageConditionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsStatuspageConditionEntity) UnmarshalBinary(b []byte) error {
	var res IntegrationsStatuspageConditionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
