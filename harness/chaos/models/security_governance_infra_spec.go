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

// SecurityGovernanceInfraSpec security governance infra spec
//
// swagger:model security_governance.InfraSpec
type SecurityGovernanceInfraSpec struct {

	// infra ids
	InfraIds []string `json:"infraIds"`

	// operator
	Operator SecurityGovernanceOperator `json:"operator,omitempty"`
}

// Validate validates this security governance infra spec
func (m *SecurityGovernanceInfraSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGovernanceInfraSpec) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	if err := m.Operator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operator")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operator")
		}
		return err
	}

	return nil
}

// ContextValidate validate this security governance infra spec based on the context it is used
func (m *SecurityGovernanceInfraSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGovernanceInfraSpec) contextValidateOperator(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	if err := m.Operator.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operator")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operator")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGovernanceInfraSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGovernanceInfraSpec) UnmarshalBinary(b []byte) error {
	var res SecurityGovernanceInfraSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
