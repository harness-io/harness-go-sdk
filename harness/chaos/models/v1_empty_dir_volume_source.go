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

// V1EmptyDirVolumeSource v1 empty dir volume source
//
// swagger:model v1.EmptyDirVolumeSource
type V1EmptyDirVolumeSource struct {

	// What type of storage medium should back this directory.
	// The default is "" which means to use the node's default medium.
	// Must be an empty string (default) or Memory.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// +optional
	Medium struct {
		V1StorageMedium
	} `json:"medium,omitempty"`

	// Total amount of local storage required for this EmptyDir volume.
	// The size limit is also applicable for memory medium.
	// The maximum usage on memory medium EmptyDir would be the minimum value between
	// the SizeLimit specified here and the sum of memory limits of all containers in a pod.
	// The default is nil which means that the limit is undefined.
	// More info: http://kubernetes.io/docs/user-guide/volumes#emptydir
	// +optional
	SizeLimit struct {
		ResourceQuantity
	} `json:"sizeLimit,omitempty"`
}

// Validate validates this v1 empty dir volume source
func (m *V1EmptyDirVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMedium(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizeLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EmptyDirVolumeSource) validateMedium(formats strfmt.Registry) error {
	if swag.IsZero(m.Medium) { // not required
		return nil
	}

	return nil
}

func (m *V1EmptyDirVolumeSource) validateSizeLimit(formats strfmt.Registry) error {
	if swag.IsZero(m.SizeLimit) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this v1 empty dir volume source based on the context it is used
func (m *V1EmptyDirVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMedium(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSizeLimit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EmptyDirVolumeSource) contextValidateMedium(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1EmptyDirVolumeSource) contextValidateSizeLimit(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *V1EmptyDirVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EmptyDirVolumeSource) UnmarshalBinary(b []byte) error {
	var res V1EmptyDirVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
