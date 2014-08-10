package routes

import (
	"log"
	"net/http"

	"github.com/elos/server/models"
	"github.com/elos/server/util"
)

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		postHandler(w, r)
	default:
		util.InvalidMethod(w)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	user, err := models.CreateUser(r.FormValue("name"))

	if err != nil {
		log.Printf("An error occurred while creating the user, err: %s", err)
		util.ServerError(w, err)
	} else {
		if *Verbose {
			log.Print("User was successfully created: %v", user)
		}

		util.ResourceResponse(w, 201, user)
	}
}
