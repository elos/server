package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	data.DB
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
	log.Printf("THE SERVER's STOP CHANNEL IS %v", stop)
	for {
		select {
		case _ = <-stop:
			log.Print("STOP CHANNEL RECIEVED SOMETHING")
			s.shutdown()
			break
		}
	}
}

func (s *Server) startup() {
	s.Core.Startup()

	if err := mongo.StartDatabaseServer(); err != nil {
		log.Fatal("Failed to start mongo, server can not start")
	}

	SetupLog(s.verbose)
	s.SetupDB("localhost")
	s.SetupModels()
	s.SetupRoutes()
	s.SetupServices()
	s.Sandbox()

	go StartServer(s.host, s.port)
}

func (s *Server) shutdown() {
	log.Print("shutting down")
	logging.Log.Logs("Shutting down server")
	// ShutdownDB()
	// mongo.StopDatabaseServer(sig)
	s.Core.Shutdown()
	os.Exit(0)
}

func StartServer(host string, port int) {
	serving_url := fmt.Sprintf("%s:%d", host, port)

	logging.Log.Logsf("Serving at http://%s", serving_url)

	log.Fatal(http.ListenAndServe(serving_url, logging.LogRequest(http.DefaultServeMux)))
}
