package plugins

import (
	"time"
	//"fmt"
	//"os/exec"
	//"strings"
	//"regexp"
)

type LinuxCPU struct {

}

func (l LinuxCPU) Name() string {
	return "linux-cpu"
}

func (l LinuxCPU) Type() string {
	return "text"
}

func (l LinuxCPU) Interval() time.Duration {
	return time.Second * 30
}

func (l LinuxCPU) CollectData() string {
	/*out, err := exec.Command("top", "-h").Output()
	if err != nil {
		fmt.Println(err)
	}
	ram := fmt.Sprintf("%s", out)
	memSplit := strings.SplitAfter(ram, "\n")
	rowSplit := regexp.MustCompile("\\s+").Split(memSplit[1], -1)
	return rowSplit[2] + " / " + rowSplit[1]*/
	return "N/A"
}
	