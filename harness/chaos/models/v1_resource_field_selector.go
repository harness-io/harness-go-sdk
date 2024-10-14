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

// V1ResourceFieldSelector v1 resource field selector
//
// swagger:model v1.ResourceFieldSelector
type V1ResourceFieldSelector struct {

	// Container name: required for volumes, optional for env vars
	// +optional
	ContainerName string `json:"containerName,omitempty"`

	// Specifies the output format of the exposed resources, defaults to "1"
	// +optional
	Divisor struct {
		ResourceQuantity
	} `json:"divisor,omitempty"`

	// Required: resource to select
	Resource string `json:"resource,omitempty"`
}

// Validate validates this v1 resource field selector
func (m *V1ResourceFieldSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDivisor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ResourceFieldSelector) validateDivisor(formats strfmt.Registry) error {
	if swag.IsZero(m.Divisor) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this v1 resource field selector based on the context it is used
func (m *V1ResourceFieldSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDivisor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ResourceFieldSelector) contextValidateDivisor(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *V1ResourceFieldSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ResourceFieldSelector) UnmarshalBinary(b []byte) error {
	var res V1ResourceFieldSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
