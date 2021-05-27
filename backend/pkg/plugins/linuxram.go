package plugins

import (
	"lethe.se/vito/cuore/pkg/model"
)

type LinuxRam struct {

}

func (l LinuxRam) Identifier() string {
	return "linux-ram"
}

func (l LinuxRam) CollectData() model.Datapoint {
	return model.Datapoint{Name:"test", Status:"teststatus", Type:"testtype"}
}