package data

/*
	Describes the ability to be interested in data
	    - An agent is listed as the concerns to a model
		- A unique identifier is the only requirement
*/
type Agent interface {
	GetID() ID
}
