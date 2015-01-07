package event

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/test"
)

var DefaultDatabase = test.NewDB()

var db data.DB = DefaultDatabase

func SetDB(database data.DB) {
	if database != nil {
		db = database
	}
}
