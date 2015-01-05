package routes

import (
	"net/http"

	"github.com/elos/server/models/user"
)

var DefaultUsersPostHandler RouteHandler = usersPost
var usersPostHandler = DefaultUsersPostHandler

func SetUsersPostHandler(handler RouteHandler) {
	if handler != nil {
		usersPostHandler = handler
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case POST:
		usersPostHandler(w, r)
	default:
		invalidMethodHandler(w)
	}
}

func usersPost(w http.ResponseWriter, r *http.Request) {
	user, err := user.Create(r.FormValue("name"))

	if err != nil {
		logf("An error occurred while creating the user, err: %s", err)
		serverErrorHandler(w, err)
	} else {
		logf("User was successfully created: %v", user)
		resourceResponseHandler(w, 201, user)
	}
}
