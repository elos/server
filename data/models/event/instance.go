package event

import (
	"github.com/elos/server/data"
)

func (e *Event) Save() error {
	return db.Save(e)
}

func (e *Event) Concerned() []data.ID {
	a := make([]data.ID, 1)
	a[0] = e.UserId
	return a
}

func (e *Event) SetUser(userId data.ID) error {
	if err := data.CheckID(userId); err != nil {
		return err
	}

	if e.UserId == userId {
		return nil
	}

	e.UserId = userId

	return e.Save()
}
