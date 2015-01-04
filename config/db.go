package config

import (
	"log"

	"github.com/elos/server/data"
	"github.com/elos/server/data/mongo"
)

/*
	The PrimaryConnection maintiained between the server and the database
		- Theoretically multiple connections could be created.
*/
var DB data.DB

/*
	Establishes a connection to the database package
*/
func SetupDB(addr string) data.DB {
	if DB != nil {
		ShutdownDB()
	}

	var err error
	DB, err = mongo.NewDB(addr)

	if err != nil {
		log.Fatal(err)
	} else {
		Log("Database connection established")
	}

	return DB
}

/*
	Closes the connection to the database package
*/
func ShutdownDB() {
	// DBConnection.Close()
	// needs word
	DB = nil
}
