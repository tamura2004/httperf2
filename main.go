package main

import (
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra"
	"github.com/tamura2004/httperf2/interfaces/presenters"
	"github.com/tamura2004/httperf2/usecase"
)

func main() {
	infra.InitLogger()

	go infra.Netstat()

	config := infra.NewConfig()
	client := infra.NewClient(config.Client)
	m := domain.NewManager(config, client)

	reporter := usecase.NewReporter(
		infra.NewFileFactory(),
		presenters.NewResultEncoder(),
		presenters.NewTpEncoder(),
		infra.GetHostname(),
	)

	collector := usecase.NewCollector(
		m.Result,
		reporter,
	)

	manager := usecase.NewManager(
		m,
		collector,
	)

	manager.Run()
}
