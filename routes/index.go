package routes

import (
	"net/http"

	"github.com/elos/server/util"
)

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		indexGetHandler(w, r)
	default:
		util.InvalidMethod(w)
	}
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
