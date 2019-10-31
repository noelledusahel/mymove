// Code generated by go-swagger; DO NOT EDIT.

package backup_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// IndexServiceMemberBackupContactsOKCode is the HTTP code returned for type IndexServiceMemberBackupContactsOK
const IndexServiceMemberBackupContactsOKCode int = 200

/*IndexServiceMemberBackupContactsOK list of service member backup contacts

swagger:response indexServiceMemberBackupContactsOK
*/
type IndexServiceMemberBackupContactsOK struct {

	/*
	  In: Body
	*/
	Payload internalmessages.IndexServiceMemberBackupContactsPayload `json:"body,omitempty"`
}

// NewIndexServiceMemberBackupContactsOK creates IndexServiceMemberBackupContactsOK with default headers values
func NewIndexServiceMemberBackupContactsOK() *IndexServiceMemberBackupContactsOK {

	return &IndexServiceMemberBackupContactsOK{}
}

// WithPayload adds the payload to the index service member backup contacts o k response
func (o *IndexServiceMemberBackupContactsOK) WithPayload(payload internalmessages.IndexServiceMemberBackupContactsPayload) *IndexServiceMemberBackupContactsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the index service member backup contacts o k response
func (o *IndexServiceMemberBackupContactsOK) SetPayload(payload internalmessages.IndexServiceMemberBackupContactsPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndexServiceMemberBackupContactsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = internalmessages.IndexServiceMemberBackupContactsPayload{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IndexServiceMemberBackupContactsBadRequestCode is the HTTP code returned for type IndexServiceMemberBackupContactsBadRequest
const IndexServiceMemberBackupContactsBadRequestCode int = 400

/*IndexServiceMemberBackupContactsBadRequest invalid request

swagger:response indexServiceMemberBackupContactsBadRequest
*/
type IndexServiceMemberBackupContactsBadRequest struct {
}

// NewIndexServiceMemberBackupContactsBadRequest creates IndexServiceMemberBackupContactsBadRequest with default headers values
func NewIndexServiceMemberBackupContactsBadRequest() *IndexServiceMemberBackupContactsBadRequest {

	return &IndexServiceMemberBackupContactsBadRequest{}
}

// WriteResponse to the client
func (o *IndexServiceMemberBackupContactsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IndexServiceMemberBackupContactsUnauthorizedCode is the HTTP code returned for type IndexServiceMemberBackupContactsUnauthorized
const IndexServiceMemberBackupContactsUnauthorizedCode int = 401

/*IndexServiceMemberBackupContactsUnauthorized request requires user authentication

swagger:response indexServiceMemberBackupContactsUnauthorized
*/
type IndexServiceMemberBackupContactsUnauthorized struct {
}

// NewIndexServiceMemberBackupContactsUnauthorized creates IndexServiceMemberBackupContactsUnauthorized with default headers values
func NewIndexServiceMemberBackupContactsUnauthorized() *IndexServiceMemberBackupContactsUnauthorized {

	return &IndexServiceMemberBackupContactsUnauthorized{}
}

// WriteResponse to the client
func (o *IndexServiceMemberBackupContactsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// IndexServiceMemberBackupContactsForbiddenCode is the HTTP code returned for type IndexServiceMemberBackupContactsForbidden
const IndexServiceMemberBackupContactsForbiddenCode int = 403

/*IndexServiceMemberBackupContactsForbidden user is not authorized to see this backup contact

swagger:response indexServiceMemberBackupContactsForbidden
*/
type IndexServiceMemberBackupContactsForbidden struct {
}

// NewIndexServiceMemberBackupContactsForbidden creates IndexServiceMemberBackupContactsForbidden with default headers values
func NewIndexServiceMemberBackupContactsForbidden() *IndexServiceMemberBackupContactsForbidden {

	return &IndexServiceMemberBackupContactsForbidden{}
}

// WriteResponse to the client
func (o *IndexServiceMemberBackupContactsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// IndexServiceMemberBackupContactsNotFoundCode is the HTTP code returned for type IndexServiceMemberBackupContactsNotFound
const IndexServiceMemberBackupContactsNotFoundCode int = 404

/*IndexServiceMemberBackupContactsNotFound contact not found

swagger:response indexServiceMemberBackupContactsNotFound
*/
type IndexServiceMemberBackupContactsNotFound struct {
}

// NewIndexServiceMemberBackupContactsNotFound creates IndexServiceMemberBackupContactsNotFound with default headers values
func NewIndexServiceMemberBackupContactsNotFound() *IndexServiceMemberBackupContactsNotFound {

	return &IndexServiceMemberBackupContactsNotFound{}
}

// WriteResponse to the client
func (o *IndexServiceMemberBackupContactsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// IndexServiceMemberBackupContactsInternalServerErrorCode is the HTTP code returned for type IndexServiceMemberBackupContactsInternalServerError
const IndexServiceMemberBackupContactsInternalServerErrorCode int = 500

/*IndexServiceMemberBackupContactsInternalServerError internal server error

swagger:response indexServiceMemberBackupContactsInternalServerError
*/
type IndexServiceMemberBackupContactsInternalServerError struct {
}

// NewIndexServiceMemberBackupContactsInternalServerError creates IndexServiceMemberBackupContactsInternalServerError with default headers values
func NewIndexServiceMemberBackupContactsInternalServerError() *IndexServiceMemberBackupContactsInternalServerError {

	return &IndexServiceMemberBackupContactsInternalServerError{}
}

// WriteResponse to the client
func (o *IndexServiceMemberBackupContactsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
