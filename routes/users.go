package routes

import (
	"net/http"

	"github.com/elos/server/data/models/user"
)

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

var UsersPost = FunctionHandler(usersPost)
