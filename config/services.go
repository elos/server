package config

import (
	"github.com/elos/models/user"
	"github.com/elos/server/agents"
)

var Outfitter *agents.Outfitter

func (s *Server) SetupServices() {
	if s.Store == nil {
		return
	}
	Outfitter = agents.NewOutfitter()
	go Outfitter.Run()

	iter, err := s.Store.NewQuery(UserKind).Execute()
	if err != nil {
	}

	u, _ := user.New(s.Store)

	for iter.Next(u) {
		agents.OutfitUser(Outfitter, s.Store, u)
	}

	if err := iter.Close(); err != nil {
	}
}
