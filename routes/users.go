package routes

import (
	"net/http"

	"github.com/elos/server/data/models/user"
)

func usersPost(w http.ResponseWriter, r *http.Request, Error func(error) http.Handler, Resource func(int, interface{}) http.Handler) {
	user, err := user.Create(r.FormValue("name"))

	if err != nil {
		logf("An error occurred while creating the user, err: %s", err)
		Error(err).ServeHTTP(w, r)
	} else {
		logf("User was successfully created: %v", user)
		Resource(201, user).ServeHTTP(w, r)
	}
}

var UsersPost = Route(
	func(w http.ResponseWriter, r *http.Request) {
		usersPost(w, r, Error, Resource)
	},
)
