// Code generated by go-swagger; DO NOT EDIT.

package office_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateOfficeUserHandlerFunc turns a function with the right signature into a create office user handler
type CreateOfficeUserHandlerFunc func(CreateOfficeUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateOfficeUserHandlerFunc) Handle(params CreateOfficeUserParams) middleware.Responder {
	return fn(params)
}

// CreateOfficeUserHandler interface for that can handle valid create office user params
type CreateOfficeUserHandler interface {
	Handle(CreateOfficeUserParams) middleware.Responder
}

// NewCreateOfficeUser creates a new http.Handler for the create office user operation
func NewCreateOfficeUser(ctx *middleware.Context, handler CreateOfficeUserHandler) *CreateOfficeUser {
	return &CreateOfficeUser{Context: ctx, Handler: handler}
}

/*CreateOfficeUser swagger:route POST /office_users office_users createOfficeUser

create an office user

creates and returns an office user record

*/
type CreateOfficeUser struct {
	Context *middleware.Context
	Handler CreateOfficeUserHandler
}

func (o *CreateOfficeUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateOfficeUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
