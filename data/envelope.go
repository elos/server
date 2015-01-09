package data

// Inbound
type Envelope struct {
	Action string                          `json:"action"`
	Data   map[Kind]map[string]interface{} `json:"data"`
}

// Outbound
type Package struct {
	Action string
	Data   map[Kind]Model
}
