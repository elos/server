package user

import (
	"errors"
	"github.com/elos/server/data/models"
)

var NoNameError = errors.New("Error: user must have a name")
var NoKeyError = errors.New("Error: user must have a key")

func Validate(u models.User) (bool, error) {
	if u.GetName() == "" {
		return false, NoNameError
	}

	if u.GetKey() == "" {
		return false, NoKeyError
	}

	return true, nil
}
