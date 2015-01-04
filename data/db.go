package data

type DB interface {
	// Management
	GetUpdatesChannel() *chan Model

	// Persistence
	Save(Model) error
	PopulateById(Model) error
	PopulateByField(string, interface{}, Model) error
}
