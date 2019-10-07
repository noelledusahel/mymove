// Code generated by go-swagger; DO NOT EDIT.

package service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteServiceItemParams creates a new DeleteServiceItemParams object
// no default values defined in spec.
func NewDeleteServiceItemParams() DeleteServiceItemParams {

	return DeleteServiceItemParams{}
}

// DeleteServiceItemParams contains all the bound params for the delete service item operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteServiceItem
type DeleteServiceItemParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of move order to use
	  Required: true
	  In: path
	*/
	MoveTaskOrderID string
	/*ID of line item to use
	  Required: true
	  In: path
	*/
	ServiceItemID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteServiceItemParams() beforehand.
func (o *DeleteServiceItemParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rMoveTaskOrderID, rhkMoveTaskOrderID, _ := route.Params.GetOK("moveTaskOrderID")
	if err := o.bindMoveTaskOrderID(rMoveTaskOrderID, rhkMoveTaskOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	rServiceItemID, rhkServiceItemID, _ := route.Params.GetOK("serviceItemID")
	if err := o.bindServiceItemID(rServiceItemID, rhkServiceItemID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMoveTaskOrderID binds and validates parameter MoveTaskOrderID from path.
func (o *DeleteServiceItemParams) bindMoveTaskOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MoveTaskOrderID = raw

	return nil
}

// bindServiceItemID binds and validates parameter ServiceItemID from path.
func (o *DeleteServiceItemParams) bindServiceItemID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ServiceItemID = raw

	return nil
}
