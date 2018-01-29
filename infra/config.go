package infra

import (
	"github.com/BurntSushi/toml"
	"github.com/tamura2004/httperf2/domain"
	"log"
)

type Config struct {
	Client   domain.Client
	Target   domain.Target
	Scinario domain.Scinario
}

var config Config

func NewConfig() Config {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
