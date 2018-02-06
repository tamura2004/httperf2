package infra_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/infra"
)

func ExampleNewConfigClient() {
	infra.InitConfig("config.test")

	fmt.Println(infra.Config.Client)
	fmt.Println(infra.Config.Scinario)
	fmt.Println(infra.Config.Target)
	// Output:
	// {true 1024 http://127.0.0.1:8888/ 13 dummy_agent/2.0 dummy_user dummy_pass}
	// {3 5 {110ns} {7ms} {0s}}
	// {http://example.net}

}
