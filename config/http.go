package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/server/agents"
	"github.com/elos/server/routes"
	"github.com/elos/server/util/logging"
	"github.com/elos/transfer"
	"github.com/julienschmidt/httprouter"
)

type HTTPServer struct {
	host string
	port int

	*autonomous.Core
	*autonomous.AgentHub
	data.Store

	SocketRequests chan *agents.ClientDataAgent
}

func NewHTTPServer(host string, port int, s data.Store) *HTTPServer {
	return &HTTPServer{
		host:           host,
		port:           port,
		Core:           autonomous.NewCore(),
		AgentHub:       autonomous.NewAgentHub(),
		Store:          s,
		SocketRequests: make(chan *agents.ClientDataAgent, 10),
	}
}

func (s *HTTPServer) Run() {
	s.startup()
	stopChannel := s.Core.StopChannel()

	for {
		select {
		case a := <-s.SocketRequests:
			s.AgentHub.StartAgent(a)
		case _ = <-*stopChannel:
			s.shutdown()
			break
		}
	}
}

func (a *HTTPServer) startup() {
	a.Core.Startup()
	r := a.SetupRoutes()
	go a.Listen(r)
}

func (a *HTTPServer) shutdown() {
	a.Core.Shutdown()
}

func (s *HTTPServer) SetupRoutes() *httprouter.Router {
	router := httprouter.New()

	router.POST("/v1/users/", routes.Auth(routes.Post(UserKind, routes.Values("name")), s.Store))

	router.POST("/v1/events/", routes.Auth(routes.Post(EventKind, routes.Values("name")), s.Store))

	router.GET("/v1/authenticate", routes.Auth(routes.WebSocket(transfer.DefaultWebSocketUpgrader, s.SocketRequests), s.Store))

	return router
}

func (a *HTTPServer) Listen(r *httprouter.Router) {
	serving_url := fmt.Sprintf("%s:%d", a.host, a.port)

	log.Print("Serving at http://%s", serving_url)

	log.Fatal(http.ListenAndServe(serving_url, logging.LogRequest(http.DefaultServeMux)))
}
