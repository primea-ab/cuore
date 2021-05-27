package main

import (
	"fmt"
	"lethe.se/vito/cuore/pkg/serv"
)

func main() {
	cfg := serv.FetchConfig()

	host := cfg.Server.Host
	port := cfg.Server.Port

	fmt.Println("Starting send of data to host '" + host + "' using port '" + port + "'")
}

