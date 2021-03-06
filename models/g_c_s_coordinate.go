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

// GCSCoordinate g c s coordinate
//
// swagger:model GCSCoordinate
type GCSCoordinate struct {

	// latitude specifies the North/South position of a coordinate, as measured from the equator.
	// Example: 40.740776
	// Maximum: 90
	// Minimum: -90
	Latitude *float64 `json:"latitude,omitempty"`

	// longitude specifies the East/West position of a coordinate, as measured from the International Reference Meridian.
	// Example: -74.001852
	// Maximum: 180
	// Minimum: -180
	Longitude *float64 `json:"longitude,omitempty"`
}

// Validate validates this g c s coordinate
func (m *GCSCoordinate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLatitude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLongitude(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GCSCoordinate) validateLatitude(formats strfmt.Registry) error {
	if swag.IsZero(m.Latitude) { // not required
		return nil
	}

	if err := validate.Minimum("latitude", "body", *m.Latitude, -90, false); err != nil {
		return err
	}

	if err := validate.Maximum("latitude", "body", *m.Latitude, 90, false); err != nil {
		return err
	}

	return nil
}

func (m *GCSCoordinate) validateLongitude(formats strfmt.Registry) error {
	if swag.IsZero(m.Longitude) { // not required
		return nil
	}

	if err := validate.Minimum("longitude", "body", *m.Longitude, -180, false); err != nil {
		return err
	}

	if err := validate.Maximum("longitude", "body", *m.Longitude, 180, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this g c s coordinate based on context it is used
func (m *GCSCoordinate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GCSCoordinate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCSCoordinate) UnmarshalBinary(b []byte) error {
	var res GCSCoordinate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
