package routes

import (
	"net/http"

	"github.com/elos/server/models/event"
)

var DefaultEventsPostHandler RouteHandler = eventsPost
var eventsPostHandler = DefaultEventsPostHandler

func SetEventsPostHandler(handler RouteHandler) {
	if handler != nil {
		eventsPostHandler = handler
	}
}

func Events(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case POST:
		eventsPostHandler(w, r)
	default:
		invalidMethodHandler(w)
	}
}

func eventsPost(w http.ResponseWriter, r *http.Request) {
	event, err := event.Create(r.FormValue("name"), r.FormValue("user_id"))

	if err != nil {
		logf("An error occurred while create the event, err: %s", err)

		serverErrorHandler(w, err)
	} else {
		logf("Event was successfully created: %v", event)

		resourceResponseHandler(w, 201, event)
	}
}
