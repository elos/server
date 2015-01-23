package data

type DB interface {
	// Management
	Connect(string) error

	// Persistence
	Save(Record) error
	Delete(Record) error
	PopulateById(Record) error
	PopulateByField(string, interface{}, Record) error

	NewQuery(Kind) Query

	RegisterForUpdates(Agent) *chan *Package

	Type() string
}

type DBConnection interface {
	Close()
}
