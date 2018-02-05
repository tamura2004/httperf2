package main

import (
	"github.com/tamura2004/httperf2/infra"
	"github.com/tamura2004/httperf2/interfaces/presenters"
	"github.com/tamura2004/httperf2/usecase"
)

func main() {
	infra.Init()
	presenters.Init()
	usecase.InitResultChan()

	go infra.Netstat()

	usecase.Run()
}
