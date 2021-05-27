package main

import (
	"fmt"
	"lethe.se/vito/cuore/pkg/serv"
	"lethe.se/vito/cuore/pkg/plugins"
	"time"
)

func main() {
	cfg := serv.FetchConfig()

	host := cfg.Sender.Host
	port := cfg.Combined.Port
	plg := cfg.Sender.Plugins

	loadPlugins(plg)

	fmt.Println("Starting send of data to host '" + host + "' using port '" + port + "'")
}

func loadPlugins(plg []string) {
	loadedPlugins := make([]plugins.Plugin, len(plg))
	for i, p := range plg {
		fmt.Println("Loading plugin " + p)
		loadedPlugin := plugins.GetPlugin(p)
		loadedPlugins[i] = loadedPlugin
	}
	var asdf int64 = 0

	for {
		callPlugins(loadedPlugins, asdf)
		asdf = asdf + 1
		time.Sleep(time.Millisecond * 2)
	}
}

func callPlugins(loadedPlugins []plugins.Plugin, i int64) {
	for _, p := range loadedPlugins {
		fmt.Println(p.Identifier())
	}
	fmt.Println(i)
}

