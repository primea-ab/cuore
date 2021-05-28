package plugins

import (
	"time"
)

type Plugin interface {
	Name() string
	Type() string
	Interval() time.Duration
	CollectData() string
}