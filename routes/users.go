package routes

import (
	"log"
	"net/http"

	"github.com/elos/server/models"
	"github.com/elos/server/util"
)

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch r.Method {
	case "POST":
		postHandler(w, r)
	default:
		util.InvalidMethod(w)
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	user, err := models.CreateUser(r.FormValue("name"))

	if err != nil {
		log.Print(err)
		util.ServerError(w, err)
	} else {
		bytes, _ := util.ToJson(user)

		// Default status is 200
		w.Write(bytes)
	}
}
