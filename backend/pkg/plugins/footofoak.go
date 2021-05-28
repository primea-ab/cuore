package plugins

import (
	"time"
)

type FootOfOak struct {

}

func (l FootOfOak) Name() string {
	return "foot-of-oak"
}

func (l FootOfOak) Type() string {
	return "text"
}

func (l FootOfOak) Interval() time.Duration {
	return time.Second * 2
}

func (l FootOfOak) CollectData() string {
	return "No new foot yet :("
}