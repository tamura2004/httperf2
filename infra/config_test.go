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
	// {true 1024  0}
	// {0 3 {100ns} {3ms}}
	// {https://ogisui.azurewebsites.net}

}
