// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1WebhooksWebhookIDReader is a Reader for the PatchV1WebhooksWebhookID structure.
type PatchV1WebhooksWebhookIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1WebhooksWebhookIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1WebhooksWebhookIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1WebhooksWebhookIDOK creates a PatchV1WebhooksWebhookIDOK with default headers values
func NewPatchV1WebhooksWebhookIDOK() *PatchV1WebhooksWebhookIDOK {
	return &PatchV1WebhooksWebhookIDOK{}
}

/*
PatchV1WebhooksWebhookIDOK describes a response with status code 200, with default header values.

Update a specific webhook
*/
type PatchV1WebhooksWebhookIDOK struct {
	Payload *models.WebhookEntity
}

// IsSuccess returns true when this patch v1 webhooks webhook Id o k response has a 2xx status code
func (o *PatchV1WebhooksWebhookIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 webhooks webhook Id o k response has a 3xx status code
func (o *PatchV1WebhooksWebhookIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 webhooks webhook Id o k response has a 4xx status code
func (o *PatchV1WebhooksWebhookIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 webhooks webhook Id o k response has a 5xx status code
func (o *PatchV1WebhooksWebhookIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 webhooks webhook Id o k response a status code equal to that given
func (o *PatchV1WebhooksWebhookIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1WebhooksWebhookIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/webhooks/{webhook_id}][%d] patchV1WebhooksWebhookIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1WebhooksWebhookIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/webhooks/{webhook_id}][%d] patchV1WebhooksWebhookIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1WebhooksWebhookIDOK) GetPayload() *models.WebhookEntity {
	return o.Payload
}

func (o *PatchV1WebhooksWebhookIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
