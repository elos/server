package routes

import (
	"net/http"

	"github.com/elos/server/autonomous"
	"github.com/elos/server/autonomous/agents"
	"github.com/elos/server/autonomous/managers"
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models/user"
)

var DefaultClientDataHub = managers.NewNullHub()

func WebSocketUpgradeHandler(w http.ResponseWriter, r *http.Request, a data.Agent, upgrader conn.WebSocketUpgrader, hub autonomous.Manager) {
	connection, err := upgrader.Upgrade(w, r, a)

	if err != nil {
		logf("An error occurred while upgrading to the websocket protocol, err: %s", err)
		// gorilla/websocket handles response to client
		return
	}

	logf("Agent with id %s just connected over websocket", a.GetID())

	agent := agents.NewClientDataAgent(connection, user.DefaultDatabase)
	go hub.StartAgent(agent)
}

var AuthenticateGet = NewAuthenticationHandler(DefaultAuthenticator,
	NewErrorHandler,
	NewUnauthorizedHandler,
	AuthenticatedHandlerFunc(func(w http.ResponseWriter, r *http.Request, a data.Agent) {
		WebSocketUpgradeHandler(w, r, a, conn.DefaultWebSocketUpgrader, DefaultClientDataHub)
	}),
)
