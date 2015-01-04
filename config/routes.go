package config

import (
	"net/http"

	"github.com/elos/server/db"
	"github.com/elos/server/routes"
)

func SetupRoutes(db db.DB) {
	http.HandleFunc("/v1/users", routes.Users)
	http.HandleFunc("/v1/events", routes.Events)
	http.HandleFunc("/v1/authenticate", routes.Authenticate)
}
