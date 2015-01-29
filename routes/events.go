package routes

import (
	"net/http"

	"github.com/elos/data"
	"github.com/elos/server/models/event"
)

type EventsPostHandler struct {
	data.Store
}

func (h *EventsPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	EventsPostFunction(w, r, NewErrorHandler, NewResourceHandler, h.Store)
}

func EventsPostFunction(w http.ResponseWriter, r *http.Request, Error ErrorHandlerConstructor, Resource ResourceHandlerConstructor, s data.Store) {
	event, err := event.Create(s, r.FormValue("name"), r.FormValue("user_id"))

	if err != nil {
		logf("An error occurred while create the event, err: %s", err)
		Error(err).ServeHTTP(w, r)
	} else {
		logf("Event was successfully created: %v", event)
		Resource(201, event).ServeHTTP(w, r)
	}
}
