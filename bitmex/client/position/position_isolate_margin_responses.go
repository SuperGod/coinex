// Code generated by go-swagger; DO NOT EDIT.

package position

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sumorf/coinex/bitmex/models"
)

// PositionIsolateMarginReader is a Reader for the PositionIsolateMargin structure.
type PositionIsolateMarginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PositionIsolateMarginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPositionIsolateMarginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPositionIsolateMarginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPositionIsolateMarginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPositionIsolateMarginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPositionIsolateMarginOK creates a PositionIsolateMarginOK with default headers values
func NewPositionIsolateMarginOK() *PositionIsolateMarginOK {
	return &PositionIsolateMarginOK{}
}

/*PositionIsolateMarginOK handles this case with default header values.

Request was successful
*/
type PositionIsolateMarginOK struct {
	Payload *models.Position
}

func (o *PositionIsolateMarginOK) Error() string {
	return fmt.Sprintf("[POST /position/isolate][%d] positionIsolateMarginOK  %+v", 200, o.Payload)
}

func (o *PositionIsolateMarginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Position)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPositionIsolateMarginBadRequest creates a PositionIsolateMarginBadRequest with default headers values
func NewPositionIsolateMarginBadRequest() *PositionIsolateMarginBadRequest {
	return &PositionIsolateMarginBadRequest{}
}

/*PositionIsolateMarginBadRequest handles this case with default header values.

Parameter Error
*/
type PositionIsolateMarginBadRequest struct {
	Payload *models.Error
}

func (o *PositionIsolateMarginBadRequest) Error() string {
	return fmt.Sprintf("[POST /position/isolate][%d] positionIsolateMarginBadRequest  %+v", 400, o.Payload)
}

func (o *PositionIsolateMarginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPositionIsolateMarginUnauthorized creates a PositionIsolateMarginUnauthorized with default headers values
func NewPositionIsolateMarginUnauthorized() *PositionIsolateMarginUnauthorized {
	return &PositionIsolateMarginUnauthorized{}
}

/*PositionIsolateMarginUnauthorized handles this case with default header values.

Unauthorized
*/
type PositionIsolateMarginUnauthorized struct {
	Payload *models.Error
}

func (o *PositionIsolateMarginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /position/isolate][%d] positionIsolateMarginUnauthorized  %+v", 401, o.Payload)
}

func (o *PositionIsolateMarginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPositionIsolateMarginNotFound creates a PositionIsolateMarginNotFound with default headers values
func NewPositionIsolateMarginNotFound() *PositionIsolateMarginNotFound {
	return &PositionIsolateMarginNotFound{}
}

/*PositionIsolateMarginNotFound handles this case with default header values.

Not Found
*/
type PositionIsolateMarginNotFound struct {
	Payload *models.Error
}

func (o *PositionIsolateMarginNotFound) Error() string {
	return fmt.Sprintf("[POST /position/isolate][%d] positionIsolateMarginNotFound  %+v", 404, o.Payload)
}

func (o *PositionIsolateMarginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
