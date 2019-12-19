// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateMTOServiceItemstatusHandlerFunc turns a function with the right signature into a update m t o service itemstatus handler
type UpdateMTOServiceItemstatusHandlerFunc func(UpdateMTOServiceItemstatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateMTOServiceItemstatusHandlerFunc) Handle(params UpdateMTOServiceItemstatusParams) middleware.Responder {
	return fn(params)
}

// UpdateMTOServiceItemstatusHandler interface for that can handle valid update m t o service itemstatus params
type UpdateMTOServiceItemstatusHandler interface {
	Handle(UpdateMTOServiceItemstatusParams) middleware.Responder
}

// NewUpdateMTOServiceItemstatus creates a new http.Handler for the update m t o service itemstatus operation
func NewUpdateMTOServiceItemstatus(ctx *middleware.Context, handler UpdateMTOServiceItemstatusHandler) *UpdateMTOServiceItemstatus {
	return &UpdateMTOServiceItemstatus{Context: ctx, Handler: handler}
}

/*UpdateMTOServiceItemstatus swagger:route PATCH /move-task-orders/{moveTaskOrderID}/service-items/{mtoServiceItemID}/status mtoServiceItem updateMTOServiceItemstatus

Change the status of a line item for a move order by ID

Changes the status of a line item for a move order by ID

*/
type UpdateMTOServiceItemstatus struct {
	Context *middleware.Context
	Handler UpdateMTOServiceItemstatusHandler
}

func (o *UpdateMTOServiceItemstatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateMTOServiceItemstatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
