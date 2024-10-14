// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesExperimentV2 types experiment v2
//
// swagger:model types.ExperimentV2
type TypesExperimentV2 struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// created by
	CreatedBy *GithubComHarnessHceSaasGraphqlServerGraphModelUserDetails `json:"createdBy,omitempty"`

	// cron syntax
	CronSyntax string `json:"cronSyntax,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// events metadata
	EventsMetadata []*ModelEventMetadata `json:"eventsMetadata"`

	// identifiers
	Identifiers *GithubComHarnessHceSaasGraphqlServerGraphModelIdentifiers `json:"identifiers,omitempty"`

	// infra
	Infra *ChaosInfrastructureV2ChaosInfraV2 `json:"infra,omitempty"`

	// is cron enabled
	IsCronEnabled bool `json:"isCronEnabled,omitempty"`

	// is custom workflow
	IsCustomWorkflow bool `json:"isCustomWorkflow,omitempty"`

	// is removed
	IsRemoved bool `json:"isRemoved,omitempty"`

	// last executed at
	LastExecutedAt string `json:"lastExecutedAt,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// recent workflow run details
	RecentWorkflowRunDetails []*ModelRecentWorkflowRun `json:"recentWorkflowRunDetails"`

	// recommendation
	Recommendation *ChaosExperimentRecommendation `json:"recommendation,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// target network service
	TargetNetworkService []*TargetserviceServiceMetadata `json:"targetNetworkService"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`

	// updated by
	UpdatedBy *GithubComHarnessHceSaasGraphqlServerGraphModelUserDetails `json:"updatedBy,omitempty"`

	// weightages
	Weightages []*TypesWeightages `json:"weightages"`

	// workflow ID
	WorkflowID string `json:"workflowID,omitempty"`

	// workflow manifest
	WorkflowManifest string `json:"workflowManifest,omitempty"`

	// workflow type
	WorkflowType string `json:"workflowType,omitempty"`
}

// Validate validates this types experiment v2
func (m *TypesExperimentV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventsMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfra(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecentWorkflowRunDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecommendation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetNetworkService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesExperimentV2) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *TypesExperimentV2) validateEventsMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.EventsMetadata) { // not required
		return nil
	}

	for i := 0; i < len(m.EventsMetadata); i++ {
		if swag.IsZero(m.EventsMetadata[i]) { // not required
			continue
		}

		if m.EventsMetadata[i] != nil {
			if err := m.EventsMetadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eventsMetadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("eventsMetadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TypesExperimentV2) validateIdentifiers(formats strfmt.Registry) error {
	if swag.IsZero(m.Identifiers) { // not required
		return nil
	}

	if m.Identifiers != nil {
		if err := m.Identifiers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identifiers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identifiers")
			}
			return err
		}
	}

	return nil
}

func (m *TypesExperimentV2) validateInfra(formats strfmt.Registry) error {
	if swag.IsZero(m.Infra) { // not required
		return nil
	}

	if m.Infra != nil {
		if err := m.Infra.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infra")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infra")
			}
			return err
		}
	}

	return nil
}

func (m *TypesExperimentV2) validateRecentWorkflowRunDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.RecentWorkflowRunDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.RecentWorkflowRunDetails); i++ {
		if swag.IsZero(m.RecentWorkflowRunDetails[i]) { // not required
			continue
		}

		if m.RecentWorkflowRunDetails[i] != nil {
			if err := m.RecentWorkflowRunDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recentWorkflowRunDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recentWorkflowRunDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TypesExperimentV2) validateRecommendation(formats strfmt.Registry) error {
	if swag.IsZero(m.Recommendation) { // not required
		return nil
	}

	if m.Recommendation != nil {
		if err := m.Recommendation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recommendation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recommendation")
			}
			return err
		}
	}

	return nil
}

func (m *TypesExperimentV2) validateTargetNetworkService(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetNetworkService) { // not required
		return nil
	}

	for i := 0; i < len(m.TargetNetworkService); i++ {
		if swag.IsZero(m.TargetNetworkService[i]) { // not required
			continue
		}

		if m.TargetNetworkService[i] != nil {
			if err := m.TargetNetworkService[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targetNetworkService" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("targetNetworkService" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TypesExperimentV2) validateUpdatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedBy) { // not required
		return nil
	}

	if m.UpdatedBy != nil {
		if err := m.UpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updatedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *TypesExperimentV2) validateWeightages(formats strfmt.Registry) error {
	if swag.IsZero(m.Weightages) { // not required
		return nil
	}

	for i := 0; i < len(m.Weightages); i++ {
		if swag.IsZero(m.Weightages[i]) { // not required
			continue
		}

		if m.Weightages[i] != nil {
			if err := m.Weightages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weightages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("weightages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this types experiment v2 based on the context it is used
func (m *TypesExperimentV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventsMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInfra(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecentWorkflowRunDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecommendation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetNetworkService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeightages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesExperimentV2) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *TypesExperimentV2) contextValidateEventsMetadata(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EventsMetadata); i++ {

		if m.EventsMetadata[i] != nil {

			if swag.IsZero(m.EventsMetadata[i]) { // not required
				return nil
			}

			if err := m.EventsMetadata[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eventsMetadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("eventsMetadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TypesExperimentV2) contextValidateIdentifiers(ctx context.Context, formats strfmt.Registry) error {

	if m.Identifiers != nil {

		if swag.IsZero(m.Identifiers) { // not required
			return nil
		}

		if err := m.Identifiers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identifiers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identifiers")
			}
			return err
		}
	}

	return nil
}

func (m *TypesExperimentV2) contextValidateInfra(ctx context.Context, formats strfmt.Registry) error {

	if m.Infra != nil {

		if swag.IsZero(m.Infra) { // not required
			return nil
		}

		if err := m.Infra.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infra")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infra")
			}
			return err
		}
	}

	return nil
}

func (m *TypesExperimentV2) contextValidateRecentWorkflowRunDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RecentWorkflowRunDetails); i++ {

		if m.RecentWorkflowRunDetails[i] != nil {

			if swag.IsZero(m.RecentWorkflowRunDetails[i]) { // not required
				return nil
			}

			if err := m.RecentWorkflowRunDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recentWorkflowRunDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recentWorkflowRunDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TypesExperimentV2) contextValidateRecommendation(ctx context.Context, formats strfmt.Registry) error {

	if m.Recommendation != nil {

		if swag.IsZero(m.Recommendation) { // not required
			return nil
		}

		if err := m.Recommendation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recommendation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recommendation")
			}
			return err
		}
	}

	return nil
}

func (m *TypesExperimentV2) contextValidateTargetNetworkService(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TargetNetworkService); i++ {

		if m.TargetNetworkService[i] != nil {

			if swag.IsZero(m.TargetNetworkService[i]) { // not required
				return nil
			}

			if err := m.TargetNetworkService[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targetNetworkService" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("targetNetworkService" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TypesExperimentV2) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedBy != nil {

		if swag.IsZero(m.UpdatedBy) { // not required
			return nil
		}

		if err := m.UpdatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updatedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *TypesExperimentV2) contextValidateWeightages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Weightages); i++ {

		if m.Weightages[i] != nil {

			if swag.IsZero(m.Weightages[i]) { // not required
				return nil
			}

			if err := m.Weightages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weightages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("weightages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TypesExperimentV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesExperimentV2) UnmarshalBinary(b []byte) error {
	var res TypesExperimentV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
