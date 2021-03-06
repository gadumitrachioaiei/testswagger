// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TestrawjsonHandlerFunc turns a function with the right signature into a testrawjson handler
type TestrawjsonHandlerFunc func(TestrawjsonParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TestrawjsonHandlerFunc) Handle(params TestrawjsonParams) middleware.Responder {
	return fn(params)
}

// TestrawjsonHandler interface for that can handle valid testrawjson params
type TestrawjsonHandler interface {
	Handle(TestrawjsonParams) middleware.Responder
}

// NewTestrawjson creates a new http.Handler for the testrawjson operation
func NewTestrawjson(ctx *middleware.Context, handler TestrawjsonHandler) *Testrawjson {
	return &Testrawjson{Context: ctx, Handler: handler}
}

/* Testrawjson swagger:route GET /rawjson test testrawjson

An endpoint that returns json data but the generated client does not deserialize

*/
type Testrawjson struct {
	Context *middleware.Context
	Handler TestrawjsonHandler
}

func (o *Testrawjson) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTestrawjsonParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
