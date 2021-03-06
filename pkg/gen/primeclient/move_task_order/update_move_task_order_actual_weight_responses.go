// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMoveTaskOrderActualWeightReader is a Reader for the UpdateMoveTaskOrderActualWeight structure.
type UpdateMoveTaskOrderActualWeightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMoveTaskOrderActualWeightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMoveTaskOrderActualWeightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMoveTaskOrderActualWeightBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateMoveTaskOrderActualWeightUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMoveTaskOrderActualWeightForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMoveTaskOrderActualWeightNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMoveTaskOrderActualWeightInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateMoveTaskOrderActualWeightOK creates a UpdateMoveTaskOrderActualWeightOK with default headers values
func NewUpdateMoveTaskOrderActualWeightOK() *UpdateMoveTaskOrderActualWeightOK {
	return &UpdateMoveTaskOrderActualWeightOK{}
}

/*UpdateMoveTaskOrderActualWeightOK handles this case with default header values.

Successfully retrieved move task order
*/
type UpdateMoveTaskOrderActualWeightOK struct {
	Payload *primemessages.MoveTaskOrder
}

func (o *UpdateMoveTaskOrderActualWeightOK) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-actual-weight][%d] updateMoveTaskOrderActualWeightOK  %+v", 200, o.Payload)
}

func (o *UpdateMoveTaskOrderActualWeightOK) GetPayload() *primemessages.MoveTaskOrder {
	return o.Payload
}

func (o *UpdateMoveTaskOrderActualWeightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.MoveTaskOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderActualWeightBadRequest creates a UpdateMoveTaskOrderActualWeightBadRequest with default headers values
func NewUpdateMoveTaskOrderActualWeightBadRequest() *UpdateMoveTaskOrderActualWeightBadRequest {
	return &UpdateMoveTaskOrderActualWeightBadRequest{}
}

/*UpdateMoveTaskOrderActualWeightBadRequest handles this case with default header values.

The request payload is invalid
*/
type UpdateMoveTaskOrderActualWeightBadRequest struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderActualWeightBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-actual-weight][%d] updateMoveTaskOrderActualWeightBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateMoveTaskOrderActualWeightBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderActualWeightBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderActualWeightUnauthorized creates a UpdateMoveTaskOrderActualWeightUnauthorized with default headers values
func NewUpdateMoveTaskOrderActualWeightUnauthorized() *UpdateMoveTaskOrderActualWeightUnauthorized {
	return &UpdateMoveTaskOrderActualWeightUnauthorized{}
}

/*UpdateMoveTaskOrderActualWeightUnauthorized handles this case with default header values.

The request was denied
*/
type UpdateMoveTaskOrderActualWeightUnauthorized struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderActualWeightUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-actual-weight][%d] updateMoveTaskOrderActualWeightUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateMoveTaskOrderActualWeightUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderActualWeightUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderActualWeightForbidden creates a UpdateMoveTaskOrderActualWeightForbidden with default headers values
func NewUpdateMoveTaskOrderActualWeightForbidden() *UpdateMoveTaskOrderActualWeightForbidden {
	return &UpdateMoveTaskOrderActualWeightForbidden{}
}

/*UpdateMoveTaskOrderActualWeightForbidden handles this case with default header values.

The request was denied
*/
type UpdateMoveTaskOrderActualWeightForbidden struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderActualWeightForbidden) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-actual-weight][%d] updateMoveTaskOrderActualWeightForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMoveTaskOrderActualWeightForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderActualWeightForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderActualWeightNotFound creates a UpdateMoveTaskOrderActualWeightNotFound with default headers values
func NewUpdateMoveTaskOrderActualWeightNotFound() *UpdateMoveTaskOrderActualWeightNotFound {
	return &UpdateMoveTaskOrderActualWeightNotFound{}
}

/*UpdateMoveTaskOrderActualWeightNotFound handles this case with default header values.

The requested resource wasn't found
*/
type UpdateMoveTaskOrderActualWeightNotFound struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderActualWeightNotFound) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-actual-weight][%d] updateMoveTaskOrderActualWeightNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMoveTaskOrderActualWeightNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderActualWeightNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderActualWeightInternalServerError creates a UpdateMoveTaskOrderActualWeightInternalServerError with default headers values
func NewUpdateMoveTaskOrderActualWeightInternalServerError() *UpdateMoveTaskOrderActualWeightInternalServerError {
	return &UpdateMoveTaskOrderActualWeightInternalServerError{}
}

/*UpdateMoveTaskOrderActualWeightInternalServerError handles this case with default header values.

A server error occurred
*/
type UpdateMoveTaskOrderActualWeightInternalServerError struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderActualWeightInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-actual-weight][%d] updateMoveTaskOrderActualWeightInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMoveTaskOrderActualWeightInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderActualWeightInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateMoveTaskOrderActualWeightBody update move task order actual weight body
swagger:model UpdateMoveTaskOrderActualWeightBody
*/
type UpdateMoveTaskOrderActualWeightBody struct {

	// actual weight
	ActualWeight int64 `json:"actualWeight,omitempty"`
}

// Validate validates this update move task order actual weight body
func (o *UpdateMoveTaskOrderActualWeightBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateMoveTaskOrderActualWeightBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateMoveTaskOrderActualWeightBody) UnmarshalBinary(b []byte) error {
	var res UpdateMoveTaskOrderActualWeightBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
