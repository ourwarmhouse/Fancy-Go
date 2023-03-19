// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gotify/go-api-client/v2/models"
)

// GetAppMessagesReader is a Reader for the GetAppMessages structure.
type GetAppMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAppMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAppMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAppMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAppMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAppMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppMessagesOK creates a GetAppMessagesOK with default headers values
func NewGetAppMessagesOK() *GetAppMessagesOK {
	return &GetAppMessagesOK{}
}

/*GetAppMessagesOK handles this case with default header values.

Ok
*/
type GetAppMessagesOK struct {
	Payload *models.PagedMessages
}

func (o *GetAppMessagesOK) Error() string {
	return fmt.Sprintf("[GET /application/{id}/message][%d] getAppMessagesOK  %+v", 200, o.Payload)
}

func (o *GetAppMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PagedMessages)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppMessagesBadRequest creates a GetAppMessagesBadRequest with default headers values
func NewGetAppMessagesBadRequest() *GetAppMessagesBadRequest {
	return &GetAppMessagesBadRequest{}
}

/*GetAppMessagesBadRequest handles this case with default header values.

Bad Request
*/
type GetAppMessagesBadRequest struct {
	Payload *models.Error
}

func (o *GetAppMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /application/{id}/message][%d] getAppMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppMessagesUnauthorized creates a GetAppMessagesUnauthorized with default headers values
func NewGetAppMessagesUnauthorized() *GetAppMessagesUnauthorized {
	return &GetAppMessagesUnauthorized{}
}

/*GetAppMessagesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAppMessagesUnauthorized struct {
	Payload *models.Error
}

func (o *GetAppMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /application/{id}/message][%d] getAppMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppMessagesForbidden creates a GetAppMessagesForbidden with default headers values
func NewGetAppMessagesForbidden() *GetAppMessagesForbidden {
	return &GetAppMessagesForbidden{}
}

/*GetAppMessagesForbidden handles this case with default header values.

Forbidden
*/
type GetAppMessagesForbidden struct {
	Payload *models.Error
}

func (o *GetAppMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /application/{id}/message][%d] getAppMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetAppMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppMessagesNotFound creates a GetAppMessagesNotFound with default headers values
func NewGetAppMessagesNotFound() *GetAppMessagesNotFound {
	return &GetAppMessagesNotFound{}
}

/*GetAppMessagesNotFound handles this case with default header values.

Not Found
*/
type GetAppMessagesNotFound struct {
	Payload *models.Error
}

func (o *GetAppMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /application/{id}/message][%d] getAppMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetAppMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
