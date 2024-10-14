// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InfraV2ProxyConfiguration infra v2 proxy configuration
//
// swagger:model infra_v2.ProxyConfiguration
type InfraV2ProxyConfiguration struct {

	// http proxy
	HTTPProxy string `json:"httpProxy,omitempty"`

	// https proxy
	HTTPSProxy string `json:"httpsProxy,omitempty"`

	// no proxy
	NoProxy string `json:"noProxy,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this infra v2 proxy configuration
func (m *InfraV2ProxyConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this infra v2 proxy configuration based on context it is used
func (m *InfraV2ProxyConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfraV2ProxyConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfraV2ProxyConfiguration) UnmarshalBinary(b []byte) error {
	var res InfraV2ProxyConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
