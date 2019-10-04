// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	adminmessages "github.com/transcom/mymove/pkg/gen/adminmessages"
)

// IndexAdminUsersOKCode is the HTTP code returned for type IndexAdminUsersOK
const IndexAdminUsersOKCode int = 200

/*IndexAdminUsersOK success

swagger:response indexAdminUsersOK
*/
type IndexAdminUsersOK struct {
	/*Used for pagination

	 */
	ContentRange string `json:"Content-Range"`

	/*
	  In: Body
	*/
	Payload adminmessages.AdminUsers `json:"body,omitempty"`
}

// NewIndexAdminUsersOK creates IndexAdminUsersOK with default headers values
func NewIndexAdminUsersOK() *IndexAdminUsersOK {

	return &IndexAdminUsersOK{}
}

// WithContentRange adds the contentRange to the index admin users o k response
func (o *IndexAdminUsersOK) WithContentRange(contentRange string) *IndexAdminUsersOK {
	o.ContentRange = contentRange
	return o
}

// SetContentRange sets the contentRange to the index admin users o k response
func (o *IndexAdminUsersOK) SetContentRange(contentRange string) {
	o.ContentRange = contentRange
}

// WithPayload adds the payload to the index admin users o k response
func (o *IndexAdminUsersOK) WithPayload(payload adminmessages.AdminUsers) *IndexAdminUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the index admin users o k response
func (o *IndexAdminUsersOK) SetPayload(payload adminmessages.AdminUsers) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndexAdminUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Range

	contentRange := o.ContentRange
	if contentRange != "" {
		rw.Header().Set("Content-Range", contentRange)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = adminmessages.AdminUsers{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IndexAdminUsersBadRequestCode is the HTTP code returned for type IndexAdminUsersBadRequest
const IndexAdminUsersBadRequestCode int = 400

/*IndexAdminUsersBadRequest invalid request

swagger:response indexAdminUsersBadRequest
*/
type IndexAdminUsersBadRequest struct {
}

// NewIndexAdminUsersBadRequest creates IndexAdminUsersBadRequest with default headers values
func NewIndexAdminUsersBadRequest() *IndexAdminUsersBadRequest {

	return &IndexAdminUsersBadRequest{}
}

// WriteResponse to the client
func (o *IndexAdminUsersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IndexAdminUsersUnauthorizedCode is the HTTP code returned for type IndexAdminUsersUnauthorized
const IndexAdminUsersUnauthorizedCode int = 401

/*IndexAdminUsersUnauthorized request requires user authentication

swagger:response indexAdminUsersUnauthorized
*/
type IndexAdminUsersUnauthorized struct {
}

// NewIndexAdminUsersUnauthorized creates IndexAdminUsersUnauthorized with default headers values
func NewIndexAdminUsersUnauthorized() *IndexAdminUsersUnauthorized {

	return &IndexAdminUsersUnauthorized{}
}

// WriteResponse to the client
func (o *IndexAdminUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// IndexAdminUsersNotFoundCode is the HTTP code returned for type IndexAdminUsersNotFound
const IndexAdminUsersNotFoundCode int = 404

/*IndexAdminUsersNotFound office not found

swagger:response indexAdminUsersNotFound
*/
type IndexAdminUsersNotFound struct {
}

// NewIndexAdminUsersNotFound creates IndexAdminUsersNotFound with default headers values
func NewIndexAdminUsersNotFound() *IndexAdminUsersNotFound {

	return &IndexAdminUsersNotFound{}
}

// WriteResponse to the client
func (o *IndexAdminUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// IndexAdminUsersInternalServerErrorCode is the HTTP code returned for type IndexAdminUsersInternalServerError
const IndexAdminUsersInternalServerErrorCode int = 500

/*IndexAdminUsersInternalServerError server error

swagger:response indexAdminUsersInternalServerError
*/
type IndexAdminUsersInternalServerError struct {
}

// NewIndexAdminUsersInternalServerError creates IndexAdminUsersInternalServerError with default headers values
func NewIndexAdminUsersInternalServerError() *IndexAdminUsersInternalServerError {

	return &IndexAdminUsersInternalServerError{}
}

// WriteResponse to the client
func (o *IndexAdminUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
