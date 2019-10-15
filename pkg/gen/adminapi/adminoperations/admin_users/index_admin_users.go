// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// IndexAdminUsersHandlerFunc turns a function with the right signature into a index admin users handler
type IndexAdminUsersHandlerFunc func(IndexAdminUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IndexAdminUsersHandlerFunc) Handle(params IndexAdminUsersParams) middleware.Responder {
	return fn(params)
}

// IndexAdminUsersHandler interface for that can handle valid index admin users params
type IndexAdminUsersHandler interface {
	Handle(IndexAdminUsersParams) middleware.Responder
}

// NewIndexAdminUsers creates a new http.Handler for the index admin users operation
func NewIndexAdminUsers(ctx *middleware.Context, handler IndexAdminUsersHandler) *IndexAdminUsers {
	return &IndexAdminUsers{Context: ctx, Handler: handler}
}

/*IndexAdminUsers swagger:route GET /admin_users admin_users indexAdminUsers

List admin users

Returns a list of admin users

*/
type IndexAdminUsers struct {
	Context *middleware.Context
	Handler IndexAdminUsersHandler
}

func (o *IndexAdminUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewIndexAdminUsersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}