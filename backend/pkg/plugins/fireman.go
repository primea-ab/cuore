package plugins

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type FireMan struct {
}

func (l FireMan) Name() string {
	return "fire-man"
}

func (l FireMan) Type() string {
	return "bar"
}

func (l FireMan) Interval() time.Duration {
	return time.Second * 10
}

func (l FireMan) CollectData() string {
	resp, err := http.Get("http://localhost:8000/fire-man-status")
	if err != nil {
		return "error"
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	players := string(bodyBytes)
	return players
}
