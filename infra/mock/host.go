package mock

import (
	"github.com/tamura2004/httperf2/interfaces/presenters"
)

type host struct{}

func (*host) Name() string {
	return "dummy"
}

func InitHost() {
	presenters.InitHost(&host{})
}
