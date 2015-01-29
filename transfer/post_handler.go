package transfer

import (
	"github.com/elos/data"
	"github.com/elos/server/conn"
	"github.com/elos/server/util"
)

func PostHandler(e *Envelope, s data.Store, c conn.Connection) {
	// Reminder
	var kind data.Kind
	var info data.AttrMap

	for kind, info = range e.Data {
		m, _ := s.Unmarshal(kind, info)

		if !m.ID().Valid() {
			m.SetID(s.NewObjectID())
		}

		if err := s.Save(m); err != nil {
			c.WriteJSON(util.ApiError{400, 400, "Error saving the model", "Check yoself"})
			return
		}
	}
}
