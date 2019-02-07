// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutV1IncidentsIncidentIDEventsEventIDStarsReader is a Reader for the PutV1IncidentsIncidentIDEventsEventIDStars structure.
type PutV1IncidentsIncidentIDEventsEventIDStarsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1IncidentsIncidentIDEventsEventIDStarsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutV1IncidentsIncidentIDEventsEventIDStarsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutV1IncidentsIncidentIDEventsEventIDStarsOK creates a PutV1IncidentsIncidentIDEventsEventIDStarsOK with default headers values
func NewPutV1IncidentsIncidentIDEventsEventIDStarsOK() *PutV1IncidentsIncidentIDEventsEventIDStarsOK {
	return &PutV1IncidentsIncidentIDEventsEventIDStarsOK{}
}

/*PutV1IncidentsIncidentIDEventsEventIDStarsOK handles this case with default header values.

Star an incident event
*/
type PutV1IncidentsIncidentIDEventsEventIDStarsOK struct {
}

func (o *PutV1IncidentsIncidentIDEventsEventIDStarsOK) Error() string {
	return fmt.Sprintf("[PUT /v1/incidents/{incident_id}/events/{event_id}/stars][%d] putV1IncidentsIncidentIdEventsEventIdStarsOK ", 200)
}

func (o *PutV1IncidentsIncidentIDEventsEventIDStarsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}