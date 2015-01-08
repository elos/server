package routes

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/services/agents"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getHandler(e *data.Envelope, db data.DB, cda agents.ClientDataAgent) {
	// kind is db.Kind
	// data is map[string]interface{}
	for kind, data := range e.Data {
		model, err := models.ModelFor(kind)

		if err != nil {
			cda.WriteJSON(util.ApiError{400, 400, "Oh shit", ""})
			return
		}

		err = models.PopulateModel(model, &data)

		if id := model.GetId(); id == bson.ObjectId("") {
			cda.WriteJSON(util.ApiError{400, 400, "Invalid ID", ""})
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
