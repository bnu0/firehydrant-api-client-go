// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutV1PostMortemsReportsReportIDReasonsOrderReader is a Reader for the PutV1PostMortemsReportsReportIDReasonsOrder structure.
type PutV1PostMortemsReportsReportIDReasonsOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1PostMortemsReportsReportIDReasonsOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutV1PostMortemsReportsReportIDReasonsOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutV1PostMortemsReportsReportIDReasonsOrderOK creates a PutV1PostMortemsReportsReportIDReasonsOrderOK with default headers values
func NewPutV1PostMortemsReportsReportIDReasonsOrderOK() *PutV1PostMortemsReportsReportIDReasonsOrderOK {
	return &PutV1PostMortemsReportsReportIDReasonsOrderOK{}
}

/*PutV1PostMortemsReportsReportIDReasonsOrderOK handles this case with default header values.

Reorder a reason in the post mortem reasons list
*/
type PutV1PostMortemsReportsReportIDReasonsOrderOK struct {
}

func (o *PutV1PostMortemsReportsReportIDReasonsOrderOK) Error() string {
	return fmt.Sprintf("[PUT /v1/post_mortems/reports/{report_id}/reasons/order][%d] putV1PostMortemsReportsReportIdReasonsOrderOK ", 200)
}

func (o *PutV1PostMortemsReportsReportIDReasonsOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}