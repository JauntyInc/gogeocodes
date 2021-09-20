// Code generated by go-swagger; DO NOT EDIT.

package zip_code

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/JauntyInc/gogeocodes/models"
)

// V1AddressZIPCodeReader is a Reader for the V1AddressZIPCode structure.
type V1AddressZIPCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AddressZIPCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AddressZIPCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AddressZIPCodeOK creates a V1AddressZIPCodeOK with default headers values
func NewV1AddressZIPCodeOK() *V1AddressZIPCodeOK {
	return &V1AddressZIPCodeOK{}
}

/* V1AddressZIPCodeOK describes a response with status code 200, with default header values.

OK
*/
type V1AddressZIPCodeOK struct {
	Payload *models.ZIPCode
}

func (o *V1AddressZIPCodeOK) Error() string {
	return fmt.Sprintf("[GET /v1/address/zip_code][%d] v1AddressZipCodeOK  %+v", 200, o.Payload)
}
func (o *V1AddressZIPCodeOK) GetPayload() *models.ZIPCode {
	return o.Payload
}

func (o *V1AddressZIPCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZIPCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}