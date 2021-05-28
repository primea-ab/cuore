package plugins

import (
	"time"
	"fmt"
	"os/exec"
	"strings"
	"regexp"
)

type LinuxRam struct {

}

func (l LinuxRam) Name() string {
	return "linux-ram"
}

func (l LinuxRam) Type() string {
	return "bar"
}

func (l LinuxRam) Interval() time.Duration {
	return time.Second * 10
}

func (l LinuxRam) CollectData() string {
	cmd := exec.Command("free")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	ram := fmt.Sprintf("%s", out)
	memSplit := strings.SplitAfter(ram, "\n")
	rowSplit := regexp.MustCompile("\\s+").Split(memSplit[1], -1)
	return rowSplit[2] + " / " + rowSplit[1]
}
	