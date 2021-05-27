package model

type Datapoint struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
}