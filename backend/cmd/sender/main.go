package main

import (
	"fmt"
	"lethe.se/vito/cuore/pkg/serv"
	"lethe.se/vito/cuore/pkg/plugins"
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
	for _, p := range plg {
		fmt.Println("Loading plugin " + p)
		loadedPlugin := plugins.GetPlugin(p)
		fmt.Println(loadedPlugin.CollectData())
	}
}

