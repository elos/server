package sockets

import (
	"github.com/elos/server/models"
	"github.com/elos/server/util"
)

func postHandler(e *Envelope) {
	// kind is db.Kind
	// data is map[string]interface{}
	for kind, data := range e.Data {
		model, err := models.ModelFor(kind)

		if err != nil {
			PrimaryHub.SendJSON(e.Source.Agent, util.ApiError{401, 400, "Unrecognized type", ""})
			return
		}

		if err := models.PopulateModel(model, &data); err != nil {
			PrimaryHub.SendJSON(e.Source.Agent, util.ApiError{400, 400, "Error populating model with json data", "I need to check maself"})
			return
		}

		if err = model.Save(); err != nil {
			PrimaryHub.SendJSON(e.Source.Agent, util.ApiError{400, 400, "Error saving the model", "Check yoself"})
			return
		}

		// Model will be broadcasted as a sucessful save through ModelUpdate channel
	}
}
