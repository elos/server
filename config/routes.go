package config

import "github.com/elos/server/routes"

func (s *Server) SetupRoutes() {
	if s.Store == nil {
		return
	}

	UsersPost := &routes.UsersPostHandler{Store: s.Store}
	EventsPost := &routes.EventsPostHandler{Store: s.Store}
	AuthenticateGet := &routes.AuthenticateGetHandler{Store: s.Store}

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
	routes.DefaultClientDataHub = s
}
