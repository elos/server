package data

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
	Action string                          `json:"action"`
	Data   map[Kind]map[string]interface{} `json:"data"`
}

func NewEnvelope(action string, data map[Kind]map[string]interface{}) *Envelope {
	return &Envelope{
		Action: action,
		Data:   data,
	}
}

// Outbound
type Package struct {
	Action string
	Data   map[Kind]Record
}

func NewPackage(action string, data map[Kind]Record) *Package {
	return &Package{
		Action: action,
		Data:   data,
	}
}
