package routes

import (
	"net/http"

	"github.com/elos/server/data"
	"github.com/elos/server/data/models/event"
)

type EventsPostHandler struct {
	data.DB
}

func (h *EventsPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	EventsPostFunction(w, r, NewErrorHandler, NewResourceHandler, h.DB)
}

func EventsPostFunction(w http.ResponseWriter, r *http.Request, Error ErrorHandlerConstructor, Resource ResourceHandlerConstructor, db data.DB) {
	event, err := event.Create(db, r.FormValue("name"), r.FormValue("user_id"))

	if err != nil {
		logf("An error occurred while create the event, err: %s", err)
		Error(err).ServeHTTP(w, r)
	} else {
		logf("Event was successfully created: %v", event)
		Resource(201, event).ServeHTTP(w, r)
	}
}
