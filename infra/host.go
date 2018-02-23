package infra

import (
	"github.com/tamura2004/httperf2/interfaces/presenters"
	"log"
	"os"
)

type host struct{}

var Host host

func (*host) Name() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func InitHost() {
	presenters.InitHost(&host{})
	log.Println("init host")

}
