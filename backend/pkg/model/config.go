package model

type Config struct {
	Combined struct {
		Port string `json:"port"`
	} `json:"Combined"`
	Sender struct {
		Name string `json:"name"`
		Receiver string `json:"receiver"`
		Plugins []string `json:"plugins"`
	} `json:"Sender"`

}
