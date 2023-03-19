// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gotify/go-api-client/v2/models"
)

// CreateClientReader is a Reader for the CreateClient structure.
type CreateClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateClientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateClientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateClientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateClientOK creates a CreateClientOK with default headers values
func NewCreateClientOK() *CreateClientOK {
	return &CreateClientOK{}
}

/*CreateClientOK handles this case with default header values.

Ok
*/
type CreateClientOK struct {
	Payload *models.Client
}

func (o *CreateClientOK) Error() string {
	return fmt.Sprintf("[POST /client][%d] createClientOK  %+v", 200, o.Payload)
}

func (o *CreateClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Client)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClientBadRequest creates a CreateClientBadRequest with default headers values
func NewCreateClientBadRequest() *CreateClientBadRequest {
	return &CreateClientBadRequest{}
}

/*CreateClientBadRequest handles this case with default header values.

Bad Request
*/
type CreateClientBadRequest struct {
	Payload *models.Error
}

func (o *CreateClientBadRequest) Error() string {
	return fmt.Sprintf("[POST /client][%d] createClientBadRequest  %+v", 400, o.Payload)
}

func (o *CreateClientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClientUnauthorized creates a CreateClientUnauthorized with default headers values
func NewCreateClientUnauthorized() *CreateClientUnauthorized {
	return &CreateClientUnauthorized{}
}

/*CreateClientUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateClientUnauthorized struct {
	Payload *models.Error
}

func (o *CreateClientUnauthorized) Error() string {
	return fmt.Sprintf("[POST /client][%d] createClientUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateClientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClientForbidden creates a CreateClientForbidden with default headers values
func NewCreateClientForbidden() *CreateClientForbidden {
	return &CreateClientForbidden{}
}

/*CreateClientForbidden handles this case with default header values.

Forbidden
*/
type CreateClientForbidden struct {
	Payload *models.Error
}

func (o *CreateClientForbidden) Error() string {
	return fmt.Sprintf("[POST /client][%d] createClientForbidden  %+v", 403, o.Payload)
}

func (o *CreateClientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
