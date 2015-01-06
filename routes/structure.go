package routes

import (
	"fmt"
	"net/http"
)

// HTTPMethods {{{

const POST string = "POST"
const GET string = "GET"

var HTTPMethods = map[string]bool{
	POST: true,
	GET:  true,
}

// HTTPMethods }}}

// HandlerMap {{{

type HandlerMap map[string]http.Handler

func (h HandlerMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log("You shouldn't be calling ServeHTTP on a handler map")
}

// HandlerMap }}}

// HTTPMethodHandler {{{

// Redirects http requests based on the requests HTTP method
type httpMethodHandler struct {
	Methods map[string]http.Handler
}

// Satisfies http.Handler interface, will dispatch ServeHTTP to
// one of it's method handlers, if one doesn't exist for the
// specified method then it handles the response with the invalidMethodHandler
func (h *httpMethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := h.Methods[r.Method]
	if ok {
		handler.ServeHTTP(w, r)
	} else {
		InvalidMethod(r.Method).ServeHTTP(w, r)
	}
}

// Registers a handler for a method
func (h *httpMethodHandler) Handle(method string, handler http.Handler) {
	h.Methods[method] = handler
}

// Creates a new httpMethodHandler
func HTTPMethodHandler() *httpMethodHandler {
	return &httpMethodHandler{
		Methods: make(map[string]http.Handler),
	}
}

// HTTPMethodHandler }}}

// SetupRoutes {{{

// joins a route prefix with the route
// e.g., join("/hey", "ho") => "/hey/ho"
func join(prefix string, route string) string {
	return fmt.Sprintf("%s/%s", prefix, route)
}

// Exported SetupRoutes calls the recursively defined setupRoutes helper
func SetupRoutes(hm HandlerMap) {
	setupRoutes(hm, "")
}

// recursively sets up the routes
func setupRoutes(hm HandlerMap, prefix string) {
	methodHandler := HTTPMethodHandler()
	for routeName, handler := range hm {
		// type assert
		subHM, ok := handler.(HandlerMap)

		// We are being pointed to another handler map
		if ok {
			setupRoutes(subHM, join(prefix, routeName))
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

// SetupRoutes }}}
