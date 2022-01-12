// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1TeamsTeamIDReader is a Reader for the DeleteV1TeamsTeamID structure.
type DeleteV1TeamsTeamIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1TeamsTeamIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1TeamsTeamIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1TeamsTeamIDNoContent creates a DeleteV1TeamsTeamIDNoContent with default headers values
func NewDeleteV1TeamsTeamIDNoContent() *DeleteV1TeamsTeamIDNoContent {
	return &DeleteV1TeamsTeamIDNoContent{}
}

/* DeleteV1TeamsTeamIDNoContent describes a response with status code 204, with default header values.

Archive a team
*/
type DeleteV1TeamsTeamIDNoContent struct {
}

func (o *DeleteV1TeamsTeamIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/teams/{team_id}][%d] deleteV1TeamsTeamIdNoContent ", 204)
}

func (o *DeleteV1TeamsTeamIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}