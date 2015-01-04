package sockets

import (
	"github.com/elos/server/data"
)

// Inbound
type Envelope struct {
	Source *Connection                          `json:"agent,omitempty"`
	Action string                               `json:"action"`
	Data   map[data.Kind]map[string]interface{} `json:"data"`
}

// Outbound
type Package struct {
	Action string                   `json:"action"`
	Data   map[data.Kind]data.Model `json:"data"`
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
		Logf("Action not recognized")
	}
}

func deleteHandler(e *Envelope) {
	PrimaryHub.SendJSON(e.Source.Agent, e) // Echo
}
