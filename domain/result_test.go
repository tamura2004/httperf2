package domain_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra/mock"
)

func ExampleNewResult() {
	mock.InitTime()
	j := domain.NewJob(2, 1)
	r := domain.NewResult(j, domain.START)

	fmt.Println(r.Job)
	fmt.Println(r.Check)
	fmt.Println(r.Time)
	fmt.Println(r.Duration)
	// Output:
	// {2 1 2018-02-01 08:00:00 +0000 UTC}
	// START
	// 2018-02-01 08:00:00 +0000 UTC
	// 10s

}
