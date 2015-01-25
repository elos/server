package transfer

import (
	"github.com/elos/data"
	"github.com/elos/server/conn"
)

func EchoHandler(e *Envelope, db data.DB, c conn.Connection) {
	c.WriteJSON(e)
}
