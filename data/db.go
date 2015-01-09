package data

type DB interface {
	// Management
	Connect(string) error

	// Persistence
	Save(Model) error
	PopulateById(Model) error
	PopulateByField(string, interface{}, Model) error

	NewQuery(Kind) Query

	RegisterForUpdates(Agent) *chan *Package
}

type DBConnection interface {
	Close()
}
