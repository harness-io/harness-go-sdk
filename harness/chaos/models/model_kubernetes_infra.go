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

// ModelKubernetesInfra model kubernetes infra
//
// swagger:model model.KubernetesInfra
type ModelKubernetesInfra struct {

	// Cluster type Indicates the type on infrastructure (Kubernetes/openshift)
	ClusterType struct {
		ModelClusterType
	} `json:"clusterType,omitempty"`

	// Timestamp when the infra was created
	CreatedAt string `json:"createdAt,omitempty"`

	// User who created the infra
	CreatedBy struct {
		GithubComHarnessHceSaasGraphqlServerGraphModelUserDetails
	} `json:"createdBy,omitempty"`

	// Description of the infra
	Description string `json:"description,omitempty"`

	// Environment ID for the infra
	EnvironmentID string `json:"environmentID,omitempty"`

	// ID of the infra
	InfraID string `json:"infraID,omitempty"`

	// Namespace where the infra is being installed
	InfraNamespace string `json:"infraNamespace,omitempty"`

	// Bool value indicating whether infra ns used already exists on infra or not
	InfraNsExists bool `json:"infraNsExists,omitempty"`

	// Bool value indicating whether service account used already exists on infra or not
	InfraSaExists bool `json:"infraSaExists,omitempty"`

	// Scope of the infra : ns or cluster
	InfraScope struct {
		ModelInfraScope
	} `json:"infraScope,omitempty"`

	// Type of the infrastructure
	InfraType struct {
		ModelInfrastructureType
	} `json:"infraType,omitempty"`

	// InstallationType connector/manifest
	InstallationType struct {
		ModelInstallationType
	} `json:"installationType,omitempty"`

	// Boolean value indicating if chaos infrastructure is active or not
	IsActive bool `json:"isActive,omitempty"`

	// Boolean value indicating if chaos infrastructure is confirmed or not
	IsInfraConfirmed bool `json:"isInfraConfirmed,omitempty"`

	// Boolean value indicating if chaos infrastructure is removed or not
	IsRemoved bool `json:"isRemoved,omitempty"`

	// Tune secret for infra
	IsSecretEnabled bool `json:"isSecretEnabled,omitempty"`

	// K8sConnectorID
	K8sConnectorID string `json:"k8sConnectorID,omitempty"`

	// Last Heartbeat status sent by the infra
	LastHeartbeat string `json:"lastHeartbeat,omitempty"`

	// Timestamp of the last workflow run in the infra
	LastWorkflowTimestamp string `json:"lastWorkflowTimestamp,omitempty"`

	// Name of the infra
	Name string `json:"name,omitempty"`

	// Number of schedules created in the infra
	NoOfSchedules int64 `json:"noOfSchedules,omitempty"`

	// Number of workflows run in the infra
	NoOfWorkflows int64 `json:"noOfWorkflows,omitempty"`

	// Infra Platform Name eg. GKE,AWS, Others
	PlatformName string `json:"platformName,omitempty"`

	// set the user group for security context in pod
	RunAsGroup int64 `json:"runAsGroup,omitempty"`

	// set the user for security context in pod
	RunAsUser int64 `json:"runAsUser,omitempty"`

	// Name of service account used by infra
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// Timestamp when the infra got connected
	StartTime string `json:"startTime,omitempty"`

	// Tags of the infra
	Tags []string `json:"tags"`

	// Token used to verify and retrieve the infra manifest
	Token string `json:"token,omitempty"`

	// update status of infra
	UpdateStatus struct {
		ModelUpdateStatus
	} `json:"updateStatus,omitempty"`

	// Timestamp when the infra was last updated
	UpdatedAt string `json:"updatedAt,omitempty"`

	// User who has updated the infra
	UpdatedBy struct {
		GithubComHarnessHceSaasGraphqlServerGraphModelUserDetails
	} `json:"updatedBy,omitempty"`

	// Upgrade struct for the chaos infrastructure
	Upgrade struct {
		ModelUpgrade
	} `json:"upgrade,omitempty"`

	// Version of the infra
	Version string `json:"version,omitempty"`
}

// Validate validates this model kubernetes infra
func (m *ModelKubernetesInfra) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfraScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfraType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgrade(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelKubernetesInfra) validateClusterType(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterType) { // not required
		return nil
	}

	return nil
}

func (m *ModelKubernetesInfra) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	return nil
}

func (m *ModelKubernetesInfra) validateInfraScope(formats strfmt.Registry) error {
	if swag.IsZero(m.InfraScope) { // not required
		return nil
	}

	return nil
}

func (m *ModelKubernetesInfra) validateInfraType(formats strfmt.Registry) error {
	if swag.IsZero(m.InfraType) { // not required
		return nil
	}

	return nil
}

func (m *ModelKubernetesInfra) validateInstallationType(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallationType) { // not required
		return nil
	}

	return nil
}

func (m *ModelKubernetesInfra) validateUpdateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateStatus) { // not required
		return nil
	}

	return nil
}

func (m *ModelKubernetesInfra) validateUpdatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedBy) { // not required
		return nil
	}

	return nil
}

func (m *ModelKubernetesInfra) validateUpgrade(formats strfmt.Registry) error {
	if swag.IsZero(m.Upgrade) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this model kubernetes infra based on the context it is used
func (m *ModelKubernetesInfra) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInfraScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInfraType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpgrade(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelKubernetesInfra) contextValidateClusterType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ModelKubernetesInfra) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ModelKubernetesInfra) contextValidateInfraScope(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ModelKubernetesInfra) contextValidateInfraType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ModelKubernetesInfra) contextValidateInstallationType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ModelKubernetesInfra) contextValidateUpdateStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ModelKubernetesInfra) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ModelKubernetesInfra) contextValidateUpgrade(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *ModelKubernetesInfra) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelKubernetesInfra) UnmarshalBinary(b []byte) error {
	var res ModelKubernetesInfra
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
