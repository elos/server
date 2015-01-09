package transfer

import (
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"log"
)

func Route(e *data.Envelope, db data.DB, c conn.Connection) {
	switch e.Action {
	case "POST":
		go postHandler(e, db, c)
	case "GET":
		go getHandler(e, db, c)
	case "DELETE":
		go deleteHandler(e, db, c)
	default:
		log.Print("Action not recognized")
	}
}

func deleteHandler(e *data.Envelope, db data.DB, c conn.Connection) {
	c.WriteJSON(e) // Echo
}
