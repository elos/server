package routes

import (
	"net/http"

	"github.com/elos/server/data/models/user"
)

func UsersPostFunction(w http.ResponseWriter, r *http.Request, Error ErrorHandlerConstructor, Resource ResourceHandlerConstructor) {
	user, err := user.Create(r.FormValue("name"))

	if err != nil {
		logf("An error occurred while creating the user, err: %s", err)
		Error(err).ServeHTTP(w, r)
	} else {
		logf("User was successfully created: %v", user)
		Resource(201, user).ServeHTTP(w, r)
	}
}

var UsersPost = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		UsersPostFunction(w, r, NewErrorHandler, NewResourceHandler)
	},
)
