package data

type DB interface {
	// Management
	Connect(string) error
	GetUpdatesChannel() *chan Model

	// Persistence
	Save(Model) error
	PopulateById(Model) error
	PopulateByField(string, interface{}, Model) error

	NewQuery(Kind) Query
}

type DBConnection interface {
	Close()
}
