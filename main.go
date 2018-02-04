package main

import (
	"github.com/tamura2004/httperf2/infra"
	"github.com/tamura2004/httperf2/interfaces/presenters"
	"github.com/tamura2004/httperf2/usecase"
)

func main() {
	infra.Init()

	usecase.InitResultChan()

	presenters.InitResultEncoder()
	presenters.InitTpEncoder()
	presenters.InitAveragePrinter()

	go infra.Netstat()

	reporter := usecase.NewReporter()
	collector := usecase.NewCollector(reporter)
	manager := usecase.NewManager(collector)

	manager.Run()
}
