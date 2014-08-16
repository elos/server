package routes

import (
	"log"
	"net/http"

	"github.com/elos/server/models/user"
	"github.com/elos/server/util"
)

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		usersPostHandler(w, r)
	default:
		util.InvalidMethod(w)
	}
}

func usersPostHandler(w http.ResponseWriter, r *http.Request) {
	user, err := user.Create(r.FormValue("name"))

	if err != nil {
		log.Printf("An error occurred while creating the user, err: %s", err)
		util.ServerError(w, err)
	} else {
		util.Logf("User was successfully created: %v", user)

		util.ResourceResponse(w, 201, user)
	}
}
