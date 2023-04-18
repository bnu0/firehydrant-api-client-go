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

// DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDReader is a Reader for the DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigID structure.
type DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK creates a DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK with default headers values
func NewDeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK() *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK {
	return &DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK{}
}

/*
DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK describes a response with status code 200, with default header values.

Archive configuration for a ticketing project
*/
type DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK struct {
	Payload *models.TicketingProjectConfigEntity
}

// IsSuccess returns true when this delete v1 ticketing projects ticketing project Id provider project configurations config Id o k response has a 2xx status code
func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 ticketing projects ticketing project Id provider project configurations config Id o k response has a 3xx status code
func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 ticketing projects ticketing project Id provider project configurations config Id o k response has a 4xx status code
func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 ticketing projects ticketing project Id provider project configurations config Id o k response has a 5xx status code
func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 ticketing projects ticketing project Id provider project configurations config Id o k response a status code equal to that given
func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/ticketing/projects/{ticketing_project_id}/provider_project_configurations/{config_id}][%d] deleteV1TicketingProjectsTicketingProjectIdProviderProjectConfigurationsConfigIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) String() string {
	return fmt.Sprintf("[DELETE /v1/ticketing/projects/{ticketing_project_id}/provider_project_configurations/{config_id}][%d] deleteV1TicketingProjectsTicketingProjectIdProviderProjectConfigurationsConfigIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) GetPayload() *models.TicketingProjectConfigEntity {
	return o.Payload
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TicketingProjectConfigEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
