// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1MetricsUserInvolvementsReader is a Reader for the GetV1MetricsUserInvolvements structure.
type GetV1MetricsUserInvolvementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1MetricsUserInvolvementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1MetricsUserInvolvementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1MetricsUserInvolvementsOK creates a GetV1MetricsUserInvolvementsOK with default headers values
func NewGetV1MetricsUserInvolvementsOK() *GetV1MetricsUserInvolvementsOK {
	return &GetV1MetricsUserInvolvementsOK{}
}

/*
GetV1MetricsUserInvolvementsOK describes a response with status code 200, with default header values.

Returns a report with time bucketed analytics data
*/
type GetV1MetricsUserInvolvementsOK struct {
	Payload *models.MetricsMetricsEntity
}

// IsSuccess returns true when this get v1 metrics user involvements o k response has a 2xx status code
func (o *GetV1MetricsUserInvolvementsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 metrics user involvements o k response has a 3xx status code
func (o *GetV1MetricsUserInvolvementsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 metrics user involvements o k response has a 4xx status code
func (o *GetV1MetricsUserInvolvementsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 metrics user involvements o k response has a 5xx status code
func (o *GetV1MetricsUserInvolvementsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 metrics user involvements o k response a status code equal to that given
func (o *GetV1MetricsUserInvolvementsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1MetricsUserInvolvementsOK) Error() string {
	return fmt.Sprintf("[GET /v1/metrics/user_involvements][%d] getV1MetricsUserInvolvementsOK  %+v", 200, o.Payload)
}

func (o *GetV1MetricsUserInvolvementsOK) String() string {
	return fmt.Sprintf("[GET /v1/metrics/user_involvements][%d] getV1MetricsUserInvolvementsOK  %+v", 200, o.Payload)
}

func (o *GetV1MetricsUserInvolvementsOK) GetPayload() *models.MetricsMetricsEntity {
	return o.Payload
}

func (o *GetV1MetricsUserInvolvementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetricsMetricsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
