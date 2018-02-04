package domain_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra"
	"testing"
)

func TestNewJob(t *testing.T) {
	infra.InitMockTime()

	j := domain.NewJob(2)

	got := j.Id
	want := 2

	if got != want {
		t.Errorf("bad job id got %s want %s", got, want)
	}
}

func ExampleNewJob() {
	infra.InitMockTime()
	j := domain.NewJob(3)
	fmt.Println(j)
	// Output:
	// &{3 2018-02-01 08:00:00 +0000 UTC {0 {{false 0  0} {0 0 {0s} {0s}} {}} <nil> <nil> <nil>}}
}
