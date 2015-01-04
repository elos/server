package db

import (
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
}

/*
	Creates a db.Connection to the database. Will fail hard, see log.Fatal
*/
func Connect(addr string) (*Connection, error) {
	session, err := mgo.Dial(addr)

	if err != nil {
		return nil, err
	}

	connection := &Connection{Session: session}

	return connection, nil
}
