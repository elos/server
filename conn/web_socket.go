package conn

import (
	"github.com/elos/server/data"
	"github.com/gorilla/websocket"
	"net/http"
)

const WebSocketProtocolHeader = "Sec-WebSocket-Protocol"

// the utility a route will use to upgrade a request to a websocket
type WebSocketUpgrader interface {
	Upgrade(http.ResponseWriter, *http.Request, data.Agent) (Connection, error)
}

// NullUpgrader {{{

type NullUpgrader struct {
	Upgraded   map[*http.Request]bool
	Connection Connection
	Error      error
}

func (u *NullUpgrader) Upgrade(w http.ResponseWriter, r *http.Request, a data.Agent) (Connection, error) {
	if u.Error != nil {
		return nil, u.Error
	}
	u.Upgraded[r] = true
	return u.Connection, nil
}

func (u *NullUpgrader) Reset() *NullUpgrader {
	u.Upgraded = make(map[*http.Request]bool)
	u.Error = nil
	return u
}

func NewNullUpgrader(c Connection) WebSocketUpgrader {
	return (&NullUpgrader{Connection: c}).Reset()
}

// }}}

// Gorilla Upgrader {{{
// wrapper for mgo upgrader
type GorillaUpgrader struct {
	Upgrader *websocket.Upgrader
}

func (u *GorillaUpgrader) Upgrade(w http.ResponseWriter, r *http.Request, a data.Agent) (Connection, error) {
	wc, err := u.Upgrader.Upgrade(w, r, ExtractProtocolHeader(r))
	if err != nil {
		return NewNullConnection(a), err
	}

	return NewGorillaConnection(wc, a), nil
}

func NewGorillaUpgrader(rbs int, wbs int, checkO bool) WebSocketUpgrader {
	u := &websocket.Upgrader{
		ReadBufferSize:  rbs,
		WriteBufferSize: wbs,
		CheckOrigin: func(r *http.Request) bool {
			return checkO
		},
	}
	return &GorillaUpgrader{
		Upgrader: u,
	}
}

// Gorilla Upgrader }}}

// A good default upgrader from gorilla/socket
var DefaultWebSocketUpgrader WebSocketUpgrader = NewGorillaUpgrader(1024, 1024, true)

func ExtractProtocolHeader(r *http.Request) http.Header {
	header := http.Header{}
	header.Add(WebSocketProtocolHeader, r.Header.Get(WebSocketProtocolHeader))
	return header
}
