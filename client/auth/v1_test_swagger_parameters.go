// Code generated by go-swagger; DO NOT EDIT.

package auth

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
)

// NewV1TestParams creates a new V1TestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1TestParams() *V1TestParams {
	return &V1TestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1TestParamsWithTimeout creates a new V1TestParams object
// with the ability to set a timeout on a request.
func NewV1TestParamsWithTimeout(timeout time.Duration) *V1TestParams {
	return &V1TestParams{
		timeout: timeout,
	}
}

// NewV1TestParamsWithContext creates a new V1TestParams object
// with the ability to set a context for a request.
func NewV1TestParamsWithContext(ctx context.Context) *V1TestParams {
	return &V1TestParams{
		Context: ctx,
	}
}

// NewV1TestParamsWithHTTPClient creates a new V1TestParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1TestParamsWithHTTPClient(client *http.Client) *V1TestParams {
	return &V1TestParams{
		HTTPClient: client,
	}
}

/* V1TestParams contains all the parameters to send to the API endpoint
   for the v1 test operation.

   Typically these are written to a http.Request.
*/
type V1TestParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1TestParams) WithDefaults() *V1TestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1TestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 test params
func (o *V1TestParams) WithTimeout(timeout time.Duration) *V1TestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 test params
func (o *V1TestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 test params
func (o *V1TestParams) WithContext(ctx context.Context) *V1TestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 test params
func (o *V1TestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 test params
func (o *V1TestParams) WithHTTPClient(client *http.Client) *V1TestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 test params
func (o *V1TestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1TestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
