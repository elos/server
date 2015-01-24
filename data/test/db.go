package test

import (
	"fmt"
	"github.com/elos/server/data"
	"github.com/elos/server/data/mongo"
)

type TestDB struct {
	ModelUpdates     chan *data.Package
	Saved            []data.Record
	Deleted          []data.Record
	PopulatedById    []data.Record
	PopulatedByField []data.Record

	Error bool
}

var TestDBError error = fmt.Errorf("TestDB Error")

const TestDBType = "test"

func NewDB() *TestDB {
	db := &TestDB{}
	db.Reset()
	return db
}

func (db *TestDB) Type() string {
	return TestDBType
}

func (db *TestDB) Reset() {
	db.ModelUpdates = make(chan *data.Package)
	db.Saved = make([]data.Record, 0)
	db.Deleted = make([]data.Record, 0)
	db.PopulatedById = make([]data.Record, 0)
	db.PopulatedByField = make([]data.Record, 0)
	db.Error = false
}

func (db *TestDB) Connect(addr string) error {
	if db.Error {
		return TestDBError
	}

	return nil
}

func (db *TestDB) RegisterForUpdates(a data.Identifiable) *chan *data.Package {
	return &db.ModelUpdates
}

func (db *TestDB) NewObjectID() data.ID {
	return mongo.NewObjectID()
}

func (db *TestDB) CheckID(id data.ID) error {
	if db.Error {
		return TestDBError
	}

	return nil
}

func (db *TestDB) Save(m data.Record) error {
	if db.Error {
		return TestDBError
	}

	db.Saved = append(db.Saved, m)
	return nil
}

func (db *TestDB) Delete(m data.Record) error {
	if db.Error {
		return TestDBError
	}

	db.Deleted = append(db.Deleted, m)
	return nil
}

func (db *TestDB) PopulateById(m data.Record) error {
	if db.Error {
		return TestDBError
	}

	db.PopulatedById = append(db.PopulatedById, m)
	return nil
}

func (db *TestDB) PopulateByField(field string, value interface{}, m data.Record) error {
	if db.Error {
		return TestDBError
	}

	db.PopulatedByField = append(db.PopulatedByField, m)
	return nil
}

func (db *TestDB) NewQuery(k data.Kind) data.Query {
	return nil
}
