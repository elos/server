package routes

import (
	"github.com/elos/server/data"
	"github.com/elos/server/services/agents"
)

func Route(e *data.Envelope, db data.DB, cda agents.ClientDataAgent) {
	switch e.Action {
	case "POST":
		go postHandler(e, db, cda)
	case "GET":
		go getHandler(e, db, cda)
	case "DELETE":
		go deleteHandler(e, db, cda)
	default:
		logf("Action not recognized")
	}
}

func deleteHandler(e *data.Envelope, db data.DB, cda agents.ClientDataAgent) {
	cda.WriteJSON(e) // Echo
}
