package data

import (
	"gopkg.in/mgo.v2/bson"
)

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

/*

func (id ObjectID) String() string {
	return bson.ObjectId(id).String()
}

func (id ObjectID) Hex() string {
	return bson.ObjectId(id).Hex()
}

func (id ObjectID) MarshalJSON() ([]byte, error) {
	return bson.ObjectId(id).MarshalJSON()
}

func (id ObjectID) UnmarshalJSON(bytes []byte) error {
	bie := bson.ObjectId(id)
	return (&bie).UnmarshalJSON(bytes)
}

func (id ObjectID) Valid() bool {
	return bson.ObjectId(id).Valid()
}

*/
