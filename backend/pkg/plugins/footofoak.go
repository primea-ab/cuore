package plugins

import (
	"lethe.se/vito/cuore/pkg/model"
)

type FootOfOak struct {

}

func (l FootOfOak) Identifier() string {
	return "foot-of-oak"
}

func (l FootOfOak) CollectData() model.Datapoint {
	return model.Datapoint{Name:"name", Status:"No new foot yet :(", Type:"oaky-foot"}
}