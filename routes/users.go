package routes

import (
	"net/http"

	"github.com/elos/data"
	"github.com/elos/models/user"
	"github.com/elos/transfer"
)

type UsersPostHandler struct {
	data.Store
}

func (h *UsersPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	UsersPostFunction(w, r, NewErrorHandler, NewResourceHandler, h.Store)
}

func UsersPostFunction(w http.ResponseWriter, r *http.Request, Error ErrorHandlerConstructor, Resource ResourceHandlerConstructor, s data.Store) {
	attrs := data.AttrMap{
		"name": r.FormValue("name"),
	}

	hack, _ := user.New(s)
	c := transfer.NewHTTPConnection(w, r, hack)
	e := transfer.New(c, transfer.POST, "user", attrs)
	go transfer.Route(e, s)
}
