package routes

import (
	"net/http"

	"github.com/elos/server/autonomous"
	"github.com/elos/server/autonomous/agents"
	"github.com/elos/server/autonomous/managers"
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/data/test"
)

var DefaultClientDataHub autonomous.Manager = managers.NewNullHub()

type AuthenticateGetHandler struct {
	data.DB
}

func (h *AuthenticateGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	NewAuthenticationHandler(h.DB, DefaultAuthenticator,
		NewErrorHandler,
		NewUnauthorizedHandler,
		AuthenticatedHandlerFunc(func(w http.ResponseWriter, r *http.Request, a data.Identifiable) {
			WebSocketUpgradeHandler(w, r, a, conn.DefaultWebSocketUpgrader, DefaultClientDataHub, test.NewDB())
		})).ServeHTTP(w, r)
}

func WebSocketUpgradeHandler(w http.ResponseWriter, r *http.Request, a data.Identifiable, upgrader conn.WebSocketUpgrader, hub autonomous.Manager, db data.DB) {
	connection, err := upgrader.Upgrade(w, r, a)

	if err != nil {
		logf("An error occurred while upgrading to the websocket protocol, err: %s", err)
		// gorilla/websocket handles response to client
		return
	}

	logf("Agent with id %s just connected over websocket", a.GetID())

	agent := agents.NewClientDataAgent(connection, db)
	go hub.StartAgent(agent)
}
