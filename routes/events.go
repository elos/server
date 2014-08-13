package routes

import (
	"log"
	"net/http"

	"github.com/elos/server/models"
	"github.com/elos/server/util"
)

func Events(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		eventsPostHandler(w, r)
	default:
		util.InvalidMethod(w)
	}
}

func eventsPostHandler(w http.ResponseWriter, r *http.Request) {
	event, err := models.CreateEvent(r.FormValue("name"), r.FormValue("user_id"))

	if err != nil {
		log.Printf("[Routes] An error occurred while create the event, err: %s", err)
		util.ServerError(w, err)
	} else {
		util.Logf("[Routes] Event was successfully created: %v", event)

		util.ResourceResponse(w, 201, event)
	}
}
