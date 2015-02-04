package transfer

import (
	"github.com/elos/conn"
	"github.com/elos/data"
	"github.com/elos/server/util"
)

func PostHandler(e *Envelope, s data.Store, c conn.Connection) {
	// Reminder
	var kind data.Kind
	var info data.AttrMap

	for kind, info = range e.Data {
		m, err := s.Unmarshal(kind, info)
		if err != nil {
			c.WriteJSON(util.ApiError{400, 400, "Error", "error"})
		}

		if !m.ID().Valid() {
			m.SetID(s.NewID())
		}

		if err := s.Save(m); err != nil {
			c.WriteJSON(util.ApiError{400, 400, "Error saving the model", "Check yoself"})
			return
		}
	}
}
