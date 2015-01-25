package models

import (
	"github.com/elos/data"
	"github.com/elos/schema"
)

type ModelConstructor func() schema.Model

var RegisteredModels = make(map[data.Kind]ModelConstructor)

func Register(k data.Kind, c ModelConstructor) {
	RegisteredModels[k] = c
}
