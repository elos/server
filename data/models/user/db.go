package user

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/test"
	"log"
)

var DefaultDatabase = test.NewDB()

var db data.DB = DefaultDatabase

func SetDB(database data.DB) {
	log.Println("USER DEFAULT DB IS DEPRECATED")
	if database != nil {
		db = database
	}
}
