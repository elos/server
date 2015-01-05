package routes

import (
	"net/http"
	"strings"

	"github.com/elos/server/models/user"
	"github.com/elos/server/sockets"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case GET:
		authenticateGetHandler(w, r)
	default:
		invalidMethodHandler(w)
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
		log("The length of the tokens extrapolated from Sec-Websocket-Protocol was not 2")
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

func authenticateGetHandler(w http.ResponseWriter, r *http.Request) {
	// Use the WebSocket protocol header to identify and authenticate the user
	id, key := ExtractCredentials(r)

	if id == "" || key == "" {
		unauthorizedHandler(w)
		return
	}

	// Prevents an error that should be dealt with within AuthenticateUser
	// The empty string breaks the authenticate function
	if id == "" {
		log("The id extrapolated from the Sec-Websocket-Protocol was: \"\"")

		unauthorizedHandler(w)
		return
	} // this now appears redundant?

	user, authenticated, err := user.Authenticate(id, key)

	if err != nil {
		logf("An error occurred during authentication, err: %s", err)
		serverErrorHandler(w, err)
		return
	}

	if authenticated {
		logf("User with id %s authenticated", id)

		ws, err := webSocketUpgrader.Upgrade(w, r, *ExtractProtocolHeader(r))

		if err != nil {
			logf("An error occurred while upgrading to the websocket protocol, err: %s", err)
			// gorilla/websocket handles response to client
			return
		}

		logf("User with id %s just connected over websocket", id)

		sockets.NewConnection(user, ws)
	} else {
		logf("User with id %s failed authentication", id)

		unauthorizedHandler(w)
		return
	}

}
