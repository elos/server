package sockets

import (
	"github.com/elos/server/db"
	"github.com/elos/server/util"
)

type Envelope struct {
	Source *Connection                        `json:"agent,omitempty"`
	Action string                             `json:"action"`
	Data   map[db.Kind]map[string]interface{} `json:"data"`
}

type OutboundEnvelope struct {
	Action string               `json:"action"`
	Data   map[db.Kind]db.Model `json:"data"`
}

func Route(e *Envelope) {
	switch e.Action {
	case "POST":
		go postHandler(e)
	case "GET":
		go getHandler(e)
	case "DELETE":
		go deleteHandler(e)
	default:
		util.Logf("[Hub] Action not recognized")
	}
}

func deleteHandler(e *Envelope) {
	PrimaryHub.SendJSON(e.Source.Agent, e) // Echo
}
