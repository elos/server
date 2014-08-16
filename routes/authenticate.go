package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/elos/server/models/user"
	"github.com/elos/server/sockets"
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
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Use the WebSocket protocol header to identify and authenticate the user
	tokens := strings.Split(r.Header.Get("Sec-WebSocket-Protocol"), "-")

	if len(tokens) != 2 {
		util.Log("The length of the tokens extrapolated from Sec-Websocket-Protocol was not 2")

		util.Unauthorized(w)
		return
	}

	var (
		id  string = tokens[0]
		key string = tokens[1]
	)

	// Query Parameter Method of Authentication
	/*
		id := r.FormValue("id")
		key := r.FormValue("key")
	*/

	// Prevents an error that should be dealt with within AuthenticateUser
	// The empty string breaks the authenticate function
	if id == "" {
		util.Log("The id extrapolated from the Sec-Websocket-Protocol was: \"\"")

		util.Unauthorized(w)
		return
	}

	user, authenticated, err := user.Authenticate(id, key)

	if err != nil {
		log.Printf("An error occurred during authentication, err: %s", err)
		util.ServerError(w, err)
		return
	}

	protocol := http.Header{
		"Sec-WebSocket-Protocol": []string{r.Header.Get("Sec-WebSocket-Protocol")},
	}

	if authenticated {
		util.Logf("User with id %s was authenticated", id)

		ws, err := upgrader.Upgrade(w, r, protocol)

		if err != nil {
			log.Printf("An error occurred while upgrading to the websocket protocol, err: %s", err)
			// gorilla/websocket handles response to client
			return
		}

		util.Logf("User with id %s just connected over websocket", id)

		sockets.NewConnection(user, ws)
	} else {
		util.Logf("User with id %s failed authentication", id)

		util.Unauthorized(w)
		return
	}

}
