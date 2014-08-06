package config

import (
	"net/http"

	"github.com/elos/server/routes"
)

func SetupRoutes() {
	http.HandleFunc("/v1/authenticate", routes.Authenticate)
}
