// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatabaseCronConfig database cron config
//
// swagger:model database.CronConfig
type DatabaseCronConfig struct {

	// expression
	Expression string `json:"expression,omitempty"`
}

// Validate validates this database cron config
func (m *DatabaseCronConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database cron config based on context it is used
func (m *DatabaseCronConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseCronConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseCronConfig) UnmarshalBinary(b []byte) error {
	var res DatabaseCronConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
