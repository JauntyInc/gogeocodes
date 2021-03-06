// Code generated by go-swagger; DO NOT EDIT.

package zip_code

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewV1AddressReverseZIPCodeParams creates a new V1AddressReverseZIPCodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1AddressReverseZIPCodeParams() *V1AddressReverseZIPCodeParams {
	return &V1AddressReverseZIPCodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1AddressReverseZIPCodeParamsWithTimeout creates a new V1AddressReverseZIPCodeParams object
// with the ability to set a timeout on a request.
func NewV1AddressReverseZIPCodeParamsWithTimeout(timeout time.Duration) *V1AddressReverseZIPCodeParams {
	return &V1AddressReverseZIPCodeParams{
		timeout: timeout,
	}
}

// NewV1AddressReverseZIPCodeParamsWithContext creates a new V1AddressReverseZIPCodeParams object
// with the ability to set a context for a request.
func NewV1AddressReverseZIPCodeParamsWithContext(ctx context.Context) *V1AddressReverseZIPCodeParams {
	return &V1AddressReverseZIPCodeParams{
		Context: ctx,
	}
}

// NewV1AddressReverseZIPCodeParamsWithHTTPClient creates a new V1AddressReverseZIPCodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1AddressReverseZIPCodeParamsWithHTTPClient(client *http.Client) *V1AddressReverseZIPCodeParams {
	return &V1AddressReverseZIPCodeParams{
		HTTPClient: client,
	}
}

/* V1AddressReverseZIPCodeParams contains all the parameters to send to the API endpoint
   for the v1 address reverse zip code operation.

   Typically these are written to a http.Request.
*/
type V1AddressReverseZIPCodeParams struct {

	/* Latitude.

	   The latitude of the query
	*/
	Latitude float64

	/* Longitude.

	   The longitude of the query
	*/
	Longitude float64

	/* Strict.

	   If true, will only return a result if the given point is within a zip code tabulation area, and will not return the closest one.
	*/
	Strict *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 address reverse zip code params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1AddressReverseZIPCodeParams) WithDefaults() *V1AddressReverseZIPCodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 address reverse zip code params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1AddressReverseZIPCodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) WithTimeout(timeout time.Duration) *V1AddressReverseZIPCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) WithContext(ctx context.Context) *V1AddressReverseZIPCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) WithHTTPClient(client *http.Client) *V1AddressReverseZIPCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLatitude adds the latitude to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) WithLatitude(latitude float64) *V1AddressReverseZIPCodeParams {
	o.SetLatitude(latitude)
	return o
}

// SetLatitude adds the latitude to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) SetLatitude(latitude float64) {
	o.Latitude = latitude
}

// WithLongitude adds the longitude to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) WithLongitude(longitude float64) *V1AddressReverseZIPCodeParams {
	o.SetLongitude(longitude)
	return o
}

// SetLongitude adds the longitude to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) SetLongitude(longitude float64) {
	o.Longitude = longitude
}

// WithStrict adds the strict to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) WithStrict(strict *bool) *V1AddressReverseZIPCodeParams {
	o.SetStrict(strict)
	return o
}

// SetStrict adds the strict to the v1 address reverse zip code params
func (o *V1AddressReverseZIPCodeParams) SetStrict(strict *bool) {
	o.Strict = strict
}

// WriteToRequest writes these params to a swagger request
func (o *V1AddressReverseZIPCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param latitude
	qrLatitude := o.Latitude
	qLatitude := swag.FormatFloat64(qrLatitude)
	if qLatitude != "" {

		if err := r.SetQueryParam("latitude", qLatitude); err != nil {
			return err
		}
	}

	// query param longitude
	qrLongitude := o.Longitude
	qLongitude := swag.FormatFloat64(qrLongitude)
	if qLongitude != "" {

		if err := r.SetQueryParam("longitude", qLongitude); err != nil {
			return err
		}
	}

	if o.Strict != nil {

		// query param strict
		var qrStrict bool

		if o.Strict != nil {
			qrStrict = *o.Strict
		}
		qStrict := swag.FormatBool(qrStrict)
		if qStrict != "" {

			if err := r.SetQueryParam("strict", qStrict); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
