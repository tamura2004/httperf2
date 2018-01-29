package main

import (
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infrastructure"
	"github.com/tamura2004/httperf2/interfaces/controllers"
)

func main() {

	client := infrastructure.NewClient(
		domain.Client{
			InsecureSkipVerify:  true,
			MaxIdleConnsPerHost: 2018,
			Proxy:               "http://127.0.0.1:8888/",
			Bps:                 128,
		},
	)

	jc := controllers.NewJobController(client, *domain.NewJob(domain.Target{
		Url: "http://login.microsoftonline.com/",
	}))

	jc.Run()

}
