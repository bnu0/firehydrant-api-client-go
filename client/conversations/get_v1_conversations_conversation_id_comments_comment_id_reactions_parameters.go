// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetV1ConversationsConversationIDCommentsCommentIDReactionsParams creates a new GetV1ConversationsConversationIDCommentsCommentIDReactionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ConversationsConversationIDCommentsCommentIDReactionsParams() *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	return &GetV1ConversationsConversationIDCommentsCommentIDReactionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithTimeout creates a new GetV1ConversationsConversationIDCommentsCommentIDReactionsParams object
// with the ability to set a timeout on a request.
func NewGetV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithTimeout(timeout time.Duration) *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	return &GetV1ConversationsConversationIDCommentsCommentIDReactionsParams{
		timeout: timeout,
	}
}

// NewGetV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithContext creates a new GetV1ConversationsConversationIDCommentsCommentIDReactionsParams object
// with the ability to set a context for a request.
func NewGetV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithContext(ctx context.Context) *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	return &GetV1ConversationsConversationIDCommentsCommentIDReactionsParams{
		Context: ctx,
	}
}

// NewGetV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithHTTPClient creates a new GetV1ConversationsConversationIDCommentsCommentIDReactionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithHTTPClient(client *http.Client) *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	return &GetV1ConversationsConversationIDCommentsCommentIDReactionsParams{
		HTTPClient: client,
	}
}

/*
GetV1ConversationsConversationIDCommentsCommentIDReactionsParams contains all the parameters to send to the API endpoint

	for the get v1 conversations conversation Id comments comment Id reactions operation.

	Typically these are written to a http.Request.
*/
type GetV1ConversationsConversationIDCommentsCommentIDReactionsParams struct {

	// CommentID.
	//
	// Format: int32
	CommentID int32

	// ConversationID.
	//
	// Format: int32
	ConversationID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 conversations conversation Id comments comment Id reactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithDefaults() *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 conversations conversation Id comments comment Id reactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 conversations conversation Id comments comment Id reactions params
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithTimeout(timeout time.Duration) *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 conversations conversation Id comments comment Id reactions params
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 conversations conversation Id comments comment Id reactions params
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithContext(ctx context.Context) *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 conversations conversation Id comments comment Id reactions params
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 conversations conversation Id comments comment Id reactions params
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithHTTPClient(client *http.Client) *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 conversations conversation Id comments comment Id reactions params
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentID adds the commentID to the get v1 conversations conversation Id comments comment Id reactions params
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithCommentID(commentID int32) *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetCommentID(commentID)
	return o
}

// SetCommentID adds the commentId to the get v1 conversations conversation Id comments comment Id reactions params
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetCommentID(commentID int32) {
	o.CommentID = commentID
}

// WithConversationID adds the conversationID to the get v1 conversations conversation Id comments comment Id reactions params
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithConversationID(conversationID int32) *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the get v1 conversations conversation Id comments comment Id reactions params
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetConversationID(conversationID int32) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param comment_id
	if err := r.SetPathParam("comment_id", swag.FormatInt32(o.CommentID)); err != nil {
		return err
	}

	// path param conversation_id
	if err := r.SetPathParam("conversation_id", swag.FormatInt32(o.ConversationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
