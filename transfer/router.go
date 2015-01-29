package transfer

import (
	"github.com/elos/data"
	"github.com/elos/server/conn"
	"github.com/elos/server/util"
)

type Router struct {
	Handlers map[string]ActionHandler
	Actions  map[string]chan *Envelope
}

func (r *Router) Handler(action string, handler ActionHandler) {
}

func (r *Router) Route(e *Envelope, s data.Store, c conn.Connection) {
}

type ActionHandler func(*Envelope, data.Store, conn.Connection)

func Route(e *Envelope, s data.Store, c conn.Connection) {
	switch e.Action {
	case POST:
		go PostHandler(e, s, c)
	case GET:
		go GetHandler(e, s, c)
	case DELETE:
		go DeleteHandler(e, s, c)
	case SYNC:
		go SyncHandler(e, s, c)
	case ECHO:
		go EchoHandler(e, s, c)
	default:
		c.WriteJSON(util.NewInvalidMethodError())
	}
}
