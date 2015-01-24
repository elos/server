package data

/*
	Model type or class
	Should correspond with the model name, generally lowercase
*/
type Kind string

// A generic interface for using data IDs
// (This is partially specific to fulfill the bson spec)
// needs work
type ID interface {
	String() string
	Hex() string
	Valid() bool
}

type Identifiable interface {
	SetID(ID)
	GetID() ID
}

type Persistable interface {
	Identifiable
	Kind() Kind
}

type Record interface {
	Persistable

	Save(DB) error
	Concerned() []ID // for model updates
}

type DB interface {
	// Management
	Connect(string) error

	// Persistence
	NewObjectID() ID
	CheckID(ID) error
	Save(Record) error
	Delete(Record) error
	PopulateById(Record) error
	PopulateByField(string, interface{}, Record) error

	NewQuery(Kind) Query

	RegisterForUpdates(Identifiable) *chan *Package

	Type() string
}

type DBConnection interface {
	Close()
}

type AttrMap map[string]interface{}
type KindMap map[Kind]Record

type Query interface {
	Execute() (RecordIterator, error)

	Select(AttrMap) Query
	Limit(int) Query
	Skip(int) Query
	Batch(int) Query
}

type RecordIterator interface {
	Next(Record) bool
	Close() error
}
