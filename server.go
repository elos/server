package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/elos/server/config"
)

// Request wrapper that each request
func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	programName := filepath.Base(os.Args[0])

	var (
		host = flag.String("h", "127.0.0.1", "IP Address to bind to")
		port = flag.Int("p", 8000, "Port to listen on")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, " %s [-h=ADDR] [-p=NUM]\n", programName)
		flag.PrintDefaults()
	}

	// Setup Server
	config.SetupRoutes()

	config.SetupRedis()
	defer config.ShutdownRedis()

	// Start serving requests
	serving_url := fmt.Sprintf("%s:%d", *host, *port)

	log.Printf("Serving at http://%s", serving_url)

	log.Fatal(http.ListenAndServe(serving_url, LogRequest(http.DefaultServeMux)))
}
