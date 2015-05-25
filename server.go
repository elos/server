package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/elos/app"
	"github.com/elos/app/middleware"
	"github.com/elos/app/routes"
	"github.com/elos/autonomous"
	"github.com/elos/data/builtin/mongo"
	"github.com/elos/ehttp/builtin"
	emiddleware "github.com/elos/ehttp/middleware"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
	"github.com/gorilla/context"
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

	db, err := models.MongoDB("localhost")
	if err != nil {
		log.Fatal(err)
	}

	//api := api.New(db, hub)
	//apiServer := serve.New(&serve.Opts{Handler: api})

	sessions := builtin.NewSessions()
	app := app.New(&app.Middleware{
		Log:         emiddleware.LogRequest,
		SessionAuth: middleware.NewSessionAuth(db, sessions, routes.SessionsSignIn),
	}, &app.Services{
		Agents:   hub,
		DB:       db,
		Sessions: sessions,
	})

	appServer := serve.New(&serve.Opts{
		Handler: context.ClearHandler(app),
		Port:    8080,
	})

	hub.StartAgent(appServer)

	u := &models.User{}
	u.SetID(db.NewID())
	u.Key = "password"
	if err := db.Save(u); err != nil {
		log.Fatal(err)
	}

	go autonomous.HandleIntercept(hub.Stop)
	hub.WaitStop()
}
