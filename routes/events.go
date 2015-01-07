package routes

import (
	"net/http"

	"github.com/elos/server/data/models/event"
)

func eventsPost(w http.ResponseWriter, r *http.Request, Error ErrorHandlerConstructor, Resource ResourceHandlerConstructor) {
	event, err := event.Create(r.FormValue("name"), r.FormValue("user_id"))

	if err != nil {
		logf("An error occurred while create the event, err: %s", err)
		Error(err).ServeHTTP(w, r)
	} else {
		logf("Event was successfully created: %v", event)
		Resource(201, event).ServeHTTP(w, r)
	}
}

var EventsPost = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		eventsPost(w, r, NewErrorHandler, NewResourceHandler)
	},
)
