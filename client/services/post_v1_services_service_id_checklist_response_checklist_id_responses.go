// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1ServicesServiceIDChecklistResponseChecklistIDReader is a Reader for the PostV1ServicesServiceIDChecklistResponseChecklistID structure.
type PostV1ServicesServiceIDChecklistResponseChecklistIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ServicesServiceIDChecklistResponseChecklistIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ServicesServiceIDChecklistResponseChecklistIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1ServicesServiceIDChecklistResponseChecklistIDCreated creates a PostV1ServicesServiceIDChecklistResponseChecklistIDCreated with default headers values
func NewPostV1ServicesServiceIDChecklistResponseChecklistIDCreated() *PostV1ServicesServiceIDChecklistResponseChecklistIDCreated {
	return &PostV1ServicesServiceIDChecklistResponseChecklistIDCreated{}
}

/* PostV1ServicesServiceIDChecklistResponseChecklistIDCreated describes a response with status code 201, with default header values.

created ChecklistResponse
*/
type PostV1ServicesServiceIDChecklistResponseChecklistIDCreated struct {
}

func (o *PostV1ServicesServiceIDChecklistResponseChecklistIDCreated) Error() string {
	return fmt.Sprintf("[POST /v1/services/{service_id}/checklist_response/{checklist_id}][%d] postV1ServicesServiceIdChecklistResponseChecklistIdCreated ", 201)
}

func (o *PostV1ServicesServiceIDChecklistResponseChecklistIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}