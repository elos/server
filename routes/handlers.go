package routes

import (
	"github.com/elos/server/data"
	"github.com/elos/server/util"
	"github.com/elos/server/util/auth"
	"net/http"
)

// NullHandler (Testing) {{{

type NullHandler struct {
	Handled map[*http.Request]bool
}

func (h *NullHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handled[r] = true
}

func (h *NullHandler) Reset() *NullHandler {
	h.Handled = make(map[*http.Request]bool)
	return h
}

func NewNullHandler() *NullHandler {
	return (&NullHandler{}).Reset()
}

// NullHandler (Testing) }}}

//  ErrorHandler {{{

// Allows route to handle an error
type ErrorHandlerConstructor func(error) http.Handler

// underlying information needed to form an error response
type ErrorHandler struct {
	Err error
}

// implement http.Handler
func (h *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.ServerError(w, h.Err)
}

// Returns a handler capable of serving the error information
func NewErrorHandler(err error) http.Handler {
	return &ErrorHandler{
		Err: err,
	}
}

// }}}

// ResourceHandler {{{

// Allows route to handle returning a json resource
type ResourceHandlerConstructor func(int, interface{}) http.Handler

// underlying information needed to return a json resource
type ResourceHandler struct {
	Code     int
	Resource interface{}
}

// implements http.Handler
func (h *ResourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.WriteResourceResponse(w, h.Code, h.Resource)
}

// Returns a handler capable of serving the resource
func NewResourceHandler(code int, resource interface{}) http.Handler {
	return &ResourceHandler{
		Code:     code,
		Resource: resource,
	}
}

// }}}

// BadMethodHandler {{{

/*
	Allows route to handle a suspected invalid method
	- Should only be used by HTTPMethodHandler
*/
type BadMethodHandlerConstructor func(*http.Request) http.Handler

// underlying information need to notify user of invalid method
type BadMethodHandler struct {
	RequestedMethod string
}

// implemens http.Handler
func (h *BadMethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.InvalidMethod(w)
}

// Returns a handler capable of notifying the user of the invalid method
func NewBadMethodHandler(r *http.Request) http.Handler {
	return &BadMethodHandler{
		RequestedMethod: r.Method,
	}
}

// InvalidMethodHandler }}}

// UnauthorizedHandler {{{

// Allows a route to indicate the agent is unauthorized
type UnauthorizedHandlerConstructor func(string) http.Handler

/*
	underlying information necessary to inform client of
	lack of credentials
*/
type UnauthorizedHandler struct {
	Reason string
}

// implements http.Handler
func (h *UnauthorizedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.Unauthorized(w)
}

// Returns a handler capable of serving the unauthorized error
func NewUnauthorizedHandler(reason string) http.Handler {
	return &UnauthorizedHandler{
		Reason: reason,
	}
}

// UnauthorizedHandler }}}

// Authenticators {{{

var DefaultAuthenticator auth.RequestAuthenticator = auth.AuthenticateRequest

// Authenticators }}}

// AuthenticationHandler {{{

type AuthenticationHandler struct {
	Authenticator          auth.RequestAuthenticator
	NewErrorHandler        ErrorHandlerConstructor
	NewUnauthorizedHandler UnauthorizedHandlerConstructor
	TransferFunc           AuthenticatedHandlerFunc
}

func NewAuthenticationHandler(a auth.RequestAuthenticator, eh ErrorHandlerConstructor,
	uh UnauthorizedHandlerConstructor, t AuthenticatedHandlerFunc) http.Handler {
	return &AuthenticationHandler{
		Authenticator:          a,
		NewErrorHandler:        eh,
		NewUnauthorizedHandler: uh,
		TransferFunc:           t,
	}
}

func (h *AuthenticationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	agent, authenticated, err := h.Authenticator(r)

	if err != nil {
		logf("An error occurred during authentication, err: %s", err)
		h.NewErrorHandler(err).ServeHTTP(w, r)
		return
	}

	if authenticated {
		NewAuthenticatedHandler(agent, h.TransferFunc).ServeHTTP(w, r)
		logf("Agent with id %s authenticated", agent.GetId())
	} else {
		h.NewUnauthorizedHandler("Not authenticated").ServeHTTP(w, r)
	}
}

// AuthenticationHandler }}}

// AuthenticatedHandlerFunc {{{

type AuthenticatedHandlerFunc func(http.ResponseWriter, *http.Request, data.Agent)

type AuthenticatedHandler struct {
	Agent data.Agent
	Fn    AuthenticatedHandlerFunc
}

func (h *AuthenticatedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Fn(w, r, h.Agent)
}

func NewAuthenticatedHandler(agent data.Agent, fn AuthenticatedHandlerFunc) http.Handler {
	return &AuthenticatedHandler{
		Agent: agent,
		Fn:    fn,
	}
}

// AuthenticatedHandlerFunc }}}
