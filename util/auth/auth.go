package auth

import (
	"net/http"
	"strings"

	"github.com/elos/server/data"
	"github.com/elos/server/models/user"
)

type RequestAuthenticator func(*http.Request) (data.Agent, bool, error)

func AuthenticateRequest(r *http.Request) (data.Agent, bool, error) {
	// Use the WebSocket protocol header to identify and authenticate the user
	id, key := ExtractCredentials(r)

	if id == "" || key == "" {
		return nil, false, nil
	}

	return user.Authenticate(id, key)
}

func ExtractCredentials(r *http.Request) (string, string) {
	tokens := strings.Split(r.Header.Get("Sec-WebSocket-Protocol"), "-")
	// Query Parameter Method of Authentication
	/*
		id := r.FormValue("id")
		key := r.FormValue("key")
	*/
	if len(tokens) != 2 {
		return "", ""
	} else {
		return tokens[0], tokens[1]
	}
}

func ExtractProtocolHeader(r *http.Request) *http.Header {
	protocol := http.Header{
		"Sec-WebSocket-Protocol": []string{r.Header.Get("Sec-WebSocket-Protocol")},
	}

	return &protocol
}
