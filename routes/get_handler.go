package routes

import (
	"fmt"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/services/agents"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2"
)

func getHandler(e *data.Envelope, db data.DB, cda agents.ClientDataAgent) {
	// kind is db.Kind
	// info is map[string]interface{}
	for kind, info := range e.Data {
		model, err := models.ModelFor(kind)

		if err != nil {
			cda.WriteJSON(util.ApiError{400, 400, "Oh shit", ""})
			return
		}

		err = models.PopulateModel(model, &info)

		if err := data.CheckID(model.GetID()); err != nil {
			cda.WriteJSON(util.ApiError{400, 400, "Invalid ID", fmt.Sprintf("%s", err)})
			return
		}

		// FIXME: INJECT!
		err = db.PopulateById(model)

		if err != nil {
			if err == mgo.ErrNotFound {
				// Handle the error here
				cda.WriteJSON(util.ApiError{404, 404, "Not Found", "Bad id?"})
				return
			}
			// Otherwise we don't know
			cda.WriteJSON(util.ApiError{400, 400, "Oh shit", ""})
			return
		}

		cda.WriteJSON(model)
	}
}
