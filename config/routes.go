package config

import (
	"net/http"

	"github.com/elos/server/routes"
)

func SetupRoutes() {
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/static/", routes.Static)
	http.HandleFunc("/v1/users", routes.Users)
	http.HandleFunc("/v1/events", routes.Events)
	http.HandleFunc("/v1/authenticate", routes.Authenticate)
}
