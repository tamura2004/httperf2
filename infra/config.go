package infra

import (
	"github.com/BurntSushi/toml"
	"github.com/tamura2004/httperf2/domain"
	"log"
)

var config domain.Config

func NewConfig() domain.Config {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
