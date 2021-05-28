package model

type Datapoint struct {
	Identifier string `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	Timestamp  string `json:"timestamp"`
}