package config

import (
	"log"

	"github.com/elos/data"
	"github.com/elos/data/mongo"
)

func (s *Server) SetupDB(addr string) data.DB {
	db, err := mongo.NewDB(addr)

	if err != nil {
		log.Fatal(err)
	} else {
		Log("Database connection established")
	}

	return db
}
