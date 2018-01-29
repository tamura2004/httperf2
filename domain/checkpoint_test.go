package domain_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
)

func ExampleCheckPoint() {
	cs := []domain.CheckPoint{
		domain.START,
		domain.CONNECT,
		domain.FINISH,
		domain.CheckPoint(4),
		domain.CheckPoint(0),
	}
	fmt.Println(cs)
	// Output:
	// [START CONNECT FINISH undefined undefined]
}
