// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ObjectFieldSelector v1 object field selector
//
// swagger:model v1.ObjectFieldSelector
type V1ObjectFieldSelector struct {

	// Version of the schema the FieldPath is written in terms of, defaults to "v1".
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`

	// Path of the field to select in the specified API version.
	FieldPath string `json:"fieldPath,omitempty"`
}

// Validate validates this v1 object field selector
func (m *V1ObjectFieldSelector) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 object field selector based on context it is used
func (m *V1ObjectFieldSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ObjectFieldSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ObjectFieldSelector) UnmarshalBinary(b []byte) error {
	var res V1ObjectFieldSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
