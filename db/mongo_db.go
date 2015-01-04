package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type MongoDB struct {
	Connection   *Connection
	ModelUpdates *chan Model
}

func NewMongoDB(addr string) (DB, error) {
	db := &MongoDB{}
	connection, err := Connect(addr)
	if err != nil {
		return db, err
	}

	db.Connection = connection
	updates := make(chan Model)
	db.ModelUpdates = &updates
	return db, nil
}

func (db *MongoDB) Save(m Model) error {
	return save(db, m)
}

func (db *MongoDB) PopulateById(m Model) error {
	return populateById(db, m)
}

func (db *MongoDB) PopulateByField(field string, value interface{}, m Model) error {
	return populateByField(db, m, field, value)
}

func (db *MongoDB) GetConnection() *Connection {
	return db.Connection
}

func (db *MongoDB) GetUpdatesChannel() *chan Model {
	return db.ModelUpdates
}

/*
	Forks the session of the primary connection
		- If the PrimaryConnection does not exist, this returns a nil session
	Note: newSession is not exported, it should not be used by another package!
	    - this is an attempt to enforce db/server agnostic
*/
func newSession(db DB) (*mgo.Session, error) {
	connection := db.GetConnection()
	if connection != nil {
		return connection.Session.Copy(), nil
	} else {
		return nil, fmt.Errorf("Primary connection does not exist")
	}
}
