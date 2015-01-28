package config

import "github.com/elos/server/routes"

func (s *Server) SetupRoutes() {
	if s.DB == nil {
		return
	}

	UsersPost := &routes.UsersPostHandler{DB: s.DB}
	EventsPost := &routes.EventsPostHandler{DB: s.DB}
	AuthenticateGet := &routes.AuthenticateGetHandler{DB: s.DB}

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
