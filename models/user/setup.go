package user

import (
	"github.com/elos/data"
)

var CurrentUserKind data.Kind
var CurrentUserSchema data.Schema
var CurrentUserVersion int

func SetupModel(s data.Schema, k data.Kind, version int) {
	CurrentUserKind = k
	CurrentUserSchema = s
	CurrentUserVersion = version
}
