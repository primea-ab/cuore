package plugins

import (
	"log"
)

func GetPlugin(identifier string) Plugin {
	switch identifier {
		case "linux-ram":
			return LinuxRam{}
		case "foot-of-oak":
			return FootOfOak{}
		default:
			log.Fatalf("Could not find plugin %v", identifier)
			return nil
	}
}