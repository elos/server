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

	if e.GetName() == "" {
		return false, NoNameError
	}

	if e.GetStartTime().IsZero() {
		return false, NoStartTimeError
	}

	if e.GetEndTime().IsZero() {
		return false, NoEndTimeError
	}

	switch e.(type) {
	case *MongoEvent:
		if e.(*MongoEvent).UserID == "" {
			return false, NoUserError
		}
	}

	return true, nil
}
