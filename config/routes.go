package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/routes"
)

var RoutesMap = routes.HandlerMap{
	"v1": routes.HandlerMap{
		"users": routes.HandlerMap{
			routes.POST: routes.UsersPost,
		},
		"events": routes.HandlerMap{
			routes.POST: routes.EventsPost,
		},
		"authenticate": routes.HandlerMap{
			routes.GET: routes.AuthenticateGet,
		},
	},
}

func SetupRoutes(db data.DB) {
	routes.SetupRoutes(RoutesMap)
	routes.SetDB(db)
}
