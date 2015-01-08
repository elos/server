package routes

import (
	"net/http"

	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/hub"
)

func WebSocketUpgradeHandler(w http.ResponseWriter, r *http.Request, a data.Agent, upgrader conn.WebSocketUpgrader) {
	ws, err := upgrader.Upgrade(w, r, a)

	if err != nil {
		logf("An error occurred while upgrading to the websocket protocol, err: %s", err)
		// gorilla/websocket handles response to client
		return
	}

	logf("Agent with id %s just connected over websocket", a.GetId())

	hub.NewConnection(ws)
}

var AuthenticateGet = NewAuthenticationHandler(DefaultAuthenticator,
	NewErrorHandler,
	NewUnauthorizedHandler,
	AuthenticatedHandlerFunc(func(w http.ResponseWriter, r *http.Request, a data.Agent) {
		WebSocketUpgradeHandler(w, r, a, conn.DefaultWebSocketUpgrader)
	}),
)
