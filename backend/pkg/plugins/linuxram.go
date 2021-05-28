package plugins

import (
	"time"
)

type LinuxRam struct {

}

func (l LinuxRam) Name() string {
	return "linux-ram"
}

func (l LinuxRam) Type() string {
	return "ram"
}

func (l LinuxRam) Interval() time.Duration {
	return time.Second * 5
}


func (l LinuxRam) CollectData() string {
	return "ram baby"
}
