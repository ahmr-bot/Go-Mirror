package routers

import (
	"encoding/json"
	"github.com/ahmr-bot/MirrorsAPI/pkg"
)

var redisClient = pkg.RedisClient()

type Location struct {
	ServerLocation any `json:"server_location"`
}
type Files struct {
	Name string
	Url  string
	Size int64
}
type Config struct {
	Directories []struct {
		Path        string `json:"path"`
		Description string `json:"description"`
		Image       string `json:"image"`
	} `json:"directories"`
	ServerAddr     any `json:"server_addr"`
	ListenPort     any `json:"listen_port"`
	ServerLocation any `json:"server_location"`
}
type Dir struct {
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Path        string   `json:"path"`
	Files       []Files  `json:"files"`
	Directories []string `json:"directories"`
	location    any      `json:"server_location"`
}

func (c Dir) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}
