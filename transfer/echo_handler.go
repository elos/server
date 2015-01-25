package transfer

import (
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
)

func EchoHandler(e *Envelope, db data.DB, c conn.Connection) {
	c.WriteJSON(e)
}
