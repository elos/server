package test

import (
	"fmt"
	"github.com/elos/server/data"
)

type TestDB struct {
	ModelUpdates     chan *data.Package
	Saved            []data.Model
	Deleted          []data.Model
	PopulatedById    []data.Model
	PopulatedByField []data.Model

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
	db.Saved = make([]data.Model, 0)
	db.Deleted = make([]data.Model, 0)
	db.PopulatedById = make([]data.Model, 0)
	db.PopulatedByField = make([]data.Model, 0)
	db.Error = false
}

func (db *TestDB) Connect(addr string) error {
	if db.Error {
		return TestDBError
	}

	return nil
}

func (db *TestDB) RegisterForUpdates(a data.Agent) *chan *data.Package {
	return &db.ModelUpdates
}

func (db *TestDB) Save(m data.Model) error {
	if db.Error {
		return TestDBError
	}

	db.Saved = append(db.Saved, m)
	return nil
}

func (db *TestDB) Delete(m data.Model) error {
	if db.Error {
		return TestDBError
	}

	db.Deleted = append(db.Deleted, m)
	return nil
}

func (db *TestDB) PopulateById(m data.Model) error {
	if db.Error {
		return TestDBError
	}

	db.PopulatedById = append(db.PopulatedById, m)
	return nil
}

func (db *TestDB) PopulateByField(field string, value interface{}, m data.Model) error {
	if db.Error {
		return TestDBError
	}

	db.PopulatedByField = append(db.PopulatedByField, m)
	return nil
}

func (db *TestDB) NewQuery(k data.Kind) data.Query {
	return nil
}
