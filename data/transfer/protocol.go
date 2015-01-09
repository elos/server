package transfer

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
