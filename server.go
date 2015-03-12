package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/elos/autonomous"
	"github.com/elos/httpserver"
	"github.com/elos/mongo"
	"github.com/elos/stack"
)

var (
	addr                           = flag.String("h", "127.0.0.1", "IP Address to bind to")
	port                           = flag.Int("p", 8000, "Port to listen on")
	programName                    = filepath.Base(os.Args[0])
	hub         autonomous.Manager = autonomous.NewHub()
)

func setFlagUsage() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, " %s [-h=ADDR] [-p=NUM] \n", programName)
		flag.PrintDefaults()
	}
}

func main() {
	setFlagUsage()

	go hub.Start()
	hub.WaitStart()

	mongo.Runner.ConfigFile = "mongo.conf"
	go hub.StartAgent(mongo.Runner)

	store := stack.SetupStore("localhost")

	server := httpserver.New(*addr, *port, store)
	go hub.StartAgent(server)
	server.WaitStart()

	stack.Sandbox(store)
	stack.SetupAgents(store)

	go autonomous.HandleIntercept(hub.Stop)
	hub.WaitStop()
}
