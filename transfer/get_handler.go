package transfer

import (
	"github.com/elos/data"
	"github.com/elos/server/conn"
	"github.com/elos/server/models"
	"github.com/elos/server/util"
)

func GetHandler(e *Envelope, db data.DB, c conn.Connection) {
	// kind is db.Kind
	// info is map[string]interface{}
	for kind, info := range e.Data {
		model, err := models.ModelFor(kind)

		if err != nil {
			c.WriteJSON(util.ApiError{400, 400, "Oh shit", ""})
			return
		}

		err = models.PopulateModel(model, &info)

		err = db.PopulateByID(model)

		if err != nil {
			if err == data.ErrNotFound {
				c.WriteJSON(util.ApiError{404, 404, "Not Found", "Bad id?"})
				return
			}
			// Otherwise we don't know
			c.WriteJSON(util.ApiError{400, 400, "Oh shit", ""})
			return
		}

		c.WriteJSON(model)
	}
}
