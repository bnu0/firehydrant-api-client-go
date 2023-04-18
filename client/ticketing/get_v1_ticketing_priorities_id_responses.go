// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1TicketingPrioritiesIDReader is a Reader for the GetV1TicketingPrioritiesID structure.
type GetV1TicketingPrioritiesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TicketingPrioritiesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TicketingPrioritiesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1TicketingPrioritiesIDOK creates a GetV1TicketingPrioritiesIDOK with default headers values
func NewGetV1TicketingPrioritiesIDOK() *GetV1TicketingPrioritiesIDOK {
	return &GetV1TicketingPrioritiesIDOK{}
}

/*
GetV1TicketingPrioritiesIDOK describes a response with status code 200, with default header values.

Retrieve a single ticketing priority by ID
*/
type GetV1TicketingPrioritiesIDOK struct {
	Payload *models.TicketingPriorityEntity
}

// IsSuccess returns true when this get v1 ticketing priorities Id o k response has a 2xx status code
func (o *GetV1TicketingPrioritiesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 ticketing priorities Id o k response has a 3xx status code
func (o *GetV1TicketingPrioritiesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 ticketing priorities Id o k response has a 4xx status code
func (o *GetV1TicketingPrioritiesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 ticketing priorities Id o k response has a 5xx status code
func (o *GetV1TicketingPrioritiesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 ticketing priorities Id o k response a status code equal to that given
func (o *GetV1TicketingPrioritiesIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1TicketingPrioritiesIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/ticketing/priorities/{id}][%d] getV1TicketingPrioritiesIdOK  %+v", 200, o.Payload)
}

func (o *GetV1TicketingPrioritiesIDOK) String() string {
	return fmt.Sprintf("[GET /v1/ticketing/priorities/{id}][%d] getV1TicketingPrioritiesIdOK  %+v", 200, o.Payload)
}

func (o *GetV1TicketingPrioritiesIDOK) GetPayload() *models.TicketingPriorityEntity {
	return o.Payload
}

func (o *GetV1TicketingPrioritiesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TicketingPriorityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}