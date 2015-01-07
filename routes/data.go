package routes

/*
	Data
*/

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/test"
)

// The default database is the null test database
var DefaultDB data.DB = test.NewDB()

// Always have a db, private: set with SetDB
var db data.DB = DefaultDB

// Set the database with which the routes look for data
func SetDB(newDB data.DB) {
	if newDB != nil {
		db = newDB
	}
}
