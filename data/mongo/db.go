package mongo

import (
	"fmt"
	"github.com/elos/server/data"
	"gopkg.in/mgo.v2"
	"sync"
)

type MongoDB struct {
	Connections []*MongoConnection
	Subscribers map[data.ID][]*chan data.Model
	m           sync.Mutex
}

func NewDB(addr string) (data.DB, error) {
	db := &MongoDB{}
	db.Connections = make([]*MongoConnection, 0)
	db.Connect(addr)
	db.Subscribers = make(map[data.ID][]*chan data.Model)
	return db, nil
}

func (db *MongoDB) Connect(addr string) error {
	connection, err := Connect(addr)

	if err != nil {
		return err
	}

	db.Connections = append(db.Connections, connection)
	return nil
}

func (db *MongoDB) Save(m data.Model) error {
	s, err := newSession(db)
	if err != nil {
		log(err)
		return err
	}

	defer s.Close()

	if err = save(s, m); err != nil {
		logf("Error saving record of kind %s, err: %s", m.Kind(), err)
		return err
	} else {
		db.NotifyConcerned(m)
		return nil
	}
}

func (db *MongoDB) PopulateById(m data.Model) error {
	s, err := newSession(db)
	if err != nil {
		log(err)
		return err
	}

	defer s.Close()

	if err = populateById(s, m); err != nil {
		logf("There was an error populating the %s model, error: %v", m.Kind(), err)
		return err
	} else {
		return nil
	}
}

func (db *MongoDB) PopulateByField(field string, value interface{}, m data.Model) error {
	s, err := newSession(db)
	if err != nil {
		log(err)
		return err
	}

	defer s.Close()

	if err = populateByField(s, m, field, value); err != nil {
		logf("There was an error populating the %s model, error: %v", m.Kind(), err)
		return err
	} else {
		return nil
	}
}

func (db *MongoDB) GetConnection() *MongoConnection {
	return db.Connections[0]
}

func (db *MongoDB) RegisterForUpdates(a data.Agent) *chan data.Model {
	db.m.Lock()
	defer db.m.Unlock()

	id := a.GetID()
	c := make(chan data.Model)

	db.Subscribers[id] = append(db.Subscribers[id], &c)

	return &c
}

func (db *MongoDB) NotifyConcerned(m data.Model) {
	concerned := m.Concerned()
	for _, concernedId := range concerned {
		channels := db.Subscribers[concernedId]
		for _, channel := range channels {
			go nonblockingchannelsend(*channel, m)
		}
	}
}

func nonblockingchannelsend(c chan data.Model, m data.Model) {
	c <- m
}

/*
	Forks the session of the primary connection
		- If the PrimaryConnection does not exist, this returns a nil session
	Note: newSession is not exported, it should not be used by another package!
	    - this is an attempt to enforce db/server agnostic
*/
func newSession(db *MongoDB) (*mgo.Session, error) {
	connection := db.GetConnection()
	if connection != nil {
		return connection.Session.Copy(), nil
	} else {
		return nil, fmt.Errorf("Primary connection does not exist")
	}
}
