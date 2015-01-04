package db

type DB interface {
	// Management
	GetConnection() *Connection
	GetUpdatesChannel() *chan Model

	// Persistence
	Save(Model) error
	PopulateById(Model) error
	PopulateByField(string, interface{}, Model) error
}

// Every saved mode is broadcasted over this channel
// Evenutally remove me!
var ModelUpdates chan Model = make(chan Model)
