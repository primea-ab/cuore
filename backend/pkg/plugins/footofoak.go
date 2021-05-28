package plugins

import (
	"time"
	"os/exec"
	"fmt"
	"strings"
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
	out, err := exec.Command("uptime", "-p").Output()
	if err != nil {
		fmt.Println(err)
	}
	uptime := fmt.Sprintf("%s", out)

	return "No new foot yet (" + strings.ReplaceAll(uptime, "\n", "") + ")"
}