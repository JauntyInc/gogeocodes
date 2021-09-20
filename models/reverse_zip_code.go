// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReverseZIPCode reverse ZIP code
//
// swagger:model ReverseZIPCode
type ReverseZIPCode struct {

	// The distance in meters that the point was from the ZIP Code. If it was within the ZIP code, this will be 0. Calculated using the distance between the input point and the ZIP Code's geometry.
	// Example: 0
	Distance float64 `json:"distance,omitempty"`

	// The result object containing the ZIP Code. Will always be present if a lookup was successful.
	// Required: true
	ZIPCode *ZIPCode `json:"zip_code"`
}

// Validate validates this reverse ZIP code
func (m *ReverseZIPCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateZIPCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReverseZIPCode) validateZIPCode(formats strfmt.Registry) error {

	if err := validate.Required("zip_code", "body", m.ZIPCode); err != nil {
		return err
	}

	if m.ZIPCode != nil {
		if err := m.ZIPCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zip_code")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this reverse ZIP code based on the context it is used
func (m *ReverseZIPCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateZIPCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReverseZIPCode) contextValidateZIPCode(ctx context.Context, formats strfmt.Registry) error {

	if m.ZIPCode != nil {
		if err := m.ZIPCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zip_code")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReverseZIPCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReverseZIPCode) UnmarshalBinary(b []byte) error {
	var res ReverseZIPCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
