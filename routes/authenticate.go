package routes

import (
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
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHandler(w, r)
	default:
		util.InvalidMethod(w)
	}
}

func ExtractCredentials(r *http.Request) (string, string) {
	tokens := strings.Split(r.Header.Get("Sec-WebSocket-Protocol"), "-")
	// Query Parameter Method of Authentication
	/*
		id := r.FormValue("id")
		key := r.FormValue("key")
	*/
	if len(tokens) != 2 {
		Log("The length of the tokens extrapolated from Sec-Websocket-Protocol was not 2")
		return "", ""
	} else {
		return tokens[0], tokens[1]
	}
}

func ExtractProtocolHeader(r *http.Request) *http.Header {
	protocol := http.Header{
		"Sec-WebSocket-Protocol": []string{r.Header.Get("Sec-WebSocket-Protocol")},
	}

	return &protocol
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Use the WebSocket protocol header to identify and authenticate the user
	id, key := ExtractCredentials(r)

	if id == "" || key == "" {
		util.Unauthorized(w)
		return
	}

	// Prevents an error that should be dealt with within AuthenticateUser
	// The empty string breaks the authenticate function
	if id == "" {
		Log("The id extrapolated from the Sec-Websocket-Protocol was: \"\"")

		util.Unauthorized(w)
		return
	} // this now appears redundant?

	user, authenticated, err := user.Authenticate(id, key)

	if err != nil {
		Logf("An error occurred during authentication, err: %s", err)
		util.ServerError(w, err)
		return
	}

	if authenticated {
		Logf("User with id %s authenticated", id)

		ws, err := upgrader.Upgrade(w, r, *ExtractProtocolHeader(r))

		if err != nil {
			Logf("An error occurred while upgrading to the websocket protocol, err: %s", err)
			// gorilla/websocket handles response to client
			return
		}

		Logf("User with id %s just connected over websocket", id)

		sockets.NewConnection(user, ws)
	} else {
		Logf("User with id %s failed authentication", id)

		util.Unauthorized(w)
		return
	}

}
