package event

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/schema"
)

var CurrentEventKind data.Kind
var CurrentEventSchema schema.Schema
var CurrentEventVersion int

func SetupModel(s schema.Schema, k data.Kind, v int) {
	CurrentEventKind = k
	CurrentEventSchema = s
	CurrentEventVersion = v
}
