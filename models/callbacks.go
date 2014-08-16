package models

import "github.com/elos/server/db"

var ModelUpdates chan db.Model = make(chan db.Model)

/*

func (u *User) DidSave() {
	ModelUpdates <- u
}

func (e *Event) DidSave() {
	ModelUpdates <- e
}
*/
