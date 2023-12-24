package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	BaseURL    string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	ServerPort string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
}

var cfg Config

func New() *Config {
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
		return nil
	}
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "address and port to run server")
	flag.StringVar(&cfg.ServerPort, "a", cfg.ServerPort, "address and port for result url")
	flag.Parse()
	return &cfg
}

func (c *Config) GetServerPort() string {
	return c.ServerPort
}

func (c *Config) GetBaseURL() string {
	return c.BaseURL
}

type Getter interface {
	GetServerPort() string
	GetBaseURL() string
}
