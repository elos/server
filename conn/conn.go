package conn

import (
	"errors"
	"github.com/elos/data"
)

var ConnectionClosedError = errors.New("SocketConnection is closed")

type AnonConnection interface {
	WriteJSON(interface{}) error
	ReadJSON(interface{}) error
	Close() error
}

type Connection interface {
	AnonConnection
	Agent() data.Identifiable
}
