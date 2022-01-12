// Code generated by go-swagger; DO NOT EDIT.

package nunc_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1NuncConnectionsNuncConnectionIDSubscribersReader is a Reader for the GetV1NuncConnectionsNuncConnectionIDSubscribers structure.
type GetV1NuncConnectionsNuncConnectionIDSubscribersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1NuncConnectionsNuncConnectionIDSubscribersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1NuncConnectionsNuncConnectionIDSubscribersOK creates a GetV1NuncConnectionsNuncConnectionIDSubscribersOK with default headers values
func NewGetV1NuncConnectionsNuncConnectionIDSubscribersOK() *GetV1NuncConnectionsNuncConnectionIDSubscribersOK {
	return &GetV1NuncConnectionsNuncConnectionIDSubscribersOK{}
}

/* GetV1NuncConnectionsNuncConnectionIDSubscribersOK describes a response with status code 200, with default header values.

Retrieves the list of subscribers for a status page.
*/
type GetV1NuncConnectionsNuncConnectionIDSubscribersOK struct {
	Payload *models.NuncEmailSubscribersEntity
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) Error() string {
	return fmt.Sprintf("[GET /v1/nunc_connections/{nunc_connection_id}/subscribers][%d] getV1NuncConnectionsNuncConnectionIdSubscribersOK  %+v", 200, o.Payload)
}
func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) GetPayload() *models.NuncEmailSubscribersEntity {
	return o.Payload
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NuncEmailSubscribersEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}