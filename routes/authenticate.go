package routes

import (
	"net/http"

	"github.com/elos/server/data"
	"github.com/elos/server/sockets"
)

func websocketUpgrade(w http.ResponseWriter, r *http.Request, a data.Agent) {
	ws, err := webSocketUpgrader.Upgrade(w, r, *ExtractProtocolHeader(r))

	if err != nil {
		logf("An error occurred while upgrading to the websocket protocol, err: %s", err)
		// gorilla/websocket handles response to client
		return
	}

	logf("Agent with id %s just connected over websocket", a.GetId())

	sockets.NewConnection(a, ws)
}

var AuthenticateGet = NewAuthenticationHandler(DefaultAuthenticator, NewErrorHandler, NewUnauthorizedHandler, websocketUpgrade)
