package user

import (
	"github.com/elos/server/data"
	"github.com/elos/server/schema"
)

var CurrentUserKind data.Kind
var CurrentUserSchema schema.Schema
var CurrentUserVersion int

func SetupModel(s schema.Schema, k data.Kind, version int) {
	CurrentUserKind = k
	CurrentUserSchema = s
	CurrentUserVersion = version
}
