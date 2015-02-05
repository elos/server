package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/elos/autonomous"
	"github.com/elos/data/mongo"
	"github.com/elos/stack"
)

func main() {
	programName := filepath.Base(os.Args[0])

	var (
		addr = flag.String("h", "127.0.0.1", "IP Address to bind to")
		port = flag.Int("p", 8000, "Port to listen on")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, " %s [-h=ADDR] [-p=NUM] [-v=FLAG]\n", programName)
		flag.PrintDefaults()
	}

	if err := mongo.StartDatabaseServer(); err != nil {
		log.Fatal("Failed to start mongo, server can not start")
	}

	store := stack.SetupStore("localhost")

	manager := autonomous.NewAgentHub()

	httpserver := stack.NewHTTPServer(*addr, *port, store)

	go manager.StartAgent(httpserver)
	go HandleSignals(manager.Stop)

	stack.Sandbox(store)
	stack.SetupServices(store)

	manager.Run()
}

func HandleSignals(end func()) {
	// Intercept sigterm
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	// Block on this channel
	/*sig*/ _ = <-signalChannel

	end()
}
