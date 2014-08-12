package hub

import (
	"fmt"

	"github.com/elos/server/db"
	"github.com/elos/server/models"
	"github.com/elos/server/util"
	"gopkg.in/mgo.v2/bson"
)

type Envelope struct {
	Agent  Agent
	Action string                  `json:"action"`
	Data   map[db.Kind]interface{} `json:"data"`
}

/*
type Serializer func(data []byte) []interface{}

func UserSerializer(data []byte) []*models.User {
	var u models.User
	json.Unmarshal(data, &u)

	var a []*models.User
	a[0] = &u

	return a
}

func UsersSerializer(data []byte) []*models.User {
	var v []interface{}
	json.Unmarshal(data, &v)

	var a []*models.User

	for _, userData := range v {
		bytes, _ := json.Marshal(userData)
		user := UserSerializer(bytes)
		a = append(a, user...)
	}

	return a
}
*/

func Route(e *Envelope, hc *HubConnection) error {
	switch e.Action {
	case "POST":
		return postHandler(e, hc)
	case "GET":
		return getHandler(e, hc)
	case "DELETE":
		return nil
		// haha, can't delete >:)
	default:
		return fmt.Errorf("Action not recognized")
	}
}

func postHandler(e *Envelope, hc *HubConnection) error {
	// Echo
	PrimaryHub.SendJson(hc.Agent, e)
	return nil
}

func getHandler(e *Envelope, hc *HubConnection) error {
	for kind, data := range e.Data {
		id := bson.ObjectIdHex(data.(map[string]interface{})["id"].(string))
		var (
			err   error
			model db.Model
		)

		util.Logf("Id determined: %v", id)

		switch kind {
		case models.UserKind:
			util.Log("Determined to be user")
			model = &models.User{
				Id: id,
			}

			err = db.PopulateById(models.UserKind, model)

			util.Logf("User looks like %#v", model)
		case models.EventKind:
			util.Log("Determined to be event")
			model = &models.Event{
				Id: id,
			}

			err = db.PopulateById(models.EventKind, model)
		default:
			util.Log("Determined to be undetermined")
			PrimaryHub.SendJson(hc.Agent, util.ApiError{400, 400, "Bad", "Stuff"})
		}

		if err != nil {
			util.Log("Error populating the data for getHandler")
			PrimaryHub.SendJson(hc.Agent, util.ApiError{400, 400, "Bad", "Stuff"})
			return fmt.Errorf("There was a error, deal with it")
		}

		util.Logf("MODEL: %#v", model)

		PrimaryHub.SendJson(hc.Agent, model)
	}

	// Echo
	PrimaryHub.SendJson(hc.Agent, e)
	return nil
}
