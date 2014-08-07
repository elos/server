package db

import "gopkg.in/mgo.v2"

var PrimarySession *mgo.Session

func Connect(addr string) (*mgo.Session, error) {
	Close()
	return mgo.Dial(addr)
}

func Close() {
	if PrimarySession != nil {
		PrimarySession.Close()
	}
}

func NewSession() *mgo.Session {
	return PrimarySession.Copy()
}
