package serv

import (
	"encoding/json"
	"lethe.se/vito/cuore/pkg/model"
	"log"
	"os"
)

func FetchConfig() model.Config {
	f, err := os.Open("config/default.json")
	if err != nil {
		log.Fatal("Could not open config file")
	}
	defer f.Close()

	var cfg model.Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal("Could not decode config file")
	}

	return cfg
}