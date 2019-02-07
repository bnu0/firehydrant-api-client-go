// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDReader is a Reader for the DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentID structure.
type DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDNoContent creates a DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDNoContent with default headers values
func NewDeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDNoContent() *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDNoContent {
	return &DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDNoContent{}
}

/*DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDNoContent handles this case with default header values.

Remove an component from the report
*/
type DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDNoContent struct {
}

func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/post_mortems/reports/{report_id}/affected_components/{affected_component_id}][%d] deleteV1PostMortemsReportsReportIdAffectedComponentsAffectedComponentIdNoContent ", 204)
}

func (o *DeleteV1PostMortemsReportsReportIDAffectedComponentsAffectedComponentIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}