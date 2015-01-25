package transfer

import (
	"github.com/elos/data"
)

/*
	Data structures for the transfer of data
	For implementations of this functionality see elos/server/data/transfer
*/

const POST = "POST"
const GET = "GET"
const DELETE = "DELETE"
const SYNC = "SYNC"
const ECHO = "ECHO"

// Actions a server can send to a client
var ServerActions = map[string]bool{
	POST:   true,
	DELETE: true,
}

// Actions a client can send to a server
var ClientActions = map[string]bool{
	POST:   true,
	GET:    true,
	DELETE: true,
	SYNC:   true,
	ECHO:   true,
}

// Inbound
type Envelope struct {
	Action string                     `json:"action"`
	Data   map[data.Kind]data.AttrMap `json:"data"`
}

func NewEnvelope(action string, data map[data.Kind]data.AttrMap) *Envelope {
	return &Envelope{
		Action: action,
		Data:   data,
	}
}

// Outbound
type Package struct {
	Action string       `json:"action"`
	Data   data.KindMap `json:"data"`
}

func NewPackage(action string, data map[data.Kind]data.Record) *Package {
	return &Package{
		Action: action,
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
