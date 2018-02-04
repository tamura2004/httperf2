package usecase

import (
	"fmt"
	"github.com/tamura2004/httperf2/infra"
)

func ExampleInitFileFactory() {
	infra.InitFileFactory()

	fmt.Println(fileFactory)
	// Output:
	// -
}
