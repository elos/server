package transfer

import (
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/util"
)

type Router struct {
	Handlers map[string]ActionHandler
	Actions  map[string]chan *Envelope
}

func (r *Router) Handler(action string, handler ActionHandler) {
}

func (r *Router) Route(e *Envelope, db data.DB, c conn.Connection) {
}

type ActionHandler func(*Envelope, data.DB, conn.Connection)

func Route(e *Envelope, db data.DB, c conn.Connection) {
	switch e.Action {
	case POST:
		go PostHandler(e, db, c)
	case GET:
		go GetHandler(e, db, c)
	case DELETE:
		go DeleteHandler(e, db, c)
	case SYNC:
		go SyncHandler(e, db, c)
	case ECHO:
		go EchoHandler(e, db, c)
	default:
		c.WriteJSON(util.NewInvalidMethodError())
	}
}
