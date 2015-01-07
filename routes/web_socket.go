package routes

import (
	"github.com/gorilla/websocket"
	"net/http"
)

const WebSocketProtocolHeader = "Sec-WebSocket-Protocol"

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

func ExtractProtocolHeader(r *http.Request) http.Header {
	header := http.Header{}
	header.Add(WebSocketProtocolHeader, r.Header.Get(WebSocketProtocolHeader))
	return header
}
