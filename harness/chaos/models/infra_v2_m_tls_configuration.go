// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InfraV2MTLSConfiguration infra v2 m TLS configuration
//
// swagger:model infra_v2.MTLSConfiguration
type InfraV2MTLSConfiguration struct {

	// cert path
	CertPath string `json:"certPath,omitempty"`

	// key path
	KeyPath string `json:"keyPath,omitempty"`

	// secret name
	SecretName string `json:"secretName,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this infra v2 m TLS configuration
func (m *InfraV2MTLSConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this infra v2 m TLS configuration based on context it is used
func (m *InfraV2MTLSConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfraV2MTLSConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraV2MTLSConfiguration) UnmarshalBinary(b []byte) error {
	var res InfraV2MTLSConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
