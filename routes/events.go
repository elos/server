package routes

import (
	"net/http"

	"github.com/elos/server/data/models/event"
)

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

var EventsPost = FunctionHandler(eventsPost)
