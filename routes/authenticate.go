package routes

import (
	"log"
	"net/http"

	"github.com/elos/server/models"
	"github.com/elos/server/util"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Authenticate(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		getHandler(w, r)
	default:
		util.InvalidMethod(w)
		return
	}

}

func getHandler(w http.ResponseWriter, r *http.Request) {

	id := r.Header.Get("Elos-ID")
	key := r.Header.Get("Elos-Key")

	user, authenticated, err := models.AuthenticateUser(id, key)

	if err != nil {
		log.Printf("%s", err)
		util.ServerError(w, err)
	}

	if authenticated {
		w.WriteHeader(200)

		bytes, _ := util.ToJson(user)
		w.Write(bytes)

		return
		/*
			var ws *websocket.Conn

			if ws, err := upgrader.Upgrade(w, r, nil); err != nil {
				log.Println(err)
				return
			}
		*/

		// send to hub
	} else {
		w.WriteHeader(401)
		return
	}

}
