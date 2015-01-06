package routes

import (
	"github.com/gorilla/websocket"
	"net/http"
)

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
