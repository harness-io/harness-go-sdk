// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExperimentSaveChaosExperimentResponse experiment save chaos experiment response
//
// swagger:model experiment.SaveChaosExperimentResponse
type ExperimentSaveChaosExperimentResponse struct {

	// audit action
	AuditAction string `json:"auditAction,omitempty"`

	// response
	Response string `json:"response,omitempty"`
}

// Validate validates this experiment save chaos experiment response
func (m *ExperimentSaveChaosExperimentResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this experiment save chaos experiment response based on context it is used
func (m *ExperimentSaveChaosExperimentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExperimentSaveChaosExperimentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExperimentSaveChaosExperimentResponse) UnmarshalBinary(b []byte) error {
	var res ExperimentSaveChaosExperimentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
