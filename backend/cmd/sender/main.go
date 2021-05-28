package main

import (
	"fmt"
	"lethe.se/vito/cuore/pkg/serv"
	"lethe.se/vito/cuore/pkg/plugins"
	"lethe.se/vito/cuore/pkg/model"
	"time"
	"encoding/json"
	"net/http"
	"bytes"
)

func main() {
	cfg := serv.FetchConfig()

	identifier := cfg.Sender.Name
	receiver := cfg.Sender.Receiver
	port := cfg.Combined.Port
	plg := cfg.Sender.Plugins

	loadPlugins(plg, identifier, receiver, port)

	fmt.Println("Starting send of data to receiver '" + receiver + "' using port '" + port + "' using identifier '" + identifier + "'")
}

func loadPlugins(plg []string, identifier string, receiver string, port string) {
	for _, p := range plg {
		fmt.Println("Loading plugin " + p)
		loadedPlugin := plugins.GetPlugin(p)
		go runPlugin(loadedPlugin, identifier, receiver, port)
	}

	fmt.Println("All plugins loaded")
	select {}
}

func runPlugin(plugin plugins.Plugin, identifier string, receiver string, port string) {
	for {
		response := plugin.CollectData()
		datapoint := model.Datapoint{Identifier: identifier, Name: plugin.Name(), Status: response, Type: plugin.Type()}
		jsonResp, _ := json.Marshal(datapoint)

		res, err := http.Post("http://" + receiver + ":" + port + "/api", "application/json", bytes.NewBuffer(jsonResp))
		if err != nil {
			fmt.Println("Error posting plugin data to server ", err)
		}

		time.Sleep(plugin.Interval())
	}	
}

