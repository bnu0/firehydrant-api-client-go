// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// DeleteV1IncidentsIncidentIDEventsEventIDReader is a Reader for the DeleteV1IncidentsIncidentIDEventsEventID structure.
type DeleteV1IncidentsIncidentIDEventsEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1IncidentsIncidentIDEventsEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1IncidentsIncidentIDEventsEventIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1IncidentsIncidentIDEventsEventIDOK creates a DeleteV1IncidentsIncidentIDEventsEventIDOK with default headers values
func NewDeleteV1IncidentsIncidentIDEventsEventIDOK() *DeleteV1IncidentsIncidentIDEventsEventIDOK {
	return &DeleteV1IncidentsIncidentIDEventsEventIDOK{}
}

/*
DeleteV1IncidentsIncidentIDEventsEventIDOK describes a response with status code 200, with default header values.

Delete a single event for an incident
*/
type DeleteV1IncidentsIncidentIDEventsEventIDOK struct {
	Payload *models.IncidentEventEntity
}

// IsSuccess returns true when this delete v1 incidents incident Id events event Id o k response has a 2xx status code
func (o *DeleteV1IncidentsIncidentIDEventsEventIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 incidents incident Id events event Id o k response has a 3xx status code
func (o *DeleteV1IncidentsIncidentIDEventsEventIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 incidents incident Id events event Id o k response has a 4xx status code
func (o *DeleteV1IncidentsIncidentIDEventsEventIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 incidents incident Id events event Id o k response has a 5xx status code
func (o *DeleteV1IncidentsIncidentIDEventsEventIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 incidents incident Id events event Id o k response a status code equal to that given
func (o *DeleteV1IncidentsIncidentIDEventsEventIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteV1IncidentsIncidentIDEventsEventIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/incidents/{incident_id}/events/{event_id}][%d] deleteV1IncidentsIncidentIdEventsEventIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1IncidentsIncidentIDEventsEventIDOK) String() string {
	return fmt.Sprintf("[DELETE /v1/incidents/{incident_id}/events/{event_id}][%d] deleteV1IncidentsIncidentIdEventsEventIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1IncidentsIncidentIDEventsEventIDOK) GetPayload() *models.IncidentEventEntity {
	return o.Payload
}

func (o *DeleteV1IncidentsIncidentIDEventsEventIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentEventEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
