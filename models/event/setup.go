package event

import (
	"github.com/elos/data"
	"github.com/elos/schema"
)

var CurrentEventKind data.Kind
var CurrentEventSchema schema.Schema
var CurrentEventVersion int

func SetupModel(s schema.Schema, k data.Kind, v int) {
	CurrentEventKind = k
	CurrentEventSchema = s
	CurrentEventVersion = v
}