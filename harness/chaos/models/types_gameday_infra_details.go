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

// TypesGamedayInfraDetails types gameday infra details
//
// swagger:model types.GamedayInfraDetails
type TypesGamedayInfraDetails struct {

	// environment ID
	EnvironmentID string `json:"environmentID,omitempty"`

	// infra ID
	InfraID string `json:"infraID,omitempty"`

	// infra type
	InfraType ModelInfrastructureType `json:"infraType,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this types gameday infra details
func (m *TypesGamedayInfraDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfraType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesGamedayInfraDetails) validateInfraType(formats strfmt.Registry) error {
	if swag.IsZero(m.InfraType) { // not required
		return nil
	}

	if err := m.InfraType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("infraType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("infraType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this types gameday infra details based on the context it is used
func (m *TypesGamedayInfraDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfraType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesGamedayInfraDetails) contextValidateInfraType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.InfraType) { // not required
		return nil
	}

	if err := m.InfraType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("infraType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("infraType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TypesGamedayInfraDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesGamedayInfraDetails) UnmarshalBinary(b []byte) error {
	var res TypesGamedayInfraDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
