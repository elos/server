package config

import (
	"github.com/elos/server/managers"
	"github.com/elos/server/models/user"
)

var Outfitter *managers.Outfitter

func (s *Server) SetupServices() {
	if s.DB == nil {
		return
	}
	Outfitter = managers.NewOutfitter()
	go Outfitter.Run()

	iter, err := s.DB.NewQuery(UserKind).Execute()
	if err != nil {
	}

	u := user.New()

	for iter.Next(u) {
		managers.OutfitUser(Outfitter, s.DB, u)
	}

	if err := iter.Close(); err != nil {
	}
}
