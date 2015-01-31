package transfer

import (
	"github.com/elos/data"
)

type Router interface {
	Route(Actionable)
}

type Handler interface {
	Handle(Actionable)
}

type AMux struct {
	routes map[Action]Handler
}

func (m *AMux) Route(a Actionable) {
	m.routes[a.Action()].Handle(a)
}

type Action string

const (
	POST   Action = "POST"
	GET    Action = "GET"
	DELETE Action = "DELETE"
	SYNC   Action = "SYNC"
	ECHO   Action = "ECHO"
)

type Actionable interface {
	Action() Action
}

// Actions a server can send to a client
var ServerActions = map[Action]bool{
	POST:   true,
	DELETE: true,
}

// Actions a client can send to a server
var ClientActions = map[Action]bool{
	POST:   true,
	GET:    true,
	DELETE: true,
	SYNC:   true,
	ECHO:   true,
}

// Inbound
type Envelope struct {
	Action Action                     `json:"action"`
	Data   map[data.Kind]data.AttrMap `json:"data"`
}

func NewEnvelope(a Action, data map[data.Kind]data.AttrMap) *Envelope {
	return &Envelope{
		Action: a,
		Data:   data,
	}
}

// Outbound
type Package struct {
	Action Action       `json:"action"`
	Data   data.KindMap `json:"data"`
}

func NewPackage(a Action, data map[data.Kind]data.Record) *Package {
	return &Package{
		Action: a,
		Data:   data,
	}
}

/*
	Returns a map like:
	{ user: { Name: "Nick Landolfi"} }
	of form:
	{ <db.Kind>: <db.Model>}
*/
func Map(m data.Record) data.KindMap {
	return data.KindMap{
		m.Kind(): m,
	}
}
