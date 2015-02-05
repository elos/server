package routes

import (
	"net/http"

	"github.com/elos/data"
	"github.com/elos/models/user"
	"github.com/elos/transfer"
)

type EventsPostHandler struct {
	data.Store
}

func (h *EventsPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	EventsPostFunction(w, r, NewErrorHandler, NewResourceHandler, h.Store)
}

func EventsPostFunction(w http.ResponseWriter, r *http.Request, Error ErrorHandlerConstructor, Resource ResourceHandlerConstructor, s data.Store) {
	attrs := data.AttrMap{
		"name":    r.FormValue("name"),
		"user_id": r.FormValue("user_id"),
	}

	hack, _ := user.New(s)
	c := transfer.NewHTTPConnection(w, r, hack)
	e := transfer.New(c, transfer.POST, "event", attrs)
	go transfer.Route(e, s)
}
