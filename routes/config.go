package routes

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/test"
	"github.com/elos/server/util"
	"github.com/elos/server/util/logging"
	"github.com/gorilla/websocket"
	"net/http"
)

/*
	HTTP Methods
*/

// Post method string
const POST string = "POST"

// Get method string
const GET string = "GET"

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

/*
	Logging
*/

// The name of this package as a service for the server
const ServiceName string = "Routes"

func log(v ...interface{}) {
	logging.Log.Logs(ServiceName, v...)
}

func logf(format string, v ...interface{}) {
	logging.Log.Logsf(ServiceName, format, v...)
}

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
