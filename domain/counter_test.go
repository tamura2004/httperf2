package domain_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra/mock"
	"time"
)

func ExampleNewCounter() {
	mock.InitTime()

	c := domain.NewCounter()
	c.AddDuration(10 * time.Second)
	c.IncTp()
	c.IncMulti()

	fmt.Println(c)
	// Output:
	// &{map[2018-02-01,08:00:00:1] map[2018-02-01,08:00:1] 1 1 10s}
}
