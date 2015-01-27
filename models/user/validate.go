package user

import (
	"errors"
	"github.com/elos/server/models"
)

var NoNameError = errors.New("Error: user must have a name")
var NoKeyError = errors.New("Error: user must have a key")

func Validate(u models.User) (bool, error) {
	if u.Name() == "" {
		return false, NoNameError
	}

	if u.Key() == "" {
		return false, NoKeyError
	}

	return true, nil
}
