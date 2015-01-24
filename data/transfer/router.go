package transfer

import (
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/util"
)

type Router struct {
	Handlers map[string]ActionHandler
	Actions  map[string]chan *data.Envelope
}

func (r *Router) Handler(action string, handler ActionHandler) {
}

func (r *Router) Route(e *data.Envelope, db data.DB, c conn.Connection) {
}

type ActionHandler func(*data.Envelope, data.DB, conn.Connection)

func Route(e *data.Envelope, db data.DB, c conn.Connection) {
	switch e.Action {
	case data.POST:
		go PostHandler(e, db, c)
	case data.GET:
		go GetHandler(e, db, c)
	case data.DELETE:
		go DeleteHandler(e, db, c)
	case data.SYNC:
		go SyncHandler(e, db, c)
	case data.ECHO:
		go EchoHandler(e, db, c)
	default:
		c.WriteJSON(util.NewInvalidMethodError())
	}
}
