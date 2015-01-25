package config

import (
	"github.com/elos/data"
	"github.com/elos/server/routes"
)

func SetupRoutes(db data.DB) {

	UsersPost := &routes.UsersPostHandler{DB: db}
	EventsPost := &routes.EventsPostHandler{DB: db}
	AuthenticateGet := &routes.AuthenticateGetHandler{DB: db}

	var RoutesMap = routes.HandlerMap{
		"v1": routes.HandlerMap{
			"users": routes.HandlerMap{
				routes.POST: UsersPost,
			},
			"events": routes.HandlerMap{
				routes.POST: EventsPost,
			},
			"authenticate": routes.HandlerMap{
				routes.GET: AuthenticateGet,
			},
		},
	}

	routes.SetupHTTPRoutes(RoutesMap)
	routes.DefaultClientDataHub = ClientDataHub
}
