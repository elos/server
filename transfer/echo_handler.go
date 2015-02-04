package transfer

import (
	"github.com/elos/conn"
	"github.com/elos/data"
)

func EchoHandler(e *Envelope, db data.DB, c conn.Connection) {
	c.WriteJSON(e)
}
