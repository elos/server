package sockets

import (
	"log"

	"github.com/elos/server/db"
)

type Envelope struct {
	Agent  Agent
	Action string                             `json:"action"`
	Data   map[db.Kind]map[string]interface{} `json:"data"`
}

func Route(e *Envelope, hc *Connection) {
	switch e.Action {
	case "POST":
		go postHandler(e, hc)
	case "GET":
		go getHandler(e, hc)
	case "DELETE":
		go deleteHandler(e, hc)
	default:
		log.Printf("Action not recognized")
	}
}

func deleteHandler(e *Envelope, hc *Connection) {
	PrimaryHub.SendJSON(hc.Agent, e) // Echo
}
