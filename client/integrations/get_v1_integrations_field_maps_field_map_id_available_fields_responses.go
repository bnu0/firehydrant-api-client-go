// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsReader is a Reader for the GetV1IntegrationsFieldMapsFieldMapIDAvailableFields structure.
type GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK creates a GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK with default headers values
func NewGetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK() *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK {
	return &GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK{}
}

/*
GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK describes a response with status code 200, with default header values.

Get a description of the fields to which data can be mapped
*/
type GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK struct {
	Payload *models.FieldMappingMappableFieldEntity
}

// IsSuccess returns true when this get v1 integrations field maps field map Id available fields o k response has a 2xx status code
func (o *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 integrations field maps field map Id available fields o k response has a 3xx status code
func (o *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 integrations field maps field map Id available fields o k response has a 4xx status code
func (o *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 integrations field maps field map Id available fields o k response has a 5xx status code
func (o *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 integrations field maps field map Id available fields o k response a status code equal to that given
func (o *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK) Error() string {
	return fmt.Sprintf("[GET /v1/integrations/field_maps/{field_map_id}/available_fields][%d] getV1IntegrationsFieldMapsFieldMapIdAvailableFieldsOK  %+v", 200, o.Payload)
}

func (o *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK) String() string {
	return fmt.Sprintf("[GET /v1/integrations/field_maps/{field_map_id}/available_fields][%d] getV1IntegrationsFieldMapsFieldMapIdAvailableFieldsOK  %+v", 200, o.Payload)
}

func (o *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK) GetPayload() *models.FieldMappingMappableFieldEntity {
	return o.Payload
}

func (o *GetV1IntegrationsFieldMapsFieldMapIDAvailableFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FieldMappingMappableFieldEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
