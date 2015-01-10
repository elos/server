package data

/*
	Data structures for the transfer of data
	For implementations of this functionality see elos/server/data/transfer
*/

// Inbound
type Envelope struct {
	Action string                          `json:"action"`
	Data   map[Kind]map[string]interface{} `json:"data"`
}

func NewEnvelope(action string, data map[Kind]map[string]interface{}) *Envelope {
	return &Envelope{
		Action: action,
		Data:   data,
	}
}

// Outbound
type Package struct {
	Action string
	Data   map[Kind]Model
}

func NewPackage(action string, data map[Kind]Model) *Package {
	return &Package{
		Action: action,
		Data:   data,
	}
}
