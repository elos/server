package event

import (
	"errors"
	"github.com/elos/server/models"
)

var NoNameError = errors.New("Error: event must have a name")
var NoStartTimeError = errors.New("Error: event must have a start time")
var NoEndTimeError = errors.New("Error: event must have an end time")
var NoUserError = errors.New("Error: event must have a user")

func Validate(e models.Event) (bool, error) {

	if e.Name() == "" {
		return false, NoNameError
	}

	if e.StartTime().IsZero() {
		return false, NoStartTimeError
	}

	if e.EndTime().IsZero() {
		return false, NoEndTimeError
	}

	switch e.(type) {
	case *MongoEvent:
		if !e.(*MongoEvent).UserID.Valid() {
			return false, NoUserError
		}
	}

	return true, nil
}
