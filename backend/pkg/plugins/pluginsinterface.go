package plugins

import (
	"lethe.se/vito/cuore/pkg/model"
)

type Plugin interface {
	Identifier() string
	CollectData() model.Datapoint
}