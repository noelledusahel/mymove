// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// ShowOrdersOKCode is the HTTP code returned for type ShowOrdersOK
const ShowOrdersOKCode int = 200

/*ShowOrdersOK the instance of the order

swagger:response showOrdersOK
*/
type ShowOrdersOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.Orders `json:"body,omitempty"`
}

// NewShowOrdersOK creates ShowOrdersOK with default headers values
func NewShowOrdersOK() *ShowOrdersOK {

	return &ShowOrdersOK{}
}

// WithPayload adds the payload to the show orders o k response
func (o *ShowOrdersOK) WithPayload(payload *internalmessages.Orders) *ShowOrdersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show orders o k response
func (o *ShowOrdersOK) SetPayload(payload *internalmessages.Orders) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowOrdersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowOrdersBadRequestCode is the HTTP code returned for type ShowOrdersBadRequest
const ShowOrdersBadRequestCode int = 400

/*ShowOrdersBadRequest invalid request

swagger:response showOrdersBadRequest
*/
type ShowOrdersBadRequest struct {
}

// NewShowOrdersBadRequest creates ShowOrdersBadRequest with default headers values
func NewShowOrdersBadRequest() *ShowOrdersBadRequest {

	return &ShowOrdersBadRequest{}
}

// WriteResponse to the client
func (o *ShowOrdersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ShowOrdersUnauthorizedCode is the HTTP code returned for type ShowOrdersUnauthorized
const ShowOrdersUnauthorizedCode int = 401

/*ShowOrdersUnauthorized request requires user authentication

swagger:response showOrdersUnauthorized
*/
type ShowOrdersUnauthorized struct {
}

// NewShowOrdersUnauthorized creates ShowOrdersUnauthorized with default headers values
func NewShowOrdersUnauthorized() *ShowOrdersUnauthorized {

	return &ShowOrdersUnauthorized{}
}

// WriteResponse to the client
func (o *ShowOrdersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ShowOrdersForbiddenCode is the HTTP code returned for type ShowOrdersForbidden
const ShowOrdersForbiddenCode int = 403

/*ShowOrdersForbidden user is not authorized

swagger:response showOrdersForbidden
*/
type ShowOrdersForbidden struct {
}

// NewShowOrdersForbidden creates ShowOrdersForbidden with default headers values
func NewShowOrdersForbidden() *ShowOrdersForbidden {

	return &ShowOrdersForbidden{}
}

// WriteResponse to the client
func (o *ShowOrdersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ShowOrdersNotFoundCode is the HTTP code returned for type ShowOrdersNotFound
const ShowOrdersNotFoundCode int = 404

/*ShowOrdersNotFound order is not found

swagger:response showOrdersNotFound
*/
type ShowOrdersNotFound struct {
}

// NewShowOrdersNotFound creates ShowOrdersNotFound with default headers values
func NewShowOrdersNotFound() *ShowOrdersNotFound {

	return &ShowOrdersNotFound{}
}

// WriteResponse to the client
func (o *ShowOrdersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ShowOrdersInternalServerErrorCode is the HTTP code returned for type ShowOrdersInternalServerError
const ShowOrdersInternalServerErrorCode int = 500

/*ShowOrdersInternalServerError internal server error

swagger:response showOrdersInternalServerError
*/
type ShowOrdersInternalServerError struct {
}

// NewShowOrdersInternalServerError creates ShowOrdersInternalServerError with default headers values
func NewShowOrdersInternalServerError() *ShowOrdersInternalServerError {

	return &ShowOrdersInternalServerError{}
}

// WriteResponse to the client
func (o *ShowOrdersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
