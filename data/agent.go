package data

import (
	"gopkg.in/mgo.v2/bson"
)

/*
	Describes the ability to be interested in data
	    - An agent is listed as the concerns to a model
		- A unique identifier is the only requirement
*/
type Agent interface {
	GetId() bson.ObjectId
}
