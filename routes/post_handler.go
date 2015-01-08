package routes

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/models"
	"github.com/elos/server/services/agents"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
)

func postHandler(e *data.Envelope, db data.DB, cda agents.ClientDataAgent) {
	// Reminder
	var kind data.Kind
	var data map[string]interface{}

	for kind, data = range e.Data {

		model, err := models.ModelFor(kind)

		if err != nil {
			cda.WriteJSON(util.ApiError{401, 400, "Unrecognized type", ""})
			return
		}

		if err := models.PopulateModel(model, &data); err != nil {
			cda.WriteJSON(util.ApiError{400, 400, "Error populating model with json data", "I need to check maself"})
			return
		}

		if !model.GetId().Valid() {
			model.SetId(bson.NewObjectId())
		}

		if err = model.Save(); err != nil {
			cda.WriteJSON(util.ApiError{400, 400, "Error saving the model", "Check yoself"})
			return
		}

		// Model will be broadcasted as a sucessful save through ModelUpdate channel
	}
}
