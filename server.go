package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/elos/autonomous"
	"github.com/elos/httpserver"
	"github.com/elos/mongo"
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
		fmt.Fprintf(os.Stderr, " %s [-h=ADDR] [-p=NUM] \n", programName)
		flag.PrintDefaults()
	}

	var hub autonomous.Manager = autonomous.NewHub()

	go hub.Start()
	hub.WaitStart()

	mongo.Runner.ConfigFile = "mongo.conf"
	go hub.StartAgent(mongo.Runner)

	store := stack.SetupStore("localhost")

	server := httpserver.New(*addr, *port, store)
	go hub.StartAgent(server)
	server.WaitStart()

	log.Print("HTTPServer started")

	go autonomous.HandleIntercept(hub.Stop)

	stack.Sandbox(store)
	stack.SetupServices(store)

	hub.WaitStop()
}
