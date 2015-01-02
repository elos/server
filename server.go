package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/elos/server/config"
	"github.com/elos/server/db"
	"github.com/elos/server/util"
)

func main() {
	programName := filepath.Base(os.Args[0])

	var (
		host    = flag.String("h", "127.0.0.1", "IP Address to bind to")
		port    = flag.Int("p", 8000, "Port to listen on")
		verbose = flag.Bool("v", true, "Whether to print verbose logs")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, " %s [-h=ADDR] [-p=NUM] [-v=FLAG]\n", programName)
		flag.PrintDefaults()
	}

	go HandleSignals()
	if err := db.StartMongo(); err != nil {
		log.Fatal("Failed to start mongo, server can not start")
	}

	config.SetVerbosity(*verbose)
	config.SetupDB("localhost")
	config.SetupRoutes()
	config.SetupSockets()

	StartServer(*host, *port)
}

func StartServer(host string, port int) {
	serving_url := fmt.Sprintf("%s:%d", host, port)

	log.Printf("Serving at http://%s", serving_url)

	log.Fatal(http.ListenAndServe(serving_url, util.LogRequest(http.DefaultServeMux)))
}

func HandleSignals() {
	// Intercept sigterm
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	sig := <-signalChannel
	Shutdown(sig)
}

func Shutdown(sig os.Signal) {
	log.Printf("Shutting down server")
	config.ShutdownDB()
	config.ShutdownSockets()
	// db.StopMongo(sig)
	os.Exit(0)
}
