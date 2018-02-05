package infra

import (
	"github.com/BurntSushi/toml"
	"github.com/tamura2004/httperf2/usecase"
	"log"
)

type config struct {
	Client   clientConfig
	Scinario usecase.ScinarioConfig
	Target   usecase.TargetConfig
}

var Config config

func InitConfig(fileName string) {
	_, err := toml.DecodeFile(fileName, &Config)
	if err != nil {
		log.Fatal(err)
	}

	ClientConfig = Config.Client
	usecase.InitScinarioConfig(Config.Scinario)
	usecase.InitTargetConfig(Config.Target)

	log.Println("init config")
	log.Println(Config.Client)
	log.Println(Config.Scinario)
	log.Println(Config.Target)
}
