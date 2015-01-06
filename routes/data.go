package routes

/*
	Data
*/

import (
	"github.com/elos/server/data"
	"github.com/elos/server/data/test"
)

// The default database is the null test database
var DefaultDatabase data.DB = test.NewDB()

// Always have a db, private: set with SetDatabase
var db data.DB = DefaultDatabase

// Set the database with which the routes look for data
func SetDatabase(newDB data.DB) {
	if newDB != nil {
		db = newDB
	}
}
