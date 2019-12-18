// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateMTOServiceItemParams creates a new CreateMTOServiceItemParams object
// no default values defined in spec.
func NewCreateMTOServiceItemParams() CreateMTOServiceItemParams {

	return CreateMTOServiceItemParams{}
}

// CreateMTOServiceItemParams contains all the bound params for the create m t o service item operation
// typically these are obtained from a http.Request
//
// swagger:parameters createMTOServiceItem
type CreateMTOServiceItemParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of the rate engine services
	  Required: true
	  In: body
	*/
	CreateMTOServiceItemBody CreateMTOServiceItemBody
	/*ID of move order to use
	  Required: true
	  In: path
	*/
	MoveTaskOrderID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateMTOServiceItemParams() beforehand.
func (o *CreateMTOServiceItemParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body CreateMTOServiceItemBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("createMTOServiceItemBody", "body"))
			} else {
				res = append(res, errors.NewParseError("createMTOServiceItemBody", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.CreateMTOServiceItemBody = body
			}
		}
	} else {
		res = append(res, errors.Required("createMTOServiceItemBody", "body"))
	}
	rMoveTaskOrderID, rhkMoveTaskOrderID, _ := route.Params.GetOK("moveTaskOrderID")
	if err := o.bindMoveTaskOrderID(rMoveTaskOrderID, rhkMoveTaskOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMoveTaskOrderID binds and validates parameter MoveTaskOrderID from path.
func (o *CreateMTOServiceItemParams) bindMoveTaskOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("moveTaskOrderID", "path", "strfmt.UUID", raw)
	}
	o.MoveTaskOrderID = *(value.(*strfmt.UUID))

	if err := o.validateMoveTaskOrderID(formats); err != nil {
		return err
	}

	return nil
}

// validateMoveTaskOrderID carries on validations for parameter MoveTaskOrderID
func (o *CreateMTOServiceItemParams) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if err := validate.FormatOf("moveTaskOrderID", "path", "uuid", o.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}
	return nil
}