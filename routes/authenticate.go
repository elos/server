package routes

import (
	"net/http"

	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/server/agents"
	"github.com/elos/server/conn"
)

var DefaultClientDataHub autonomous.Manager = autonomous.NewNullHub()

type AuthenticateGetHandler struct {
	data.Store
}

func (h *AuthenticateGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	NewAuthenticationHandler(h.Store, DefaultAuthenticator,
		NewErrorHandler,
		NewUnauthorizedHandler,
		AuthenticatedHandlerFunc(func(w http.ResponseWriter, r *http.Request, a data.Identifiable) {
			WebSocketUpgradeHandler(w, r, a, conn.DefaultWebSocketUpgrader, DefaultClientDataHub, h.Store)
		})).ServeHTTP(w, r)
}

func WebSocketUpgradeHandler(w http.ResponseWriter, r *http.Request, a data.Identifiable, upgrader conn.WebSocketUpgrader, hub autonomous.Manager, s data.Store) {
	connection, err := upgrader.Upgrade(w, r, a)

	if err != nil {
		logf("An error occurred while upgrading to the websocket protocol, err: %s", err)
		// gorilla/websocket handles response to client
		return
	}

	logf("Agent with id %s just connected over websocket", a.ID())

	agent := agents.NewClientDataAgent(connection, s)
	go hub.StartAgent(agent)
}
