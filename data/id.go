package data

// A generic interface for using data IDs
// (This is partially specific to fulfill the bson spec)

type ID interface {
	String() string
	Hex() string
	Valid() bool
}
