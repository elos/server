package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/elos/api"
	apimiddleware "github.com/elos/api/middleware"
	apiroutes "github.com/elos/api/routes"
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/data/builtin/mongo"
	emiddleware "github.com/elos/ehttp/middleware"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
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

	Sandbox(db)

	api := api.New(&api.Middleware{
		Cors: new(apimiddleware.Cors),
		Log:  emiddleware.LogRequest,
		SessionAuth: &apimiddleware.SessionAuth{
			DB:                  db,
			UnauthorizedHandler: apiroutes.Unauthorized,
		},
	}, &api.Services{
		DB:     db,
		Agents: hub,
	})
	apiServer := serve.New(
		&serve.Opts{
			Handler: api,
		},
	)

	/*
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
	*/

	//hub.StartAgent(appServer)
	hub.StartAgent(apiServer)

	go autonomous.HandleIntercept(hub.Stop)
	hub.WaitStop()
}

// --- Sandbox {{{
func Sandbox(db data.DB) {
	u := models.NewUser()
	u.SetID(db.NewID())

	c := models.NewCredential()
	c.SetID(db.NewID())
	c.Public = "public"
	c.Private = "private"

	c.SetOwner(u)
	u.IncludeCredential(c)

	if err := db.Save(u); err != nil {
		log.Fatal(err)
	}

	if err := db.Save(c); err != nil {
		log.Fatal(err)
	}

	p := models.NewPerson()
	p.SetID(db.NewID())
	if err := db.Save(p); err != nil {
		log.Fatal(err)
	}

	cal, err := p.CalendarOrCreate(db)
	if err != nil {
		log.Fatal(err)
	}

	b, err := cal.BaseScheduleOrCreate(db)
	if err != nil {
		log.Fatal(err)
	}

	f := models.NewFixture()
	f.SetID(db.NewID())
	f.Name = "Test Base Fixture"
	f.StartTime = time.Now()
	f.EndTime = time.Now().Add(1 * time.Hour)
	if err := db.Save(f); err != nil {
		log.Fatal(err)
	}
	b.IncludeFixture(f)

	if err := db.Save(b); err != nil {
		log.Fatal(err)
	}

	if err := db.Save(cal); err != nil {
		log.Fatal(err)
	}
}

// --- }}}
