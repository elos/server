package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/elos/server/hub"
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
	tokens := strings.Split(r.Header.Get("Sec-WebSocket-Protocol"), "-")

	if len(tokens) != 2 {
		return
	}

	id := tokens[0]
	key := tokens[1]

	log.Print(r.FormValue("foo"))

	/*
		id := r.FormValue("id")
		key := r.FormValue("key")
	*/

	if id == "" {
		return
	}

	user, authenticated, err := models.AuthenticateUser(id, key)

	if err != nil {
		log.Printf("%s", err)
		return
		// util.ServerError(w, err)
	}

	if authenticated {
		log.Print("authenticated")
		ws, err := upgrader.Upgrade(w, r, r.Header)

		if err != nil {
			log.Println(err)
			return
		}

		hub.NewConnection(user, ws)
	} else {
		return
		// util.Unauthorized(w)
	}

}
