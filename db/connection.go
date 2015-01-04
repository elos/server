package db

import (
	"fmt"

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
	Creates a db.Connection to the database. Will fail hard, see log.Fatal
*/
func Connect(addr string) (*Connection, error) {
	session, err := mgo.Dial(addr)

	if err != nil {
		return nil, err
	}

	connection := &Connection{Session: session}

	if PrimaryConnection == nil {
		PrimaryConnection = connection
	}

	return connection, nil
}

/*
	Forks the session of the primary connection
		- If the PrimaryConnection does not exist, this returns a nil session
	Note: newSession is not exported, it should not be used by another package!
	    - this is an attempt to enforce db/server agnostic
*/
func newSession() (*mgo.Session, error) {
	if PrimaryConnection != nil {
		return PrimaryConnection.Session.Copy(), nil
	} else {
		return nil, fmt.Errorf("Primary connection does not exist")
	}
}
