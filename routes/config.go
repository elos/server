package routes

import (
	"fmt"
	"github.com/elos/server/data"
	"github.com/elos/server/data/test"
	"github.com/elos/server/util"
	"github.com/elos/server/util/auth"
	"net/http"
)

type HandlerMap map[string]http.Handler

func (h HandlerMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := h[r.Method]
	if ok {
		handler.ServeHTTP(w, r)
	} else {
		http.NotFoundHandler().ServeHTTP(w, r)
	}
}

const POST string = "POST"
const GET string = "GET"

var HTTPMethods map[string]bool = map[string]bool{
	POST: true,
	GET:  true,
}

type httpMethodHandler struct {
	Methods map[string]http.Handler
}

func (h *httpMethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := h.Methods[r.Method]
	if ok {
		handler.ServeHTTP(w, r)
	} else {
		invalidMethodHandler(w)
	}
}

func (h *httpMethodHandler) Handle(method string, handler http.Handler) {
	h.Methods[method] = handler
}

func HTTPMethodHandler() *httpMethodHandler {
	return &httpMethodHandler{
		Methods: make(map[string]http.Handler),
	}
}

func join(prefix string, route string) string {
	return fmt.Sprintf("%s/%s", prefix, route)
}

func SetupRoutes(hm HandlerMap, prefix string) {
	methodHandler := HTTPMethodHandler()
	for routeName, handler := range hm {
		// type assert
		subHM, ok := handler.(HandlerMap)

		// We are being pointed to another handler map
		if ok {
			SetupRoutes(subHM, join(prefix, routeName))
		} else { // this is a handler
			_, isHTTPMethod := HTTPMethods[routeName]
			if isHTTPMethod {
				methodHandler.Handle(routeName, handler)
			} else {
				log("this functionality is not defined")
			}
		}
	}
	if prefix != "" {
		http.Handle(prefix, methodHandler)
	}
}

var DefaultAuthenticator auth.RequestAuthenticator = auth.AuthenticateRequest

/*
	InvalidMethodHandler
*/

// The function a route will write to if it can't handle it's method
type ResponseHandler func(http.ResponseWriter)

// The default invalidMethodHandler is the one defined in server/util
var DefaultInvalidMethodHandler ResponseHandler = util.InvalidMethod

// Always have a invalidMethodHandler, private: set with SetInvalidMethodHandler

var invalidMethodHandler ResponseHandler = DefaultInvalidMethodHandler

// The entire routes package uses the same one
func SetInvalidMethodHandler(handler ResponseHandler) {
	if handler != nil {
		invalidMethodHandler = handler
	}
}

// The default Unauthorized handler is the one provided by util
var DefaultUnauthorizedHandler ResponseHandler = util.Unauthorized

// Always start with an unauthorized response handler
var unauthorizedHandler ResponseHandler = DefaultUnauthorizedHandler

func SetUnauthorizedHandler(handler ResponseHandler) {
	if handler != nil {
		unauthorizedHandler = handler
	}
}

type ResourceResponseHandler func(http.ResponseWriter, int, interface{})

var DefaultResourceResponseHandler ResourceResponseHandler = util.WriteResourceResponse

var resourceResponseHandler = DefaultResourceResponseHandler

func SetResourceResponseHandler(handler ResourceResponseHandler) {
	if handler != nil {
		resourceResponseHandler = handler
	}
}

type AuthRouteHandler func(http.ResponseWriter, *http.Request, auth.RequestAuthenticator)

/*
	Data
*/

// The default database is the null test database
var DefaultDatabase data.DB = test.NewDB()

// Always have a db, private: set with SetDatabase
var db data.DB = DefaultDatabase

// Set the database with which the routes look for data
func SetDatabase(newDB data.DB) {
	if newDB != nil {
		db = newDB
	}
}
