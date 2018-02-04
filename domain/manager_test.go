package domain_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra"
)

func ExampleNewManager() {
	infra.InitMockTime()

	config := domain.Config{}
	client := infra.NewClient(config.Client)
	m := domain.NewManager(config, client)

	fmt.Println(m.Config)
	// Output:
	// {{false 0  0} {0 0 {0s} {0s}} {}}
}
