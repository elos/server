package routes

import (
	"net/http"

	"github.com/elos/server/models/user"
	"github.com/elos/server/util"
)

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case POST:
		usersPostHandler(w, r)
	default:
		invalidMethodHandler(w)
	}
}

func usersPostHandler(w http.ResponseWriter, r *http.Request) {
	user, err := user.Create(r.FormValue("name"))

	if err != nil {
		logf("An error occurred while creating the user, err: %s", err)
		serverErrorHandler(w, err)
	} else {
		logf("User was successfully created: %v", user)

		util.WriteResourceResponse(w, 201, user)
	}
}
