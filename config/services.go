package config

import (
	"github.com/elos/models/user"
	"github.com/elos/server/managers"
)

var Outfitter *managers.Outfitter

func (s *Server) SetupServices() {
	if s.Store == nil {
		return
	}
	Outfitter = managers.NewOutfitter()
	go Outfitter.Run()

	iter, err := s.Store.NewQuery(UserKind).Execute()
	if err != nil {
	}

	u, _ := user.New(s.Store)

	for iter.Next(u) {
		managers.OutfitUser(Outfitter, s.Store, u)
	}

	if err := iter.Close(); err != nil {
	}
}
