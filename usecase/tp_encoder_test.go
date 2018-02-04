package usecase

import (
	"fmt"
	"github.com/tamura2004/httperf2/infra"
)

func ExampleInitTpEncoder() {
	infra.InitTpEncoder()

	fmt.Println(tpEncoder)
	// Output:
	// -
}
