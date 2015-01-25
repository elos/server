package routes

import (
	"net/http"

	"github.com/elos/server/data"
	"github.com/elos/server/models/user"
)

type UsersPostHandler struct {
	data.DB
}

func (h *UsersPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	UsersPostFunction(w, r, NewErrorHandler, NewResourceHandler, h.DB)
}

func UsersPostFunction(w http.ResponseWriter, r *http.Request, Error ErrorHandlerConstructor, Resource ResourceHandlerConstructor, db data.DB) {
	user, err := user.Create(db, r.FormValue("name"))

	if err != nil {
		logf("An error occurred while creating the user, err: %s", err)
		Error(err).ServeHTTP(w, r)
	} else {
		logf("User was successfully created: %v", user)
		Resource(201, user).ServeHTTP(w, r)
	}
}
