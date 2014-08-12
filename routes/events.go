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
		log.Printf("An error occurred while create the event, err: %s", err)
		util.ServerError(w, err)
	} else {
		util.Logf("Event was successfully created: %v", event)
		util.Log("Event user", event.GetUser())

		util.ResourceResponse(w, 201, event)
	}
}
