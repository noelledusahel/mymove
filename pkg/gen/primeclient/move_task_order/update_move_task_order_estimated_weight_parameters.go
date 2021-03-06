// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateMoveTaskOrderEstimatedWeightParams creates a new UpdateMoveTaskOrderEstimatedWeightParams object
// with the default values initialized.
func NewUpdateMoveTaskOrderEstimatedWeightParams() *UpdateMoveTaskOrderEstimatedWeightParams {
	var ()
	return &UpdateMoveTaskOrderEstimatedWeightParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMoveTaskOrderEstimatedWeightParamsWithTimeout creates a new UpdateMoveTaskOrderEstimatedWeightParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMoveTaskOrderEstimatedWeightParamsWithTimeout(timeout time.Duration) *UpdateMoveTaskOrderEstimatedWeightParams {
	var ()
	return &UpdateMoveTaskOrderEstimatedWeightParams{

		timeout: timeout,
	}
}

// NewUpdateMoveTaskOrderEstimatedWeightParamsWithContext creates a new UpdateMoveTaskOrderEstimatedWeightParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMoveTaskOrderEstimatedWeightParamsWithContext(ctx context.Context) *UpdateMoveTaskOrderEstimatedWeightParams {
	var ()
	return &UpdateMoveTaskOrderEstimatedWeightParams{

		Context: ctx,
	}
}

// NewUpdateMoveTaskOrderEstimatedWeightParamsWithHTTPClient creates a new UpdateMoveTaskOrderEstimatedWeightParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMoveTaskOrderEstimatedWeightParamsWithHTTPClient(client *http.Client) *UpdateMoveTaskOrderEstimatedWeightParams {
	var ()
	return &UpdateMoveTaskOrderEstimatedWeightParams{
		HTTPClient: client,
	}
}

/*UpdateMoveTaskOrderEstimatedWeightParams contains all the parameters to send to the API endpoint
for the update move task order estimated weight operation typically these are written to a http.Request
*/
type UpdateMoveTaskOrderEstimatedWeightParams struct {

	/*Body*/
	Body UpdateMoveTaskOrderEstimatedWeightBody
	/*MoveTaskOrderID
	  ID of move order to use

	*/
	MoveTaskOrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update move task order estimated weight params
func (o *UpdateMoveTaskOrderEstimatedWeightParams) WithTimeout(timeout time.Duration) *UpdateMoveTaskOrderEstimatedWeightParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update move task order estimated weight params
func (o *UpdateMoveTaskOrderEstimatedWeightParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update move task order estimated weight params
func (o *UpdateMoveTaskOrderEstimatedWeightParams) WithContext(ctx context.Context) *UpdateMoveTaskOrderEstimatedWeightParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update move task order estimated weight params
func (o *UpdateMoveTaskOrderEstimatedWeightParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update move task order estimated weight params
func (o *UpdateMoveTaskOrderEstimatedWeightParams) WithHTTPClient(client *http.Client) *UpdateMoveTaskOrderEstimatedWeightParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update move task order estimated weight params
func (o *UpdateMoveTaskOrderEstimatedWeightParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update move task order estimated weight params
func (o *UpdateMoveTaskOrderEstimatedWeightParams) WithBody(body UpdateMoveTaskOrderEstimatedWeightBody) *UpdateMoveTaskOrderEstimatedWeightParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update move task order estimated weight params
func (o *UpdateMoveTaskOrderEstimatedWeightParams) SetBody(body UpdateMoveTaskOrderEstimatedWeightBody) {
	o.Body = body
}

// WithMoveTaskOrderID adds the moveTaskOrderID to the update move task order estimated weight params
func (o *UpdateMoveTaskOrderEstimatedWeightParams) WithMoveTaskOrderID(moveTaskOrderID string) *UpdateMoveTaskOrderEstimatedWeightParams {
	o.SetMoveTaskOrderID(moveTaskOrderID)
	return o
}

// SetMoveTaskOrderID adds the moveTaskOrderId to the update move task order estimated weight params
func (o *UpdateMoveTaskOrderEstimatedWeightParams) SetMoveTaskOrderID(moveTaskOrderID string) {
	o.MoveTaskOrderID = moveTaskOrderID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMoveTaskOrderEstimatedWeightParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param moveTaskOrderID
	if err := r.SetPathParam("moveTaskOrderID", o.MoveTaskOrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
