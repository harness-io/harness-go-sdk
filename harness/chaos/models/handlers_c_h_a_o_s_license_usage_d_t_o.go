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

// HandlersCHAOSLicenseUsageDTO handlers c h a o s license usage d t o
//
// swagger:model handlers.CHAOSLicenseUsageDTO
type HandlersCHAOSLicenseUsageDTO struct {

	// account identifier
	AccountIdentifier string `json:"accountIdentifier,omitempty"`

	// experiment runs per month
	ExperimentRunsPerMonth *HandlersUsageDataDTO `json:"experimentRunsPerMonth,omitempty"`

	// module
	Module string `json:"module,omitempty"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`
}

// Validate validates this handlers c h a o s license usage d t o
func (m *HandlersCHAOSLicenseUsageDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExperimentRunsPerMonth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HandlersCHAOSLicenseUsageDTO) validateExperimentRunsPerMonth(formats strfmt.Registry) error {
	if swag.IsZero(m.ExperimentRunsPerMonth) { // not required
		return nil
	}

	if m.ExperimentRunsPerMonth != nil {
		if err := m.ExperimentRunsPerMonth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("experimentRunsPerMonth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("experimentRunsPerMonth")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this handlers c h a o s license usage d t o based on the context it is used
func (m *HandlersCHAOSLicenseUsageDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExperimentRunsPerMonth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HandlersCHAOSLicenseUsageDTO) contextValidateExperimentRunsPerMonth(ctx context.Context, formats strfmt.Registry) error {

	if m.ExperimentRunsPerMonth != nil {

		if swag.IsZero(m.ExperimentRunsPerMonth) { // not required
			return nil
		}

		if err := m.ExperimentRunsPerMonth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("experimentRunsPerMonth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("experimentRunsPerMonth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HandlersCHAOSLicenseUsageDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HandlersCHAOSLicenseUsageDTO) UnmarshalBinary(b []byte) error {
	var res HandlersCHAOSLicenseUsageDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
