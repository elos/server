package hub

type Envelope struct {
	Agent  Agent
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

/*
type Serializer func(data []byte) []interface{}

func UserSerializer(data []byte) []*models.User {
	var u models.User
	json.Unmarshal(data, &u)

	var a []*models.User
	a[0] = &u

	return a
}

func UsersSerializer(data []byte) []*models.User {
	var v []interface{}
	json.Unmarshal(data, &v)

	var a []*models.User

	for _, userData := range v {
		bytes, _ := json.Marshal(userData)
		user := UserSerializer(bytes)
		a = append(a, user...)
	}

	return a
}
*/

func Route(e Envelope, hc HubConnection) {
	/*
		Serializers := map[string]Serializer{
			"user":  UserSerializer,
			"users": UsersSerializer,
		}

		log.Printf("Envelope: %v", e)

		var v []interface{}
		for key, _ := range e.Data {
			bytes, _ := json.Marshal(e.Data[key])
			v := Serializers[key](bytes)
		}

		log.Printf("Structured data: %v", v)
		PrimaryHub.SendJson(hc.User, v)
	*/

	// Echo
	PrimaryHub.SendJson(hc.Agent, e)
}
