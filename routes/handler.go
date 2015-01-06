package routes

import (
	"github.com/elos/server/data"
	"github.com/elos/server/util"
	"github.com/elos/server/util/auth"
	"net/http"
)

// ServerError {{{

type serverErrorHandler struct {
	err error
}

func (h *serverErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.ServerError(w, h.err)
}

type ErrorHandler func(error) http.Handler

func ServerError(err error) http.Handler {
	return &serverErrorHandler{
		err: err,
	}
}

// }}}

// ResourceHandler {{{
type resourceHandler struct {
	code     int
	resource interface{}
}

func (h *resourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.WriteResourceResponse(w, h.code, h.resource)
}

type ResourceHandler func(int, interface{}) http.Handler

func Resource(code int, resource interface{}) http.Handler {
	return &resourceHandler{
		code:     code,
		resource: resource,
	}
}

// }}}

type AuthenticationHandler struct {
	authenticator       auth.RequestAuthenticator
	errorHandler        ErrorHandler
	unauthorizedHandler ErrorHandler
	transferFunc        AuthenticatedHandlerFunc
}

type AuthenticatedHandlerFunc func(http.ResponseWriter, *http.Request, data.Agent)

type authenticatedHandler struct {
	agent data.Agent
	fn    AuthenticatedHandlerFunc
}

func (h *authenticatedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.fn(w, r, h.agent)
}

func AuthenticatedHandler(agent data.Agent, fn AuthenticatedHandlerFunc) http.Handler {
	return &authenticatedHandler{
		agent: agent,
		fn:    fn,
	}
}

func (h *AuthenticationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	agent, authenticated, err := h.authenticator(r)

	if err != nil {
		logf("An error occurred during authentication, err: %s", err)
		h.errorHandler(err).ServeHTTP(w, r)
		return
	}

	if authenticated {
		AuthenticatedHandler(agent, h.transferFunc).ServeHTTP(w, r)
		logf("Agent with id %s authenticated", agent.GetId())
	} else {
		h.unauthorizedHandler(nil).ServeHTTP(w, r)
	}
}

func AuthenticateRoute(auth auth.RequestAuthenticator, Error ErrorHandler, Unauth ErrorHandler, Transfer AuthenticatedHandlerFunc) http.Handler {
	return &AuthenticationHandler{
		authenticator:       auth,
		errorHandler:        Error,
		unauthorizedHandler: Unauth,
		transferFunc:        Transfer,
	}
}

func FunctionHandler(fn http.HandlerFunc) http.Handler {
	return &functionHandler{
		fn: fn,
	}
}

type functionHandler struct {
	fn http.HandlerFunc
}

func (fh *functionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fh.fn(w, r)
}

func Route(fn http.HandlerFunc) http.Handler {
	return FunctionHandler(fn)
}
