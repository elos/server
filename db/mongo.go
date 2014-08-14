package db

import (
	"gopkg.in/mgo.v2"
)

var PrimarySession *mgo.Session

func Connect(addr string) (*mgo.Session, error) {
	// Close any existing connection
	Close()

	var err error
	PrimarySession, err = mgo.Dial(addr)

	return PrimarySession, err
}

func Close() {
	if PrimarySession != nil {
		PrimarySession.Close()
	}
}

func NewSession() *mgo.Session {
	return PrimarySession.Copy()
}
