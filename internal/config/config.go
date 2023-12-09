package config

import (
	"flag"
	"fmt"
)

type Config struct {
	BaseURL    string
	ServerPort string
}

var cfg Config

func New() *Config {
	flag.StringVar(&cfg.BaseURL, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&cfg.ServerPort, "b", "localhost:8080", "address and port for result url")
	flag.Parse()
	fmt.Println(cfg.BaseURL, cfg.ServerPort)
	return &cfg
}

func (c *Config) GetPort() string {
	return c.ServerPort
}

func (c *Config) GetBaseURL() string {
	return c.BaseURL
}

type Getter interface {
	GetPort() string
	GetBaseURL() string
}
