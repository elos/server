package config

import (
	"log"

	"github.com/elos/server/db"
)

/*
	The PrimaryConnection maintiained between the server and the database
		- Theoretically multiple connections could be created.
*/
var DBConnection *db.Connection

/*
	Establishes a connection to the database package
*/
func SetupDB(addr string) *db.Connection {
	if DBConnection != nil {
		ShutdownDB()
	}

	DBConnection, err := db.Connect(addr)

	if err != nil {
		log.Fatal(err)
	} else {
		Log("Database connection established")
	}

	return DBConnection
}

/*
	Closes the connection to the database package
*/
func ShutdownDB() {
	// DBConnection.Close()
	DBConnection = nil
}
