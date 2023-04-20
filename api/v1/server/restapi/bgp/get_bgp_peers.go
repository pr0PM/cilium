// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package bgp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBgpPeersHandlerFunc turns a function with the right signature into a get bgp peers handler
type GetBgpPeersHandlerFunc func(GetBgpPeersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBgpPeersHandlerFunc) Handle(params GetBgpPeersParams) middleware.Responder {
	return fn(params)
}

// GetBgpPeersHandler interface for that can handle valid get bgp peers params
type GetBgpPeersHandler interface {
	Handle(GetBgpPeersParams) middleware.Responder
}

// NewGetBgpPeers creates a new http.Handler for the get bgp peers operation
func NewGetBgpPeers(ctx *middleware.Context, handler GetBgpPeersHandler) *GetBgpPeers {
	return &GetBgpPeers{Context: ctx, Handler: handler}
}

/*
	GetBgpPeers swagger:route GET /bgp/peers bgp getBgpPeers

# Lists operational state of BGP peers

Retrieves current operational state of BGP peers created by
Cilium BGP virtual router. This includes session state, uptime,
information per address family, etc.
*/
type GetBgpPeers struct {
	Context *middleware.Context
	Handler GetBgpPeersHandler
}

func (o *GetBgpPeers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBgpPeersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
