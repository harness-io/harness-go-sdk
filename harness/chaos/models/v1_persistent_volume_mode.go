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

// V1PersistentVolumeMode v1 persistent volume mode
//
// swagger:model v1.PersistentVolumeMode
type V1PersistentVolumeMode string

func NewV1PersistentVolumeMode(value V1PersistentVolumeMode) *V1PersistentVolumeMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1PersistentVolumeMode.
func (m V1PersistentVolumeMode) Pointer() *V1PersistentVolumeMode {
	return &m
}

const (

	// V1PersistentVolumeModeBlock captures enum value "Block"
	V1PersistentVolumeModeBlock V1PersistentVolumeMode = "Block"

	// V1PersistentVolumeModeFilesystem captures enum value "Filesystem"
	V1PersistentVolumeModeFilesystem V1PersistentVolumeMode = "Filesystem"
)

// for schema
var v1PersistentVolumeModeEnum []interface{}

func init() {
	var res []V1PersistentVolumeMode
	if err := json.Unmarshal([]byte(`["Block","Filesystem"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1PersistentVolumeModeEnum = append(v1PersistentVolumeModeEnum, v)
	}
}

func (m V1PersistentVolumeMode) validateV1PersistentVolumeModeEnum(path, location string, value V1PersistentVolumeMode) error {
	if err := validate.EnumCase(path, location, value, v1PersistentVolumeModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 persistent volume mode
func (m V1PersistentVolumeMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1PersistentVolumeModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 persistent volume mode based on context it is used
func (m V1PersistentVolumeMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
