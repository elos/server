package routes

import (
	"fmt"
	"github.com/elos/server/data"
	"github.com/elos/server/data/test"
	"github.com/elos/server/util"
	"github.com/elos/server/util/auth"
	"github.com/gorilla/websocket"
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
		log("Hey")
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

type RouteName string

var UsersRoute RouteName = "users"
var EventsRoute RouteName = "events"
var AuthenticateRoute RouteName = "authenticate"

type RouteHandler struct {
	RouteName RouteName
}

func (h *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	possibleRoutes := Routes[h]
	handlerFunc, exists := possibleRoutes[r.Method]
	if exists {
		handlerFunc(w, r)
	} else {
		invalidMethodHandler(w)
	}
}

var UsersHandler = &RouteHandler{RouteName: UsersRoute}
var EventsHandler = &RouteHandler{RouteName: EventsRoute}
var AuthenticateHandler = &RouteHandler{RouteName: AuthenticateRoute}

var Routes map[*RouteHandler]map[string]http.HandlerFunc = map[*RouteHandler]map[string]http.HandlerFunc{
	UsersHandler: {
		POST: usersPost,
	},
	EventsHandler: {
		POST: eventsPost,
	},
	AuthenticateHandler: {
		GET: tempAuth,
	},
}

var DefaultAuthenticator auth.RequestAuthenticator = auth.AuthenticateRequest

func tempAuth(w http.ResponseWriter, r *http.Request) {
	authenticateGet(w, r, DefaultAuthenticator)
}

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

// A response handler that also takes an error
type ErrorResponseHandler func(http.ResponseWriter, error)

// The default server error handler is the one provided by server/util
var DefaultServerErrorHandler ErrorResponseHandler = util.ServerError

// Always start with a serverErrorHandler
var serverErrorHandler ErrorResponseHandler = DefaultServerErrorHandler

func SetServerErrorHandler(handler ErrorResponseHandler) {
	if handler != nil {
		serverErrorHandler = handler
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

/*
	Web Sockets
*/

// the utility a route will use to upgrade a request to a websocket
type WebSocketUpgrader interface {
	Upgrade(http.ResponseWriter, *http.Request, http.Header) (*websocket.Conn, error)
}

// A good default upgrader from gorilla/socket
var DefaultWebSocketUpgrader WebSocketUpgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Always start with an upgrader, private: set with SetWebSocketUpgrader
var webSocketUpgrader WebSocketUpgrader = DefaultWebSocketUpgrader

// Sets the websocket upgrader to be used by a route attempting an upgrade
func SetWebSocketUpgrader(u WebSocketUpgrader) {
	if u != nil {
		webSocketUpgrader = u
	}
}
