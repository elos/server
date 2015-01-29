package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/data/mongo"
	"github.com/elos/server/util/logging"
)

type Server struct {
	*autonomous.AgentHub
	*autonomous.Core

	host    string
	port    int
	verbose bool

	data.Store
}

func NewServer(host string, port int, verbose bool) *Server {
	s := &Server{
		Core:     autonomous.NewCore(),
		AgentHub: autonomous.NewAgentHub(),
		host:     host,
		port:     port,
		verbose:  verbose,
	}

	go s.AgentHub.Run()

	return s
}

func (s *Server) Run() {
	s.startup()
	stop := *s.Core.StopChannel()
Run:
	for {
		select {
		case _ = <-stop:
			s.shutdown()
			break Run
		}
	}
}

func (s *Server) startup() {
	s.Core.Startup()

	if err := mongo.StartDatabaseServer(); err != nil {
		log.Fatal("Failed to start mongo, server can not start")
	}

	SetupLog(s.verbose)
	db := s.SetupDB("localhost")
	sc := s.SetupModels()
	s.Store = data.NewStore(db, sc)
	s.SetupRoutes()
	s.SetupServices()
	// s.Sandbox()

	go StartServer(s.host, s.port)
}

func (s *Server) shutdown() {
	log.Print("Shutting down server")
	// ShutdownDB()
	// mongo.StopDatabaseServer(sig)
	s.Core.Shutdown()
}

func StartServer(host string, port int) {
	serving_url := fmt.Sprintf("%s:%d", host, port)

	logging.Log.Logsf("Serving at http://%s", serving_url)

	log.Fatal(http.ListenAndServe(serving_url, logging.LogRequest(http.DefaultServeMux)))
}
