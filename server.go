package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/elos/autonomous"
	"github.com/elos/data/mongo"
	"github.com/elos/server/config"
	"github.com/elos/server/managers"
	"github.com/elos/server/util/logging"
)

var programName string

func main() {
	programName = filepath.Base(os.Args[0])

	var (
		_ = flag.String("h", "127.0.0.1", "IP Address to bind to")
		_ = flag.Int("p", 8000, "Port to listen on")
		_ = flag.Bool("v", true, "Whether to print verbose logs")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, " %s [-h=ADDR] [-p=NUM] [-v=FLAG]\n", programName)
		flag.PrintDefaults()
	}

	server := NewServer("127.0.0.1", 8000, true)
	manager := managers.NewAgentHub()
	go manager.StartAgent(server)
	manager.Run()
}

type Server struct {
	*autonomous.BaseAgent
	host    string
	port    int
	verbose bool
}

func NewServer(host string, port int, verbose bool) *Server {
	return &Server{
		BaseAgent: autonomous.NewBaseAgent(),
		host:      host,
		port:      port,
		verbose:   verbose,
	}
}

func (s *Server) Run() {
	s.startup()
	stopChannel := s.BaseAgent.StopChannel()

	for {
		select {
		case _ = <-*stopChannel:
			s.shutdown()
			break
		}
	}
}

func (s *Server) startup() {
	s.BaseAgent.Startup()
	go HandleSignals(s)

	if err := mongo.StartDatabaseServer(); err != nil {
		log.Fatal("Failed to start mongo, server can not start")
	}

	config.SetupLog(s.verbose)
	config.SetupDB("localhost")
	config.SetupClientDataHub()
	config.SetupModels(config.DB)
	config.SetupRoutes(config.DB)
	config.SetupServices(config.DB)
	config.Sandbox(config.DB)

	StartServer(s.host, s.port)
}

func (s *Server) shutdown() {
	logging.Log.Logs(programName, "Shutting down server")
	config.ShutdownDB()
	// config.ShutdownClientDataHub()
	// mongo.StopDatabaseServer(sig)
	os.Exit(0)
	s.BaseAgent.Shutdown()
}

func StartServer(host string, port int) {
	serving_url := fmt.Sprintf("%s:%d", host, port)

	logging.Log.Logsf(programName, "Serving at http://%s", serving_url)

	log.Fatal(http.ListenAndServe(serving_url, logging.LogRequest(http.DefaultServeMux)))
}

func HandleSignals(s *Server) {
	// Intercept sigterm
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	/*sig*/ _ = <-signalChannel
	//Shutdown(sig)
	s.Stop()
	os.Exit(0)
}
