package data

import (
	"gopkg.in/mgo.v2/bson"
)

// A generic interface for using data IDs
// This is specifc to fulfill the bson spec:

type ID interface {
	String() string
	Hex() string
	Valid() bool
}

func NewObjectID() ID {
	return bson.NewObjectId()
}

func NewObjectIDFromHex(hex string) ID {
	return bson.ObjectIdHex(hex)
}

func IsObjectIDHex(hex string) bool {
	return bson.IsObjectIdHex(hex)
}
