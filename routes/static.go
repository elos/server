package routes

import (
	"net/http"

	"github.com/elos/server/util"
)

func Static(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		staticGetHandler(w, r)
	default:
		util.InvalidMethod(w)
	}
}

func staticGetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
