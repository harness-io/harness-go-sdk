// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ConfigMapEnvSource v1 config map env source
//
// swagger:model v1.ConfigMapEnvSource
type V1ConfigMapEnvSource struct {

	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// TODO: Add other useful fields. apiVersion, kind, uid?
	// +optional
	Name string `json:"name,omitempty"`

	// Specify whether the ConfigMap must be defined
	// +optional
	Optional bool `json:"optional,omitempty"`
}

// Validate validates this v1 config map env source
func (m *V1ConfigMapEnvSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 config map env source based on context it is used
func (m *V1ConfigMapEnvSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ConfigMapEnvSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ConfigMapEnvSource) UnmarshalBinary(b []byte) error {
	var res V1ConfigMapEnvSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
