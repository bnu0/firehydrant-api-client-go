// Code generated by go-swagger; DO NOT EDIT.

package scim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1ScimV2UsersIDReader is a Reader for the GetV1ScimV2UsersID structure.
type GetV1ScimV2UsersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ScimV2UsersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ScimV2UsersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1ScimV2UsersIDOK creates a GetV1ScimV2UsersIDOK with default headers values
func NewGetV1ScimV2UsersIDOK() *GetV1ScimV2UsersIDOK {
	return &GetV1ScimV2UsersIDOK{}
}

/*
GetV1ScimV2UsersIDOK describes a response with status code 200, with default header values.

SCIM endpoint that lists a User
*/
type GetV1ScimV2UsersIDOK struct {
}

// IsSuccess returns true when this get v1 scim v2 users Id o k response has a 2xx status code
func (o *GetV1ScimV2UsersIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 scim v2 users Id o k response has a 3xx status code
func (o *GetV1ScimV2UsersIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 scim v2 users Id o k response has a 4xx status code
func (o *GetV1ScimV2UsersIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 scim v2 users Id o k response has a 5xx status code
func (o *GetV1ScimV2UsersIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 scim v2 users Id o k response a status code equal to that given
func (o *GetV1ScimV2UsersIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1ScimV2UsersIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/scim/v2/Users/{id}][%d] getV1ScimV2UsersIdOK ", 200)
}

func (o *GetV1ScimV2UsersIDOK) String() string {
	return fmt.Sprintf("[GET /v1/scim/v2/Users/{id}][%d] getV1ScimV2UsersIdOK ", 200)
}

func (o *GetV1ScimV2UsersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}