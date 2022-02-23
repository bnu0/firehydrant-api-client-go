// Code generated by go-swagger; DO NOT EDIT.

package priorities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1PrioritiesReader is a Reader for the GetV1Priorities structure.
type GetV1PrioritiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1PrioritiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1PrioritiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1PrioritiesOK creates a GetV1PrioritiesOK with default headers values
func NewGetV1PrioritiesOK() *GetV1PrioritiesOK {
	return &GetV1PrioritiesOK{}
}

/* GetV1PrioritiesOK describes a response with status code 200, with default header values.

Lists priorities
*/
type GetV1PrioritiesOK struct {
	Payload *models.PriorityEntity
}

func (o *GetV1PrioritiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/priorities][%d] getV1PrioritiesOK  %+v", 200, o.Payload)
}
func (o *GetV1PrioritiesOK) GetPayload() *models.PriorityEntity {
	return o.Payload
}

func (o *GetV1PrioritiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PriorityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
