package domain_test

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/infra/mock"
	"testing"
)

func TestNewJob(t *testing.T) {
	mock.InitTime()

	j := domain.NewJob(1, 2)

	got := j.JobId
	want := 2

	if got != want {
		t.Errorf("bad job id got %s want %s", got, want)
	}
}

func ExampleNewJob() {
	mock.InitTime()
	j := domain.NewJob(2, 3)
	fmt.Println(j)
	// Output:
	// {2 3 2018-02-01 08:00:00 +0000 UTC}
}
