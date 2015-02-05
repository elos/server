package config

import (
	"github.com/elos/data"
	"github.com/elos/models/user"
	"github.com/elos/server/agents"
)

var Outfitter *agents.Outfitter

func SetupServices(s data.Store) {
	Outfitter = agents.NewOutfitter()
	go Outfitter.Run()

	iter, err := s.NewQuery(UserKind).Execute()
	if err != nil {
	}

	u, _ := user.New(s)

	for iter.Next(u) {
		agents.OutfitUser(Outfitter, s, u)
	}

	if err := iter.Close(); err != nil {
	}
}
