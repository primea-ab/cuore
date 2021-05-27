package model

type Config struct {
	Combined struct {
		Port string `json:"port"`
	} `json:"Combined"`
	Sender struct {
		Host string `json:"host"`
		Plugins []string `json:"plugins"`
	} `json:"Sender"`

}
