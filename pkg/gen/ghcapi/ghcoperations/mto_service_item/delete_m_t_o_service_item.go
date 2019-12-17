// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteMTOServiceItemHandlerFunc turns a function with the right signature into a delete m t o service item handler
type DeleteMTOServiceItemHandlerFunc func(DeleteMTOServiceItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMTOServiceItemHandlerFunc) Handle(params DeleteMTOServiceItemParams) middleware.Responder {
	return fn(params)
}

// DeleteMTOServiceItemHandler interface for that can handle valid delete m t o service item params
type DeleteMTOServiceItemHandler interface {
	Handle(DeleteMTOServiceItemParams) middleware.Responder
}

// NewDeleteMTOServiceItem creates a new http.Handler for the delete m t o service item operation
func NewDeleteMTOServiceItem(ctx *middleware.Context, handler DeleteMTOServiceItemHandler) *DeleteMTOServiceItem {
	return &DeleteMTOServiceItem{Context: ctx, Handler: handler}
}

/*DeleteMTOServiceItem swagger:route DELETE /move-task-orders/{moveTaskOrderID}/service-items/{mtoServiceItemID} mtoServiceItem deleteMTOServiceItem

Deletes a line item by ID for a move order by ID

Deletes a line item by ID for a move order by ID

*/
type DeleteMTOServiceItem struct {
	Context *middleware.Context
	Handler DeleteMTOServiceItemHandler
}

func (o *DeleteMTOServiceItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteMTOServiceItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
