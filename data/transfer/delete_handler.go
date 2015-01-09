package transfer

import (
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
)

func DeleteHandler(e *data.Envelope, db data.DB, c conn.Connection) {
	c.WriteJSON(e) // Echo
}
