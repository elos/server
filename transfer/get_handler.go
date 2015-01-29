package transfer

import (
	"github.com/elos/data"
	"github.com/elos/server/conn"
	"github.com/elos/server/util"
)

func GetHandler(e *Envelope, s data.Store, c conn.Connection) {
	// kind is s.Kind
	// info is map[string]interface{}
	for kind, info := range e.Data {
		m, _ := s.Unmarshal(kind, info)

		err := s.PopulateByID(m)

		if err != nil {
			if err == data.ErrNotFound {
				c.WriteJSON(util.ApiError{404, 404, "Not Found", "Bad id?"})
				return
			}
			// Otherwise we don't know
			c.WriteJSON(util.ApiError{400, 400, "Oh shit", ""})
			return
		}

		c.WriteJSON(m)
	}
}
