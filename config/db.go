package config

import (
	"log"

	"github.com/elos/data/mongo"
)

func (s *Server) SetupDB(addr string) {
	db, err := mongo.NewDB(addr)

	if err != nil {
		log.Fatal(err)
	} else {
		Log("Database connection established")
	}

	s.DB = db
}

func (s *Server) StopDB() {
	s.DB = nil
}
