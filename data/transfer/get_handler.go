package transfer

import (
	"fmt"
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2"
)

func getHandler(e *data.Envelope, db data.DB, c conn.Connection) {
	// kind is db.Kind
	// info is map[string]interface{}
	for kind, info := range e.Data {
		model, err := models.ModelFor(kind)

		if err != nil {
			c.WriteJSON(util.ApiError{400, 400, "Oh shit", ""})
			return
		}

		err = models.PopulateModel(model, &info)

		if err := data.CheckID(model.GetID()); err != nil {
			c.WriteJSON(util.ApiError{400, 400, "Invalid ID", fmt.Sprintf("%s", err)})
			return
		}

		// FIXME: INJECT!
		err = db.PopulateById(model)

		if err != nil {
			if err == mgo.ErrNotFound {
				// Handle the error here
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
