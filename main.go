package main

import (
	"fmt"

	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra"
	"github.com/tamura2004/httperf2/interfaces/controllers"
	"github.com/tamura2004/httperf2/usecase"
)

func main() {
	infra.InitLogger()

	config := infra.NewConfig()
	client := infra.NewClient(config.Client)

	manager := usecase.NewManager()
	manager.Run()

	jc := controllers.NewJobRunner(client, domain.NewJob(&config.Target))

	ch := make(chan *domain.Result)

	go jc.Run(ch)

	for result := range ch {
		fmt.Println(result.Check)
	}
}
