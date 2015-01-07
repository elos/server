package routes

import (
	"github.com/elos/server/data"
	"github.com/elos/server/util"
	"github.com/elos/server/util/auth"
	"net/http"
)

// NullHandler (Testing) {{{

type NullHandlerType struct {
	Handled map[*http.Request]bool
}

func (h *NullHandlerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handled[r] = true
}

func NullHandler() http.Handler {
	return &NullHandlerType{Handled: make(map[*http.Request]bool)}
}

// NullHandler (Testing) }}}

//  ErrorHandler {{{

// Allows route to handle an error
type ErrorHandler func(error) http.Handler

// underlying information needed to form an error response
type serverErrorHandler struct {
	err error
}

// implement http.Handler
func (h *serverErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.ServerError(w, h.err)
}

// Returns a handler capable of serving the error information
func Error(err error) http.Handler {
	return &serverErrorHandler{
		err: err,
	}
}

// }}}

// ResourceHandler {{{

// Allows route to handle returning a json resource
type ResourceHandler func(int, interface{}) http.Handler

// underlying information needed to return a json resource
type resourceHandler struct {
	code     int
	resource interface{}
}

// implements http.Handler
func (h *resourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.WriteResourceResponse(w, h.code, h.resource)
}

// Returns a handler capable of serving the resource
func Resource(code int, resource interface{}) http.Handler {
	return &resourceHandler{
		code:     code,
		resource: resource,
	}
}

// }}}

// InvalidMethodHandler {{{

/*
	Allows route to handle a suspected invalid method
	- Should only be used by HTTPMethodHandler
*/
type InvalidMethodHandler func(string) http.Handler

// underlying information need to notify user of invalid method
type invalidMethodHandler struct {
	requestedMethod string
}

// implemens http.Handler
func (h *invalidMethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.InvalidMethod(w)
}

// Returns a handler capable of notifying the user of the invalid method
func InvalidMethod(method string) http.Handler {
	return &invalidMethodHandler{
		requestedMethod: method,
	}
}

// InvalidMethodHandler }}}

// UnauthorizedHandler {{{

// Allows a route to indicate the agent is unauthorized
type UnauthorizedHandler func(string) http.Handler

/*
	underlying information necessary to inform client of
	lack of credentials
*/
type unauthorizedHandler struct {
	reason string
}

// implements http.Handler
func (h *unauthorizedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.Unauthorized(w)
}

// Returns a handler capable of serving the unauthorized error
func Unauthorized(reason string) http.Handler {
	return &unauthorizedHandler{
		reason: reason,
	}
}

// UnauthorizedHandler }}}

// Authenticators {{{

var DefaultAuthenticator auth.RequestAuthenticator = auth.AuthenticateRequest

// Authenticators }}}

// AuthenticationHandler {{{

type authenticationHandler struct {
	authenticator       auth.RequestAuthenticator
	errorHandler        ErrorHandler
	unauthorizedHandler UnauthorizedHandler
	transferFunc        AuthenticatedHandlerFunc
}

func AuthenticateRoute(a auth.RequestAuthenticator, eh ErrorHandler,
	uh UnauthorizedHandler, t AuthenticatedHandlerFunc) http.Handler {
	return &authenticationHandler{
		authenticator:       a,
		errorHandler:        eh,
		unauthorizedHandler: uh,
		transferFunc:        t,
	}
}

func (h *authenticationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		h.unauthorizedHandler("Not authenticated").ServeHTTP(w, r)
	}
}

// AuthenticationHandler }}}

// AuthenticatedHandlerFunc {{{

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

// AuthenticatedHandlerFunc }}}
