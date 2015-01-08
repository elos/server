package conn

import (
	"errors"
	"github.com/elos/server/data"
)

var ConnectionClosedError = errors.New("SocketConnection is closed")

type Connection interface {
	Agent() data.Agent
	WriteJSON(interface{}) error
	ReadJSON(interface{}) error
	Close() error
}
