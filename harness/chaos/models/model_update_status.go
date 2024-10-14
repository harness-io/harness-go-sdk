// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ModelUpdateStatus model update status
//
// swagger:model model.UpdateStatus
type ModelUpdateStatus string

func NewModelUpdateStatus(value ModelUpdateStatus) *ModelUpdateStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ModelUpdateStatus.
func (m ModelUpdateStatus) Pointer() *ModelUpdateStatus {
	return &m
}

const (

	// ModelUpdateStatusAVAILABLE captures enum value "AVAILABLE"
	ModelUpdateStatusAVAILABLE ModelUpdateStatus = "AVAILABLE"

	// ModelUpdateStatusMANDATORY captures enum value "MANDATORY"
	ModelUpdateStatusMANDATORY ModelUpdateStatus = "MANDATORY"

	// ModelUpdateStatusNOTREQUIRED captures enum value "NOT_REQUIRED"
	ModelUpdateStatusNOTREQUIRED ModelUpdateStatus = "NOT_REQUIRED"
)

// for schema
var modelUpdateStatusEnum []interface{}

func init() {
	var res []ModelUpdateStatus
	if err := json.Unmarshal([]byte(`["AVAILABLE","MANDATORY","NOT_REQUIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelUpdateStatusEnum = append(modelUpdateStatusEnum, v)
	}
}

func (m ModelUpdateStatus) validateModelUpdateStatusEnum(path, location string, value ModelUpdateStatus) error {
	if err := validate.EnumCase(path, location, value, modelUpdateStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this model update status
func (m ModelUpdateStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateModelUpdateStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this model update status based on context it is used
func (m ModelUpdateStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
