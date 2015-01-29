package event

import (
	"github.com/elos/data"
)

var CurrentEventKind data.Kind
var CurrentEventSchema data.Schema
var CurrentEventVersion int

func SetupModel(s data.Schema, k data.Kind, v int) {
	CurrentEventKind = k
	CurrentEventSchema = s
	CurrentEventVersion = v
}
