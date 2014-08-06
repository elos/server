package routes

import (
	"net/http"

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
		http.Error(w, "Method not allowed", 405)
		return
	}

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// var ws *websocket.Conn

	/*
		if ws, err := upgrader.Upgrade(w, r, nil); err != nil {
			log.Println(err)
			return
		}
	*/

	// authenticate

	// send to hub
}
