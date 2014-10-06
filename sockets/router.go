package sockets

import (
	"github.com/elos/server/db"
	"github.com/elos/server/util"
)

type Envelope struct {
	Agent  Agent                              `json:"agent,omitempty"`
	Action string                             `json:"action"`
	Data   map[db.Kind]map[string]interface{} `json:"data"`
}

type Package struct {
	Action string               `json:"action"`
	Data   map[db.Kind]db.Model `json:"data"`
}

func Route(e *Envelope, c *Connection) {
	switch e.Action {
	case "POST":
		go postHandler(e, c)
	case "GET":
		go getHandler(e, c)
	case "DELETE":
		go deleteHandler(e, c)
	default:
		util.Logf("[Hub] Action not recognized")
	}
}

func deleteHandler(e *Envelope, c *Connection) {
	PrimaryHub.SendJSON(c.Agent, e) // Echo
}
