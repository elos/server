package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

/*
	Moderately abstract data type for managing relationship
	between database and the rest of the server
*/
type Connection struct {
	Session *mgo.Session
}

// Closes the connection
func (c *Connection) Close() {
	c.Session.Close()

	if c == PrimaryConnection {
		PrimaryConnection = nil
	}
}

// The primary connection used for forking sessions from the database
var PrimaryConnection *Connection

/*
	Creates a db.Connection to the database. Will fail hard
*/
func Connect(addr string) (*Connection, error) {
	session, err := mgo.Dial(addr)

	if err != nil {
		log.Fatal(err)
	}

	connection := &Connection{Session: session}

	if PrimaryConnection == nil {
		PrimaryConnection = connection
	}

	return connection, err
}

/*
	Forks the session of the primary connection
		- If the PrimaryConnection does not exist, this returns a nil session
*/
func NewSession() *mgo.Session {
	if PrimaryConnection != nil {
		return PrimaryConnection.Session.Copy()
	} else {
		return nil
	}
}
