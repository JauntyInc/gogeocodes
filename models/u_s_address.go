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

// USAddress u s address
//
// swagger:model USAddress
type USAddress struct {

	// The city of the address.
	// Example: New York
	// Required: true
	City *string `json:"city"`

	// A fully-formatted version of the address.
	// Example: 111 8th Ave\nNew York, NY 10011
	// Required: true
	Formatted *string `json:"formatted"`

	// The 2-letter state code of the address. Note that in addition to the 50 states, you might also see 2 letter codes for American territories, or the District of Columbia.
	// Example: NY
	// Required: true
	StateCode *string `json:"state_code"`

	// The structured street address parsed from the input.
	// Required: true
	StreetAddress *USStreetAddress `json:"street_address"`

	// The ZIP code and optional plus-four code for the address.
	// Required: true
	ZIPCode *ZIPCode `json:"zip_code"`
}

// Validate validates this u s address
func (m *USAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormatted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreetAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZIPCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *USAddress) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *USAddress) validateFormatted(formats strfmt.Registry) error {

	if err := validate.Required("formatted", "body", m.Formatted); err != nil {
		return err
	}

	return nil
}

func (m *USAddress) validateStateCode(formats strfmt.Registry) error {

	if err := validate.Required("state_code", "body", m.StateCode); err != nil {
		return err
	}

	return nil
}

func (m *USAddress) validateStreetAddress(formats strfmt.Registry) error {

	if err := validate.Required("street_address", "body", m.StreetAddress); err != nil {
		return err
	}

	if m.StreetAddress != nil {
		if err := m.StreetAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("street_address")
			}
			return err
		}
	}

	return nil
}

func (m *USAddress) validateZIPCode(formats strfmt.Registry) error {

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

// ContextValidate validate this u s address based on the context it is used
func (m *USAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStreetAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZIPCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *USAddress) contextValidateStreetAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.StreetAddress != nil {
		if err := m.StreetAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("street_address")
			}
			return err
		}
	}

	return nil
}

func (m *USAddress) contextValidateZIPCode(ctx context.Context, formats strfmt.Registry) error {

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
func (m *USAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *USAddress) UnmarshalBinary(b []byte) error {
	var res USAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
