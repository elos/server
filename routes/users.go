package routes

import (
	"net/http"

	"github.com/elos/server/models"
)

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.Method {
	case "POST":
		postHandler(w, r)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	user, err := models.CreateUser()

	if err != nil {
		w.WriteHeader(404)
	} else {
		bytes, _ := user.ToJson()

		// Default status is 200
		w.Write(bytes)
	}
}
