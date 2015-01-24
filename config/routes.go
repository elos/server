package config

import (
	"github.com/elos/server/data"
	"github.com/elos/server/routes"
)

func SetupRoutes(db data.DB) {

	UsersPost := &routes.UsersPostHandler{}
	EventsPost := &routes.UsersPostHandler{}
	AuthenticateGet := &routes.UsersPostHandler{}
	UsersPost.DB = db
	EventsPost.DB = db
	AuthenticateGet.DB = db

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
