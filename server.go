package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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

	config.SetVerbosity(*verbose)
	config.SetupDB("localhost")
	config.SetupRoutes()
	config.SetupSockets()

	defer config.ShutdownDB()
	defer config.ShutdownSockets()
	defer db.StopMongo()

	StartServer(*host, *port)
}

func StartServer(host string, port int) {
	serving_url := fmt.Sprintf("%s:%d", host, port)

	log.Printf("Serving at http://%s", serving_url)

	log.Fatal(http.ListenAndServe(serving_url, util.LogRequest(http.DefaultServeMux)))
}
