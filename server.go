package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/elos/autonomous"
	"github.com/elos/server/config"
)

func main() {
	programName := filepath.Base(os.Args[0])

	var (
		addr = flag.String("h", "127.0.0.1", "IP Address to bind to")
		port = flag.Int("p", 8000, "Port to listen on")
		verb = flag.Bool("v", true, "Whether to print verbose logs")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, " %s [-h=ADDR] [-p=NUM] [-v=FLAG]\n", programName)
		flag.PrintDefaults()
	}

	server := config.NewServer(*addr, *port, *verb)
	manager := autonomous.NewAgentHub()

	go manager.StartAgent(server)
	go HandleSignals(manager.Stop)

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
