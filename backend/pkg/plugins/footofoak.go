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
	return "uptime"
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

	return strings.ReplaceAll(strings.ReplaceAll(uptime, "\n", ""), "up ", "")
}