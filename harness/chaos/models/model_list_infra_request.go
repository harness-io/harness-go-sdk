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

// ModelListInfraRequest model list infra request
//
// swagger:model model.ListInfraRequest
type ModelListInfraRequest struct {

	// Environment ID
	EnvironmentIDs []string `json:"environmentIDs"`

	// Details for fetching filtered data
	Filter struct {
		GithubComHarnessHceSaasGraphqlServerGraphModelInfraFilterInput
	} `json:"filter,omitempty"`

	// Array of infra IDs for which details will be fetched
	InfraIDs []string `json:"infraIDs"`

	// Connector ID
	K8sConnectorIDs []string `json:"k8sConnectorIDs"`

	// Details for fetching paginated data
	Pagination struct {
		GithubComHarnessHceSaasGraphqlServerGraphModelPagination
	} `json:"pagination,omitempty"`
}

// Validate validates this model list infra request
func (m *ModelListInfraRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelListInfraRequest) validateFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	return nil
}

func (m *ModelListInfraRequest) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this model list infra request based on the context it is used
func (m *ModelListInfraRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelListInfraRequest) contextValidateFilter(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ModelListInfraRequest) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *ModelListInfraRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelListInfraRequest) UnmarshalBinary(b []byte) error {
	var res ModelListInfraRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
