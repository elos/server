package transfer

import (
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/util"
)

func postHandler(e *data.Envelope, db data.DB, c conn.Connection) {
	// Reminder
	var kind data.Kind
	var info map[string]interface{}

	for kind, info = range e.Data {

		model, err := models.ModelFor(kind)

		if err != nil {
			c.WriteJSON(util.ApiError{401, 400, "Unrecognized type", ""})
			return
		}

		if err := models.PopulateModel(model, &info); err != nil {
			c.WriteJSON(util.ApiError{400, 400, "Error populating model with json data", "I need to check maself"})
			return
		}

		if !model.GetID().Valid() {
			model.SetID(data.NewObjectID())
		}

		if err = model.Save(); err != nil {
			c.WriteJSON(util.ApiError{400, 400, "Error saving the model", "Check yoself"})
			return
		}

		// Model will be broadcasted as a sucessful save through ModelUpdate channel
	}
}