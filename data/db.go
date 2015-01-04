package data

type DB interface {
	// Management
	Connect(string) error
	GetUpdatesChannel() *chan Model

	// Persistence
	Save(Model) error
	PopulateById(Model) error
	PopulateByField(string, interface{}, Model) error
}

type DBConnection interface {
	Close()
}
