package event

import (
	"github.com/elos/server/data"
)

func (e *Event) GetID() data.ID {
	return e.ID
}

func (e *Event) SetID(id data.ID) {
	e.ID = id
}

func (e *Event) Kind() data.Kind {
	return Kind
}
