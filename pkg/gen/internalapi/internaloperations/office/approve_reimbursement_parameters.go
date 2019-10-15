// Code generated by go-swagger; DO NOT EDIT.

package office

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewApproveReimbursementParams creates a new ApproveReimbursementParams object
// no default values defined in spec.
func NewApproveReimbursementParams() ApproveReimbursementParams {

	return ApproveReimbursementParams{}
}

// ApproveReimbursementParams contains all the bound params for the approve reimbursement operation
// typically these are obtained from a http.Request
//
// swagger:parameters approveReimbursement
type ApproveReimbursementParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*UUID of the reimbursement being approved
	  Required: true
	  In: path
	*/
	ReimbursementID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewApproveReimbursementParams() beforehand.
func (o *ApproveReimbursementParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rReimbursementID, rhkReimbursementID, _ := route.Params.GetOK("reimbursementId")
	if err := o.bindReimbursementID(rReimbursementID, rhkReimbursementID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindReimbursementID binds and validates parameter ReimbursementID from path.
func (o *ApproveReimbursementParams) bindReimbursementID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("reimbursementId", "path", "strfmt.UUID", raw)
	}
	o.ReimbursementID = *(value.(*strfmt.UUID))

	if err := o.validateReimbursementID(formats); err != nil {
		return err
	}

	return nil
}

// validateReimbursementID carries on validations for parameter ReimbursementID
func (o *ApproveReimbursementParams) validateReimbursementID(formats strfmt.Registry) error {

	if err := validate.FormatOf("reimbursementId", "path", "uuid", o.ReimbursementID.String(), formats); err != nil {
		return err
	}
	return nil
}