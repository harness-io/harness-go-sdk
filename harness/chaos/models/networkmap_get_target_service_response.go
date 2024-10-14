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

// NetworkmapGetTargetServiceResponse networkmap get target service response
//
// swagger:model networkmap.GetTargetServiceResponse
type NetworkmapGetTargetServiceResponse struct {

	// chaos details
	ChaosDetails *TargetserviceTargetService `json:"chaosDetails,omitempty"`

	// discovery details
	DiscoveryDetails *DatabaseServiceCollection `json:"discoveryDetails,omitempty"`

	// workload details
	WorkloadDetails *DatabaseDiscoveredServiceKubernetesSpec `json:"workloadDetails,omitempty"`
}

// Validate validates this networkmap get target service response
func (m *NetworkmapGetTargetServiceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChaosDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscoveryDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkmapGetTargetServiceResponse) validateChaosDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ChaosDetails) { // not required
		return nil
	}

	if m.ChaosDetails != nil {
		if err := m.ChaosDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chaosDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chaosDetails")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkmapGetTargetServiceResponse) validateDiscoveryDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscoveryDetails) { // not required
		return nil
	}

	if m.DiscoveryDetails != nil {
		if err := m.DiscoveryDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discoveryDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discoveryDetails")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkmapGetTargetServiceResponse) validateWorkloadDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkloadDetails) { // not required
		return nil
	}

	if m.WorkloadDetails != nil {
		if err := m.WorkloadDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workloadDetails")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this networkmap get target service response based on the context it is used
func (m *NetworkmapGetTargetServiceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChaosDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiscoveryDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkloadDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkmapGetTargetServiceResponse) contextValidateChaosDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ChaosDetails != nil {

		if swag.IsZero(m.ChaosDetails) { // not required
			return nil
		}

		if err := m.ChaosDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chaosDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chaosDetails")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkmapGetTargetServiceResponse) contextValidateDiscoveryDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.DiscoveryDetails != nil {

		if swag.IsZero(m.DiscoveryDetails) { // not required
			return nil
		}

		if err := m.DiscoveryDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discoveryDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discoveryDetails")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkmapGetTargetServiceResponse) contextValidateWorkloadDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkloadDetails != nil {

		if swag.IsZero(m.WorkloadDetails) { // not required
			return nil
		}

		if err := m.WorkloadDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workloadDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkmapGetTargetServiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkmapGetTargetServiceResponse) UnmarshalBinary(b []byte) error {
	var res NetworkmapGetTargetServiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
