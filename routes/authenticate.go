package routes

import (
	"net/http"

	"github.com/elos/server/sockets"
	"github.com/elos/server/util/auth"
)

func authenticateGet(w http.ResponseWriter, r *http.Request, a auth.RequestAuthenticator) {
	agent, authenticated, err := DefaultAuthenticator(r)

	if err != nil {
		logf("An error occurred during authentication, err: %s", err)
		serverErrorHandler(w, err)
		return
	}

	if authenticated {
		logf("Agent with id %s authenticated", agent.GetId())

		ws, err := webSocketUpgrader.Upgrade(w, r, *auth.ExtractProtocolHeader(r))

		if err != nil {
			logf("An error occurred while upgrading to the websocket protocol, err: %s", err)
			// gorilla/websocket handles response to client
			return
		}

		logf("Agent with id %s just connected over websocket", agent.GetId())

		sockets.NewConnection(agent, ws)
	} else {
		logf("Agent with id %s failed authentication", agent.GetId())

		unauthorizedHandler(w)
		return
	}

}

var AuthenticateGet = FunctionHandler(func(w http.ResponseWriter, r *http.Request) { authenticateGet(w, r, DefaultAuthenticator) })
