// Code generated by go-swagger; DO NOT EDIT.

package service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// DeleteServiceItemOKCode is the HTTP code returned for type DeleteServiceItemOK
const DeleteServiceItemOKCode int = 200

/*DeleteServiceItemOK Successfully deleted move task order

swagger:response deleteServiceItemOK
*/
type DeleteServiceItemOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MoveTaskOrder `json:"body,omitempty"`
}

// NewDeleteServiceItemOK creates DeleteServiceItemOK with default headers values
func NewDeleteServiceItemOK() *DeleteServiceItemOK {

	return &DeleteServiceItemOK{}
}

// WithPayload adds the payload to the delete service item o k response
func (o *DeleteServiceItemOK) WithPayload(payload *ghcmessages.MoveTaskOrder) *DeleteServiceItemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service item o k response
func (o *DeleteServiceItemOK) SetPayload(payload *ghcmessages.MoveTaskOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceItemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteServiceItemBadRequestCode is the HTTP code returned for type DeleteServiceItemBadRequest
const DeleteServiceItemBadRequestCode int = 400

/*DeleteServiceItemBadRequest The request payload is invalid

swagger:response deleteServiceItemBadRequest
*/
type DeleteServiceItemBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteServiceItemBadRequest creates DeleteServiceItemBadRequest with default headers values
func NewDeleteServiceItemBadRequest() *DeleteServiceItemBadRequest {

	return &DeleteServiceItemBadRequest{}
}

// WithPayload adds the payload to the delete service item bad request response
func (o *DeleteServiceItemBadRequest) WithPayload(payload interface{}) *DeleteServiceItemBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service item bad request response
func (o *DeleteServiceItemBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceItemBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteServiceItemUnauthorizedCode is the HTTP code returned for type DeleteServiceItemUnauthorized
const DeleteServiceItemUnauthorizedCode int = 401

/*DeleteServiceItemUnauthorized The request was denied

swagger:response deleteServiceItemUnauthorized
*/
type DeleteServiceItemUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteServiceItemUnauthorized creates DeleteServiceItemUnauthorized with default headers values
func NewDeleteServiceItemUnauthorized() *DeleteServiceItemUnauthorized {

	return &DeleteServiceItemUnauthorized{}
}

// WithPayload adds the payload to the delete service item unauthorized response
func (o *DeleteServiceItemUnauthorized) WithPayload(payload interface{}) *DeleteServiceItemUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service item unauthorized response
func (o *DeleteServiceItemUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceItemUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteServiceItemForbiddenCode is the HTTP code returned for type DeleteServiceItemForbidden
const DeleteServiceItemForbiddenCode int = 403

/*DeleteServiceItemForbidden The request was denied

swagger:response deleteServiceItemForbidden
*/
type DeleteServiceItemForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteServiceItemForbidden creates DeleteServiceItemForbidden with default headers values
func NewDeleteServiceItemForbidden() *DeleteServiceItemForbidden {

	return &DeleteServiceItemForbidden{}
}

// WithPayload adds the payload to the delete service item forbidden response
func (o *DeleteServiceItemForbidden) WithPayload(payload interface{}) *DeleteServiceItemForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service item forbidden response
func (o *DeleteServiceItemForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceItemForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteServiceItemNotFoundCode is the HTTP code returned for type DeleteServiceItemNotFound
const DeleteServiceItemNotFoundCode int = 404

/*DeleteServiceItemNotFound The requested resource wasn't found

swagger:response deleteServiceItemNotFound
*/
type DeleteServiceItemNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteServiceItemNotFound creates DeleteServiceItemNotFound with default headers values
func NewDeleteServiceItemNotFound() *DeleteServiceItemNotFound {

	return &DeleteServiceItemNotFound{}
}

// WithPayload adds the payload to the delete service item not found response
func (o *DeleteServiceItemNotFound) WithPayload(payload interface{}) *DeleteServiceItemNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service item not found response
func (o *DeleteServiceItemNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceItemNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteServiceItemInternalServerErrorCode is the HTTP code returned for type DeleteServiceItemInternalServerError
const DeleteServiceItemInternalServerErrorCode int = 500

/*DeleteServiceItemInternalServerError A server error occurred

swagger:response deleteServiceItemInternalServerError
*/
type DeleteServiceItemInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteServiceItemInternalServerError creates DeleteServiceItemInternalServerError with default headers values
func NewDeleteServiceItemInternalServerError() *DeleteServiceItemInternalServerError {

	return &DeleteServiceItemInternalServerError{}
}

// WithPayload adds the payload to the delete service item internal server error response
func (o *DeleteServiceItemInternalServerError) WithPayload(payload interface{}) *DeleteServiceItemInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete service item internal server error response
func (o *DeleteServiceItemInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteServiceItemInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
